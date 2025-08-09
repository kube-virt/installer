#!/bin/bash

# Copyright AppsCode Inc. and Contributors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

KLUSTER_MANAGER_CLUSTER_AUTH_TAG=${KLUSTER_MANAGER_CLUSTER_AUTH_TAG:-v0.4.0}
KLUSTER_MANAGER_CLUSTER_PROFILE_TAG=${KLUSTER_MANAGER_CLUSTER_PROFILE_TAG:-v0.9.0}
KLUSTER_MANAGER_FLUXCD_ADDON_TAG=${KLUSTER_MANAGER_FLUXCD_ADDON_TAG:-v0.0.7}
OPEN_CLUSTER_MANAGEMENT_IO_API_TAG=${OPEN_CLUSTER_MANAGEMENT_IO_API_TAG:-v1.0.0}
PROMETHEUS_OPERATOR_PROMETHEUS_OPERATOR=${PROMETHEUS_OPERATOR_PROMETHEUS_OPERATOR:-v0.81.0}

crd-importer \
    --input=https://github.com/prometheus-operator/prometheus-operator/raw/${PROMETHEUS_OPERATOR_PROMETHEUS_OPERATOR}/example/prometheus-operator-crd/monitoring.coreos.com_servicemonitors.yaml \
    --out=./charts/cluster-auth/crds

crd-importer \
    --input=https://github.com/kluster-manager/cluster-auth/raw/${KLUSTER_MANAGER_CLUSTER_AUTH_TAG}/crds/authentication.k8s.appscode.com_accounts.yaml \
    --input=https://github.com/kluster-manager/cluster-auth/raw/${KLUSTER_MANAGER_CLUSTER_AUTH_TAG}/crds/authorization.k8s.appscode.com_managedclusterrolebindings.yaml \
    --input=https://github.com/kluster-manager/cluster-auth/raw/${KLUSTER_MANAGER_CLUSTER_AUTH_TAG}/crds/authorization.k8s.appscode.com_managedclusterroles.yaml \
    --input=https://github.com/kluster-manager/cluster-auth/raw/${KLUSTER_MANAGER_CLUSTER_AUTH_TAG}/crds/authorization.k8s.appscode.com_managedclustersetrolebindings.yaml \
    --input=https://github.com/open-cluster-management-io/api/raw/${OPEN_CLUSTER_MANAGEMENT_IO_API_TAG}/addon/v1alpha1/0000_00_addon.open-cluster-management.io_clustermanagementaddons.crd.yaml \
    --input=https://github.com/open-cluster-management-io/api/raw/${OPEN_CLUSTER_MANAGEMENT_IO_API_TAG}/cluster/v1beta1/0000_02_clusters.open-cluster-management.io_placements.crd.yaml \
    --input=https://github.com/open-cluster-management-io/api/raw/${OPEN_CLUSTER_MANAGEMENT_IO_API_TAG}/cluster/v1beta2/0000_01_clusters.open-cluster-management.io_managedclustersetbindings.crd.yaml \
    --input=https://github.com/open-cluster-management-io/api/raw/${OPEN_CLUSTER_MANAGEMENT_IO_API_TAG}/work/v1alpha1/0000_00_work.open-cluster-management.io_manifestworkreplicasets.crd.yaml \
    --out=./charts/cluster-auth-manager/crds

crd-importer \
    --input=https://github.com/kluster-manager/cluster-gateway/raw/master/hack/crd/bases/config.gateway.open-cluster-management.io_clustergatewayconfigurations.yaml \
    --input=https://github.com/open-cluster-management-io/api/raw/${OPEN_CLUSTER_MANAGEMENT_IO_API_TAG}/addon/v1alpha1/0000_00_addon.open-cluster-management.io_clustermanagementaddons.crd.yaml \
    --input=https://github.com/open-cluster-management-io/api/raw/${OPEN_CLUSTER_MANAGEMENT_IO_API_TAG}/cluster/v1beta1/0000_02_clusters.open-cluster-management.io_placements.crd.yaml \
    --input=https://github.com/open-cluster-management-io/api/raw/${OPEN_CLUSTER_MANAGEMENT_IO_API_TAG}/cluster/v1beta2/0000_01_clusters.open-cluster-management.io_managedclustersetbindings.crd.yaml \
    --out=./charts/cluster-gateway-manager/crds

crd-importer \
    --input=https://github.com/open-cluster-management-io/api/raw/${OPEN_CLUSTER_MANAGEMENT_IO_API_TAG}/operator/v1/0000_01_operator.open-cluster-management.io_clustermanagers.crd.yaml \
    --out=./charts/cluster-manager-hub/crds

crd-importer \
    --input=https://github.com/open-cluster-management-io/api/raw/${OPEN_CLUSTER_MANAGEMENT_IO_API_TAG}/operator/v1/0000_00_operator.open-cluster-management.io_klusterlets.crd.yaml \
    --out=./charts/cluster-manager-spoke/crds

