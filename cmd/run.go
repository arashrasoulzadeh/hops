/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/

package cmd

import (
    "fmt"
    lua "github.com/yuin/gopher-lua"
    "hops/engine"
    "github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
    Use:   "run [module] [function]",
    Short: "Run a module command",
    Long:  `This command will run a specified function from a given module.`,
    Run: func(cmd *cobra.Command, args []string) {
        if len(args) < 2 {
            fmt.Println("Error: module and function name are required.")
            return
        }

        moduleName := args[0]
        functionName := args[1]

        // Check if the module exists
        pkg, ok := engine.FunctionMap[moduleName]
        if !ok {
            fmt.Printf("Error: module '%s' not found.\n", moduleName)
            return
        }

        // Check if the function exists within the module
        f, ok := pkg[functionName]
        if !ok {
            fmt.Printf("Error: function '%s' not found in module '%s'.\n", functionName, moduleName)
            return
        }

        // Create a Lua value for the function argument
        a1 := lua.LString("arash") // Modify this based on your use case

        // Execute the function
        err := engine.ExecuteFunction(f, a1)
        if err != nil {
            fmt.Printf("Error executing function: %v\n", err)
        }
    },
}

func init() {
    rootCmd.AddCommand(runCmd)

    // Define any flags and configuration settings here
    // runCmd.PersistentFlags().String("foo", "", "A help for foo")
    // runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
