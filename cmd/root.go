package cmd

import (
	"hops/engine"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/spf13/cobra"
	lua "github.com/yuin/gopher-lua"
	"go.uber.org/zap"
)

// rootCmd represents the base command
var rootCmd = &cobra.Command{
	Use:   "[module] [function]",
	Short: "Run a module command",
	Long:  `This command will run a specified function from a given module.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Retrieve the value of the watch flag
		watch, err := cmd.Flags().GetBool("watch")
		if err != nil {
			cmd.PrintErrf("Error retrieving watch flag: %v\n", err)
			return
		}

		if len(args) == 0 {
			cmd.Println("Usage:\nhops [module] [command]\n\nAvailable Modules:")
			for key := range engine.FunctionMap {
				cmd.Println("  " + strings.Replace(key, "modules/", "", 1))
			}
			return
		}

		moduleName := args[0]
		path := "modules/" + moduleName

		// Check if the module exists
		pkg, ok := engine.FunctionMap[path]
		if !ok {
			cmd.Printf("Error: module '%s' not found.\n", moduleName)
			return
		}
		meta, ok := engine.LuaMetaMap["modules/"+moduleName]
		if !ok {
			cmd.Printf("Error: module '%s' metadata not found.\n", moduleName)
			return
		}

		// If the function name is not provided, list available functions
		if len(args) == 1 {
			cmd.Printf("Usage:\nhops %s [command]\n\nAvailable Commands:\n", moduleName)

			// Define the width for the command and description
			const commandWidth = 25

			for functionName := range pkg {
				// Format the output to align the command and description
				cmd.Printf("%-*s\t%s\n", commandWidth, meta.Functions[functionName], meta.Comments[functionName])
			}
			return
		}

		functionName := args[1]

		// Check if the function exists within the module

		f, ok := pkg[functionName]
		if !ok {
			cmd.Printf("Error: function '%s' not found in module '%s'.\n", functionName, moduleName)
			return
		}

		// Create a Lua value for the function argument
		argsToPass := make([]lua.LValue, len(meta.Variables[functionName]))
		variables, hasVariables := meta.Variables[functionName]

		if hasVariables && len(variables) > 0 && variables[0] != "" {
			// Loop through the variables for the function and map arguments (or prompt if not enough args)
			for i, k := range variables {
				// If an argument exists (i.e., i + 2 is within bounds of args), use it; otherwise, prompt the user
				if len(args) > 2 {
					argsToPass[i] = lua.LString(args[i+2])
				} else {
					// If not enough arguments are passed, prompt the user interactively using Scan
					argsToPass[i] = lua.LString(engine.Scan(cmd, args, k))
				}
			}
		}
		// Define execution loop based on the 'watch' flag
		for {

			// Execute the function
			response, err := engine.ExecuteFunction(f, argsToPass...)
			if err != nil {
				cmd.PrintErrf("Error executing function: %v\n", err)
			}

			zap.L().Debug(response)

			// If 'watch' mode is not enabled, break after the first run
			if !watch {
				break
			}

			// Add a delay to avoid tight loops in watch mode
			time.Sleep(1 * time.Second) // Adjust the delay as needed
			// If in watch mode, clear the terminal before each run
			if watch {
				clearTerminal()
			}
		}
	},
}

// clearTerminal clears the terminal screen in a cross-platform manner
func clearTerminal() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func init() {
	// Adding the watch flag to the root command
	rootCmd.Flags().BoolP("watch", "w", false, "Enable watch mode for continuous execution with terminal clearing")
}
