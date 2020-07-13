package local

import (
	"github.com/spf13/cobra"

	"github.com/confluentinc/cli/internal/pkg/cmd"
	"github.com/confluentinc/cli/internal/pkg/examples"
)

func NewCurrentCommand(prerunner cmd.PreRunner) *cobra.Command {
	c := NewLocalCommand(
		&cobra.Command{
			Use:   "current",
			Args:  cobra.NoArgs,
			Short: "Get the path of the current Confluent run.",
			Example: examples.BuildExampleString(
				examples.Example{
					Desc: "In Linux, running ``confluent local current`` should resemble the following:",
					Code: "/tmp/confluent.SpBP4fQi",
				},
				examples.Example{
					Desc: "In macOS, running ``confluent local current`` should resemble the following:",
					Code: "/var/folders/cs/1rndf6593qb3kb6r89h50vgr0000gp/T/confluent.000000",
				},
			),
		}, prerunner)

	c.Command.RunE = cmd.NewCLIRunE(c.runCurrentCommand)
	return c.Command
}

func (c *Command) runCurrentCommand(command *cobra.Command, _ []string) error {
	dir, err := c.cc.GetCurrentDir()
	if err != nil {
		return err
	}

	command.Println(dir)
	return nil
}
