package main

import (
	"fmt"
	crdv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"

	"log"
	"os"
)

type calicoCRDs []crdv1beta1.CustomResourceDefinition

const crdKind = "CustomResourceDefinition"
const crdApiVersion = "apiextensions.k8s.io/v1beta1"

const calicoCRDGroup = "crd.projectcalico.org"
const calicoCRDVersion = "v1"

func getCalicoCRDs() {
	crds := calicoCRDs{
		crdv1beta1.CustomResourceDefinition{
			TypeMeta: metav1.TypeMeta{
				Kind:       crdKind,
				APIVersion: crdApiVersion,
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: "blockaffinities.crd.projectcalico.org",
			},
			Spec: crdv1beta1.CustomResourceDefinitionSpec{
				Scope:   crdv1beta1.ClusterScoped,
				Group:   calicoCRDGroup,
				Version: calicoCRDVersion,
				Names: crdv1beta1.CustomResourceDefinitionNames{
					Plural:     "blockaffinities",
					Singular:   "blockaffinity",
					ShortNames: nil,
					Kind:       "BlockAffinity",
					ListKind:   "",
					Categories: nil,
				},
			},
			Status: crdv1beta1.CustomResourceDefinitionStatus{},
		},

		crdv1beta1.CustomResourceDefinition{
			TypeMeta: metav1.TypeMeta{
				Kind:       crdKind,
				APIVersion: crdApiVersion,
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: "ipamhandles.crd.projectcalico.org",
			},
			Spec: crdv1beta1.CustomResourceDefinitionSpec{
				Scope:   crdv1beta1.ClusterScoped,
				Group:   calicoCRDGroup,
				Version: calicoCRDVersion,
				Names: crdv1beta1.CustomResourceDefinitionNames{
					Plural:     "ipamhandles",
					Singular:   "ipamhandle",
					ShortNames: nil,
					Kind:       "IPAMHandle",
					ListKind:   "",
					Categories: nil,
				},
			},
			Status: crdv1beta1.CustomResourceDefinitionStatus{},
		},

		crdv1beta1.CustomResourceDefinition{
			TypeMeta: metav1.TypeMeta{
				Kind:       crdKind,
				APIVersion: crdApiVersion,
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: "ipamconfigs.crd.projectcalico.org",
			},
			Spec: crdv1beta1.CustomResourceDefinitionSpec{
				Scope:   crdv1beta1.ClusterScoped,
				Group:   calicoCRDGroup,
				Version: calicoCRDVersion,
				Names: crdv1beta1.CustomResourceDefinitionNames{
					Plural:     "ipamconfigs",
					Singular:   "ipamconfig",
					ShortNames: nil,
					Kind:       "IPAMConfig",
					ListKind:   "",
					Categories: nil,
				},
			},
			Status: crdv1beta1.CustomResourceDefinitionStatus{},
		},

		crdv1beta1.CustomResourceDefinition{
			TypeMeta: metav1.TypeMeta{
				Kind:       crdKind,
				APIVersion: crdApiVersion,
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: "bgppeers.crd.projectcalico.org",
			},
			Spec: crdv1beta1.CustomResourceDefinitionSpec{
				Scope:   crdv1beta1.ClusterScoped,
				Group:   calicoCRDGroup,
				Version: calicoCRDVersion,
				Names: crdv1beta1.CustomResourceDefinitionNames{
					Plural:     "bgppeers",
					Singular:   "bgppeer",
					ShortNames: nil,
					Kind:       "BGPPeer",
					ListKind:   "",
					Categories: nil,
				},
			},
			Status: crdv1beta1.CustomResourceDefinitionStatus{},
		},

		crdv1beta1.CustomResourceDefinition{
			TypeMeta: metav1.TypeMeta{
				Kind:       crdKind,
				APIVersion: "apiextensions.k8s.io/v1beta1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: "felixconfigurations.crd.projectcalico.org",
			},
			Spec: crdv1beta1.CustomResourceDefinitionSpec{
				Group:   calicoCRDGroup,
				Version: calicoCRDVersion,
				Names: crdv1beta1.CustomResourceDefinitionNames{
					Plural:   "felixconfigurations",
					Singular: "felixconfiguration",
					Kind:     "FelixConfiguration",
				},
				Scope: crdv1beta1.ClusterScoped,
			},
			Status: crdv1beta1.CustomResourceDefinitionStatus{},
		},

		crdv1beta1.CustomResourceDefinition{
			TypeMeta: metav1.TypeMeta{
				Kind:       crdKind,
				APIVersion: crdApiVersion,
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: "ipamblocks.crd.projectcalico.org",
			},
			Spec: crdv1beta1.CustomResourceDefinitionSpec{
				Group:   calicoCRDGroup,
				Version: calicoCRDVersion,
				Names: crdv1beta1.CustomResourceDefinitionNames{
					Plural:   "ipamblocks",
					Singular: "ipamblock",
					Kind:     "IPAMBlock",
				},
				Scope: crdv1beta1.ClusterScoped,
			},
			Status: crdv1beta1.CustomResourceDefinitionStatus{},
		},
	}

}

func main() {
	homeDir := os.Getenv("HOME")
	kubeConfPath := fmt.Sprintf("%s/.kube/kube-dev-cluster", homeDir)
	restConfig, err := clientcmd.BuildConfigFromFlags("", kubeConfPath)
	if err != nil {
		log.Fatal(err)
	}
	client, err := clientset.NewForConfig(restConfig)
	if err != nil {
		log.Fatal(err)
	}

	felixConf := calicoFelixConfCRD()
	createdCRD, err := client.ApiextensionsV1beta1().CustomResourceDefinitions().Create(felixConf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", createdCRD)

}
