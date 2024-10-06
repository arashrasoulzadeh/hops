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
			for functionName := range pkg {
				cmd.Println(meta.Functions[functionName] + "\t\t" + meta.Comments[functionName])
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
		a1 := lua.LString("arash") // Modify this based on your use case

		// Define execution loop based on the 'watch' flag
		for {

			// Execute the function
			response, err := engine.ExecuteFunction(f, a1)
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
