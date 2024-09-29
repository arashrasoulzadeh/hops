package main

import (
    "log"
    lua "github.com/yuin/gopher-lua"
)

// executeFunction executes a Lua function
func executeFunction(function *lua.LFunction) {
    l := lua.NewState()
    defer l.Close()

    // Call the Lua function
    if err := l.CallByParam(lua.P{
        Fn:      function,
        NRet:    0,
        Protect: true,
    }); err != nil {
        log.Fatalf("Error executing function: %v\n", err)
    }
}
