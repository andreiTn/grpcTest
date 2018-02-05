package templmanager

import (
	"fmt"
	"github.com/oxtoacart/bpool"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"myGrpc/config"
)

var templates map[string]*template.Template
var bufpool *bpool.BufferPool
var mainTmpl = `{{define "main" }} {{ template "base" . }} {{ end }}`

type TemplateError struct {
	s string
}

func (e *TemplateError) Error() string {
	return e.s
}

func NewError(text string) error {
	return &TemplateError{text}
}

// create a buffer pool
func init() {
	bufpool = bpool.NewBufferPool(100  * 10024)
	log.Println("buffer allocation successful")
}

func LoadTemplates(config config.Config) (err error) {
	fmt.Println("Load template....")
	if templates == nil {
		templates = make(map[string]*template.Template)
	}

	layoutFiles, err := filepath.Glob(config.TemplateLayoutPath + "*.gohtml")
	if err != nil {
		return err
	}

	includeFiles, err := filepath.Glob(config.TemplateIncludePath + "*.gohtml")
	if err != nil {
		return err
	}
	fmt.Println("TC: ", config)
	mainTemplate := template.New("main")
	mainTemplate, err = mainTemplate.Parse(mainTmpl)

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range includeFiles {
		fileName := filepath.Base(file)
		fmt.Println("NAME: ", fileName)
		files := append(layoutFiles, file)
		fmt.Println(fileName)
		templates[fileName], err = mainTemplate.Clone()
		if err != nil {
			return err
		}
		templates[fileName] = template.Must(templates[fileName].ParseFiles(files...))
	}

	log.Println("templates loading successful")
	return nil

}

func RenderTemplate(w http.ResponseWriter, name string, data interface{}) error {
	tmpl, ok := templates[name]
	if !ok {
		http.Error(w, fmt.Sprintf("The template %s does not exist.", name),
			http.StatusInternalServerError)
		err := NewError("Template doesn't exist")
		return err
	}

	buf := bufpool.Get()
	defer bufpool.Put(buf)

	err := tmpl.Execute(buf, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		err := NewError("Template execution failed")
		return err
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	buf.WriteTo(w)
	return nil
}
