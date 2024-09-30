/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
    "hops/cmd"
    "hops/engine"
)

func main() {
    engine.LoadPath("modules/intro")
    engine.LoadPath("modules/nginx")
    engine.LoadPath("modules/os")

    // Print collected metadata (functions, comments, variables)
    totalFunctions := 0
    for _, meta := range engine.LuaMetaMap {
        for _ = range meta.Functions {
            totalFunctions++
        }
    }
    //    log.Println(fmt.Sprintf("loaded %d modules and %d functions", len(engine.LuaMetaMap), totalFunctions))

    cmd.Execute()
}
