package engine

import (
	"log"

	lua "github.com/yuin/gopher-lua"
)

func ExecuteFunction(function *lua.LFunction, args ...lua.LValue) (string, error) {
	l := lua.NewState()
	defer l.Close()

	// Define how many return values you expect (e.g., 1 return value)
	nRet := 1

	// Ensure the function is protected and pass the arguments
	if err := l.CallByParam(lua.P{
		Fn:      function,
		NRet:    nRet,
		Protect: true,
	}, args...); err != nil {
		log.Printf("Error executing function: %v\n", err)
		return "", err
	}

	// Get the return value(s) from the stack
	ret := l.Get(-1) // Get the top value of the stack
	l.Pop(1)         // Remove it from the stack

	// Return the result as a string
	if retStr, ok := ret.(lua.LString); ok {
		return string(retStr), nil
	}

	return "", nil
}
