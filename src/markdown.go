package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/google/go-github/github"
)

func main() {
	mdPath := os.Args[1]
	if mdPath == "" {
		fmt.Println("Required markdown file path.")
		return
	}
	md, err := ioutil.ReadFile(mdPath)
	if err != nil {
		fmt.Println("No such as markdown file")
		return
	}

	client := github.NewClient(nil)

	html, _, err := client.Markdown(string(md), nil)
	if err != nil {
		panic(err)
	}

	bodyTag := Tag{
		Name: "body",
	}

	html = bodyTag.insert(html)
	htmlTag := Tag{
		Name: "html",
	}

	html = htmlTag.insert(html)

	cssFile, err := ioutil.ReadFile("./github.css")
	css := "<style type=\"text/css\">"
	css += string(cssFile)
	css += "</style>"

	html += css
	ioutil.WriteFile("sample.html", []byte(html), 0644)
}

type Tag struct {
	Name string
}

func (t *Tag) OpenTag() string {
	return "<" + t.Name + ">"
}

func (t *Tag) CloseTag() string {
	return "</" + t.Name + ">"
}

func (t *Tag) insert(html string) string {
	html = t.OpenTag() + "\n" + html + "\n" + t.CloseTag()
	return html
}
