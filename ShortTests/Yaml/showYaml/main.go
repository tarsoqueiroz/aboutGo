package main

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

// readYAML reads a YAML file from the given path and unmarshals its content into a map.
func readYAML(filePath string) (map[string]interface{}, error) {
	data, err := os.ReadFile(filePath) // Reads the file content using os.ReadFile.
	if err != nil {
		return nil, err // Returns an error if the file cannot be read.
	}

	var content map[string]interface{}
	if err := yaml.Unmarshal(data, &content); err != nil { // Unmarshals YAML data into a map.
		return nil, err // Returns an error if unmarshalling fails.
	}

	return content, nil // Returns the unmarshalled content.
}

// updateYAML applies the required transformations to the YAML structure.
func updateYAML(input map[string]interface{}) map[string]interface{} {
	// Update "apiVersion" to "apps/v1".
	input["apiVersion"] = "apps/v1"

	// Update "kind" to "Deployment".
	input["kind"] = "Deployment"

	if metadata, ok := input["metadata"].(map[string]interface{}); ok {
		// Remove "metadata.annotations" if it exists.
		delete(metadata, "annotations")
		// Remove "metadata.creationTimestamp" if it exists.
		delete(metadata, "creationTimestamp")
		// Remove "metadata.generation" if it exists.
		delete(metadata, "generation")
		// Remove "metadata.resourceVersion" if it exists.
		delete(metadata, "resourceVersion")
		// Remove "metadata.uid" if it exists.
		delete(metadata, "uid")
	}

	if spec, ok := input["spec"].(map[string]interface{}); ok {
		if selector, ok := spec["selector"].(map[string]interface{}); ok {
			// Remove "spec.selector.deploymentconfig" if it exists.
			delete(selector, "deploymentconfig")

			if app, ok := selector["app"]; ok {
				// Move "spec.selector.app" to "spec.selector.matchLabels.app".
				if matchLabels, ok := selector["matchLabels"].(map[string]interface{}); ok {
					matchLabels["app"] = app
				} else {
					selector["matchLabels"] = map[string]interface{}{
						"app": app,
					}
				}
				delete(selector, "app")
			}
		}

		if strategy, ok := spec["strategy"].(map[string]interface{}); ok {
			// Update "spec.strategy.type" to "RollingUpdate".
			strategy["type"] = "RollingUpdate"
		}

		if template, ok := spec["template"].(map[string]interface{}); ok {
			if tpltMetadata, ok := template["metadata"].(map[string]interface{}); ok {
				if tpltMtdtLabels, ok := tpltMetadata["labels"].(map[string]interface{}); ok {
					// Remove "spec.template.metadata.labels.deploymentconfig" if it exists.
					delete(tpltMtdtLabels, "deploymentconfig")
				}
			}
		}

		// Remove "spec.triggers" if it exists.
		delete(spec, "triggers")

	}

	// Remove "status" if it exists.
	delete(input, "status")

	return input
}

// writeYAML marshals a map into YAML format and writes it to a file.
func writeYAML(filePath string, content map[string]interface{}) error {
	data, err := yaml.Marshal(content) // Marshals the map into YAML.
	if err != nil {
		return err // Returns an error if marshalling fails.
	}

	return os.WriteFile(filePath, data, 0644) // Writes the YAML data to a file using os.WriteFile.
}

func main() {
	if len(os.Args) < 2 { // Checks if the input file argument is provided.
		fmt.Println("Usage: go run script.go <input_file.yaml>")
		os.Exit(1)
	}

	inputFile := os.Args[1]                  // Gets the input file path from the command-line arguments.
	ext := strings.LastIndex(inputFile, ".") // Finds the last occurrence of '.' to determine the file extension.
	if ext == -1 {                           // Validates that the file has an extension.
		fmt.Println("Invalid file name. Ensure it has a YAML extension.")
		os.Exit(1)
	}

	outputFile := inputFile[:ext] + "_deploy.yaml" // Constructs the output file name by appending "_deploy".

	inputData, err := readYAML(inputFile) // Reads and unmarshals the input YAML file.
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	updatedData := updateYAML(inputData) // Applies the required transformations to the YAML data.

	if err := writeYAML(outputFile, updatedData); err != nil { // Writes the updated data to the new YAML file.
		fmt.Printf("Error writing output file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Output written to %s\n", outputFile) // Prints a success message with the output file name.
}
