package lib

import (
	"html/template"
	"strings"
	"time"
)

func registerTemplateFuncs() template.FuncMap {
	funcMap := template.FuncMap{
		"spaceToDash": func(str string) string {
			// replace spaces with dashes
			// and lowercase
			return strings.ToLower(strings.Replace(str, " ", "-", -1))
		},
		"paragraphSplit": func(str string) template.HTML {
			// split string by newlines
			strByNewLine := strings.Split(str, "\n")
			return template.HTML("<p>" + strings.Join(strByNewLine, "</p><p>") + "</p>")
		},
		"toLowerCase": func(str string) string {
			return strings.ToLower(str)
		},
		"toMonthYear": func(str string) string {
			date, err := time.Parse("2006-01-02", str)
			if err != nil {
				return str
			}
			return date.Format("January 2006")
		},
		"toYear": func(str string) string {
			date, err := time.Parse("2006-01-02", str)
			if err != nil {
				return str
			}
			return date.Format("2006")
		},
		"toDayMonthYear": func(str string) string {
			date, err := time.Parse("2006-01-02", str)
			if err != nil {
				return str
			}
			return date.Format("02 January 2006")
		},
	}
	return funcMap
}
