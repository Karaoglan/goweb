package main

import (
	"log"
	"os"
	"path"
	"strings"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

// create a FuncMap to register functions.
// "uc" is what the func will be called in the template
// "uc" is the ToUpper func from package strings
// "ft" is a func I declared
// "ft" slices a string, returning the first three characters
var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 3 {
		s = s[:3]
	}
	return s
}

func init()  {
	tpl = template.Must(template.New(path.Base("03/tpl.gohtml")).Funcs(fm).ParseFiles("03/tpl.gohtml"))
}

func main() {

	sages := []string{"gandi", "MLK", "Budha"}

	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}
