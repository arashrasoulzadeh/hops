package engine

import (
    lua "github.com/yuin/gopher-lua"
    "io/ioutil"
    "log"
    "path/filepath"
    "strings"
)

// LoadPath loads Lua functions, variables, and comments from a directory
func LoadPath(folder string) {
    l := lua.NewState()
    defer l.Close()

    // Create a sub-map for the folder
    FunctionMap[folder] = make(map[string]*lua.LFunction)

    // Initialize metadata for the folder
    meta := LuaMetadata{
        Functions: make(map[string]string),
        Variables: make(map[string]string),
    }

    // Load Lua files and collect functions, comments, and variables
    files, err := ioutil.ReadDir(folder)
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

            // Run the Lua file and get the returned table
            if err := l.DoFile(filePath); err != nil {
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
