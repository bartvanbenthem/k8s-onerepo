package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	ac := GenerateAllClustersManifest()
	fmt.Println(ac)
}

func GenerateAllTeamsManifest() {}
func GenerateAllHelmManifest()  {}

func GenerateAllClustersManifest() string {
	cmd := exec.Command("./k8s-manifestgen",
		"clusters/cluster-all/values/allclusters.yaml",
		"clusters/cluster-all/templates/allclusters.yaml",
		"clusters/cluster-all/manifest")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	return strings.TrimSuffix(string(stdoutStderr), "\n")
}
