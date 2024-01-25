package model

const PluginSemVer = "v0.0.1"
const PluginName = "vision-plugin-plugin-v1"
const PluginCommand = "plugin"
const PluginOutputDir = "."

type PluginConfig struct {
	Name    string `json:"name"`
	Module  string `json:"module"`
	Command string `json:"command"`
}

type PluginData struct {
	PluginConfig PluginConfig `json:"plugin"`
}
