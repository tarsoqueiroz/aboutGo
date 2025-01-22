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

// writeYAML marshals a map into YAML format and writes it to a file.
func writeYAML(filePath string, content map[string]interface{}) error {
	data, err := yaml.Marshal(content) // Marshals the map into YAML.
	if err != nil {
		return err // Returns an error if marshalling fails.
	}

	return os.WriteFile(filePath, data, 0644) // Writes the YAML data to a file using os.WriteFile.
}

// displayAndBuildYAML recursively prints the YAML structure and builds a new map from it.
func displayAndBuildYAML(input map[string]interface{}, indent string, output map[string]interface{}) {
	for key, value := range input {
		fmt.Printf("%s%s:\n", indent, key) // Prints the key with proper indentation.
		switch v := value.(type) {
		case map[string]interface{}: // Handles nested maps.
			subMap := make(map[string]interface{}) // Creates a new sub-map for the output.
			output[key] = subMap
			displayAndBuildYAML(v, indent+"  ", subMap) // Recursively processes the nested map.
		case []interface{}: // Handles lists.
			fmt.Printf("%s- %v\n", indent+"  ", v) // Prints the list elements.
			output[key] = v                        // Adds the list to the output map.
		default: // Handles scalar values.
			fmt.Printf("%s  %v\n", indent, v) // Prints the scalar value.
			output[key] = v                   // Adds the scalar value to the output map.
		}
	}
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

	outputFile := inputFile[:ext] + "_converted.yaml" // Constructs the output file name by appending "_deploy".

	inputData, err := readYAML(inputFile) // Reads and unmarshals the input YAML file.
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	outputData := make(map[string]interface{})     // Initializes the output map.
	displayAndBuildYAML(inputData, "", outputData) // Processes the input YAML and builds the output map.

	if err := writeYAML(outputFile, outputData); err != nil { // Writes the output map to the new YAML file.
		fmt.Printf("Error writing output file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Output written to %s\n", outputFile) // Prints a success message with the output file name.
}
