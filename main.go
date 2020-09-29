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
	cfg := ReadFiles("config")
	for _, f := range cfg {
		fmt.Printf("%v\n", f.pathAndFile)
	}

	// start manigest generation for cluster-all
	fmt.Printf("Generate Cluster-all manifests...\n")
	GenerateClusterAllManifest(clAllValues, clAllTemplates, clAllConfig)
	cfgAll := ReadFiles(clAllConfig)
	for _, f := range cfgAll {
		fmt.Printf("%v\n", f.pathAndFile)
	}

	// start manigest generation for cluster-specific
	fmt.Printf("Generate Cluster-specific manifests...\n")
	GenerateClusterManifests(clSpecValues, clSpecTemplates, clSpecConfig)
	cfgSpec := ReadFiles(clSpecConfig)
	for _, f := range cfgSpec {
		fmt.Printf("%v\n", f.pathAndFile)
	}
}

func CopyHelmTemplatesToConfig() {}
func GenerateHelmValueFile()     {}

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

func CreateConfigBaseDirs(varDir, configDir string) {
	var folders []string
	valueDir := ReadFiles(varDir)
	for _, folder := range valueDir {
		folders = append(folders, folder.fileName)
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
