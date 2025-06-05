package gen

import (
	"os"
	"path/filepath"
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
		{
			name: "Testcase 2",
			args: args{
				cfg: model.Config{
					ProjectName: "mega-pay-master",
					PkgPath:     "internal/api/rest",
					PkgName:     "rest",
					OutputFile:  "server_cfg.go",
					BoolVars: map[string]bool{
						"wbxCheckRids":   true,
						"forceDiscount":  false,
						"enableDiscount": true,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Testcase 3",
			args: args{
				cfg: model.Config{
					ProjectName: "mega-pay-master",
					PkgPath:     "internal/service/position",
					PkgName:     "position",
					OutputFile:  "service_cfg.go",
					BoolVars: map[string]bool{
						"forceSkipCB": false,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Testcase 4",
			args: args{
				cfg: model.Config{
					ProjectName: "mega-pay-master",
					PkgPath:     "internal/featurebalancer",
					PkgName:     "featurebalancer",
					OutputFile:  "balancer_cfg.go",
					IntVars: map[string]int{
						"sbpWalletPercent": 0,
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
				return
			}

			outputPath := filepath.Join(tt.args.cfg.PkgPath, tt.args.cfg.OutputFile)
			if _, err := os.Stat(outputPath); err != nil {
				t.Fatalf("generated file %s missing: %v", outputPath, err)
			}

			// cleanup generated file and directory if possible
			_ = os.Remove(outputPath)
			_ = os.Remove(tt.args.cfg.PkgPath)
		})
	}
}
