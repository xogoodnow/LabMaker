package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func InstallDocker() {
	cmds := []string{
		"sudo apt install -y apt-transport-https ca-certificates curl software-properties-common",
		"curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -",
		"sudo add-apt-repository 'deb [arch=amd64] https://download.docker.com/linux/ubuntu focal stable'",
		"apt-cache policy docker-ce",
		"sudo apt install docker-ce",
		"sudo systemctl status docker",
	}
	runCommands(cmds)

}

var DockerInstallCmd = &cobra.Command{
	Use:   "docker",
	Short: "Installing docker",
	Long:  "Installing docker as CRI on the env",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Setting up Docker")
		InstallDocker()
	},
}

func init() {
	rootCmd.AddCommand(DockerInstallCmd)
}
