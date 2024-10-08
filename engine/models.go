package engine

import (
	lua "github.com/yuin/gopher-lua"
)

// LuaMetadata holds information about functions, variables, and comments from Lua
type LuaMetadata struct {
	Functions map[string]string   // Function name to comment
	Variables map[string][]string // Variable name to value
	Comments  map[string]string   // Variable to hold comments
}

// Define the type for the global function map (nested map by folder and function)
type luaFunctionMap map[string]map[string]*lua.LFunction

// FunctionMap Global function map
var FunctionMap = make(luaFunctionMap)

var LuaMetaMap = make(map[string]LuaMetadata) // Global metadata map
