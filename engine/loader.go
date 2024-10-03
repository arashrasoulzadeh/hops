package engine

import (
	"embed"
	"hops/renderer"
	"log"
	"path/filepath"
	"strings"

	lua "github.com/yuin/gopher-lua"
)

// LoadPath loads Lua functions, variables, and comments from a directory
func LoadPath(fs embed.FS, folder string) {
	l := lua.NewState()
	defer l.Close()

	// Create a sub-map for the folder
	FunctionMap[folder] = make(map[string]*lua.LFunction)

	// Initialize metadata for the folder
	meta := LuaMetadata{
		Functions: make(map[string]string),
		Variables: make(map[string]string),
		Comments:  make(map[string]string),
	}

	r := renderer.NewRenderer()

	// Read the directory using embed.FS
	files, err := fs.ReadDir(folder)
	if err != nil {
		log.Fatal(err)
	}

	// Loop through each file in the folder
	for _, file := range files {
		// Check if the file has a .lua extension
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".lua") {
			filePath := filepath.Join(folder, file.Name())

			// Collect comments and variables from Lua file and update metadata
			collectLuaMetadata(filePath, &meta)

			// Read the content of the Lua file
			content, err := fs.ReadFile(filePath)
			if err != nil {
				log.Fatalf("Error reading %s: %v\n", filePath, err)
			}
			// Render content of file , eg replace place holders , this does not compile lua
			rendered_content, err := r.Render(content)
			if err != nil {
				log.Fatalf("Error rendering %s: %v\n", filePath, err)
			}

			// Run the Lua script from the file content
			if err := l.DoString(string(rendered_content)); err != nil {
				log.Fatalf("Error loading %s: %v\n", filePath, err)
			}

			// Get the table returned by the Lua script
			if tbl := l.Get(-1); tbl.Type() == lua.LTTable {
				tbl := tbl.(*lua.LTable)

				// Iterate through the table to get the functions
				tbl.ForEach(func(key lua.LValue, value lua.LValue) {
					if fn, ok := value.(*lua.LFunction); ok {
						funcName := key.String()
						FunctionMap[folder][funcName] = fn
					}
				})
			}
		}
	}

	// After processing all files, update the global luaMetaMap for this folder
	LuaMetaMap[folder] = meta
}
