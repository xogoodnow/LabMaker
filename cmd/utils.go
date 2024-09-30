package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

func InstallDiskUtils(utilName string) {
	utilList := []string{"iostat", "dstat", "atop", "ioping"}

	if utilName == "list" {
		fmt.Println("Available disk utilities:", utilList)
		return
	}

	if utilName == "all" {
		for _, util := range utilList {
			cmd := exec.Command(ShellToUse, "-c", fmt.Sprintf("sudo apt install -y %s", util))
			err := cmd.Run()
			if err != nil {
				log.Fatalf("Failed to install %s: %v", util, err)
			}
			fmt.Printf("%s installed successfully\n", util)
		}
	} else {
		found := false
		for _, util := range utilList {
			if util == utilName {
				cmd := exec.Command(ShellToUse, "-c", fmt.Sprintf("sudo apt install -y %s", util))
				err := cmd.Run()
				if err != nil {
					log.Fatalf("Failed to install %s: %v", util, err)
				}
				fmt.Printf("%s installed successfully\n", util)
				found = true
				break
			}
		}
		if !found {
			fmt.Printf("Utility %s is not recognized. Available utilities: %v\n", utilName, utilList)
		}
	}
}

func InstallNetworkUtils(utilName string) {
	utilList := []string{"telnet", "netstat", "bmon", "ss"}

	if utilName == "list" {
		fmt.Println("Available network utilities:", utilList)
		return
	}

	if utilName == "all" {
		for _, util := range utilList {
			cmd := exec.Command(ShellToUse, "-c", fmt.Sprintf("sudo apt install -y %s", util))
			err := cmd.Run()
			if err != nil {
				log.Fatalf("Failed to install %s: %v", util, err)
			}
			fmt.Printf("%s installed successfully\n", util)
		}
	} else {
		found := false
		for _, util := range utilList {
			if util == utilName {
				cmd := exec.Command(ShellToUse, "-c", fmt.Sprintf("sudo apt install -y %s", util))
				err := cmd.Run()
				if err != nil {
					log.Fatalf("Failed to install %s: %v", util, err)
				}
				fmt.Printf("%s installed successfully\n", util)
				found = true
				break
			}
		}
		if !found {
			fmt.Printf("Utility %s is not recognized. Available utilities: %v\n", utilName, utilList)
		}
	}
}

func InstallResourceUtils(utilName string) {
	utilList := []string{"htop", "top"}

	if utilName == "list" {
		fmt.Println("Available resource utilities:", utilList)
		return
	}

	if utilName == "all" {
		for _, util := range utilList {
			cmd := exec.Command(ShellToUse, "-c", fmt.Sprintf("sudo apt install -y %s", util))
			err := cmd.Run()
			if err != nil {
				log.Fatalf("Failed to install %s: %v", util, err)
			}
			fmt.Printf("%s installed successfully\n", util)
		}
	} else {
		found := false
		for _, util := range utilList {
			if util == utilName {
				cmd := exec.Command(ShellToUse, "-c", fmt.Sprintf("sudo apt install -y %s", util))
				err := cmd.Run()
				if err != nil {
					log.Fatalf("Failed to install %s: %v", util, err)
				}
				fmt.Printf("%s installed successfully\n", util)
				found = true
				break
			}
		}
		if !found {
			fmt.Printf("Utility %s is not recognized. Available utilities: %v\n", utilName, utilList)
		}
	}
}

var diskToolsCmd = &cobra.Command{
	Use:   "disk [utility]",
	Short: "Disk analysis utilities",
	Long:  "Install or list disk analysis utilities",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		InstallDiskUtils(args[0])
	},
}

var networkToolsCmd = &cobra.Command{
	Use:   "network [utility]",
	Short: "Network analysis utilities",
	Long:  "Install or list network analysis utilities",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		InstallNetworkUtils(args[0])
	},
}

var resourceToolsCmd = &cobra.Command{
	Use:   "resources [utility]",
	Short: "Resource monitoring utilities",
	Long:  "Install or list resource monitoring utilities",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		InstallResourceUtils(args[0])
	},
}

func init() {
	rootCmd.AddCommand(diskToolsCmd)
	rootCmd.AddCommand(networkToolsCmd)
	rootCmd.AddCommand(resourceToolsCmd)
}
