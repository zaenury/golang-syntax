package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "",
		Short: "",
		Long:  ``,
	}
)

func init() {
	// rootCmd.AddCommand(scheduler.Commands())
	rootCmd.AddCommand(Commands())
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
