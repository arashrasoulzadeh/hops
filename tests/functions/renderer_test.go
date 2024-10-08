package functions

import (
	"hops/renderer"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {

	r := renderer.NewRenderer()

	result, err := r.Render([]byte("{{os.Name}}"))
	assert.NoError(t, err)

	// hardware
	fm := r.GetRenderer()
	hardwareInfo := fm["hardware"].(func() *renderer.HardwareInfo)() // Call the function to get osInfo
	result, err = r.Render([]byte("{{hardware.arch}}"))
	assert.NoError(t, err)
	assert.Equal(t, string(result), hardwareInfo.Arch())

	//os
	osInfo := fm["os"].(func() *renderer.OsInfo)()
	assert.Equal(t, string(result), osInfo.Name())

	result, err = r.Render([]byte("{{os.Version}}"))
	assert.NoError(t, err)
	assert.Equal(t, string(result), osInfo.Version())

}
