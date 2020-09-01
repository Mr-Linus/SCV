## SCV
SCV is a distributed cluster GPU sniffer. 
It can cooperate with [Yoda-Scheduler](https://github.com/NJUPT-ISL/Yoda-Scheduler) to achieve 
fine-grained GPU scheduling tasks.

![Status](https://github.com/NJUPT-ISL/SCV/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/NJUPT-ISL/SCV)](https://goreportcard.com/report/github.com/NJUPT-ISL/SCV)
### GPU metrics that SCV can monitor
- Core Frequency
- Model
- Free Memory 
- Total Memory 
- Memory Frequency
- Bandwidth
- Power
- GPU Number
### Get Started
- Ensure that the nvidia container runtime and the nvidia driver are installed on each kubernetes worker node. See [nvidia-docker](https://github.com/NVIDIA/nvidia-docker#quickstart)
for more details.
    -  Ubuntu 
    
       ```shell
       # Add the package repositories
       $ distribution=$(. /etc/os-release;echo $ID$VERSION_ID)
       $ curl -s -L https://nvidia.github.io/nvidia-docker/gpgkey | sudo apt-key add -
       $ curl -s -L https://nvidia.github.io/nvidia-docker/$distribution/nvidia-docker.list | sudo tee /etc/apt/sources.list.d/nvidia-docker.list
            
       $ sudo apt-get update && sudo apt-get install -y nvidia-container-toolkit nvidia-container-runtime
       $ sudo systemctl restart docker
        ```
    - Centos
    
        ```shell
        $ distribution=$(. /etc/os-release;echo $ID$VERSION_ID)
        $ curl -s -L https://nvidia.github.io/nvidia-docker/$distribution/nvidia-docker.repo | sudo tee /etc/yum.repos.d/nvidia-docker.repo
            
        $ sudo yum install -y nvidia-container-toolkit nvidia-container-runtime
        $ sudo systemctl restart docker
        ```
- Enable the nvidia-container-runtime as docker default runtime on each kubernetes worker node.

    You need to modify `/etc/docker/daemon.json` to the following content on each worker node：
    ```json
        {
            "default-runtime": "nvidia",
            "runtimes": {
                "nvidia": {
                    "path": "/usr/bin/nvidia-container-runtime",
                    "runtimeArgs": []
                }
            },
            "exec-opts": ["native.cgroupdriver=systemd"],
            "log-driver": "json-file",
            "log-opts": {
              "max-size": "100m"
            },
            "storage-driver": "overlay2",
            "registry-mirrors": ["https://registry.docker-cn.com"]
        }
    ```
- Deploy the SCV into your kubernetes cluster:
    ```shell
    kubectl apply -f  https://raw.githubusercontent.com/NJUPT-ISL/SCV/master/deploy/deploy.yaml
    ```

