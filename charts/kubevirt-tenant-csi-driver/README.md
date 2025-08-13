# KubeVirt tenant CSI driver

[KubeVirt tenant CSI driver by AppsCode](https://github.com/kube-virt/csi-driver) - A Helm chart for KubeVirt tenant CSI driver

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/kubevirt-tenant-csi-driver --version=v0.1.0
$ helm upgrade -i kubevirt-tenant-csi-driver appscode/kubevirt-tenant-csi-driver -n kubevirt-csi-driver --create-namespace --version=v0.1.0
```

## Introduction

This chart deploys a CSI driver on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.28+

## Installing the Chart

To install/upgrade the chart with the release name `kubevirt-tenant-csi-driver`:

```bash
$ helm upgrade -i kubevirt-tenant-csi-driver appscode/kubevirt-tenant-csi-driver -n kubevirt-csi-driver --create-namespace --version=v0.1.0
```

The command deploys a CSI driver on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `kubevirt-tenant-csi-driver`:

```bash
$ helm uninstall kubevirt-tenant-csi-driver -n kubevirt-csi-driver
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `kubevirt-tenant-csi-driver` chart and their default values.

|              Parameter              |                                                                                           Description                                                                                            |                             Default                              |
|-------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------|
| driver.repository                   |                                                                                                                                                                                                  | <code>quay.io/kubevirt/kubevirt-csi-driver</code>                |
| driver.tag                          | Overrides the image tag whose default is the chart appVersion.                                                                                                                                   | <code>"latest"</code>                                            |
| driver.pullPolicy                   | This sets the pull policy for images.                                                                                                                                                            | <code>Always</code>                                              |
| driver.securityContext              |                                                                                                                                                                                                  | <code>{"allowPrivilegeEscalation":true,"privileged":true}</code> |
| driver.resources                    |                                                                                                                                                                                                  | <code>{"requests":{"cpu":"10m","memory":"50Mi"}}</code>          |
| nodeDriverRegistrar.repository      |                                                                                                                                                                                                  | <code>quay.io/openshift/origin-csi-node-driver-registrar</code>  |
| nodeDriverRegistrar.tag             | Overrides the image tag whose default is the chart appVersion.                                                                                                                                   | <code>"latest"</code>                                            |
| nodeDriverRegistrar.pullPolicy      | This sets the pull policy for images.                                                                                                                                                            | <code>IfNotPresent</code>                                        |
| nodeDriverRegistrar.securityContext |                                                                                                                                                                                                  | <code>{}</code>                                                  |
| nodeDriverRegistrar.resources       |                                                                                                                                                                                                  | <code>{"requests":{"cpu":"5m","memory":"20Mi"}}</code>           |
| livenessprobe.repository            |                                                                                                                                                                                                  | <code>quay.io/openshift/origin-csi-livenessprobe</code>          |
| livenessprobe.tag                   | Overrides the image tag whose default is the chart appVersion.                                                                                                                                   | <code>"latest"</code>                                            |
| livenessprobe.pullPolicy            | This sets the pull policy for images.                                                                                                                                                            | <code>IfNotPresent</code>                                        |
| livenessprobe.securityContext       |                                                                                                                                                                                                  | <code>{}</code>                                                  |
| livenessprobe.resources             |                                                                                                                                                                                                  | <code>{"requests":{"cpu":"5m","memory":"20Mi"}}</code>           |
| imagePullSecrets                    | This is for the secrets for pulling an image from a private repository more information can be found here: https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/ | <code>[]</code>                                                  |
| nameOverride                        | This is to override the chart name.                                                                                                                                                              | <code>""</code>                                                  |
| fullnameOverride                    |                                                                                                                                                                                                  | <code>""</code>                                                  |
| serviceAccount.create               | Specifies whether a service account should be created                                                                                                                                            | <code>true</code>                                                |
| serviceAccount.automount            | Automatically mount a ServiceAccount's API credentials?                                                                                                                                          | <code>true</code>                                                |
| serviceAccount.annotations          | Annotations to add to the service account                                                                                                                                                        | <code>{}</code>                                                  |
| serviceAccount.name                 | The name of the service account to use. If not set and create is true, a name is generated using the fullname template                                                                           | <code>""</code>                                                  |
| podAnnotations                      | This is for setting Kubernetes Annotations to a Pod. For more information checkout: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/                               | <code>{}</code>                                                  |
| podLabels                           | This is for setting Kubernetes Labels to a Pod. For more information checkout: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/                                         | <code>{}</code>                                                  |
| podSecurityContext                  |                                                                                                                                                                                                  | <code>{}</code>                                                  |
| nodeSelector                        |                                                                                                                                                                                                  | <code>{}</code>                                                  |
| tolerations                         |                                                                                                                                                                                                  | <code>[{"operator":"Exists"}]</code>                             |
| affinity                            |                                                                                                                                                                                                  | <code>{}</code>                                                  |
| priorityClassName                   |                                                                                                                                                                                                  | <code>system-node-critical</code>                                |
| tenant.namespace                    |                                                                                                                                                                                                  | <code>""</code>                                                  |
| tenant.labels                       |                                                                                                                                                                                                  | <code>csi-driver/cluster=tenant</code>                           |
| infra.storageClassName              |                                                                                                                                                                                                  | <code>standard</code>                                            |
| infra.snapshotClassName             |                                                                                                                                                                                                  | <code>{}</code>                                                  |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i kubevirt-tenant-csi-driver appscode/kubevirt-tenant-csi-driver -n kubevirt-csi-driver --create-namespace --version=v0.1.0 --set driver.repository=quay.io/kubevirt/kubevirt-csi-driver
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i kubevirt-tenant-csi-driver appscode/kubevirt-tenant-csi-driver -n kubevirt-csi-driver --create-namespace --version=v0.1.0 --values values.yaml
```
