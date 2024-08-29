package main

import (
	"github.com/alecthomas/kong"
	"github.com/julianstephens/localbrew/pkg/commands"
	"github.com/julianstephens/localbrew/pkg/version"
)

func main() {
	cli := commands.CLI{
		Globals: commands.Globals{
			Version: commands.VersionFlag(version.VERSION),
		},
	}

	ctx := kong.Parse(&cli,
		kong.Name("localbrew"),
		kong.Description(""),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
		}),
		kong.Vars{
			"version": version.VERSION,
		})
	err := ctx.Run(&cli.Globals)
	ctx.FatalIfErrorf(err)
}
