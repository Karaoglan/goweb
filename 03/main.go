package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init()  {
	tpl = template.Must(template.ParseFiles("03/tpl.gohtml"))
}

func main() {

	sages := []string{"gandi", "MLK", "Budha"}

	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}
