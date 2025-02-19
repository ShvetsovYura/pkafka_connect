package services

import (
	"bytes"
	"fmt"
	"github/ShvetsovYura/pkafka_connect/internal/types"
	"text/template"
)

var metricsTemplate = `{{range .}}
# HELP {{.Name}} {{.Description}}
# TYPE {{.Name}} {{.Type}}
{{.Name}} {{.Value}}{{end}}`

type TemplateService struct {
	tpl *template.Template
}

func NewTemplateService() (*TemplateService, error) {
	metricsList, err := template.New("metricsList").Parse(metricsTemplate)
	if err != nil {
		return nil, err
	}
	return &TemplateService{
		tpl: metricsList,
	}, nil
}

func (t *TemplateService) Render(metrics []types.Metric) ([]byte, error) {
	var result bytes.Buffer

	if err := t.tpl.Execute(&result, metrics); err != nil {
		return nil, fmt.Errorf("ошибка при выполнении шаблона: %w", err)
	}
	return result.Bytes(), nil
}
