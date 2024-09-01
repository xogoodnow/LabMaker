import os
import subprocess

def execute_command(command):
    result = subprocess.run(command, shell=True, stdout=subprocess.PIPE, stderr=subprocess.PIPE)
    if result.returncode != 0:
        print(f"Error: {result.stderr.decode('utf-8')}")
    return result.stdout.decode('utf-8')

def install_go():
    go_url = "https://go.dev/dl/go1.23.0.linux-amd64.tar.gz"
    commands = [
        f"wget {go_url}",
        "sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.23.0.linux-amd64.tar.gz",
        "echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc",
        "bash -c 'source ~/.bashrc; go version'"
    ]
    for cmd in commands:
        output = execute_command(cmd)
        if 'go version' in cmd:
            print(output)

def deploy_standard_k8s_cluster():
    commands = [
        # Step 1: Enable iptables Bridged Traffic on all Nodes
        "sudo modprobe overlay",
        "sudo modprobe br_netfilter",
        "sudo sysctl --system",
        
        # Step 2: Disable swap
        "sudo swapoff -a",
        "(crontab -l 2>/dev/null; echo '@reboot /sbin/swapoff -a') | crontab - || true",
        
        # Step 3: Install CRI-O Runtime On All Nodes
        "sudo apt-get update -y",
        "sudo apt-get install -y software-properties-common curl apt-transport-https ca-certificates",
        "curl -fsSL https://pkgs.k8s.io/addons:/cri-o:/prerelease:/main/deb/Release.key | gpg --dearmor -o /etc/apt/keyrings/cri-o-apt-keyring.gpg",
        "echo 'deb [signed-by=/etc/apt/keyrings/cri-o-apt-keyring.gpg] https://pkgs.k8s.io/addons:/cri-o:/prerelease:/main/deb/ /' | sudo tee /etc/apt/sources.list.d/cri-o.list",
        "sudo apt-get update -y",
        "sudo apt-get install -y cri-o",
        "sudo systemctl daemon-reload",
        "sudo systemctl enable crio --now",
        "sudo systemctl start crio.service",
        
        # Step 4: Install Kubeadm & Kubelet & Kubectl on all Nodes
        "KUBERNETES_VERSION=1.29",
        "sudo mkdir -p /etc/apt/keyrings",
        "curl -fsSL https://pkgs.k8s.io/core:/stable:/v$KUBERNETES_VERSION/deb/Release.key | sudo gpg --dearmor -o /etc/apt/keyrings/kubernetes-apt-keyring.gpg",
        "echo 'deb [signed-by=/etc/apt/keyrings/kubernetes-apt-keyring.gpg] https://pkgs.k8s.io/core:/stable:/v$KUBERNETES_VERSION/deb/ /' | sudo tee /etc/apt/sources.list.d/kubernetes.list",
        "sudo apt-get update -y",
        "sudo apt-get install -y kubelet kubeadm kubectl",
        "sudo apt-mark hold kubelet kubeadm kubectl",
        
        # Step 5: Initialize Kubeadm On Master Node
        "IPADDR=$(curl ifconfig.me && echo '')",
        "NODENAME=$(hostname -s)",
        "POD_CIDR='192.168.0.0/16'",
        "sudo kubeadm init --control-plane-endpoint=$IPADDR  --apiserver-cert-extra-sans=$IPADDR  --pod-network-cidr=$POD_CIDR --node-name $NODENAME --ignore-preflight-errors Swap",
        "mkdir -p $HOME/.kube",
        "sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config",
        "sudo chown $(id -u):$(id -g) $HOME/.kube/config",
        "kubectl get po -n kube-system",
        
        # Step 6: Join Worker Nodes
        "kubeadm token create --print-join-command | bash",
        
        # Step 7: Install Calico Network Plugin for Pod Networking
        "kubectl apply -f https://docs.projectcalico.org/manifests/calico.yaml",
        
        # Step 8: Setup Kubernetes Metrics Server
        "kubectl apply -f https://raw.githubusercontent.com/techiescamp/kubeadm-scripts/main/manifests/metrics-server.yaml",
        
        # Step 9: Deploy A Sample Nginx Application
        "kubectl apply -f - <<EOF\napiVersion: apps/v1\nkind: Deployment\nmetadata:\n  name: nginx-deployment\nspec:\n  selector:\n    matchLabels:\n      app: nginx\n  replicas: 2\n  template:\n    metadata:\n      labels:\n        app: nginx\n    spec:\n      containers:\n      - name: nginx\n        image: nginx:latest\n        ports:\n        - containerPort: 80\nEOF",
        "kubectl apply -f - <<EOF\napiVersion: v1\nkind: Service\nmetadata:\n  name: nginx-service\nspec:\n  selector:\n    app: nginx\n  type: NodePort\n  ports:\n    - port: 80\n      targetPort: 80\n      nodePort: 32000\nEOF"
    ]
    for cmd in commands:
        execute_command(cmd)

def deploy_microk8s_cluster():
    commands = [
        "sudo snap install microk8s --classic",
        "microk8s enable dns",
        "microk8s enable dashboard",
        "microk8s kubectl get all --all-namespaces",
        "microk8s status --wait-ready"
    ]
    for cmd in commands:
        execute_command(cmd)

def main_menu():
    while True:
        print("Choose an option:")
        print("1. Install Go")
        print("2. Deploy Standard K8s Cluster using Kubeadm")
        print("3. Deploy MicroK8s Cluster")
        print("4. Quit")

        choice = input("Enter your choice (1, 2, 3, or 4): ")

        if choice == '1':
            install_go()
            satisfaction = input("Was the Go installation satisfactory? (yes or no): ").lower()
            if satisfaction != 'yes':
                continue
        elif choice == '2':
            deploy_standard_k8s_cluster()
            satisfaction = input("Was the standard K8s deployment satisfactory? (yes or no): ").lower()
            if satisfaction != 'yes':
                continue
        elif choice == '3':
            deploy_microk8s_cluster()
            satisfaction = input("Was the MicroK8s deployment satisfactory? (yes or no): ").lower()
            if satisfaction != 'yes':
                continue
        elif choice == '4':
            print("Exiting.")
            break
        else:
            print("Invalid choice. Please select a valid option.")
            continue

if __name__ == "__main__":
    main_menu()
