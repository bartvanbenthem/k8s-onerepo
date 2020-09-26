package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/bartvanbenthem/k8s-manifestgen/client"
)

// Declare all project path variables
const clusterAllValues string = "values/cluster-all/values"
const clusterAllTemplates string = "values/cluster-all/templates"
const clusterAllConfig string = "config/cluster-all"
const clusterSpecValues string = ""
const clusterSpecTemplates string = ""
const clusterSpecConfig string = ""
const clusterHelmValues string = ""
const clusterHelmTemplates string = ""
const clusterHelmConfig string = ""

func main() {
	// start manigest generation for cluster-all
	fmt.Printf("Generate Cluster-all manifests...\n")
	GenerateClusterAllManifest(clusterAllValues, clusterAllTemplates, clusterAllConfig)
	cfgAll := ReadFiles(clusterAllConfig)
	for _, f := range cfgAll {
		fmt.Printf("%v\n", f.pathAndFile)
	}
}

func GenerateClusterHelmManifests() {}

// needs to read all folders in valuesPath and create these folders in config
func GenerateClusterSpecificManifests() {}

func GenerateClusterAllManifest(valuesPath, templatePath, outputFolder string) {
	valueFiles := ReadFiles(valuesPath)
	templateFiles := ReadFiles(templatePath)
	for _, val := range valueFiles {
		for _, tmpl := range templateFiles {
			if tmpl.fileName == val.fileName {
				client.GenerateManifestFromValues(val.pathAndFile,
					tmpl.pathAndFile, outputFolder)
			}
		}
	}
}

type FileInfo struct {
	fileName    string
	filePath    string
	pathAndFile string
}

func ReadFiles(path string) []FileInfo {
	var file FileInfo
	var files []FileInfo

	fs, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range fs {
		file.fileName = f.Name()
		file.filePath = path
		file.pathAndFile = fmt.Sprintf("%v/%v", path, f.Name())
		files = append(files, file)
	}

	return files
}
