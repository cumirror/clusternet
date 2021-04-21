/*
Copyright 2021 The Clusternet Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package known

// label key
const (
	ClusterRegisteredByLabel  = "clusters.clusternet.io/registered-by"
	ClusterIDLabel            = "clusters.clusternet.io/cluster-id"
	ClusterNameLabel          = "clusters.clusternet.io/cluster-name"
	ClusterBootstrappingLabel = "clusters.clusternet.io/bootstrapping"
	ObjectCreatedByLabel      = "clusternet.io/created-by"
)

// label value
const (
	CredentialsAuto = "credentials-auto"
	RBACDefaults    = "rbac-defaults"

	ClusternetAgentName = "clusternet-agent"
	ClusternetHubName   = "clusternet-hub"
)