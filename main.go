package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func ParseResume(fileName string) Resume {
	var res Resume

	// read file
	f, e := ioutil.ReadFile(fileName)

	if e != nil {
		panic(e)
	}

	// parse yaml
	e = yaml.Unmarshal(f, &res)

	if e != nil {
		panic(e)
	}

	return res
}

// Reading CSS Stylesheet
func ReadCSSFile(fileName string) string {
	f, e := ioutil.ReadFile(fileName)
	css := ""
	if e != nil {
		fmt.Println("Error reading CSS file")
	} else {
		css = string(f)
	}

	return css
}

func main() {
	// parse command line flags
	infile := flag.String("i", "resume.yml", "Input file")
	css_file := flag.String("s", "resume.css", "CSS file")
	flag.Parse()

	resume := ParseResume(*infile)

	// string of html with stylesheet link
	resumeHtmlBlock := resume.ToHTML()
	html := `<!DOCTYPE html><html><head>
	<link rel="stylesheet" type="text/css" href="https://cdn.jsdelivr.net/gh/dreampulse/computer-modern-web-font@master/fonts.css">
	<style>` + ReadCSSFile(*css_file) + `</style></head><body>` + resumeHtmlBlock + `</body></html>`

	// save string to html file
	ioutil.WriteFile(resume.Header.Name+" Resume.html", []byte(html), 0644)
}
