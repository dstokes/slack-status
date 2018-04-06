package main

import (
	"log"
	"os"
	"os/user"
	"path"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/nlopes/slack"
	flag "github.com/ogier/pflag"
)

type workspace struct {
	Token  string
	Groups []string
}

var (
	f = flag.NewFlagSet("flags", flag.ExitOnError)

	// options
	awayFlag      = f.Bool("away", false, "away")
	groupFlag     = f.StringP("group", "g", "", "group")
	workspaceFlag = f.StringP("workspace", "w", "", "workspace")
)

func main() {
	f.Parse(os.Args[1:])
	args := f.Args()

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	configFile := path.Join(usr.HomeDir, ".slack-status")

	var cfg map[string]workspace
	_, err = toml.DecodeFile(configFile, &cfg)
	if err != nil {
		log.Fatal(err)
	}

	var icon, message string
	if len(args) > 0 {
		if strings.HasPrefix(args[0], ":") {
			icon = args[0]
			message = strings.Join(args[1:], " ")
		} else {
			message = strings.Join(args[0:], " ")
		}
	}

	presence := "auto"
	if *awayFlag {
		presence = "away"
	}

	for name, c := range cfg {
		if *workspaceFlag != "" && *workspaceFlag != name {
			continue
		}
		if *groupFlag != "" {
			found := false
			for _, g := range c.Groups {
				if g == *groupFlag {
					found = true
					break
				}
			}
			if found == false {
				continue
			}
		}

		api := slack.New(c.Token)

		if err := api.SetUserPresence(presence); err != nil {
			log.Fatal(err)
		}

		if err := api.SetUserCustomStatus(message, icon); err != nil {
			log.Fatal(err)
		}
	}
}
