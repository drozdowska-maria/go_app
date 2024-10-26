package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <output_path>")
		return
	}

	data := struct {
		Name      string
		RouteName string
	}{
		Name:      os.Args[1],
		RouteName: generateRouteName(os.Args[1]),
	}

	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error loading template:", err)
		return
	}
	templatePath := filepath.Join(currentDir, "libs", "templates", "controller.gotmpl")
	controllersPath := filepath.Join(currentDir, "src", "adapters", fmt.Sprintf("%sController.go", data.Name))
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		fmt.Println("Error loading template:", err)
		return
	}

	var generatedCode bytes.Buffer
	err = tmpl.Execute(&generatedCode, data)
	if err != nil {
		fmt.Println("Error generating code:", err)
		return
	}

	err = os.WriteFile(controllersPath, generatedCode.Bytes(), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("Controller generated successfully.")
}

func generateRouteName(input string) string {
	re := regexp.MustCompile("([a-z0-9])([A-Z])")
	kebab := re.ReplaceAllString(input, "${1}-${2}")

	return strings.ToLower(fmt.Sprintf("/%ss", kebab))
}
