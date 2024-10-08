package engine

import (
	"bufio"
	"embed"
	"log"
	"regexp"
	"strings"
)

// collectLuaMetadata parses Lua files for comments and variables
func collectLuaMetadata(fs *embed.FS, filePath string, metadata *LuaMetadata) {
	file, err := fs.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file: %v\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var currentFunction string
	var lastComment string

	// Regex for detecting Lua variables (both local and global)
	//varPattern := regexp.MustCompile(`^(\s*)(local\s+)?([a-zA-Z_]\w*)\s*=\s*(.+)`)
	funcPattern := regexp.MustCompile(`^(\s*)function\s+([a-zA-Z_]\w*)\s*\(`)

	for scanner.Scan() {
		line := scanner.Text()
		args := extractFunctionArguments(line)

		// Check for comments (starting with --)
		if strings.HasPrefix(strings.TrimSpace(line), "--") {
			lastComment = strings.TrimSpace(strings.TrimPrefix(line, "--"))
			continue
		}

		// Check for function declarations
		if matches := funcPattern.FindStringSubmatch(line); matches != nil {
			currentFunction = matches[2] // Get the function name

			// Store the comment in metadata for this function
			metadata.Functions[currentFunction] = matches[2]

			// Store the comment in Comments for this function
			metadata.Comments[currentFunction] = lastComment
			if len(args) > 0 {
				metadata.Variables[currentFunction] = args
			}

			lastComment = "" // Reset the comment after it's assigned
			continue
		}

		//// Check for variable declarations
		//if matches := varPattern.FindStringSubmatch(line); matches != nil {
		//	varName := matches[3]
		//	varValue := matches[4]
		//	metadata.Variables[varName] = varValue
		//}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v\n", err)
	}
}
func extractFunctionArguments(functionDeclaration string) []string {
	// (?s) enables the dot to match newlines
	var funcPattern = regexp.MustCompile(`(?s)^\s*function\s+([a-zA-Z_]\w*)\s*\(\s*([^\)]*)\s*\)`)
	matches := funcPattern.FindStringSubmatch(functionDeclaration)
	if len(matches) > 2 {
		// Split the arguments by comma and trim whitespace
		args := strings.Split(matches[2], ",")
		for i := range args {
			args[i] = strings.TrimSpace(args[i]) // Trim whitespace
		}
		return args
	}
	return nil
}
