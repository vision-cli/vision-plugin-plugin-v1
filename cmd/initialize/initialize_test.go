package initialize

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunCommandWithCorrectParameterReturnsCorrectJson(t *testing.T) {
	o := bytes.NewBufferString("")
	r := RunCommand(o)
	err := r(nil, []string{"github.com/vision-cli/vision-plugin-test-v1"})
	assert.NoError(t, err)

	expected := `{"config":{"name":"vision-plugin-plugin-v1","module":"github.com/vision-cli/vision-plugin-test-v1","command":"plugin"},"success":true}` + "\n"
	assert.Equal(t, expected, o.String())
}

func TestRunCommandWithIncorrectParametersReturnsError(t *testing.T) {
	c := InitCmd
	err := c.Args(nil, []string{})
	assert.Error(t, err)
}

func TestRunCommandWithInvalidParametersReturnsError(t *testing.T) {
	c := InitCmd
	err := c.Args(nil, []string{"vision-plugin-test-v1"})
	assert.Error(t, err)
}
