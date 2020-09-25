package main

import (
	"github.com/bartvanbenthem/k8s-manifestgen/client"
)

func main() {
	GenerateClusterAllManifest("clusters/cluster-all/values/allclusters.yaml",
		"clusters/cluster-all/templates/allclusters.yaml",
		"clusters/cluster-all/manifest")
}

func GenerateClusterSpecificManifest() {}
func GenerateClusterHelmManifest()     {}

func GenerateClusterAllManifest(valuesFile, templateFile, outputFolder string) {
	//build a foreach loop to check all value and template files within their specific dir
	client.GenerateManifestFromValues(valuesFile, templateFile, outputFolder)
}
