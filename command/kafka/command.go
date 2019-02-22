package kafka

import (
	"github.com/spf13/cobra"

	"github.com/confluentinc/cli/command/common"
	"github.com/confluentinc/cli/shared"
	"github.com/confluentinc/cli/shared/kafka"
)

type command struct {
	*cobra.Command
	config *shared.Config
}

// New returns the default command object for interacting with Kafka.
func New(config *shared.Config) (*cobra.Command, error) {
	return newCMD(config, common.GRPCLoader(kafka.Name))
}

// NewKafkaCommand returns a command object using a custom Kafka provider.
func NewKafkaCommand(config *shared.Config, provider common.Provider) (*cobra.Command, error) {
	return newCMD(config, provider)
}

// newCMD returns a command for interacting with Kafka.
func newCMD(config *shared.Config, plugin common.Provider) (*cobra.Command, error) {
	cmd := &command{
		Command: &cobra.Command{
			Use:   "kafka",
			Short: "Manage Kafka",
		},
		config: config,
	}
	err := cmd.init(plugin)
	return cmd.Command, err
}

func (c *command) init(plugin common.Provider) error {
	c.AddCommand(NewClusterCommand(c.config, plugin))
	c.AddCommand(NewTopicCommand(c.config, plugin))
	c.AddCommand(NewACLCommand(c.config, plugin))

	return nil
}
