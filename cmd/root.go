package cmd

import (
	"github.com/spf13/cobra"

	"github.com/vision-cli/vision-plugin-plugin-v1/cmd/generate"
	"github.com/vision-cli/vision-plugin-plugin-v1/cmd/info"
	"github.com/vision-cli/vision-plugin-plugin-v1/cmd/initialize"
	"github.com/vision-cli/vision-plugin-plugin-v1/cmd/model"
	"github.com/vision-cli/vision-plugin-plugin-v1/cmd/version"
)

func init() {
	rootCmd.AddCommand(initialize.InitCmd)
	rootCmd.AddCommand(info.InfoCmd)
	rootCmd.AddCommand(version.VersionCmd)
	rootCmd.AddCommand(generate.GenerateCmd)
}

var rootCmd = &cobra.Command{
	Use:                model.PluginCommand,
	FParseErrWhitelist: cobra.FParseErrWhitelist{UnknownFlags: true},
}

func Execute() {
	rootCmd.Execute()
}
