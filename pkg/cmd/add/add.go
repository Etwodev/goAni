package add

import (
	"fmt"

	"github.com/Etwodev/goAni/pkg/gelbooru"
	"github.com/spf13/cobra"
)

type AddCommand struct {
	Command *cobra.Command
	group   string
	tags    string
	pages   int
}

// Create a new instance of the add command
func New() *cobra.Command {
	addCmd := &AddCommand{
		Command: &cobra.Command{
			Use:   "add",
			Short: "adds new data to the training pool",
			Long:  "Placeholder here",
		},
	}

	addCmd.init()

	return addCmd.Command
}

func (c *AddCommand) init() {
	c.Command.PersistentFlags().StringVar(&c.group, "group", "", "the name to associate the data with")
	c.Command.PersistentFlags().StringVar(&c.tags, "tags", "", "the gelbooru tags to use")
	c.Command.PersistentFlags().IntVar(&c.pages, "pages", 1, "the number of pages to iterate over")
	c.Command.RunE = c.run
}

func (c *AddCommand) check() error {
	if c.tags == "" {
		return fmt.Errorf("check: no tags provided")
	}
	if c.group == "" {
		return fmt.Errorf("check: no group name provided")
	}
	return nil
}

func (c *AddCommand) run(cmd *cobra.Command, args []string) error {
	if err := c.check(); err != nil {
		return err
	}

	err := gelbooru.Get(c.group, c.tags, c.pages)
	if err != nil {
		return err
	}

	return nil
}
