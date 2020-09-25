package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/bartvanbenthem/k8s-manifestgen/client"
)

func main() {
	GenerateClusterAllManifest("values/cluster-all/values",
		"values/cluster-all/templates", "config/cluster-all")
}

func GenerateClusterHelmManifest() {}

// needs to read all folders in valuesPath and create these folders in config
func GenerateClusterSpecificManifest() {}

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
