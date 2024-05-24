package parser

import (
	"log"

	"gopkg.in/yaml.v3"

	"github.com/alexfilus/go-dynamic-feature-flag/internal/model"
)

func Parse(yamlContent []byte) (model.Config, error) {
	var cfg model.Config
	err := yaml.Unmarshal(yamlContent, &cfg)
	log.Println(cfg)

	return cfg, err
}
