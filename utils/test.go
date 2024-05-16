package main

import (
	"fmt"
	"regexp"

	"github.com/buger/jsonparser"
)

func main() {
	data := []byte(`{
		"response": {
		  "apiVersion": "config.openshift.io/v1",
		  "items": [
			{
			  "apiVersion": "config.openshift.io/v1",
			  "kind": "OAuth",
			  "metadata": {
				"annotations": {
				  "include.release.openshift.io/ibm-cloud-managed": "true",
				  "include.release.openshift.io/self-managed-high-availability": "true",
				  "include.release.openshift.io/single-node-developer": "true",
				  "kubectl.kubernetes.io/last-applied-configuration": "{\"apiVersion\":\"config.openshift.io/v1\",\"kind\":\"OAuth\",\"metadata\":{\"annotations\":{\"include.release.openshift.io/ibm-cloud-managed\":\"true\",\"include.release.openshift.io/self-managed-high-availability\":\"true\",\"include.release.openshift.io/single-node-developer\":\"true\"},\"name\":\"cluster\"},\"spec\":{\"identityProviders\":[{\"mappingMethod\":\"claim\",\"name\":\"Indradhanus-IAM\",\"openID\":{\"claims\":{\"email\":[\"email\"],\"name\":[\"name\"],\"preferredUsername\":[\"preferred_username\"]},\"clientID\":\"indradhanus\",\"clientSecret\":{\"name\":\"indradhanus-iam-secret\"},\"extraScopes\":[],\"issuer\":\"https://iam-dev.jio.indradhanus.com/realms/indradhanus\"},\"type\":\"OpenID\"}]}}\n",
				  "release.openshift.io/create-only": "true"
				},
				"creationTimestamp": "2024-02-07T11:54:46Z",
				"generation": 11,
				"managedFields": [
				  {
					"apiVersion": "config.openshift.io/v1",
					"fieldsType": "FieldsV1",
					"fieldsV1": {
					  "f:metadata": {
						"f:annotations": {
						  ".": {},
						  "f:include.release.openshift.io/ibm-cloud-managed": {},
						  "f:include.release.openshift.io/self-managed-high-availability": {},
						  "f:include.release.openshift.io/single-node-developer": {},
						  "f:release.openshift.io/create-only": {}
						},
						"f:ownerReferences": {
						  ".": {},
						  "k:{\"uid\":\"fbc4d5db-700d-4902-aca0-bbdd4947f1a1\"}": {}
						}
					  },
					  "f:spec": {}
					},
					"manager": "cluster-version-operator",
					"operation": "Update",
					"time": "2024-02-07T11:54:46Z"
				  },
				  {
					"apiVersion": "config.openshift.io/v1",
					"fieldsType": "FieldsV1",
					"fieldsV1": {
					  "f:metadata": {
						"f:annotations": {
						  "f:kubectl.kubernetes.io/last-applied-configuration": {}
						}
					  },
					  "f:spec": {
						"f:identityProviders": {}
					  }
					},
					"manager": "kubectl-client-side-apply",
					"operation": "Update",
					"time": "2024-04-05T20:38:47Z"
				  }
				],
				"name": "cluster",
				"ownerReferences": [
				  {
					"apiVersion": "config.openshift.io/v1",
					"kind": "ClusterVersion",
					"name": "version",
					"uid": "fbc4d5db-700d-4902-aca0-bbdd4947f1a1"
				  }
				],
				"resourceVersion": "78792121",
				"uid": "41fe9b1c-be5c-4192-ab8c-df4be327a65d"
			  },
			  "spec": {
				"identityProviders": [
				  {
					"mappingMethod": "claim",
					"name": "Indradhanus-IAM",
					"openID": {
					  "claims": {
						"email": [
						  "email"
						],
						"name": [
						  "name"
						],
						"preferredUsername": [
						  "preferred_username"
						]
					  },
					  "clientID": "indradhanus",
					  "clientSecret": {
						"name": "indradhanus-iam-secret"
					  },
					  "extraScopes": [],
					  "issuer": "https://iam-dev.jio.indradhanus.com/realms/indradhanus"
					},
					"type": "OpenID"
				  },
				  {
					"mappingMethod": "claim",
					"name": "Indradhanus-IAM2",
					"openID": {
					  "claims": {
						"email": [
						  "email"
						],
						"name": [
						  "name"
						],
						"preferredUsername": [
						  "preferred_username"
						]
					  },
					  "clientID": "indradhanus",
					  "clientSecret": {
						"name": "indradhanus-iam-secret"
					  },
					  "extraScopes": [],
					  "issuer": "https://iam-dev.jio.indradhanus.com/realms/indradhanus"
					},
					"type": "OpenID"
				  }
				]
			  }
			},
			 {
			  "apiVersion": "config.openshift.io/v1",
			  "kind": "OAuth",
			  "metadata": {
				"annotations": {
				  "include.release.openshift.io/ibm-cloud-managed": "true",
				  "include.release.openshift.io/self-managed-high-availability": "true",
				  "include.release.openshift.io/single-node-developer": "true",
				  "kubectl.kubernetes.io/last-applied-configuration": "{\"apiVersion\":\"config.openshift.io/v1\",\"kind\":\"OAuth\",\"metadata\":{\"annotations\":{\"include.release.openshift.io/ibm-cloud-managed\":\"true\",\"include.release.openshift.io/self-managed-high-availability\":\"true\",\"include.release.openshift.io/single-node-developer\":\"true\"},\"name\":\"cluster\"},\"spec\":{\"identityProviders\":[{\"mappingMethod\":\"claim\",\"name\":\"Indradhanus-IAM\",\"openID\":{\"claims\":{\"email\":[\"email\"],\"name\":[\"name\"],\"preferredUsername\":[\"preferred_username\"]},\"clientID\":\"indradhanus\",\"clientSecret\":{\"name\":\"indradhanus-iam-secret\"},\"extraScopes\":[],\"issuer\":\"https://iam-dev.jio.indradhanus.com/realms/indradhanus\"},\"type\":\"OpenID\"}]}}\n",
				  "release.openshift.io/create-only": "true"
				},
				"creationTimestamp": "2024-02-07T11:54:46Z",
				"generation": 11,
				"managedFields": [
				  {
					"apiVersion": "config.openshift.io/v1",
					"fieldsType": "FieldsV1",
					"fieldsV1": {
					  "f:metadata": {
						"f:annotations": {
						  ".": {},
						  "f:include.release.openshift.io/ibm-cloud-managed": {},
						  "f:include.release.openshift.io/self-managed-high-availability": {},
						  "f:include.release.openshift.io/single-node-developer": {},
						  "f:release.openshift.io/create-only": {}
						},
						"f:ownerReferences": {
						  ".": {},
						  "k:{\"uid\":\"fbc4d5db-700d-4902-aca0-bbdd4947f1a1\"}": {}
						}
					  },
					  "f:spec": {}
					},
					"manager": "cluster-version-operator",
					"operation": "Update",
					"time": "2024-02-07T11:54:46Z"
				  },
				  {
					"apiVersion": "config.openshift.io/v1",
					"fieldsType": "FieldsV1",
					"fieldsV1": {
					  "f:metadata": {
						"f:annotations": {
						  "f:kubectl.kubernetes.io/last-applied-configuration": {}
						}
					  },
					  "f:spec": {
						"f:identityProviders": {}
					  }
					},
					"manager": "kubectl-client-side-apply",
					"operation": "Update",
					"time": "2024-04-05T20:38:47Z"
				  }
				],
				"name": "cluster",
				"ownerReferences": [
				  {
					"apiVersion": "config.openshift.io/v1",
					"kind": "ClusterVersion",
					"name": "version",
					"uid": "fbc4d5db-700d-4902-aca0-bbdd4947f1a1"
				  }
				],
				"resourceVersion": "78792121",
				"uid": "41fe9b1c-be5c-4192-ab8c-df4be327a65d"
			  },
			  "spec": {
				"identityProviders": [
				  {
					"mappingMethod": "claim",
					"name": "Indradhanus-IAM",
					"openID": {
					  "claims": {
						"email": [
						  "email"
						],
						"name": [
						  "name"
						],
						"preferredUsername": [
						  "preferred_username"
						]
					  },
					  "clientID": "indradhanus",
					  "clientSecret": {
						"name": "indradhanus-iam-secret"
					  },
					  "extraScopes": [],
					  "issuer": "https://iam-dev.jio.indradhanus.com/realms/indradhanus"
					},
					"type": "OpenID"
				  }
				]
			  }
			}
		  ],
		  "kind": "OAuthList",
		  "metadata": {
			"continue": "",
			"resourceVersion": "175342313"
		  }
		},
		"status": "OK"
	  }`)
	fmt.Println(jsonparser.Get(data, "response"))
	jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		jsonparser.ArrayEach(value, func(providerValue []byte, dataType jsonparser.ValueType, providerOffset int, providerErr error) {
			clientID, _ := jsonparser.GetString(providerValue, "openID", "clientID")
			name, _ := jsonparser.GetString(providerValue, "name")
			fmt.Printf("Client ID: %s, Name: %s\n", clientID, name)
			matched, err := regexp.MatchString(name, "Indradhanus-IAM")
			fmt.Println(matched, err)
		}, "spec", "identityProviders")
	}, "response", "items")
	// jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {

	// 	jsonparser.ArrayEach(value, func(providerValue []byte, dataType jsonparser.ValueType, offset int, err error) {
	// 		fmt.Println("I")
	// 		clientID, _ := jsonparser.GetString(providerValue, "openID", "clientID")
	// 		name, _ := jsonparser.GetString(providerValue, "name")
	// 		fmt.Printf("Client ID: %s, Name: %s\n", clientID, name)
	// 	}, "spec", "identityProviders")
	// }, "response", "items")

}
