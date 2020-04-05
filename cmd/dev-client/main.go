package main

import (
	"fmt"
	calicocrd "github.com/kube-operators/calico/pkg/crd"
	"k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/client-go/tools/clientcmd"

	"log"
	"os"
)

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

	for _, crd := range getCalicoCRDs() {
		fmt.Printf("Create CRD: %s\n", crd.Name)
		createdCRD, err := client.ApiextensionsV1beta1().CustomResourceDefinitions().Create(&crd)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("CRD: %s has been created.\n", createdCRD.Name)
	}
}
