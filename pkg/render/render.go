package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedFiles, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedFiles.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template", err)
		return
	}
}
