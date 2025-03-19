package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/saumsy/localizer/internal"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Check what version of localizer you are using",
	Long:  "Check what version of localizer you are using",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Localizer version: %s\n", color.GreenString(internal.Version))
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
