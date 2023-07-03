package run

import (
	"embed"
	"fmt"

	"github.com/vision-cli/common/execute"
	"github.com/vision-cli/common/file"
	"github.com/vision-cli/common/tmpl"
	"github.com/vision-cli/vision-plugin-plugin-v1/placeholders"
)

const (
	goTemplateDir = "_templates"
)

//go:embed all:_templates
var templateFiles embed.FS

func Create(p *placeholders.Placeholders, executor execute.Executor, t tmpl.TmplWriter) error {
	var err error

	if file.Exists(p.Directory) {
		return fmt.Errorf("directory %q already exists", p.Name)
	}

	if err = tmpl.GenerateFS(templateFiles, goTemplateDir, p.Directory, p, false, t); err != nil {
		return fmt.Errorf("generating structure from the template: %w", err)
	}

	return nil
}
