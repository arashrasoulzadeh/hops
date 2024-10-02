package cmd

import (
	"fmt"
	"hops/engine"
	"strings"

	"github.com/spf13/cobra"
	lua "github.com/yuin/gopher-lua"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run [module] [function]",
	Short: "Run a module command",
	Long:  `This command will run a specified function from a given module.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 && len(args) < 2 {
			fmt.Println("Usage:\nhops run [module] [command]\n\nAvailable Modules:")
			for key := range engine.FunctionMap {
				fmt.Println("  " + strings.Replace(key, "modules/", "", 1))
			}
			return
		}

		moduleName := args[0]

		path := "modules/" + moduleName
		// Check if the module exists
		pkg, ok := engine.FunctionMap[path]
		if !ok {
			fmt.Printf("Error: module '%s' not found.\n", moduleName)
			return
		}
		meta, ok := engine.LuaMetaMap["modules/"+moduleName]
		if !ok {
			fmt.Printf("Error: module '%s' not found.\n", moduleName)
			return
		}

		// If the function name is not provided, list available functions
		if len(args) < 2 {
			fmt.Printf("Usage:\nhops run %s [command]\n\nAvailable Commands:\n", moduleName)

			for functionName := range pkg {
				fmt.Println(meta.Functions[functionName] + "\t\t" + meta.Comments[functionName])
			}
			return
		}

		functionName := args[1]

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
