package lib

import (
	"html/template"
	"log"
	"os"
)

func GenerateResume(json map[string]interface{}, templateDir string, outputFile string) {
	// render html template
	template := template.New("")
	template.Funcs(registerTemplateFuncs())
	template.ParseGlob(templateDir + "/*.tmpl")
	template.ParseGlob(templateDir + "/**/*.tmpl")
	// create output file if it doesn't exist
	outputfile, err := os.Create(outputFile)
	if err != nil {
		panic(err)
	}
	err = template.ExecuteTemplate(outputfile, "resume.tmpl", json)
	if err != nil {
		log.Println("Error executing template: ", err)
	}
}
