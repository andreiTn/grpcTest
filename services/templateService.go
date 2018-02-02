package services

import (
	"html/template"
)

var Tpl *template.Template

func init() {
	Tpl = template.Must(template.ParseGlob("web/templates/*"))
}