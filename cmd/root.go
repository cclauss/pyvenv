package cmd

import (
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var noColor bool

var rootCmd = &cobra.Command{
	Use:     "pyvenv",
	Short:   "Personal tool to manage Python virtual environments.",
	Version: Version,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(setColorOutput)
	rootCmd.PersistentFlags().BoolVar(&noColor, "no-color", false, "disable color output")
}

func setColorOutput() {
	if noColor {
		color.NoColor = true
	}
}
