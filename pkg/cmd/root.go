package cmd

import (
	"github.com/Etwodev/goAni/pkg/cmd/add"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	root := &cobra.Command{
		Use:   "go-ani",
		Short: "A tool that trains and detects anime characters",
	}

	root.AddCommand(
		add.New(),
	)

	return root
}
