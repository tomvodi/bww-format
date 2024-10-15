package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

func NewRootCmd(opts *Options) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "bww-format [path]",
		Short: "A command line tool for formatting bagpipe tunes in .bww and .bmw format",
		RunE:  newRootRunFunc(opts),
	}

	addVerbose(rootCmd, opts)

	return rootCmd
}

func Execute() {
	opts := &Options{}

	rootCmd := NewRootCmd(opts)

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func newRootRunFunc(opts *Options) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			err := cmd.Help()
			if err != nil {
				return err
			}
			os.Exit(0)
		}

		return nil
	}
}
