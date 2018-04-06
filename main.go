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
	Token string
}

var (
	f = flag.NewFlagSet("flags", flag.ExitOnError)

	// options
	awayFlag = f.Bool("away", false, "away")
)

func main() {
	f.Parse(os.Args[1:])
	args := f.Args()

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	configFile := path.Join(usr.HomeDir, ".slack-status")

	var spaces map[string]workspace
	_, err = toml.DecodeFile(configFile, &spaces)
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

	for _, cfg := range spaces {
		api := slack.New(cfg.Token)

		if err := api.SetUserPresence(presence); err != nil {
			log.Fatal(err)
		}

		if err := api.SetUserCustomStatus(message, icon); err != nil {
			log.Fatal(err)
		}
	}
}
