# LoadBalancer Controller

[LoadBalancer Controller by AppsCode](https://github.com/kube-virt/installer) - A Helm chart for the LoadBalancer Controller (loadbalancer-controller + keepalived-grpc)

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/loadbalancer-controller --version=v0.1.0
$ helm upgrade -i loadbalancer-controller appscode/loadbalancer-controller -n loadbalancer-controller --create-namespace --version=v0.1.0
```

## Introduction

This chart deploys a LoadBalancer controller on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.28+

## Installing the Chart

To install/upgrade the chart with the release name `loadbalancer-controller`:

```bash
$ helm upgrade -i loadbalancer-controller appscode/loadbalancer-controller -n loadbalancer-controller --create-namespace --version=v0.1.0
```

The command deploys a LoadBalancer controller on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `loadbalancer-controller`:

```bash
$ helm uninstall loadbalancer-controller -n loadbalancer-controller
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `loadbalancer-controller` chart and their default values.

|                  Parameter                  |                                                                                                                        Description                                                                                                                        |                        Default                         |
|---------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------------------------------------------|
| nameOverride                                |                                                                                                                                                                                                                                                           | <code>""</code>                                        |
| fullnameOverride                            |                                                                                                                                                                                                                                                           | <code>""</code>                                        |
| mode                                        | ------------------------------------------------------------------- Mode: "openstack" or "harvester" In harvester mode, a kubeconfig secret is mounted; openstack-creds are NOT used. ------------------------------------------------------------------- | <code>harvester # openstack &#124; harvester</code>    |
| loadbalancerController.replicaCount         |                                                                                                                                                                                                                                                           | <code>2</code>                                         |
| loadbalancerController.image.repository     |                                                                                                                                                                                                                                                           | <code>ghcr.io/kube-virt/loadbalancer-controller</code> |
| loadbalancerController.image.tag            |                                                                                                                                                                                                                                                           | <code>v0.1.0-ac</code>                                 |
| loadbalancerController.image.pullPolicy     |                                                                                                                                                                                                                                                           | <code>IfNotPresent</code>                              |
| loadbalancerController.serviceAccountName   |                                                                                                                                                                                                                                                           | <code>loadbalancer-controller</code>                   |
| loadbalancerController.env.osTenantId       |                                                                                                                                                                                                                                                           | <code>""</code>                                        |
| loadbalancerController.env.osNetworkId      |                                                                                                                                                                                                                                                           | <code>""</code>                                        |
| loadbalancerController.env.osSubnetId       |                                                                                                                                                                                                                                                           | <code>""</code>                                        |
| loadbalancerController.env.clusterNamespace |                                                                                                                                                                                                                                                           | <code>"new"</code>                                     |
| loadbalancerController.resources            |                                                                                                                                                                                                                                                           | <code>{}</code>                                        |
| openstackCreds.osAuthUrl                    |                                                                                                                                                                                                                                                           | <code>""</code>                                        |
| openstackCreds.osUsername                   |                                                                                                                                                                                                                                                           | <code>"admin"</code>                                   |
| openstackCreds.osPassword                   |                                                                                                                                                                                                                                                           | <code>""</code>                                        |
| openstackCreds.osDomainName                 |                                                                                                                                                                                                                                                           | <code>"Default"</code>                                 |
| openstackCreds.osRegionName                 |                                                                                                                                                                                                                                                           | <code>"RegionOne"</code>                               |
| openstackCreds.osProjectId                  |                                                                                                                                                                                                                                                           | <code>""</code>                                        |
| dbCredentials.dbHost                        |                                                                                                                                                                                                                                                           | <code>""</code>                                        |
| dbCredentials.dbPort                        |                                                                                                                                                                                                                                                           | <code>"3306"</code>                                    |
| dbCredentials.dbUser                        |                                                                                                                                                                                                                                                           | <code>"lbapi"</code>                                   |
| dbCredentials.dbPassword                    |                                                                                                                                                                                                                                                           | <code>"LbApiStrongPass123!"</code>                     |
| dbCredentials.dbName                        |                                                                                                                                                                                                                                                           | <code>"lbapi"</code>                                   |
| dbCredentials.dbSslmode                     |                                                                                                                                                                                                                                                           | <code>"disable"</code>                                 |
| harvester.kubeconfigSecretName              |                                                                                                                                                                                                                                                           | <code>harvester-kubeconfig</code>                      |
| harvester.kubeconfig                        | Paste the full kubeconfig as a multiline string value, or leave empty and create the secret manually.                                                                                                                                                     | <code>""</code>                                        |
| keepalivedGrpc.image.repository             |                                                                                                                                                                                                                                                           | <code>ghcr.io/kube-virt/keepalived</code>              |
| keepalivedGrpc.image.tag                    |                                                                                                                                                                                                                                                           | <code>v0.1.0-ac</code>                                 |
| keepalivedGrpc.image.pullPolicy             |                                                                                                                                                                                                                                                           | <code>IfNotPresent</code>                              |
| keepalivedGrpc.hostNetwork                  |                                                                                                                                                                                                                                                           | <code>true</code>                                      |
| keepalivedGrpc.hostPID                      |                                                                                                                                                                                                                                                           | <code>true</code>                                      |
| keepalivedGrpc.resources.requests.cpu       |                                                                                                                                                                                                                                                           | <code>50m</code>                                       |
| keepalivedGrpc.resources.requests.memory    |                                                                                                                                                                                                                                                           | <code>64Mi</code>                                      |
| keepalivedGrpc.keepalivedConfigPath         | hostPath for keepalived config directory                                                                                                                                                                                                                  | <code>/etc/keepalived</code>                           |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i loadbalancer-controller appscode/loadbalancer-controller -n loadbalancer-controller --create-namespace --version=v0.1.0 --set mode=harvester # openstack | harvester
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i loadbalancer-controller appscode/loadbalancer-controller -n loadbalancer-controller --create-namespace --version=v0.1.0 --values values.yaml
```
