//go:build js
// +build js

package core

import (
	"github.com/ethereum/go-ethereum/accounts"
)

// scwallet is not supported in browser yet.
// Someome brave could implement it
func fromSCPath(ksLocation string, scpath string) accounts.Backend {
	return nil
}
