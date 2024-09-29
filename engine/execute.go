package engine

import (
    "log"
    lua "github.com/yuin/gopher-lua"
)

func ExecuteFunction(function *lua.LFunction, args ...lua.LValue) error {
    l := lua.NewState()
    defer l.Close()

    // Ensure the function is protected and pass the arguments
    if err := l.CallByParam(lua.P{
        Fn:      function,
        NRet:    0,
        Protect: true,
    }, args...); err != nil {
        log.Fatalf("Error executing function: %v\n", err)
        return err
    }
    return nil
}
