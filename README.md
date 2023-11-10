# replica-operator
k8s operator for controlling replica count during time

## Description
The operator watches for custom resources of a specific kind (e.g., `ReplicaControl`) in a designated namespace. These custom resources define the desired replica count at different points in time.

## Features

- **Dynamic Scaling**: Automatically adjusts the replica count of specified deployments.
- **Time-Based Control**: Define time intervals for scaling operations.
- **Configurability**: Easily configure the target deployments and scaling parameters.

## Prerequisites

- [Kubernetes](https://kubernetes.io/) cluster with version 1.19 or later.
- [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/) installed.
- [Operator SDK](https://sdk.operatorframework.io/docs/installation/install-operator-sdk/) installed.


## Getting Started
You’ll need a Kubernetes cluster to run against. You can use [KIND](https://sigs.k8s.io/kind) to get a local cluster for testing, or run against a remote cluster.
**Note:** Your controller will automatically use the current context in your kubeconfig file (i.e. whatever cluster `kubectl cluster-info` shows).

### Running on the cluster
1. Clone the repository:

```bash
git clone https://github.com/yourusername/k8s-replica-scaler-operator.git
cd k8s-replica-scaler-operator
```

2. Install Instances of Custom Resources:

```sh
kubectl apply -f config/samples/
```
3. Build and push your image to the location specified by `IMG`:

```sh
make docker-build docker-push IMG=<some-registry>/operator:tag
```

4. Deploy the controller to the cluster with the image specified by `IMG`:

```sh
make deploy IMG=<some-registry>/operator:tag
```

### Uninstall CRDs
To delete the CRDs from the cluster:

```sh
make uninstall
```

### Undeploy controller
UnDeploy the controller from the cluster:

```sh
make undeploy
```

## Contributing
We welcome contributions! Please fork the repository, create a branch, and submit a pull request.

### How it works
This project aims to follow the Kubernetes [Operator pattern](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/).

It uses [Controllers](https://kubernetes.io/docs/concepts/architecture/controller/),
which provide a reconcile function responsible for synchronizing resources until the desired state is reached on the cluster.

### Test It Out
1. Install the CRDs into the cluster:

```sh
make install
```

2. Run your controller (this will run in the foreground, so switch to a new terminal if you want to leave it running):

```sh
make run
```

**NOTE:** You can also run this in one step by running: `make install run`

### Modifying the API definitions
If you are editing the API definitions, generate the manifests such as CRs or CRDs using:

```sh
make manifests
```

**NOTE:** Run `make --help` for more information on all potential `make` targets

More information can be found via the [Kubebuilder Documentation](https://book.kubebuilder.io/introduction.html)

## License

Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.


Creating a Kubernetes Operator in Go that controls replica count over time involves several steps. Below is a basic README template to get you started. Make sure to replace placeholders with your actual project details.

---

# Kubernetes Operator for Dynamic Replica Scaling

## Overview

This Kubernetes Operator, written in Go, provides dynamic control over the replica count of your deployments over time. It leverages the Kubernetes API to monitor and adjust the replica count based on specified time intervals.

## Features

- **Dynamic Scaling**: Automatically adjusts the replica count of specified deployments.
- **Time-Based Control**: Define time intervals for scaling operations.
- **Configurability**: Easily configure the target deployments and scaling parameters.

## Prerequisites

- Kubernetes cluster (1.16 or later)
- `kubectl` configured to access your cluster
- Go installed (version X or later)
- [Operator SDK](https://sdk.operatorframework.io/docs/installation/install-operator-sdk/) installed

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/yourusername/k8s-replica-scaler-operator.git
    cd k8s-replica-scaler-operator
    ```

2. Build and deploy the operator:

    ```bash
    make install
    make run
    ```

## Configuration

Adjust the operator behavior by modifying the configuration file located at `config/default/k8sreplicascalerconfig_v1alpha1_k8sreplicascalerconfig.yaml`.

Example Configuration:

```yaml
apiVersion: k8sreplicascalerconfig.example.com/v1alpha1
kind: K8sReplicaScalerConfig
metadata:
  name: example-config
spec:
  deployments:
    - name: your-deployment
      minReplicas: 2
      maxReplicas: 10
      scaleInterval: 5m
```

- `name`: Name of the target deployment.
- `minReplicas`: Minimum replica count.
- `maxReplicas`: Maximum replica count.
- `scaleInterval`: Time interval for scaling operations.

Apply the configuration:

```bash
kubectl apply -f config/default/k8sreplicascalerconfig_v1alpha1_k8sreplicascalerconfig.yaml
```

## Usage

The operator will now watch for changes to the `K8sReplicaScalerConfig` resources and adjust the replica count of the specified deployments accordingly.

To check the operator logs:

```bash
kubectl logs -f deployment/k8s-replica-scaler-operator-controller-manager -n default
```

## Contributing

Feel free to contribute by opening issues or submitting pull requests. See [CONTRIBUTING.md](CONTRIBUTING.md) for more details.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.

---

Adjust the sections and details based on your specific project structure and requirements. Additionally, consider adding more documentation, examples, and testing information as needed.