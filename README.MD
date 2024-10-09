Lua Execution Engine
====================

[![Build Status](https://github.com/arashrasoulzadeh/hops/actions/workflows/go.yml/badge.svg)](https://github.com/arashrasoulzadeh/hops/actions/workflows/go.yml)

This Go package allows you to load and execute Lua scripts dynamically from the `modules/` directory. It provides a flexible way to run predefined Lua functions and handle arguments.

Features
--------

*   **Dynamic Lua Execution**: Load Lua files from the `modules/` directory and execute specific functions on-demand.
*   **Customizable Lua Function Arguments**: Pass arguments to Lua functions from Go or prompt for input if not provided.
*   **Error Handling**: Provides feedback on Lua function execution.

Installation
------------

To use this package, make sure you have Go installed and set up. You also need to install the dependencies:

    go get github.com/yuin/gopher-lua
    go get github.com/spf13/cobra


Clone the repository and run the `run` command with your Lua script and function name:

    git clone https://github.com/arashrasoulzadeh/hops.git
    cd hops
    go run . module_name module_function


Usage
-----

1.  **Define a Lua script**:

Write Lua functions in a Lua script and return them as a table. For example, `modules/nginx/nginx.lua`:

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


2.  **Call the Lua function from Go**:

Use the package’s command-line interface to load the Lua module and execute the desired function:

    go run . nginx restart


This will execute the `restart` function from the `modules/nginx/nginx.lua` file.

3.  **Pass arguments**:

If the Lua function requires arguments, pass them as follows. Arguments after the second one are passed to the Lua function. For example, given a Lua function `wordpress(db_engine, db_user, db_password, db_name)`:

    function wordpress(db_engine, db_user, db_password, db_name)
        print("Setting up WordPress with:")
        print("DB Engine: " .. db_engine)
        print("DB User: " .. db_user)
        print("DB Password: " .. db_password)
        print("DB Name: " .. db_name)
    end
    
    return {
        wordpress = wordpress
    }


You can run this Lua function and pass arguments like so:

    go run . wordpress mysql root password mydb


If any argument is missing, the program will prompt for the missing values interactively.

4.  **Interactive input**:

If arguments are not provided, the program will ask for them interactively:

    go run . wordpress
    for db_engine: mysql
    for db_user: root
    for db_password: password
    for db_name: mydb


Example
-------

    go run . nginx restart
    go run . wordpress mysql root password mydb


The first command will load the Lua file `modules/nginx/nginx.lua` and execute the `restart` function to restart the Nginx service. The second command will execute the `wordpress` function with the provided arguments.

Structure
---------

    .
    ├── modules/
    │   └── nginx/
    │       └── nginx.lua      # Example Lua script for Nginx


TODO
----

*   \[x\] Add support for passing multiple arguments to Lua functions.
*   \[ \] Add more examples for different types of Lua scripts.
*   \[x\] Implement more error handling in Lua execution.
*   \[ \] Add unit tests for Lua function execution.
*   \[ \] Create a Dockerfile for easy deployment.