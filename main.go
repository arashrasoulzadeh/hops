/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"embed"
	"hops/cmd"
	"hops/engine"
)

//go:embed modules
var fs embed.FS

func main() {
	engine.LoadPath(fs, "modules/intro")
	engine.LoadPath(fs, "modules/nginx")
	engine.LoadPath(fs, "modules/os")

	// Print collected metadata (functions, comments, variables)
	totalFunctions := 0
	for _, meta := range engine.LuaMetaMap {
		for range meta.Functions {
			totalFunctions++
		}
	}
	//    log.Println(fmt.Sprintf("loaded %d modules and %d functions", len(engine.LuaMetaMap), totalFunctions))

	cmd.Execute()
}
