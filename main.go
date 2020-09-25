package main

import (
	"fmt"

	"github.com/bartvanbenthem/k8s-manifestgen/client"
)

func main() {
	ac := GenerateClusterAllManifest("clusters/cluster-all/values/allclusters.yaml",
		"clusters/cluster-all/templates/allclusters.yaml",
		"clusters/cluster-all/manifest")
	fmt.Println(ac)
}

func GenerateClusterSpecificManifest() {}
func GenerateClusterHelmManifest()     {}

func GenerateClusterAllManifest(valuesFile, templateFile, outputFolder string) string {
	var cmg client.ManifestGenClient
	//build a foreach loop to check all value and template files within their specific dir
	cmg.GenerateManifestFromValues(valuesFile, templateFile, outputFolder)
}