crd-importer \
    --input=https://github.com/kluster-manager/cluster-auth/raw/${KLUSTER_MANAGER_CLUSTER_AUTH_TAG}/crds/authentication.k8s.appscode.com_accounts.yaml \
    --input=https://github.com/kluster-manager/cluster-profile/raw/${KLUSTER_MANAGER_CLUSTER_PROFILE_TAG}/crds/profile.k8s.appscode.com_managedclusterprofilebindings.yaml \
    --input=https://github.com/kluster-manager/cluster-profile/raw/${KLUSTER_MANAGER_CLUSTER_PROFILE_TAG}/crds/profile.k8s.appscode.com_managedclustersetprofiles.yaml \
    --input=https://github.com/open-cluster-management-io/api/raw/${OPEN_CLUSTER_MANAGEMENT_IO_API_TAG}/cluster/v1beta1/0000_02_clusters.open-cluster-management.io_placements.crd.yaml \
    --input=https://github.com/open-cluster-management-io/api/raw/${OPEN_CLUSTER_MANAGEMENT_IO_API_TAG}/cluster/v1beta2/0000_01_clusters.open-cluster-management.io_managedclustersetbindings.crd.yaml \
    --input=https://github.com/open-cluster-management-io/api/raw/${OPEN_CLUSTER_MANAGEMENT_IO_API_TAG}/work/v1alpha1/0000_00_work.open-cluster-management.io_manifestworkreplicasets.crd.yaml \
    --out=./charts/cluster-profile-manager/crds

crd-importer \
    --input=https://github.com/open-cluster-management-io/api/raw/${OPEN_CLUSTER_MANAGEMENT_IO_API_TAG}/addon/v1alpha1/0000_00_addon.open-cluster-management.io_clustermanagementaddons.crd.yaml \
    --input=https://github.com/open-cluster-management-io/api/raw/${OPEN_CLUSTER_MANAGEMENT_IO_API_TAG}/cluster/v1beta1/0000_02_clusters.open-cluster-management.io_placements.crd.yaml \
    --input=https://github.com/open-cluster-management-io/api/raw/${OPEN_CLUSTER_MANAGEMENT_IO_API_TAG}/cluster/v1beta2/0000_01_clusters.open-cluster-management.io_managedclustersetbindings.crd.yaml \
    --input=https://github.com/kluster-manager/cluster-proxy/raw/${KLUSTER_MANAGER_CLUSTER_PROFILE_TAG}/hack/crd/bases/proxy.open-cluster-management.io_managedproxyconfigurations.yaml \
    --input=https://github.com/kluster-manager/cluster-proxy/raw/${KLUSTER_MANAGER_CLUSTER_PROFILE_TAG}/hack/crd/bases/proxy.open-cluster-management.io_managedproxyserviceresolvers.yaml \
    --out=./charts/cluster-proxy-manager/crds

crd-importer \
    --input=https://github.com/kluster-manager/fluxcd-addon/raw/${KLUSTER_MANAGER_FLUXCD_ADDON_TAG}/crds/fluxcd.open-cluster-management.io_fluxcdconfigs.yaml \
    --input=https://github.com/open-cluster-management-io/api/raw/${OPEN_CLUSTER_MANAGEMENT_IO_API_TAG}/addon/v1alpha1/0000_00_addon.open-cluster-management.io_clustermanagementaddons.crd.yaml \
    --input=https://github.com/open-cluster-management-io/api/raw/${OPEN_CLUSTER_MANAGEMENT_IO_API_TAG}/cluster/v1beta1/0000_02_clusters.open-cluster-management.io_placements.crd.yaml \
    --input=https://github.com/open-cluster-management-io/api/raw/${OPEN_CLUSTER_MANAGEMENT_IO_API_TAG}/cluster/v1beta2/0000_01_clusters.open-cluster-management.io_managedclustersetbindings.crd.yaml \
    --out=./charts/fluxcd-manager/crds

crd-importer \
    --input=https://github.com/open-cluster-management-io/api/raw/${OPEN_CLUSTER_MANAGEMENT_IO_API_TAG}/addon/v1alpha1/0000_00_addon.open-cluster-management.io_clustermanagementaddons.crd.yaml \
    --input=https://github.com/open-cluster-management-io/api/raw/${OPEN_CLUSTER_MANAGEMENT_IO_API_TAG}/cluster/v1beta1/0000_02_clusters.open-cluster-management.io_placements.crd.yaml \
    --input=https://github.com/open-cluster-management-io/api/raw/${OPEN_CLUSTER_MANAGEMENT_IO_API_TAG}/cluster/v1beta2/0000_01_clusters.open-cluster-management.io_managedclustersetbindings.crd.yaml \
    --input=https://github.com/kluster-manager/managed-serviceaccount/raw/main/config/crd/bases/authentication.open-cluster-management.io_managedserviceaccounts.yaml \
    --out=./charts/managed-serviceaccount-manager/crds
