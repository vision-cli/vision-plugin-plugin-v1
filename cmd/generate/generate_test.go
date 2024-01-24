package generate_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vision-cli/vision-plugin-plugin-v1/cmd/model"
)

func TestOpenVisionJsonReturnsConvertedConfig(t *testing.T) {
	b := `{
		"plugin": {
			"command": "plugin",
			"module": "github.com/vision-cli/test",
			"name": "vision-plugin-plugin-v1"
		},
	  "project_name": "test"
	}`
	var jsonData model.PluginData
	err := json.Unmarshal([]byte(b), &jsonData)
	assert.NoError(t, err)
	fmt.Printf("%+v", jsonData)
	assert.Equal(t, "plugin", jsonData.PluginConfig.Command)
	assert.Equal(t, "github.com/vision-cli/test", jsonData.PluginConfig.Module)
	assert.Equal(t, "vision-plugin-plugin-v1", jsonData.PluginConfig.Name)
}
