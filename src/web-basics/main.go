package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
	"time"
)

type overwatchCharacter struct {
	Name         string
	Role         string
	HealthPoints int
	Weapon       string
	Country      string
}

//FuncMap returns a variable of type FuncMap that maps names ("uc") into function("string.ToUpper") to be used in a template
var fm = template.FuncMap{
	"uc":         strings.ToUpper,
	"formatTime": formatTime,
	"double":     doubleMyNumber,
}

//tpl is the representation of a parsed Template
//pointer of type Template of package template
var tpl *template.Template

//here we initialize our program, parsing all templates of a folder 'templates' into
//tpl variable
func init() {
	//must return a pointer of a type Template
	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("templates/*")) //1-tpl holds all my parsed template file
}

func main() {
	start := time.Now()
	defer fmt.Println("\nending now...")

	slc := []string{"Alface", "Arroz", "Carne", "Batata"}
	consoleBrands := map[string]string{
		"Nintendo":  "Switch",
		"Sony":      "PlayStation4",
		"Microsoft": "Xbox",
	}
	oC := overwatchCharacter{
		"Tracer",
		"DPS",
		200,
		"dual wield automatic pistols",
		"United Kingdom",
	}

	//2-Here we executes tpl, which holds all my template files
	err := tpl.ExecuteTemplate(os.Stdout, "header.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "index.gohtml", 42)
	if err != nil {
		log.Fatal(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "slice_on_template.gohtml", slc)
	if err != nil {
		log.Fatal(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "map_on_template.gohtml", consoleBrands)
	if err != nil {
		log.Fatal(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "struct_on_template.gohtml", oC)
	if err != nil {
		log.Fatal(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "function_on_template.gohtml", time.Now())
	if err != nil {
		log.Fatal(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "pipeline_on_template.gohtml", 3)
	if err != nil {
		log.Fatal(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "footer.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Printf("\nelapsed: %v", elapsed)
}

func formatTime(t time.Time) string {
	return t.Format("01-02-2006")
}

func doubleMyNumber(number int) int {
	return number * 2
}
