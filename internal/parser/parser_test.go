package parser

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/alexfilus/go-dynamic-feature-flag/internal/model"
)

func Test_Parse(t *testing.T) {
	expected := model.Config{
		ProjectName: "test-project",
		PkgPath:     "internal/test",
		PkgName:     "test",
		OutputFile:  "test_cfg.go",
		StringVars: map[string]string{
			"test_str1": "test1",
			"test_str2": "test2",
		},
		DurationVars: map[string]time.Duration{
			"test_duration1": 1 * time.Second,
			"test_duration2": 2 * time.Minute,
		},
		IntVars: map[string]int{
			"test_int1": 1,
			"test_int2": 2,
		},
		BoolVars: map[string]bool{
			"test_bool1": true,
			"test_bool2": false,
		},
	}
	yamlContent, err := os.ReadFile("example.yaml")
	require.NoError(t, err)

	cfg, err := Parse(yamlContent)
	require.NoError(t, err)

	require.Equal(t, expected, cfg)
}
