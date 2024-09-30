package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const ShellToUse = "bash"

var rootCmd = &cobra.Command{
	Use:   "lab-cli", // The name of the CLI tool
	Short: "Lab CLI tool",
	Long:  `A CLI tool to set up lab environments with Kubernetes, Docker, and other tools.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
