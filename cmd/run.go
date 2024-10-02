package cmd

import (
	"fmt"
	"hops/engine"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/spf13/cobra"
	lua "github.com/yuin/gopher-lua"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run [module] [function]",
	Short: "Run a module command",
	Long:  `This command will run a specified function from a given module.`,
	Run: func(cmd *cobra.Command, args []string) {

		// Retrieve the value of the watch flag
		watch, err := cmd.Flags().GetBool("watch")
		if err != nil {
			fmt.Printf("Error retrieving watch flag: %v\n", err)
			return
		}

		if len(args) < 1 || len(args) < 2 {
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

		// Define execution loop based on the 'watch' flag
		for {

			// Execute the function
			err = engine.ExecuteFunction(f, a1)
			if err != nil {
				fmt.Printf("Error executing function: %v\n", err)
			}

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
	rootCmd.AddCommand(runCmd)

	// Define any flags and configuration settings here
	runCmd.Flags().BoolP("watch", "w", false, "Enable watch mode for continuous execution with terminal clearing")
}
