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
git clone https://github.com/sadegh-msm/replica-operator.git
cd replica-operator
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
