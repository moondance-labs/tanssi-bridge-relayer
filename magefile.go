//go:build mage
// +build mage

package main

import (
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

func Build() {
	mg.Deps(BuildMain)
}

func BuildMain() error {
	err := sh.Run("./update_contract_interface.sh")
	if err != nil {
		return err
	}

	return sh.Run("go", "build", "-o", "build/tanssi-bridge-relayer", "main.go")
}

func Test() error {
	return sh.RunV("go", "test", "./...")
}

func Lint() error {
	return sh.Run("revive", "-config", "revive.toml", "./...")
}
