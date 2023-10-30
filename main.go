package main

import (
	"os"

	"github.com/Etwodev/goAni/pkg/cmd"
)

func main() {
	root := cmd.New()
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
