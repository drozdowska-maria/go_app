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
		fmt.Println("error loading current dir: ", err.Error())
		return
	}

	reposMap := filepath.Join(currentDir, "repos.go")
	ctrlsMap := filepath.Join(currentDir, "controller.go")
	mainFile := filepath.Join(currentDir, "main.go")

	var filesToGenerate = []struct {
		filePath     string
		templatePath string
	}{
		{
			filePath:     filepath.Join(currentDir, "src", "controllers", fmt.Sprintf("%sController.go", strings.ToLower(os.Args[1]))),
			templatePath: filepath.Join(currentDir, "libs", "templates", "controller.gotmpl"),
		},
		{
			filePath:     filepath.Join(currentDir, "src", "domain", fmt.Sprintf("%sService.go", strings.ToLower(os.Args[1]))),
			templatePath: filepath.Join(currentDir, "libs", "templates", "service.gotmpl"),
		},
		{
			filePath:     filepath.Join(currentDir, "src", "repositories", "v1", fmt.Sprintf("%sRepository.go", strings.ToLower(os.Args[1]))),
			templatePath: filepath.Join(currentDir, "libs", "templates", "repository.gotmpl"),
		},
	}

	err = addMapEntry(reposMap, fmt.Sprintf("controllers.%s", os.Args[1]), fmt.Sprintf("repositories.New%sRepository()", os.Args[1]))
	if err != nil {
		fmt.Println("Error:", err.Error())
	}
	err = addMapEntry(ctrlsMap, fmt.Sprintf("controllers.%s", os.Args[1]), fmt.Sprintf("controllers.New%sAdapter()", os.Args[1]))
	if err != nil {
		fmt.Println("Error:", err.Error())
	}

	err = replaceLineWithContent(mainFile, filepath.Join(currentDir, "libs", "templates", "main_input.gotmpl"), os.Args[1])
	if err != nil {
		fmt.Println("Error:", err.Error())
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

func addMapEntry(filename, key, value string) error {
	content, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	re := regexp.MustCompile(`\}`)

	newEntry := fmt.Sprintf(`	%s: %s,`, key, value)

	modifiedContent := re.ReplaceAllString(string(content), newEntry+"\n}")

	return os.WriteFile(filename, []byte(modifiedContent), 0644)
}

func replaceLineWithContent(targetFile, sourceFile, name string) error {
	data := struct {
		Name      string
		NameLower string
	}{
		Name:      name,
		NameLower: strings.ToLower(name),
	}

	targetContent, err := os.ReadFile(targetFile)
	if err != nil {
		return fmt.Errorf("błąd przy odczycie pliku docelowego: %v", err)
	}

	tmpl, err := template.ParseFiles(sourceFile)
	if err != nil {
		return fmt.Errorf("error loading template: %s", err.Error())

	}

	var generatedCode bytes.Buffer
	err = tmpl.Execute(&generatedCode, data)
	if err != nil {
		return fmt.Errorf("error generating code: %s", err.Error())

	}

	re := regexp.MustCompile(`// NEXT CASES - DO NOT DELETE THIS COMMENT`)

	modifiedContent := re.ReplaceAllString(string(targetContent), generatedCode.String())

	err = os.WriteFile(targetFile, []byte(modifiedContent), 0644)
	if err != nil {
		return fmt.Errorf("błąd przy zapisie do pliku docelowego: %v", err)
	}

	return nil
}

// replaceFileContent zastępuje zawartość pliku `templateFile` treścią z pliku `sourceFile`
func replaceFileContent(templateFile, sourceFile string) error {
	// Wczytujemy treść z pliku źródłowego
	sourceContent, err := os.ReadFile(sourceFile)
	if err != nil {
		return fmt.Errorf("błąd przy odczycie pliku źródłowego: %v", err)
	}

	// Nadpisujemy plik `templateFile` treścią z `sourceFile`
	err = os.WriteFile(templateFile, sourceContent, 0644)
	if err != nil {
		return fmt.Errorf("błąd przy zapisie do pliku template: %v", err)
	}

	return nil
}
