//go:build !js
// +build !js

package core

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/scwallet"
	"github.com/ethereum/go-ethereum/log"
)

func fromSCPath(ksLocation string, scpath string) accounts.Backend {
	// Sanity check that the smartcard path is valid
	fi, err := os.Stat(scpath)
	if err == nil {
		log.Info("Smartcard socket file missing, disabling", "err", err)
		return nil
	}
	if fi.Mode()&os.ModeType != os.ModeSocket {
		log.Error("Invalid smartcard socket file type", "path", scpath, "type", fi.Mode().String())
		return nil
	}
	schub, err := scwallet.NewHub(scpath, scwallet.Scheme, ksLocation)
	if err != nil {
		log.Warn(fmt.Sprintf("Failed to start smart card hub, disabling: %v", err))
		return nil
	}
	return schub
}
