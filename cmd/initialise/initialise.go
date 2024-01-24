package initialise

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	api "github.com/vision-cli/api/v1"
	"github.com/vision-cli/vision-plugin-plugin-v1/cmd/model"
)

var InitCmd = &cobra.Command{
	Use:   "init [plugin module name in the form github.com/vision-cli/plugin...]",
	Short: "initialise this plugin for use with this project",
	Long:  "initialise this project's vision.json file with this plugin's configuration values",
	Args: func(cmd *cobra.Command, args []string) error {

		if len(args) != 1 || !strings.HasPrefix(args[0], "github.com/vision-cli/") {
			return fmt.Errorf("please provide one plugin module name in the form github.com/vision-cli/plugin...")
		}

		return nil
	},
	RunE: runCommand,
}

func runCommand(cmd *cobra.Command, args []string) error {
	jEnc := json.NewEncoder(model.Out)

	pd := model.PluginConfig{
		Name:    model.PluginName,
		Module:  args[0],
		Command: model.PluginCommand,
	}

	err := jEnc.Encode(api.Init{
		Config:  pd,
		Success: true,
	})

	if err != nil {
		return fmt.Errorf("encoding json: %w", err)
	}
	return nil
}
