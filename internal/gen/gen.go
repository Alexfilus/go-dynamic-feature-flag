package gen

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/iancoleman/strcase"

	"github.com/alexfilus/go-dynamic-feature-flag/internal/model"
)

func Gen(cfg model.Config) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	outputDir := dir + "/" + cfg.PkgPath
	outputPath := outputDir + "/" + cfg.OutputFile
	err = os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		return err
	}
	f, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	log.Println("Generated file: " + outputPath)
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
	if len(cfg.IntVars) > 0 || len(cfg.BoolVars) > 0 {
		b.WriteString("\t\"strconv\"\n")
	}
	b.WriteString("\t\"time\"\n")
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

	b.WriteString("type RequestDynamicConfigUpdate struct {\n")

	for k := range cfg.StringVars {
		b.WriteString("\t" + strcase.ToCamel(k) + " *string `json:\"" + k + ",omitempty\"`\n")
	}

	for k := range cfg.DurationVars {
		b.WriteString("\t" + strcase.ToCamel(k) + " *time.Duration `json:\"" + k + ",omitempty\"`\n")
	}

	for k := range cfg.IntVars {
		b.WriteString("\t" + strcase.ToCamel(k) + " *int `json:\"" + k + ",omitempty\"`\n")
	}

	for k := range cfg.BoolVars {
		b.WriteString("\t" + strcase.ToCamel(k) + " *bool `json:\"" + k + ",omitempty\"`\n")
	}

	b.WriteString("}\n\n")

	b.WriteString("const (\n")
	for k := range cfg.StringVars {
		b.WriteString("\t")
		b.WriteString(cfg.ToConstStr(k))
		b.WriteString(" = \"")
		b.WriteString(cfg.ToRedisKey(k))
		b.WriteString("\"\n")
	}

	for k := range cfg.DurationVars {
		b.WriteString("\t")
		b.WriteString(cfg.ToConstDuration(k))
		b.WriteString(" = \"")
		b.WriteString(cfg.ToRedisKey(k))
		b.WriteString("\"\n")
	}

	for k := range cfg.IntVars {
		b.WriteString("\t")
		b.WriteString(cfg.ToConstInt(k))
		b.WriteString(" = \"")
		b.WriteString(cfg.ToRedisKey(k))
		b.WriteString("\"\n")
	}

	for k := range cfg.BoolVars {
		b.WriteString("\t")
		b.WriteString(cfg.ToConstBool(k))
		b.WriteString(" = \"")
		b.WriteString(cfg.ToRedisKey(k))
		b.WriteString("\"\n")
	}
	b.WriteString(")\n\n")

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
		key := cfg.ToConstStr(k)
		fieldName := strcase.ToLowerCamel(k)
		title := strcase.ToCamel(k)
		b.WriteString("\nfunc (c *DynamicConfig) " + title + "(ctx context.Context) string {\n\t")
		b.WriteString("if c.client == nil {\n\t\treturn c." + fieldName + "\n\t}\n\n")
		b.WriteString("resp, err := c.client.DoCache(\n\t\tctx,\n\t\tc.client.B().Get().Key(" + key + ").Cache(),\n\t\ttime.Minute,\n\t).ToString()\n\n")
		b.WriteString("\tif err != nil {\n\t\treturn c." + fieldName + "\n\t}\n\n")
		b.WriteString("\tc." + fieldName + " = resp\n")
		b.WriteString("\treturn c." + fieldName + "\n}\n")

		b.WriteString("\nfunc (c *DynamicConfig) Set" + title + "(value string) *DynamicConfig {\n\t")
		b.WriteString("c." + fieldName + " = value\n\treturn c\n}\n")

		b.WriteString("\nfunc (c *DynamicConfig) Store" + title + "(ctx context.Context, value string) error {\n\t")
		b.WriteString("return c.client.Do(\n\t\tctx, \n\t\tc.client.B().Set().Key(" + key + ").\n\t\t\tValue(value).\n\t\t\tBuild()).\n\t\tError()\n}\n")
	}

	for k := range cfg.DurationVars {
		key := cfg.ToConstDuration(k)
		fieldName := strcase.ToLowerCamel(k)
		title := strcase.ToCamel(k)
		b.WriteString("\nfunc (c *DynamicConfig) " + title + "(ctx context.Context) time.Duration {\n\t")
		b.WriteString("if c.client == nil {\n\t\treturn c." + fieldName + "\n\t}\n\n")
		b.WriteString("resp, err := c.client.DoCache(\n\t\tctx,\n\t\tc.client.B().Get().Key(" + key + ").Cache(),\n\t\ttime.Minute,\n\t).AsInt64()\n\n")
		b.WriteString("\tif err != nil {\n\t\treturn c." + fieldName + "\n\t}\n\n")
		b.WriteString("\tc." + fieldName + " = time.Duration(resp) * time.Millisecond\n")
		b.WriteString("\treturn c." + fieldName + "\n}\n")

		b.WriteString("\nfunc (c *DynamicConfig) Set" + title + "(value time.Duration) *DynamicConfig {\n\t")
		b.WriteString("c." + fieldName + " = value\n\treturn c\n}\n")

		b.WriteString("\nfunc (c *DynamicConfig) Store" + title + "(ctx context.Context, value time.Duration) error {\n\t")
		b.WriteString("return c.client.Do(\n\t\tctx, \n\t\tc.client.B().Set().Key(" + key + ").\n\t\t\tValue(strconv.FormatInt(value.Milliseconds(), 10)).\n\t\t\tBuild()).\n\t\t\tError()\n}\n")
	}

	for k := range cfg.IntVars {
		key := cfg.ToConstInt(k)
		fieldName := strcase.ToLowerCamel(k)
		title := strcase.ToCamel(k)

		b.WriteString("\nfunc (c *DynamicConfig) " + title + "(ctx context.Context) int {\n\t")
		b.WriteString("if c.client == nil {\n\t\treturn c." + fieldName + "\n\t}\n\n")
		b.WriteString("resp, err := c.client.DoCache(\n\t\tctx,\n\t\tc.client.B().Get().Key(" + key + ").Cache(),\n\t\ttime.Minute,\n\t).ToString()\n\n")
		b.WriteString("\tif err != nil {\n\t\treturn c." + fieldName + "\n\t}\n\n")
		b.WriteString("\trespInt, err := strconv.Atoi(resp)\n")
		b.WriteString("\tif err != nil {\n\t\treturn c." + fieldName + "\n\t}\n\n")
		b.WriteString("\tc." + fieldName + " = respInt\n")
		b.WriteString("\treturn c." + fieldName + "\n}\n")

		b.WriteString("\nfunc (c *DynamicConfig) Set" + title + "(value int) *DynamicConfig {\n\t")
		b.WriteString("c." + fieldName + " = value\n\treturn c\n}\n")

		b.WriteString("\nfunc (c *DynamicConfig) Store" + title + "(ctx context.Context, value int) error {\n\t")
		b.WriteString("return c.client.Do(\n\t\tctx, \n\t\tc.client.B().Set().Key(" + key + ").\n\t\t\tValue(strconv.Itoa(value)).\n\t\t\tBuild()).\n\t\tError()\n}\n")
	}

	for k := range cfg.BoolVars {
		key := cfg.ToConstBool(k)
		fieldName := strcase.ToLowerCamel(k)
		title := strcase.ToCamel(k)
		b.WriteString("\nfunc (c *DynamicConfig) " + title + "(ctx context.Context) bool {\n\t")
		b.WriteString("if c.client == nil {\n\t\treturn c." + fieldName + "\n\t}\n\n")
		b.WriteString("resp, err := c.client.DoCache(\n\t\tctx,\n\t\tc.client.B().Get().Key(" + key + ").Cache(),\n\t\ttime.Minute,\n\t).ToString()\n\n")
		b.WriteString("\tif err != nil {\n\t\treturn c." + fieldName + "\n\t}\n\n")
		b.WriteString("\tc." + fieldName + " = resp == \"true\"\n")
		b.WriteString("\treturn c." + fieldName + "\n}\n")

		b.WriteString("\nfunc (c *DynamicConfig) Set" + title + "(value bool) *DynamicConfig {\n\t")
		b.WriteString("c." + fieldName + " = value\n\treturn c\n}\n")

		b.WriteString("\nfunc (c *DynamicConfig) Store" + title + "(ctx context.Context, value bool) error {\n\t")
		b.WriteString("return c.client.Do(\n\t\tctx, \n\t\tc.client.B().Set().Key(" + key + ").\n\t\t\tValue(strconv.FormatBool(value)).\n\t\t\tBuild()).\n\t\tError()\n}\n")
	}

	b.WriteString("\nfunc (c *DynamicConfig) Update(ctx context.Context, req *RequestDynamicConfigUpdate) error {\n")
	for k := range cfg.StringVars {
		b.WriteString("\tif req." + strcase.ToCamel(k) + " != nil {\n")
		b.WriteString("\t\tif err := c.Store" + strcase.ToCamel(k) + "(ctx, *req." + strcase.ToCamel(k) + "); err != nil {\n")
		b.WriteString("\t\t\treturn err\n\t\t}\n\t}\n")
	}

	for k := range cfg.DurationVars {
		b.WriteString("\tif req." + strcase.ToCamel(k) + " != nil {\n")
		b.WriteString("\t\tif err := c.Store" + strcase.ToCamel(k) + "(ctx, *req." + strcase.ToCamel(k) + "); err != nil {\n")
		b.WriteString("\t\t\treturn err\n\t\t}\n\t}\n")
	}

	for k := range cfg.IntVars {
		b.WriteString("\tif req." + strcase.ToCamel(k) + " != nil {\n")
		b.WriteString("\t\tif err := c.Store" + strcase.ToCamel(k) + "(ctx, *req." + strcase.ToCamel(k) + "); err != nil {\n")
		b.WriteString("\t\t\treturn err\n\t\t}\n\t}\n")
	}

	for k := range cfg.BoolVars {
		b.WriteString("\tif req." + strcase.ToCamel(k) + " != nil {\n")
		b.WriteString("\t\tif err := c.Store" + strcase.ToCamel(k) + "(ctx, *req." + strcase.ToCamel(k) + "); err != nil {\n")
		b.WriteString("\t\t\treturn err\n\t\t}\n\t}\n")
	}

	b.WriteString("\treturn nil\n}\n")

	_, err = f.WriteString(b.String())

	return err
}
