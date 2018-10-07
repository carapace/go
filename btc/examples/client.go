package main

import (
	"fmt"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/carapace/go/btc"
	"github.com/carapace/go/btc/db"
)

func main() {
	// Create a new config
	config := btc.NewDefaultConfig()

	// Use testnet
	config.Params = &chaincfg.TestNet3Params

	// Select wallet datastore
	sqliteDatastore, _ := db.Create(config.RepoPath)
	config.DB = sqliteDatastore

	// Create the wallet
	wallet, err := btc.NewSPVWallet(config)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Start it!
	wallet.Start()
}
