//go:build !js
// +build !js

package core

// scwallet is not supported in browser yet.
// Someome brave could implement it
func fromSCPath(scpath string) accounts.Backend {
	return nil
}
