package gen

import (
	"log"
	"os"
	"strconv"
	"strings"
	
	"github.com/iancoleman/strcase"
	
	"go-dynamic-feature-flag/internal/model"
)

func Gen(cfg model.Config) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	outputPath := dir + "/../../" + cfg.PkgPath + "/" + cfg.OutputFile
	f, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer func() {
		err := f.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}()
	
	b := strings.Builder{}
	b.WriteString("package " + cfg.PkgName + "\n\n")
	b.WriteString("import (\n")
	b.WriteString("\t\"context\"\n")
	if len(cfg.DurationVars) > 0 {
		b.WriteString("\t\"time\"\n")
	}
	b.WriteString("\n\t\"github.com/redis/rueidis\"\n")
	b.WriteString(")\n\n")
	
	b.WriteString("type DynamicConfig struct {\n")
	
	b.WriteString("\tclient rueidis.Client\n")
	
	for k := range cfg.StringVars {
		b.WriteString("\t" + strcase.ToLowerCamel(k) + " string\n")
	}
	
	for k := range cfg.DurationVars {
		b.WriteString("\t" + strcase.ToLowerCamel(k) + " time.Duration\n")
	}
	
	for k := range cfg.IntVars {
		b.WriteString("\t" + strcase.ToLowerCamel(k) + " int\n")
	}
	
	for k := range cfg.BoolVars {
		b.WriteString("\t" + strcase.ToLowerCamel(k) + " bool\n")
	}
	
	b.WriteString("}\n\n")
	
	b.WriteString("func NewDynamicConfig(client rueidis.Client) *DynamicConfig {\n\treturn &DynamicConfig{\n")
	b.WriteString("\t\tclient: client,\n")
	for k, v := range cfg.StringVars {
		b.WriteString("\t\t" + strcase.ToLowerCamel(k) + ": \"" + v + "\",\n")
	}
	
	for k, v := range cfg.DurationVars {
		b.WriteString("\t\t" + strcase.ToLowerCamel(k) + ": " + strconv.FormatInt(v.Milliseconds(), 10) + " * time.Millisecond,\n")
	}
	
	for k, v := range cfg.IntVars {
		b.WriteString("\t\t" + strcase.ToLowerCamel(k) + ": " + strconv.Itoa(v) + ",\n")
	}
	
	for k, v := range cfg.BoolVars {
		b.WriteString("\t\t" + strcase.ToLowerCamel(k) + ": " + strconv.FormatBool(v) + ",\n")
	}
	
	b.WriteString("\t}\n}\n")
	
	for k := range cfg.StringVars {
		b.WriteString("\nfunc (c *DynamicConfig) " + strcase.ToCamel(k) + "(ctx context.Context) string {\n\t")
		key := cfg.ProjectName + ":" + cfg.PkgName + ":" + strcase.ToLowerCamel(k)
		b.WriteString("resp, err := c.client.DoCache(\n\t\tctx,\n\t\tc.client.B().Get().Key(\"" + key + "\").Cache(),\n\t\ttime.Minute,\n\t).ToString()\n\n")
		b.WriteString("\tif err != nil {\n\t\treturn c." + strcase.ToLowerCamel(k) + "\n\t}\n\n")
		b.WriteString("\tc." + strcase.ToLowerCamel(k) + " = resp\n")
		b.WriteString("\treturn resp\n}\n")
	}
	
	for k := range cfg.DurationVars {
		b.WriteString("\nfunc (c *DynamicConfig) " + strcase.ToCamel(k) + "(ctx context.Context) time.Duration {\n\t")
		key := cfg.ProjectName + ":" + cfg.PkgName + ":" + strcase.ToLowerCamel(k)
		b.WriteString("resp, err := c.client.DoCache(\n\t\tctx,\n\t\tc.client.B().Get().Key(\"" + key + "\").Cache(),\n\t\ttime.Minute,\n\t).ToString()\n\n")
		b.WriteString("\tif err != nil {\n\t\treturn c." + strcase.ToLowerCamel(k) + "\n\t}\n\n")
		b.WriteString("\tc." + strcase.ToLowerCamel(k) + ", _ = time.ParseDuration(resp)\n")
		b.WriteString("\treturn c." + strcase.ToLowerCamel(k) + "\n}\n")
	}
	
	for k := range cfg.IntVars {
		b.WriteString("\nfunc (c *DynamicConfig) " + strcase.ToCamel(k) + "(ctx context.Context) int {\n\t")
		key := cfg.ProjectName + ":" + cfg.PkgName + ":" + strcase.ToLowerCamel(k)
		b.WriteString("resp, err := c.client.DoCache(\n\t\tctx,\n\t\tc.client.B().Get().Key(\"" + key + "\").Cache(),\n\t\ttime.Minute,\n\t).ToInt64()\n\n")
		b.WriteString("\tif err != nil {\n\t\treturn c." + strcase.ToLowerCamel(k) + "\n\t}\n\n")
		b.WriteString("\tc." + strcase.ToLowerCamel(k) + " = int(resp)\n")
		b.WriteString("\treturn c." + strcase.ToLowerCamel(k) + "\n}\n")
	}
	
	for k := range cfg.BoolVars {
		b.WriteString("\nfunc (c *DynamicConfig) " + strcase.ToCamel(k) + "(ctx context.Context) bool {\n\t")
		key := cfg.ProjectName + ":" + cfg.PkgName + ":" + strcase.ToLowerCamel(k)
		b.WriteString("resp, err := c.client.DoCache(\n\t\tctx,\n\t\tc.client.B().Get().Key(\"" + key + "\").Cache(),\n\t\ttime.Minute,\n\t).ToBool()\n\n")
		b.WriteString("\tif err != nil {\n\t\treturn c." + strcase.ToLowerCamel(k) + "\n\t}\n\n")
		b.WriteString("\tc." + strcase.ToLowerCamel(k) + " = resp\n")
		b.WriteString("\treturn c." + strcase.ToLowerCamel(k) + "\n}\n")
	}
	
	_, err = f.WriteString(b.String())
	
	return err
}
