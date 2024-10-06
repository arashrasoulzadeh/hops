# Lua Execution Engine

This Go package allows you to load and execute Lua scripts dynamically from the `modules/` directory. It provides a
flexible way to run predefined Lua functions and handle arguments.

## Features

- **Dynamic Lua Execution**: Load Lua files from the `modules/` directory and execute specific functions on-demand.
- **Customizable Lua Function Arguments**: Pass arguments to Lua functions from Go.
- **Error Handling**: Provides feedback on Lua function execution.

## Installation

To use this package, make sure you have Go installed and set up. You also need to install the dependencies:

```bash
go get github.com/yuin/gopher-lua
go get github.com/spf13/cobra
```

Clone the repository and run the `run` command with your Lua script and function name:

```bash
git clone <repo-url>
cd <repo-folder>
go run . module_name module_function
```

## Usage

1. **Define a Lua script**:

   Write Lua functions in a Lua script and return them as a table. For example, `modules/nginx/nginx.lua`:

   ```lua
   -- Sample Lua script to restart Nginx
   function restart()
       print("Restarting Nginx service...")
       local result = os.execute("sudo systemctl restart nginx")
       if result == 0 then
           print("Nginx restarted successfully.")
       else
           print("Failed to restart Nginx.")
       end
   end

   return {
       restart = restart
   }
   ```

2. **Call the Lua function from Go**:

   Use the package’s command-line interface to load the Lua module and execute the desired function:

   ```bash
   go run .  nginx restart
   ```

   This will execute the `restart` function from the `modules/nginx/nginx.lua` file.

## Example

```bash
go run . nginx restart
```

This command will load the Lua file `modules/nginx/nginx.lua` and execute the `restart` function to restart the Nginx
service.

## Structure

```bash
.
├── cmd/
│   ├── root.go            # Root Cobra command
│   ├── run.go             # Command to run Lua functions
├── engine/
│   ├── engine.go          # Lua execution engine
├── modules/
│   └── nginx/
│       └── nginx.lua      # Example Lua script for Nginx
```

## TODO

- [ ] Add support for passing multiple arguments to Lua functions.
- [ ] Add more examples for different types of Lua scripts.
- [ ] Implement more error handling in Lua execution.
- [ ] Add unit tests for Lua function execution.
- [ ] Create a Dockerfile for easy deployment.
