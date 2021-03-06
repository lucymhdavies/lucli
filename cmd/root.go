package cmd

import (
	"github.com/skybet/cali"
)

var (
	// Define this here, then all other files in cmd can add subcommands to it
	cli = cali.NewCli("lucli")
)

func init() {
	cli.SetShort("Lu(cy) CLI")
	cli.SetLong("Named after myself, because I'm vain.\nPronounced 'loosely', because while it's cool, its utility is only loosely based in reality.")

	// TODO: global flag: creds

	cli.Flags().StringP("aws-profile", "p", "lmhd_root", "AWS Profile to use (if applicable)")
	cli.BindFlags()
}

func Execute() {
	cli.Start()
}
