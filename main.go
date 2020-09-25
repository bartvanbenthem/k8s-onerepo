package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	ac := GenerateClusterAllManifest("clusters/cluster-all/values/allclusters.yaml",
		"clusters/cluster-all/templates/allclusters.yaml",
		"clusters/cluster-all/manifest")
	fmt.Println(ac)
}

func GenerateAllTeamsManifest() {}
func GenerateAllHelmManifest()  {}

func GenerateClusterAllManifest(valuesFile, templateFile, outputFolder string) string {
	cmd := exec.Command("./k8s-manifestgen", valuesFile, templateFile, outputFolder)
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	return strings.TrimSuffix(string(stdoutStderr), "\n")
}
