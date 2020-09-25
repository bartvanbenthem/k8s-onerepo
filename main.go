package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/bartvanbenthem/k8s-manifestgen/client"
)

func main() {
	GenerateClusterAllManifest("clusters/cluster-all/values",
		"clusters/cluster-all/templates",
		"clusters/cluster-all/manifest")
}

func GenerateClusterSpecificManifest() {}
func GenerateClusterHelmManifest()     {}

func GenerateClusterAllManifest(valuesPath, templatePath, outputFolder string) {

	valueFiles := ReadFileNames(valuesPath)
	templateFiles := ReadFileNames(templatePath)

	for _, val := range valueFiles {
		for _, tmpl := range templateFiles {
			if tmpl.fileName == val.fileName {
				client.GenerateManifestFromValues(val.pathAndFile, tmpl.pathAndFile, outputFolder)
			}
		}
	}
}

type FileInfo struct {
	fileName    string
	filePath    string
	pathAndFile string
}

func ReadFileNames(path string) []FileInfo {
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
