package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/bartvanbenthem/k8s-onerepo/utils/manifestgen"
)

// Declare all project path variables
const clusterAllValues string = "var/cluster-all/values"
const clusterAllTemplates string = "var/cluster-all/templates"
const clusterAllConfig string = "config/cluster-all"
const clusterSpecValues string = "var/cluster-specific/values"
const clusterSpecTemplates string = "var/cluster-specific/templates"
const clusterSpecConfig string = "config/cluster-specific"
const clusterHelmValues string = "var/helmcharts/values"
const clusterHelmTemplates string = "var/helmcharts/templates"
const clusterHelmConfig string = "config/helmcharts"

func main() {
	// start manigest generation for cluster-all
	fmt.Printf("Generate Cluster-all manifests...\n")
	GenerateClusterAllManifest(clusterAllValues, clusterAllTemplates, clusterAllConfig)
	cfgAll := ReadFiles(clusterAllConfig)
	for _, f := range cfgAll {
		fmt.Printf("%v\n", f.pathAndFile)
	}

	// start manigest generation for cluster-specific
	fmt.Printf("Generate Cluster-specific manifests...\n")
	GenerateClusterManifests(clusterSpecValues, clusterSpecTemplates, clusterSpecConfig)
	cfgSpec := ReadFiles(clusterSpecConfig)
	for _, f := range cfgSpec {
		fmt.Printf("%v\n", f.pathAndFile)
	}
}

func GenerateClusterHelmManifests() {}

func GenerateClusterAllManifest(valuesPath, templatePath, outputFolder string) {
	var cmg manifestgen.ManifestGenClient
	valueFiles := ReadFiles(valuesPath)
	templateFiles := ReadFiles(templatePath)
	for _, val := range valueFiles {
		for _, tmpl := range templateFiles {
			if tmpl.fileName == val.fileName {
				cmg.GenerateManifestFromValues(val.pathAndFile,
					tmpl.pathAndFile, outputFolder)
			}
		}
	}
}

func GenerateClusterManifests(valuesPath, templatePath, outputFolder string) {
	var cmg manifestgen.ManifestGenClient
	var folders []string
	valueDir := ReadFiles(valuesPath)
	for _, folder := range valueDir {
		folders = append(folders, folder.fileName)
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
		valueFiles := ReadFiles(fmt.Sprintf("%v/%v", valuesPath, folder))
		templateFiles := ReadFiles(fmt.Sprintf("%v", templatePath))
		for _, val := range valueFiles {
			filename := strings.Split(val.fileName, "-")
			for _, tmpl := range templateFiles {
				if tmpl.fileName == fmt.Sprintf("%v.yaml", filename[0]) ||
					tmpl.fileName == val.fileName {
					cmg.GenerateManifestFromValues(val.pathAndFile,
						tmpl.pathAndFile, outputPathCluster)
				}
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
