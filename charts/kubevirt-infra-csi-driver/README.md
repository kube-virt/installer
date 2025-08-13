# KubeVirt infra CSI driver

[KubeVirt infra CSI driver by AppsCode](https://github.com/kube-virt/csi-driver) - A Helm chart for KubeVirt infra CSI driver

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/kubevirt-infra-csi-driver --version=v0.1.0
$ helm upgrade -i kubevirt-infra-csi-driver appscode/kubevirt-infra-csi-driver -n kubevirt-csi-driver --create-namespace --version=v0.1.0
```

## Introduction

This chart deploys a CSI driver on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.28+

## Installing the Chart

To install/upgrade the chart with the release name `kubevirt-infra-csi-driver`:

```bash
$ helm upgrade -i kubevirt-infra-csi-driver appscode/kubevirt-infra-csi-driver -n kubevirt-csi-driver --create-namespace --version=v0.1.0
```

The command deploys a CSI driver on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `kubevirt-infra-csi-driver`:

```bash
$ helm uninstall kubevirt-infra-csi-driver -n kubevirt-csi-driver
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `kubevirt-infra-csi-driver` chart and their default values.

|           Parameter           |                                                                                           Description                                                                                            |                                                                      Default                                                                       |
|-------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------|
| replicaCount                  | This will set the replicaset count more information can be found here: https://kubernetes.io/docs/concepts/workloads/controllers/replicaset/                                                     | <code>1</code>                                                                                                                                     |
| driver.repository             |                                                                                                                                                                                                  | <code>quay.io/kubevirt/kubevirt-csi-driver</code>                                                                                                  |
| driver.tag                    | Overrides the image tag whose default is the chart appVersion.                                                                                                                                   | <code>"latest"</code>                                                                                                                              |
| driver.pullPolicy             | This sets the pull policy for images.                                                                                                                                                            | <code>Always</code>                                                                                                                                |
| driver.securityContext        |                                                                                                                                                                                                  | <code>{}</code>                                                                                                                                    |
| driver.resources              |                                                                                                                                                                                                  | <code>{"requests":{"cpu":"10m","memory":"50Mi"}}</code>                                                                                            |
| provisioner.repository        |                                                                                                                                                                                                  | <code>quay.io/openshift/origin-csi-external-provisioner</code>                                                                                     |
| provisioner.tag               | Overrides the image tag whose default is the chart appVersion.                                                                                                                                   | <code>"latest"</code>                                                                                                                              |
| provisioner.pullPolicy        | This sets the pull policy for images.                                                                                                                                                            | <code>IfNotPresent</code>                                                                                                                          |
| provisioner.securityContext   |                                                                                                                                                                                                  | <code>{}</code>                                                                                                                                    |
| provisioner.resources         |                                                                                                                                                                                                  | <code>{}</code>                                                                                                                                    |
| attacher.repository           |                                                                                                                                                                                                  | <code>quay.io/openshift/origin-csi-external-attacher</code>                                                                                        |
| attacher.tag                  | Overrides the image tag whose default is the chart appVersion.                                                                                                                                   | <code>"latest"</code>                                                                                                                              |
| attacher.pullPolicy           | This sets the pull policy for images.                                                                                                                                                            | <code>IfNotPresent</code>                                                                                                                          |
| attacher.securityContext      |                                                                                                                                                                                                  | <code>{}</code>                                                                                                                                    |
| attacher.resources            |                                                                                                                                                                                                  | <code>{"requests":{"cpu":"10m","memory":"50Mi"}}</code>                                                                                            |
| livenessprobe.repository      |                                                                                                                                                                                                  | <code>quay.io/openshift/origin-csi-livenessprobe</code>                                                                                            |
| livenessprobe.tag             | Overrides the image tag whose default is the chart appVersion.                                                                                                                                   | <code>"latest"</code>                                                                                                                              |
| livenessprobe.pullPolicy      | This sets the pull policy for images.                                                                                                                                                            | <code>IfNotPresent</code>                                                                                                                          |
| livenessprobe.securityContext |                                                                                                                                                                                                  | <code>{}</code>                                                                                                                                    |
| livenessprobe.resources       |                                                                                                                                                                                                  | <code>{"requests":{"cpu":"10m","memory":"50Mi"}}</code>                                                                                            |
| snapshotter.repository        |                                                                                                                                                                                                  | <code>k8s.gcr.io/sig-storage/csi-snapshotter</code>                                                                                                |
| snapshotter.tag               | Overrides the image tag whose default is the chart appVersion.                                                                                                                                   | <code>"v4.2.1"</code>                                                                                                                              |
| snapshotter.pullPolicy        | This sets the pull policy for images.                                                                                                                                                            | <code>IfNotPresent</code>                                                                                                                          |
| snapshotter.securityContext   |                                                                                                                                                                                                  | <code>{}</code>                                                                                                                                    |
| snapshotter.resources         |                                                                                                                                                                                                  | <code>{"requests":{"cpu":"10m","memory":"20Mi"}}</code>                                                                                            |
| resizer.repository            |                                                                                                                                                                                                  | <code>registry.k8s.io/sig-storage/csi-resizer</code>                                                                                               |
| resizer.tag                   | Overrides the image tag whose default is the chart appVersion.                                                                                                                                   | <code>"v1.13.1"</code>                                                                                                                             |
| resizer.pullPolicy            | This sets the pull policy for images.                                                                                                                                                            | <code>IfNotPresent</code>                                                                                                                          |
| resizer.securityContext       |                                                                                                                                                                                                  | <code>{"capabilities":{"drop":["ALL"]}}</code>                                                                                                     |
| resizer.resources             |                                                                                                                                                                                                  | <code>{"requests":{"cpu":"10m","memory":"20Mi"}}</code>                                                                                            |
| imagePullSecrets              | This is for the secrets for pulling an image from a private repository more information can be found here: https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/ | <code>[]</code>                                                                                                                                    |
| nameOverride                  | This is to override the chart name.                                                                                                                                                              | <code>""</code>                                                                                                                                    |
| fullnameOverride              |                                                                                                                                                                                                  | <code>""</code>                                                                                                                                    |
| serviceAccount.create         | Specifies whether a service account should be created                                                                                                                                            | <code>true</code>                                                                                                                                  |
| serviceAccount.automount      | Automatically mount a ServiceAccount's API credentials?                                                                                                                                          | <code>true</code>                                                                                                                                  |
| serviceAccount.annotations    | Annotations to add to the service account                                                                                                                                                        | <code>{}</code>                                                                                                                                    |
| serviceAccount.name           | The name of the service account to use. If not set and create is true, a name is generated using the fullname template                                                                           | <code>""</code>                                                                                                                                    |
| podAnnotations                | This is for setting Kubernetes Annotations to a Pod. For more information checkout: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/                               | <code>{}</code>                                                                                                                                    |
| podLabels                     | This is for setting Kubernetes Labels to a Pod. For more information checkout: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/                                         | <code>{}</code>                                                                                                                                    |
| podSecurityContext            |                                                                                                                                                                                                  | <code>{}</code>                                                                                                                                    |
| nodeSelector                  |                                                                                                                                                                                                  | <code>{"node-role.kubernetes.io/control-plane":"true"}</code>                                                                                      |
| tolerations                   |                                                                                                                                                                                                  | <code>[{"key":"CriticalAddonsOnly","operator":"Exists"},{"effect":"NoSchedule","key":"node-role.kubernetes.io/master","operator":"Exists"}]</code> |
| affinity                      |                                                                                                                                                                                                  | <code>{}</code>                                                                                                                                    |
| priorityClassName             |                                                                                                                                                                                                  | <code>system-cluster-critical</code>                                                                                                               |
| tenant.namespace              |                                                                                                                                                                                                  | <code>""</code>                                                                                                                                    |
| tenant.labels                 |                                                                                                                                                                                                  | <code>csi-driver/cluster=tenant</code>                                                                                                             |
| tenant.kubeconfig             |                                                                                                                                                                                                  | <code>""</code>                                                                                                                                    |
| infraStorageClassEnforcement  |                                                                                                                                                                                                  | <code>{}</code>                                                                                                                                    |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i kubevirt-infra-csi-driver appscode/kubevirt-infra-csi-driver -n kubevirt-csi-driver --create-namespace --version=v0.1.0 --set replicaCount=1
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i kubevirt-infra-csi-driver appscode/kubevirt-infra-csi-driver -n kubevirt-csi-driver --create-namespace --version=v0.1.0 --values values.yaml
```
