package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init()  {
	tpl = template.Must(template.ParseFiles("02/tpl.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 42)
	if err != nil {
		log.Fatalln(err)
	}
}
