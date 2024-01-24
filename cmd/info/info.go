package info

import (
	_ "embed"
	"encoding/json"

	"github.com/spf13/cobra"
	api "github.com/vision-cli/api/v1"
	"github.com/vision-cli/vision-plugin-plugin-v1/cmd/model"
)

//go:embed info.txt
var infoOutput string

var InfoCmd = &cobra.Command{
	Use:   "info",
	Short: "return info about the plugin",
	Long:  "return detailed information about the plugin",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		json.NewEncoder(model.Out).Encode(api.Info{
			ShortDescription: "Create vision plugins",
			LongDescription:  infoOutput,
		})
	},
}