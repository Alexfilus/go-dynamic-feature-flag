package main

import (
	"flag"
	"log"
	"os"

	"go-dynamic-feature-flag/internal/gen"
	"go-dynamic-feature-flag/internal/parser"
)

var cfgPath = flag.String("cfg_path", "", "path to config yaml file")

func main() {
	flag.Parse()
	if *cfgPath == "" {
		log.Fatalln("cfg_path is required")
	}

	yamlContent, err := os.ReadFile(*cfgPath)
	if err != nil {
		log.Fatalln(err)
	}

	cfg, err := parser.Parse(yamlContent)
	if err != nil {
		log.Fatalln(err)
	}

	if err := gen.Gen(cfg); err != nil {
		log.Fatalln(err)
	}
}
