package main

import (
    "fmt"
    lua "github.com/yuin/gopher-lua"
    "log"
)

// LuaMetadata holds information about functions, variables, and comments from Lua
type LuaMetadata struct {
    Functions map[string]string // Function name to comment
    Variables map[string]string // Variable name to value
}

// Define the type for the global function map (nested map by folder and function)
type luaFunctionMap map[string]map[string]*lua.LFunction

// Global function map
var functionMap = make(luaFunctionMap)
var luaMetaMap = make(map[string]LuaMetadata) // Global metadata map

func main() {
    // Load Lua files from different directories
    LoadPath("modules/intro")
    LoadPath("modules/nginx")

    // Print collected metadata (functions, comments, variables)
    totalFunctions := 0
    for _, meta := range luaMetaMap {
        for _ = range meta.Functions {
            totalFunctions++
        }
    }
    log.Println(fmt.Sprintf("loaded %d modules and %d functions", len(luaMetaMap), totalFunctions))

}
