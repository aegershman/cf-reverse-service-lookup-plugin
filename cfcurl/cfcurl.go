package cfcurl

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/cloudfoundry/cli/plugin"
)

// Curl calls cf-curl and returns the resulting json
func Curl(cli plugin.CliConnection, path string) (map[string]interface{}, error) {
	output, err := cli.CliCommandWithoutTerminalOutput("curl", path)
	if nil != err {
		return nil, err
	}
	return parseOutput(output)
}

func parseOutput(output []string) (map[string]interface{}, error) {
	if nil == output || 0 == len(output) {
		return nil, errors.New("CF API returned no output")
	}

	data := strings.Join(output, "\n")

	if 0 == len(data) || "" == data {
		return nil, errors.New("Failed to join output")
	}

	var f interface{}
	err := json.Unmarshal([]byte(data), &f)
	return f.(map[string]interface{}), err
}
