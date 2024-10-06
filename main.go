/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"embed"
	"fmt"
	"hops/cmd"
	"hops/engine"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//go:embed modules/**/*.lua
var fs embed.FS

func main() {
	// Load Lua paths from the embedded filesystem
	if err := engine.LoadPath(fs, "modules/intro"); err != nil {
		fmt.Println("Error loading modules/intro:", err)
	}
	if err := engine.LoadPath(fs, "modules/nginx"); err != nil {
		fmt.Println("Error loading modules/nginx:", err)
	}
	if err := engine.LoadPath(fs, "modules/os"); err != nil {
		fmt.Println("Error loading modules/os:", err)
	}

	// Print collected metadata (functions, comments, variables)
	totalFunctions := 0
	for _, meta := range engine.LuaMetaMap {
		for range meta.Functions {
			totalFunctions++
		}
	}

	// Create a file append logger
	file, _ := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	// Create a encoder
	encoderConfig := zap.NewProductionEncoderConfig()
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	// Create a file logger
	core := zapcore.NewCore(encoder, zapcore.AddSync(file), zapcore.DebugLevel)
	logger := zap.New(core)

	defer logger.Sync()

	// Replace the global logger with the file logger
	zap.ReplaceGlobals(logger)

	zap.L().Debug(fmt.Sprintf("loaded %d modules and %d functions", len(engine.LuaMetaMap), totalFunctions))

	// Execute the root command
	cmd.Execute()
}
