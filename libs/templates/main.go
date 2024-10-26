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

	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("error loading template: %s", err.Error())
		return
	}

	var filesToGenerate = []struct {
		filePath     string
		templatePath string
	}{
		{
			filePath:     filepath.Join(currentDir, "src", "adapters", fmt.Sprintf("%s%s.go", strings.ToLower(os.Args[1]), "Controller")),
			templatePath: filepath.Join(currentDir, "libs", "templates", "controller.gotmpl"),
		},
		{
			filePath:     filepath.Join(currentDir, "src", "domain", fmt.Sprintf("%s%s.go", strings.ToLower(os.Args[1]), "Service")),
			templatePath: filepath.Join(currentDir, "libs", "templates", "service.gotmpl"),
		},
		{
			filePath:     filepath.Join(currentDir, "src", "ports", "v1", fmt.Sprintf("%s%s.go", strings.ToLower(os.Args[1]), "Repository")),
			templatePath: filepath.Join(currentDir, "libs", "templates", "repository.gotmpl"),
		},
	}

	for _, paths := range filesToGenerate {
		_, err := os.Stat(paths.filePath)

		if err != nil {
			err := generateFile(os.Args[1], paths.filePath, paths.templatePath)

			if err != nil {
				fmt.Println("error: ", err.Error())
				return
			}
		}

	}

}

func generateRouteName(input string) string {
	re := regexp.MustCompile("([a-z0-9])([A-Z])")
	kebab := re.ReplaceAllString(input, "${1}-${2}")

	return strings.ToLower(fmt.Sprintf("/%ss", kebab))
}

func generateFile(name string, dirPath string, templatePath string) error {

	data := struct {
		Name      string
		RouteName string
	}{
		Name:      name,
		RouteName: generateRouteName(name),
	}

	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return fmt.Errorf("error loading template: %s", err.Error())

	}

	var generatedCode bytes.Buffer
	err = tmpl.Execute(&generatedCode, data)
	if err != nil {
		return fmt.Errorf("error generating code: %s", err.Error())

	}

	err = os.WriteFile(dirPath, generatedCode.Bytes(), 0644)
	if err != nil {
		return fmt.Errorf("error writing file: %s", err.Error())

	}

	fmt.Println("File generated successfully: ", dirPath)

	return nil
}
