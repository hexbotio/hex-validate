package main

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/hashicorp/go-plugin"
	"github.com/hexbotio/hex-plugin"
)

type HexValidate struct {
}

func (g *HexValidate) Perform(args hexplugin.Arguments) (resp hexplugin.Response) {

	// initialize return values
	var output = ""
	var success = true

	// make sure exists
	if args.Command == "" {
		success = false
	}

	// match type
	if args.Config["type"] != "" {
		switch args.Config["type"] {
		case "string":
			// everything is a string
		case "int":
			_, err := strconv.ParseInt(args.Command, 10, 32)
			if err != nil {
				success = false
			}
		case "int64":
			_, err := strconv.ParseInt(args.Command, 10, 64)
			if err != nil {
				success = false
			}
		case "float":
			_, err := strconv.ParseFloat(args.Command, 64)
			if err != nil {
				success = false
			}
		case "bool":
			_, err := strconv.ParseBool(args.Command)
			if err != nil {
				success = false
			}
		}
	}

	// match regular expression
	if args.Config["match_re"] != "" {
		pattern := strings.Replace(args.Config["match_re"], "/", "", -1)
		regx := regexp.MustCompile(pattern)
		success = regx.MatchString(args.Command)
	}

	// match list
	if args.Config["match_list"] != "" {
		match := false
		for _, member := range strings.Split(args.Config["match_list"], ",") {
			member = strings.TrimSpace(member)
			if member == args.Command {
				match = true
			}
		}
		success = match
	}

	// evaluate results
	if success {
		output = args.Config["success"]
	} else {
		output = args.Config["failure"]
	}

	resp = hexplugin.Response{
		Output:  output,
		Success: success,
	}
	return resp
}

func main() {
	var pluginMap = map[string]plugin.Plugin{
		"action": &hexplugin.HexPlugin{Impl: &HexValidate{}},
	}
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: hexplugin.GetHandshakeConfig(),
		Plugins:         pluginMap,
	})
}
