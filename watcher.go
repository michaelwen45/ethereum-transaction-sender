package main

import (
	"context"
	"fmt"
	"git.ddex.io/lib/monitor"
	"github.com/HydroProtocol/nights-watch"
	"github.com/HydroProtocol/nights-watch/plugin"
	"github.com/HydroProtocol/nights-watch/structs"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"strconv"
	"sync"
	"time"
)

var redisClient *redis.Client
var lastSavedBlockNumber int
var updateBlockNumberMutex *sync.Mutex

func connectRedis() {
	opt, err := redis.ParseURL(config.RedisUrl)

	if err != nil {
		panic(err)
	}

	opt.MinIdleConns = 2
	opt.PoolSize = 2

	redisClient = redis.NewClient(opt)
	updateBlockNumberMutex = &sync.Mutex{}
}

func saveBlockNumber(blockNum int) error {
	updateBlockNumberMutex.Lock()
	defer updateBlockNumberMutex.Unlock()

	if blockNum <= lastSavedBlockNumber {
		return nil
	}

	blockNumString := strconv.Itoa(blockNum)
	err := redisClient.Set(config.RedisBlockNumberCacheKey, blockNumString, 0).Err()

	if err != nil {
		logrus.Warnf("fail when write %s = %s", config.RedisBlockNumberCacheKey, blockNumString)
	} else {
		lastSavedBlockNumber = blockNum
		logrus.Infof("save block number %s to redis", blockNumString)
	}

	return err
}

func getHighestSyncedBlock() int {
	rst, err := redisClient.Get(config.RedisBlockNumberCacheKey).Result()

	if err == redis.Nil {
		return 0
	} else if err != nil {
		panic(fmt.Sprintf("redis err when GetHighestSyncedBlock: %s", err))
	} else {
		blockNumber, err := strconv.ParseUint(rst, 10, 64)
		if err != nil {
			panic(fmt.Sprintf("invalid value of redis key(%s): %s", config.RedisBlockNumberCacheKey, rst))
		}

		logrus.Debugf("start calculate from %s: %d", config.RedisBlockNumberCacheKey, blockNumber)

		lastSavedBlockNumber = int(blockNumber)
		return int(blockNumber)
	}
}

func startNightWatch(ctx context.Context) {
	w := nights_watch.NewHttpBasedEthWatcher(ctx, config.EthereumNodeUrl)

	w.RegisterTxReceiptPlugin(plugin.NewTxReceiptPlugin(func(txAndReceipt *structs.RemovableTxAndReceipt) {
		if txAndReceipt.IsRemoved {
			return
		}

		_ = saveBlockNumber(int(txAndReceipt.Receipt.GetBlockNumber()))
		monitor.Value("block_number", float64(txAndReceipt.Receipt.GetBlockNumber()))

		var log LaunchLog
		db.Where("hash = ?", txAndReceipt.Receipt.GetTxHash()).First(&log)

		if log.ID == 0 && log.From == "" {
			return
		}

		var result string
		var err error
		gasUsed := 0
		executedAt := int(txAndReceipt.TimeStamp)

		// to get gasUsed, TODO: return gasUsed from receipt
		receipt, err := ethrpcClient.EthGetTransactionReceipt(txAndReceipt.Receipt.GetTxHash())

		if err != nil || receipt == nil || receipt.TransactionHash == "" {
			logrus.Errorf("get receipt gasUsed failed, err: %+v, receipt: %+v", err, receipt)
		} else {
			gasUsed = receipt.GasUsed
		}

		if txAndReceipt.Receipt.GetResult() {
			result = "successful"
			err = handleLaunchLogStatus(&log, true, gasUsed, executedAt)
		} else {
			result = "failed"
			err = handleLaunchLogStatus(&log, false, gasUsed, executedAt)
		}

		logrus.Infof("tx %s err: %+v result: %s", txAndReceipt.Receipt.GetTxHash(), err, result)
	}))

	for {
		err := w.RunTillExitFromBlock(uint64(getHighestSyncedBlock()))

		if err != nil {
			logrus.Errorf("watcher error: %+v", err)
			time.Sleep(1 * time.Second)
		}
	}
}
