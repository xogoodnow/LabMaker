package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

func SetupKubeAdm(version string) {
	fmt.Println("Installing CRI-O runtime...")
	installCriO()

	fmt.Printf("Installing Kubeadm, Kubelet, and Kubectl (version: %s)...\n", version)
	installKubeComponents(version)

	fmt.Println("Initializing the control plane...")
	initControlPlane()

	fmt.Println("Setting up Calico network plugin...")
	installCalico()

	fmt.Println("Kubernetes setup complete!")
}

func installCriO() {
	cmds := []string{
		"sudo apt-get update -y",
		"sudo apt-get install -y software-properties-common gpg curl apt-transport-https ca-certificates",
		"curl -fsSL https://pkgs.k8s.io/addons:/cri-o:/prerelease:/main/deb/Release.key | gpg --dearmor -o /etc/apt/keyrings/cri-o-apt-keyring.gpg",
		`echo "deb [signed-by=/etc/apt/keyrings/cri-o-apt-keyring.gpg] https://pkgs.k8s.io/addons:/cri-o:/prerelease:/main/deb/ /" | tee /etc/apt/sources.list.d/cri-o.list`,
		"sudo apt-get update -y",
		"sudo apt-get install -y cri-o",
		"sudo systemctl daemon-reload",
		"sudo systemctl enable crio --now",
		"sudo systemctl start crio.service",
	}
	runCommands(cmds)
}

func installKubeComponents(version string) {
	cmds := []string{
		"sudo mkdir -p /etc/apt/keyrings",
		fmt.Sprintf("curl -fsSL https://pkgs.k8s.io/core:/stable:/v%s/deb/Release.key | sudo gpg --dearmor -o /etc/apt/keyrings/kubernetes-apt-keyring.gpg", version),
		fmt.Sprintf(`echo "deb [signed-by=/etc/apt/keyrings/kubernetes-apt-keyring.gpg] https://pkgs.k8s.io/core:/stable:/v%s/deb/ /" | sudo tee /etc/apt/sources.list.d/kubernetes.list`, version),
		"sudo apt-get update -y",
		fmt.Sprintf("sudo apt-get install -y kubelet=%s-1.1 kubectl=%s-1.1 kubeadm=%s-1.1", version, version, version),
		"sudo apt-mark hold kubelet kubeadm kubectl",
	}
	runCommands(cmds)
}

func initControlPlane() {
	cmd := `sudo kubeadm init --pod-network-cidr=192.168.0.0/16 --node-name master-1`
	runCommand(cmd)

	kubeConfigCmds := []string{
		"mkdir -p $HOME/.kube",
		"sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config",
		"sudo chown $(id -u):$(id -g) $HOME/.kube/config",
	}
	runCommands(kubeConfigCmds)

	runCommand("kubectl get po -n kube-system")
}

func installCalico() {
	cmd := `kubectl apply -f https://docs.projectcalico.org/manifests/calico.yaml`
	runCommand(cmd)
}

func runCommand(command string) {
	cmd := exec.Command(ShellToUse, "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Failed to run command: %v\nOutput: %s", err, output)
	}
	fmt.Printf("Command Output: %s\n", output)
}

func runCommands(commands []string) {
	for _, cmd := range commands {
		runCommand(cmd)
	}
}

func SetupMiniKube() {
	cmd := []string{
		"sudo snap install microk8s --classic",
		"microk8s enable dns",
		"microk8s enable dashboard",
		"microk8s kubectl get all --all-namespaces",
		"microk8s status --wait-ready",
	}
	runCommands(cmd)
}

var kubernetesCmd = &cobra.Command{
	Use:   "kubernetes [version]",
	Short: "Set up Kubernetes",
	Long:  `Installs and configures Kubernetes in your environment.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		version := args[0]
		fmt.Printf("Setting up Kubernetes version %s...\n", version)
		SetupKubeAdm(version)
	},
}

var MiniKubeCmd = &cobra.Command{
	Use:   "minik8s",
	Short: "Set up mini k8s",
	Long:  `Installs minik8s in your environment.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Setting up Mini Kube\n")
		SetupMiniKube()
	},
}

func init() {
	rootCmd.AddCommand(kubernetesCmd)
	rootCmd.AddCommand(MiniKubeCmd)

}
