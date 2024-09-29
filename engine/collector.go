package engine

import (
    "bufio"
    "log"
    "os"
    "regexp"
    "strings"
)

// collectLuaMetadata parses Lua files for comments and variables
func collectLuaMetadata(filePath string, metadata *LuaMetadata) {
    file, err := os.Open(filePath)
    if err != nil {
        log.Fatalf("Error opening file: %v\n", err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var currentFunction string
    var lastComment string

    // Regex for detecting Lua variables (both local and global)
    varPattern := regexp.MustCompile(`^(\s*)(local\s+)?([a-zA-Z_]\w*)\s*=\s*(.+)`)
    funcPattern := regexp.MustCompile(`^(\s*)function\s+([a-zA-Z_]\w*)\s*\(`)

    for scanner.Scan() {
        line := scanner.Text()

        // Check for comments (starting with --)
        if strings.HasPrefix(strings.TrimSpace(line), "--") {
            lastComment = strings.TrimSpace(strings.TrimPrefix(line, "--"))
            continue
        }

        // Check for function declarations
        if matches := funcPattern.FindStringSubmatch(line); matches != nil {
            currentFunction = matches[2] // Get the function name
            metadata.Functions[currentFunction] = lastComment
            lastComment = "" // Reset comment after it's assigned
            continue
        }

        // Check for variable declarations
        if matches := varPattern.FindStringSubmatch(line); matches != nil {
            varName := matches[3]
            varValue := matches[4]
            metadata.Variables[varName] = varValue
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatalf("Error reading file: %v\n", err)
    }
}
