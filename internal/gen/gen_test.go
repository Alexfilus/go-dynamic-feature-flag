package gen

import (
	"testing"
	"time"

	"github.com/alexfilus/go-dynamic-feature-flag/internal/model"
)

func TestGen(t *testing.T) {
	type args struct {
		cfg model.Config
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Testcase 1",
			args: args{
				cfg: model.Config{
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
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Gen(tt.args.cfg); (err != nil) != tt.wantErr {
				t.Errorf("Gen() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
