package parser

import (
	"gopkg.in/yaml.v3"

	"go-dynamic-feature-flag/internal/model"
)

func Parse(yamlContent []byte) (model.Config, error) {
	var cfg model.Config
	err := yaml.Unmarshal(yamlContent, &cfg)

	return cfg, err
}
