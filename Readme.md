# LabMaker

This repository contains `LabMaker`, a Golang cli/Python script that automates the setup of lab environments. The script installs Go, sets up a standard Kubernetes cluster, and configures a MicroK8s cluster on a single node.

## Good for

- Network/Disk/Resource utilities.
- Sets up a standard Kubernetes cluster.
- Configures a MicroK8s cluster.

## How to Use

1. Clone this repository:

    ```bash
    git clone https://github.com/xogoodnow/LabMaker.git
    cd LabMaker
    ```

2. Run the script:

    ```bash
    go build .
    ./Labmaker ....
    ```

    OR
    ```bash
    python labmaker.py
    ```

## Requirements

- Golang 1.23
- Python 3.x
- Administrator privileges
- Internet connection
