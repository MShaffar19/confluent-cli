package auditlog

import (
	"net/http"

	"github.com/spf13/cobra"

	pcmd "github.com/confluentinc/cli/internal/pkg/cmd"
	"github.com/confluentinc/cli/internal/pkg/errors"
)

type command struct {
	*pcmd.AuthenticatedCLICommand
	prerunner pcmd.PreRunner
}

// New returns the default command object for interacting with audit logs.
func New(prerunner pcmd.PreRunner) *cobra.Command {
	cliCmd := pcmd.NewAuthenticatedWithMDSCLICommand(
		&cobra.Command{
			Use:   "audit-log",
			Short: "Manage audit log configuration.",
			Long:  "Manage which auditable events are logged, and where the events are sent.",
		}, prerunner)
	cmd := &command{
		AuthenticatedCLICommand: cliCmd,
		prerunner:               prerunner,
	}
	cmd.init()
	return cmd.Command
}

func (c *command) init() {
	c.AddCommand(NewConfigCommand(c.prerunner))
	c.AddCommand(NewRouteCommand(c.prerunner))
}

func HandleMdsAuditLogApiError(cmd *cobra.Command, err error, response *http.Response) error {
	if response != nil && response.StatusCode == http.StatusNotFound {
		return errors.NewWrapErrorWithSuggestions(err, errors.UnableToAccessEndpointErrorMsg, errors.UnableToAccessEndpointSuggestions)
	}
	return err
}
