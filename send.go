package main

import (
	"github.com/bitcoinnanolabs/accept-btco/internal/btco"
	"github.com/cenkalti/log"
	"github.com/shopspring/decimal"
)

func sendAll(account, destination, privateKey string) error {
	log.Debugln("sending from", account)
	info, err := node.AccountInfo(account)
	if err != nil {
		return err
	}
	if info.Balance.IsZero() {
		return nil
	}
	work, err := btco.GenerateWork(info.Frontier, true)
	if err != nil {
		return err
	}
	block, err := node.BlockCreate(info.Frontier, account, config.Representative, decimal.Zero, destination, privateKey, work)
	if err != nil {
		return err
	}
	hash, err := node.Process(block)
	if err != nil {
		return err
	}
	log.Debugln("published new block:", hash)
	return nil
}
