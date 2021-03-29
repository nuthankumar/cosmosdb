package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DBProvisioning is a top-level type
type DBProvisioning struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +optional
	Status DBProvisioningStatus `json:"status,omitempty"`
	// This is where you can define
	// your own custom spec
	Spec DBProvisioningSpec `json:"spec,omitempty"`
}

// custom spec
type DBProvisioningSpec struct {
	ClientID                    string `json:"clientid"`
	ClientNamespace             string `json:"clientns"`
	Keyspace                    string `json:"keyspacename"`
	Datacenter                  string `json:"datacenter"`
	KeyspaceTopology            string `json:"keyspacetopology"`
	ReplicationFactor           string `json:"replicationfactor"`
	Rolename                    string `json:"rolename"`
	Servicename                 string `json:"servicename"`
	CreateSchema                bool   `json:"createschema"`
	Configmap                   string `json:"configmapname"`
	SchemaScriptName            string `json:"schemascriptname"`
	DefaultValueScriptName      string `json:"defaultvaluescriptname"`
	CosmosCassandraContactPoint string `json:"cassandracontactpoint"`
	CosmosCassandraPassword     string `json:"cassandrapassword"`
}

// custom status
type DBProvisioningStatus struct {
	Name string
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// no client needed for list as it's been created in above
type DBProvisioningList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `son:"metadata,omitempty"`

	Items []DBProvisioning `json:"items"`
}
