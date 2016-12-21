package cmd

import (
        "github.com/spf13/cobra"
)

func init() {

}

var RootCmd = &cobra.Command{
        Use: "rebootclient --help",
        Run: func(cmd *cobra.Command, args []string) {
        },
}
