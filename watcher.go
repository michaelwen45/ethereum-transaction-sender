package main

import (
	"context"
	"github.com/HydroProtocol/nights-watch"
	"github.com/HydroProtocol/nights-watch/plugin"
	"github.com/HydroProtocol/nights-watch/structs"
	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"
)

func startNightWatch(ctx context.Context) {
	w := nights_watch.NewHttpBasedEthWatcher(ctx, config.EthereumNodeUrl)

	w.RegisterTxReceiptPlugin(plugin.NewTxReceiptPlugin(func(txAndReceipt *structs.RemovableTxAndReceipt) {
		if txAndReceipt.IsRemoved {
			return
		}

		spew.Dump(txAndReceipt)
		var log LaunchLog
		db.Where("hash = ?", txAndReceipt.Receipt.GetTxHash()).First(&log)

		if log.ID == 0 && log.From == "" {
			return
		}

		var result string

		if txAndReceipt.Receipt.GetResult() {
			result = "successful"
			handleLaunchLogStatus(&log, true)
		} else {
			result = "failed"
			handleLaunchLogStatus(&log, false)
		}

		logrus.Infof("tx %s result: %s", txAndReceipt.Receipt.GetTxHash(), result)
	}))

	_ = w.RunTillExit()
}
