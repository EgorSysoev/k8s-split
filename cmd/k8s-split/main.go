package main

import (
	"flag"
	"log"

	split "github.com/leominov/k8s-split"
)

var (
	specsFile = flag.String("f", "", "Path to file with Kubernetes specification")
	outputDir = flag.String("o", "", "Path to output directory")
)

func main() {
	flag.Parse()
	err := split.Process(*specsFile, *outputDir)
	if err != nil {
		log.Fatal(err)
	}
}
