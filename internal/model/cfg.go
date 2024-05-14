package model

import (
	"time"

	"github.com/iancoleman/strcase"
)

type Config struct {
	ProjectName  string                   `yaml:"project_name"`
	PkgPath      string                   `yaml:"pkg_path"`
	PkgName      string                   `yaml:"pkg_name"`
	OutputFile   string                   `yaml:"output_file"`
	StringVars   map[string]string        `yaml:"string_vars"`
	DurationVars map[string]time.Duration `yaml:"duration_vars"`
	IntVars      map[string]int           `yaml:"int_vars"`
	BoolVars     map[string]bool          `yaml:"bool_vars"`
}

func (c Config) ToRedisKey(key string) string {
	return c.ProjectName + ":" + c.PkgName + ":" + key
}

func (c Config) ToConstStr(key string) string {
	return "strKey" + strcase.ToCamel(key)
}

func (c Config) ToConstDuration(key string) string {
	return "durationKey" + strcase.ToCamel(key)
}

func (c Config) ToConstInt(key string) string {
	return "intKey" + strcase.ToCamel(key)
}

func (c Config) ToConstBool(key string) string {
	return "boolKey" + strcase.ToCamel(key)
}
