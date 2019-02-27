package main

import (
	"os"

	"github.com/confluentinc/cli/command"
	"github.com/confluentinc/cli/command/apikey"
	"github.com/confluentinc/cli/command/auth"
	"github.com/confluentinc/cli/command/common"
	"github.com/confluentinc/cli/command/config"
	"github.com/confluentinc/cli/command/connect"
	"github.com/confluentinc/cli/command/kafka"
	"github.com/confluentinc/cli/command/ksql"
	"github.com/confluentinc/cli/command/service-account"
	"github.com/hashicorp/go-plugin"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/confluentinc/cli/log"
	"github.com/confluentinc/cli/metric"
	"github.com/confluentinc/cli/shared"
	cliVersion "github.com/confluentinc/cli/version"
)

const cliName = "ccloud"

var (
	// Injected from linker flags like `go build -ldflags "-X main.version=$VERSION" -X ...`
	version = "v0.0.0"
	commit  = ""
	date    = ""
	host    = ""
)

func main() {
	viper.AutomaticEnv()

	logger := log.New()

	metricSink := metric.NewSink()

	var cfg *shared.Config
	{
		cfg = shared.NewConfig(&shared.Config{
			MetricSink: metricSink,
			Logger:     logger,
		})
		err := cfg.Load()
		if err != nil && err != shared.ErrNoConfig {
			logger.Errorf("unable to load config: %v", err)
		}
	}

	version := cliVersion.NewVersion(version, commit, date, host)
	factory := &common.GRPCPluginFactoryImpl{}

	defer plugin.CleanupClients()

	cli := BuildCommand(cfg, version, factory, logger)
	err := cli.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func BuildCommand(cfg *shared.Config, version *cliVersion.Version, factory common.GRPCPluginFactory, logger *log.Logger) *cobra.Command {
	cli := &cobra.Command{
		Use:   cliName,
		Short: "Welcome to the Confluent Cloud CLI",
	}
	cli.PersistentFlags().CountP("verbose", "v", "increase output verbosity")
	cli.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		if err := common.SetLoggingVerbosity(cmd, logger); err != nil {
			return common.HandleError(err, cmd)
		}
		return nil
	}

	prompt := command.NewTerminalPrompt(os.Stdin)

	cli.Version = version.Version
	cli.AddCommand(common.NewVersionCmd(version, prompt))

	conn := config.New(cfg)
	conn.Hidden = true // The config/context feature isn't finished yet, so let's hide it
	cli.AddCommand(conn)

	conn, err := common.NewCompletionCmd(cli, prompt, cliName)
	if err != nil {
		logger.Log("msg", err)
	} else {
		cli.AddCommand(conn)
	}

	cli.AddCommand(auth.New(cfg)...)

	conn, err = kafka.New(cfg, factory)
	if err != nil {
		logger.Log("msg", err)
	} else {
		cli.AddCommand(conn)
	}

	conn, err = connect.New(cfg, factory)
	if err != nil {
		logger.Log("msg", err)
	} else {
		cli.AddCommand(conn)
	}

	conn, err = ksql.New(cfg, factory)
	if err != nil {
		logger.Log("msg", err)
	} else {
		cli.AddCommand(conn)
	}

	conn, err = service_account.New(cfg, factory)
	if err != nil {
		logger.Log("msg", err)
	} else {
		cli.AddCommand(conn)
	}

	conn, err = apikey.New(cfg, factory)
	if err != nil {
		logger.Log("msg", err)
	} else {
		cli.AddCommand(conn)
	}

	return cli
}
