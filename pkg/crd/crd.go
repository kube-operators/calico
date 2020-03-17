package crd

import (
	"k8s.io/api/core/v1"
	apiext "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
)

//apiVersion: apiextensions.k8s.io/v1beta1
//kind: CustomResourceDefinition
//metadata:
//  name: felixconfigurations.crd.projectcalico.org
//spec:
//  scope: Cluster
//  group: crd.projectcalico.org
//  version: v1
//  names:
//    kind: FelixConfiguration
//    plural: felixconfigurations
//    singular: felixconfiguration

func installCalicoCRD() {
	apiext.CustomResourceColumnDefinition{}
}
