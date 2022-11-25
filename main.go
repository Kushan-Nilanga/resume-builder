package main

import (
	"io/ioutil"
	"resume-builder/models"

	"gopkg.in/yaml.v2"
)

const resume_file = "./resources/resume.yml"

func ParseResume(fileName string) models.Resume {
	var res models.Resume

	f, e := ioutil.ReadFile(resume_file)

	if e != nil {
		panic(e)
	}

	e = yaml.Unmarshal(f, &res)

	if e != nil {
		panic(e)
	}

	return res
}

func main() {
	resume := ParseResume("./resources/resume.yml")

	// string of html with stylesheet link
	resumeHtmlBlock := resume.ToHTML()
	html := `<!DOCTYPE html><html><head>
	<link rel="stylesheet" type="text/css" href="https://cdn.jsdelivr.net/gh/dreampulse/computer-modern-web-font@master/fonts.css">
	<link rel='stylesheet' href='resume.css'></head><body>` + resumeHtmlBlock + `</body></html>`

	// save string to file
	ioutil.WriteFile("./resources/resume.html", []byte(html), 0644)
}
