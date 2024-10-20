package main

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

// Define the structure for user inputs
type ProjectData struct {
	ProjectName string
}

func main() {
	// Log the start of the program
	fmt.Println("Starting project generator...")

	// Get the project name from user input
	if len(os.Args) < 2 {
		fmt.Println("Usage: go-go-gadget <project-name>")
		os.Exit(1)
	}
	projectName := os.Args[1]
	fmt.Println("Project name:", projectName)
	projectData := ProjectData{ProjectName: projectName}

	// Create the project directory
	fmt.Println("Creating project directory...")
	err := os.Mkdir(projectName, 0755)
	if err != nil {
		fmt.Println("Error creating project directory:", err)
		os.Exit(1)
	}
	fmt.Println("Project directory created:", projectName)

	// Define the templates and their destinations
	templates := map[string]string{
		"templates/main.go.tmpl":   filepath.Join(projectName, "main.go"),
		"templates/README.md.tmpl": filepath.Join(projectName, "README.md"),
		"templates/Makefile.tmpl":  filepath.Join(projectName, "Makefile"),
	}

	// Parse and generate files from templates
	for tmplPath, destPath := range templates {
		fmt.Println("Processing template:", tmplPath)
		err := generateFileFromTemplate(tmplPath, destPath, projectData)
		if err != nil {
			fmt.Println("Error generating file:", err)
			os.Exit(1)
		}
	}

	fmt.Println("Project generated successfully!")
}

// Helper function to generate files from templates
func generateFileFromTemplate(templatePath, destPath string, data ProjectData) error {
	// Log that the template is being parsed
	fmt.Println("Parsing template:", templatePath)

	// Parse the template
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}

	// Log that the destination file is being created
	fmt.Println("Creating destination file:", destPath)

	// Create the destination file
	destFile, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer destFile.Close()

	// Log that the template is being executed
	fmt.Println("Executing template...")

	// Execute the template and write to the destination file
	err = tmpl.Execute(destFile, data)
	if err != nil {
		return err
	}

	return nil
}
