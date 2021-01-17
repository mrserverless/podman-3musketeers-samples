# Podman 3 Musketeers (beta)

!(podman logo)[images/podman.png]

An alternative implementation to 3 Musketeers using:

1. Container Runtime: Podman instead of Docker. It's a drop in replace for docker following OCI standards. Has a nice [self-contained daemon-less achitecture](https://ti8m.com/blog/Why-Podman-is-worth-a-look-.htm) that is perfect for local SDLC and supports [generation of kube YAML](https://developers.redhat.com/blog/2019/01/29/podman-kubernetes-yaml/)  
2. Container Orchestration: Kubernetes Pods instead of Docker Compose. Kubernetes typically would already have a Pod YAML definition, it makes less sense to slap on another docker-compose YAML that duplicates the same definitions. 
3. Command Line: Makefile. Nothing changes here.

## Justifications

As Kubernetes becomse a standard in terms of container orchestration and deployment, it make sense for 3 Musketeers to consider a more Kube-friendly approach using the tools that are already installed on a developer's machine. 

