/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"embed"
	"fmt"
	"hops/cmd"
	"hops/engine"
	"log"
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//go:embed modules/**/*.lua
var fs embed.FS

func main() {
	// Load Lua paths from the embedded filesystem
	args := os.Args
	err := loadAllModules(args[1], fs, "modules")
	if err != nil {
		log.Fatal(err)
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

func loadAllModules(module string, fs embed.FS, modulesDir string) error {

	// Walk through the modules directory
	err := filepath.Walk(modulesDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Only process directories (skip files)
		if info.IsDir() {
			if path != modulesDir && path == "modules/"+module {
				// Load the Lua path for the current module (directory)
				if loadErr := engine.LoadPath(fs, path); loadErr != nil {
					fmt.Printf("Error loading module %s: %v\n", path, loadErr)
				}
			}
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("error walking through modules directory: %w", err)
	}

	return nil
}
