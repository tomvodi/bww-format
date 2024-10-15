package cmd

import "github.com/spf13/cobra"

type Options struct {
	Verbose bool
}

func addVerbose(cmd *cobra.Command, opts *Options) {
	cmd.PersistentFlags().BoolVarP(&opts.Verbose, "verbose", "v", false,
		"verbose output",
	)
}
