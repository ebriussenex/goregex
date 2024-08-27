package main

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/browser"

	"github.com/ebriussenex/goregex/regex"
	"github.com/ebriussenex/goregex/templates"
)

func main() {

	args := os.Args[1:]
	fmt.Println(args)
	switch args[0] {
	case "draw":
		if len(args) == 2 {
			RenderFSM(args[1])
		}

		if len(args) == 3 {

			RenderRunner(args[1], args[2])
		}
	case "out":
		if len(args) == 4 {
			OutputRunnerToFile(args[1], args[2], args[3])
		}
	default:
		panic("command not recognized")
	}
}

func RenderFSM(input string) {
	graph := regex.NewRegex(input).DebugFSM()
	html := buildFsmHtml(graph)
	outputToBrowser(html)
}

func RenderRunner(regex, input string) {
	data := buildRunnerTemplateData(regex, input)
	htmlRunner := buildRunnerHTML(data)
	outputToBrowser(htmlRunner)
}

func buildRunnerHTML(data templates.TemplateData) string {
	return renderWithTemplate(templates.RunnerTemplate, data)
}

func buildRunnerTemplateData(regexp string, input string) templates.TemplateData {
	newRegex := regex.NewRegex(regexp)
	debugSteps := newRegex.DebugMatch(input)

	var steps []templates.Step
	for _, step := range debugSteps {
		steps = append(steps, templates.Step{
			Graph:      step.RunnerDrawing,
			InputSplit: threeSplitString(input, step.CurrentCharacterIndex),
		})
	}

	data := templates.TemplateData{
		Steps: steps, Regex: regexp,
	}

	return data
}

func buildFsmHtml(graph string) string {
	return renderWithTemplate(templates.FsmTemplate, graph)
}

func renderWithTemplate(tmplt string, data any) string {
	t, err := template.New("graph").Parse(tmplt)
	if err != nil {
		panic(err)
	}
	w := bytes.Buffer{}
	err = t.Execute(&w, data)
	if err != nil {
		panic(err)
	}
	return w.String()
}

func outputToBrowser(html string) {
	reader := strings.NewReader(html)
	err := browser.OpenReader(reader)
	if err != nil {
		panic(err)
	}
}

func threeSplitString(s string, i int) []string {
	var left, middle, right string

	left = s[:i]
	if i < len(s) {
		middle = string(s[i])
		right = s[i+1:]
	}

	return []string{left, middle, right}
}

func OutputRunnerToFile(regex, input, filePath string) {
	data := buildRunnerTemplateData(regex, input)
	htmlRunner := buildRunnerHTML(data)
	outputToFile(htmlRunner, filePath)
}

func outputToFile(html, path string) {
	containingDir := filepath.Dir(path)
	err := os.MkdirAll(containingDir, 0750)
	if err != nil {
		panic(err)
	}

	if filepath.Ext(path) == "" {
		path += ".html"
	}

	if filepath.Ext(path) != ".html" {
		panic("only .html extension permitted")
	}

	err = os.WriteFile(path, []byte(html), 0750)
	if err != nil {
		panic(err)
	}
}
