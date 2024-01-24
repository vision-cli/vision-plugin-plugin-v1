package initialise

import (
	"github.com/spf13/cobra"
)

var InitCmd = &cobra.Command{
	Use:   "", // TODO: update usage
	Short: "", // TODO: update usage
	Long:  "", // TODO: update usage
	// TODO: add Args validation (see cobra documentation)
	RunE: runCommand,
}

func runCommand(cmd *cobra.Command, args []string) error {
	// TODO: fill in your init logic
	return nil
}
