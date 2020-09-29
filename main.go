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
const clAllValues string = "var/cluster-all/values"
const clAllTemplates string = "var/cluster-all/templates"
const clAllConfig string = "config/cluster-all"
const clSpecValues string = "var/cluster-specific/values"
const clSpecTemplates string = "var/cluster-specific/templates"
const clSpecConfig string = "config/cluster-specific"
const clHelmValues string = "var/helmcharts/values"
const clHelmTemplates string = "var/helmcharts/templates"
const clHelmConfig string = "config/helmcharts"

func main() {
	// create base dirs in config
	fmt.Printf("Check config base dirs...\n")
	CreateConfigBaseDirs("var", "config")
	cfg, _ := filesystem.ReadFiles("config")
	for _, f := range cfg {
		fmt.Printf("%v\n", f.PathAndFile)
	}

	// start manigest generation for cluster-all
	fmt.Printf("Generate Cluster-all manifests...\n")
	GenerateClusterAllManifest(clAllValues, clAllTemplates, clAllConfig)
	cfgAll, _ := filesystem.ReadFiles(clAllConfig)
	for _, f := range cfgAll {
		fmt.Printf("%v\n", f.PathAndFile)
	}

	// start manigest generation for cluster-specific
	fmt.Printf("Generate Cluster-specific manifests...\n")
	GenerateClusterManifests(clSpecValues, clSpecTemplates, clSpecConfig)
	cfgSpec, _ := filesystem.ReadFiles(clSpecConfig)
	for _, f := range cfgSpec {
		fmt.Printf("%v\n", f.PathAndFile)
	}

	fmt.Printf("Copy helmcharts to config...\n")
	CopyHelmTemplatesToConfig(clHelmTemplates, clHelmConfig)
	fmt.Println("finished")

}

func CopyHelmTemplatesToConfig(sourceDir, destinationDir string) {
	err := filesystem.CopyDir(sourceDir, destinationDir)
	if err != nil {
		log.Fatal(err)
	}
}

func GenerateHelmValueFile() {}

func GenerateClusterAllManifest(valuesPath, templatePath, outputFolder string) {

	var cmg manifestgen.ManifestGenClient
	valueFiles, err := filesystem.ReadFiles(valuesPath)
	if err != nil {
		log.Fatal(err)
	}

	templateFiles, err := filesystem.ReadFiles(templatePath)
	if err != nil {
		log.Fatal(err)
	}

	for _, val := range valueFiles {
		for _, tmpl := range templateFiles {
			if tmpl.FileName == val.FileName {
				cmg.GenerateManifestFromValues(val.PathAndFile,
					tmpl.PathAndFile, outputFolder)
			}
		}
	}
}

func GenerateClusterManifests(valuesPath, templatePath, outputFolder string) {

	var cmg manifestgen.ManifestGenClient
	var folders []string

	valueDir, err := filesystem.ReadFiles(valuesPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, folder := range valueDir {
		folders = append(folders, folder.FileName)
	}

	for _, folder := range folders {
		outputPathCluster := fmt.Sprintf("%v/%v", outputFolder, folder)
		if _, err := os.Stat(outputPathCluster); os.IsNotExist(err) {
			err := os.Mkdir(outputPathCluster, 0755)
			if err != nil {
				log.Fatal(err)
			}
		}
		// combine value files and template files
		valueFiles, err := filesystem.ReadFiles(fmt.Sprintf("%v/%v", valuesPath, folder))
		if err != nil {
			log.Fatal(err)
		}

		templateFiles, err := filesystem.ReadFiles(fmt.Sprintf("%v", templatePath))
		if err != nil {
			log.Fatal(err)
		}

		for _, val := range valueFiles {
			filename := strings.Split(val.FileName, "-")
			for _, tmpl := range templateFiles {
				if tmpl.FileName == fmt.Sprintf("%v.yaml", filename[0]) ||
					tmpl.FileName == val.FileName {
					cmg.GenerateManifestFromValues(val.PathAndFile,
						tmpl.PathAndFile, outputPathCluster)
				}
			}
		}
	}
}

func CreateConfigBaseDirs(varDir, configDir string) {
	var folders []string

	valueDir, err := filesystem.ReadFiles(varDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, folder := range valueDir {
		folders = append(folders, folder.FileName)
	}

	for _, folder := range folders {
		outputPathCluster := fmt.Sprintf("%v/%v", configDir, folder)
		if _, err := os.Stat(outputPathCluster); os.IsNotExist(err) {
			err := os.Mkdir(outputPathCluster, 0755)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
