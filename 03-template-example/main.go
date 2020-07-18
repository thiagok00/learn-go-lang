package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

// Person ...
type Person struct {
	Name string
	Age  int
}

func greet(name string) string {
	return fmt.Sprintf("wohoo hello %s", name)
}

func main() {

	p1 := Person{"Thiago", 26}
	p2 := Person{"Ana", 24}
	p3 := Person{"James", 30}

	var fm = template.FuncMap{"greet": greet}

	// Since the templates created by ParseFiles are named by the base names of the argument files,
	// tpl should usually have the name of one of the (base) names of the files.
	// https://golang.org/pkg/text/template/#Template.ParseFiles
	tpl := template.Must(template.New("index.gohtml").Funcs(fm).ParseFiles("template/index.gohtml"))

	err := tpl.Execute(os.Stdout, []Person{p1, p2, p3})
	if err != nil {
		log.Fatalln(err)
	}
}
