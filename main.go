package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/bartvanbenthem/k8s-onerepo/utils/filesystem"
	"github.com/bartvanbenthem/k8s-onerepo/utils/manifestgen"
)

// Declare all project path variables
const clValues string = "var/cluster/values"
const clTemplates string = "var/cluster/templates"
const clConfig string = "config/cluster"
const clHelmTemplates string = "var/helmcharts"
const clHelmConfig string = "config/helmcharts"

func main() {
	// create base dirs in config
	fmt.Printf("Check config base dirs...\n")
	CreateConfigBaseDirs(clConfig)

	// copy helmcharts to config dir
	fmt.Printf("\nCopy helmcharts to config...\n")
	CopyHelmTemplatesToConfig(clHelmTemplates, clHelmConfig)
	cfg, _ := filesystem.ReadFiles(clHelmConfig)
	for _, f := range cfg {
		fmt.Printf("%v\n", f.PathAndFile)
	}

	// start manigest generation for cluster-specific
	fmt.Printf("\nGenerate Cluster manifests...\n")
	GenerateClusterManifests(clValues, clTemplates, clConfig)
	cfg, _ = filesystem.ReadFiles(clConfig)
	for _, f := range cfg {
		fmt.Printf("%v\n", f.PathAndFile)
	}
}

func CreateConfigBaseDirs(configDir string) {
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		err := os.Mkdir(configDir, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func CopyHelmTemplatesToConfig(sourceDir, destinationDir string) {
	err := filesystem.CopyDir(sourceDir, destinationDir)
	if err != nil {
		log.Fatal(err)
	}
}

func GenerateClusterManifests(valuesPath, templatePath, outputFolder string) {
	var cmg manifestgen.ManifestGenClient

	outputPathCluster := fmt.Sprintf("%v", outputFolder)
	if _, err := os.Stat(outputPathCluster); os.IsNotExist(err) {
		err := os.Mkdir(outputPathCluster, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}

	valueFiles, err := filesystem.ReadFiles(valuesPath)
	if err != nil {
		log.Fatal(err)
	}

	templateFiles, err := filesystem.ReadFiles(templatePath)
	if err != nil {
		log.Fatal(err)
	}

	for _, val := range valueFiles {
		valfilename := strings.Split(val.FileName, "-")
		v := fmt.Sprintf("%v.yaml", valfilename[0])
		for _, tmpl := range templateFiles {
			tmplfilename := strings.Split(tmpl.FileName, "-")
			t := fmt.Sprintf("%v.yaml", tmplfilename[0])
			if t == v || tmpl.FileName == val.FileName {
				cmg.GenerateManifestFromValues(val.PathAndFile,
					tmpl.PathAndFile, outputPathCluster)
			}
		}
	}
}
