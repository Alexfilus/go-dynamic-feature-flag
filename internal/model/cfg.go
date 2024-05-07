package model

import "time"

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
