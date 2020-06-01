package feedback

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"

	"github.com/confluentinc/cli/internal/pkg/cmd"
	"github.com/confluentinc/cli/internal/pkg/config"
	"github.com/confluentinc/cli/internal/pkg/config/v3"
	"github.com/confluentinc/cli/mock"
)

func TestFeedback(t *testing.T) {
	command := cmd.BuildRootCommand()
	command.AddCommand(mockFeedbackCommand("This feedback tool is great!"))

	req := require.New(t)
	out, err := cmd.ExecuteCommand(command, "feedback")
	req.NoError(err)
	req.Contains(out, "Enter feedback: ")
	req.Contains(out, "Thanks for your feedback.")
}

func TestFeedbackEmptyMessage(t *testing.T) {
	command := cmd.BuildRootCommand()
	command.AddCommand(mockFeedbackCommand(""))

	req := require.New(t)
	out, err := cmd.ExecuteCommand(command, "feedback")
	req.NoError(err)
	req.Contains(out, "Enter feedback: ")
}

func mockFeedbackCommand(msg string) *cobra.Command {
	mockPreRunner := mock.NewPreRunnerMock(nil, nil)
	mockConfig := v3.New(&config.Params{CLIName: "ccloud"})
	mockAnalytics := mock.NewDummyAnalyticsMock()
	mockPrompt := mock.NewPromptMock(msg)
	return NewFeedbackCmdWithPrompt(mockPreRunner, mockConfig, mockAnalytics, mockPrompt)
}
