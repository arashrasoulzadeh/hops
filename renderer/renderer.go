package renderer

import (
	"bytes"
	"text/template"
)

// Renderer interface defines the methods required for rendering templates.
type Renderer interface {
	Render(content []byte) ([]byte, error)
}

// DefaultRenderer is the default implementation of the Renderer interface.
type DefaultRenderer struct {
	funcMap template.FuncMap
}

// NewRenderer initializes and returns a new DefaultRenderer with preloaded functions.
func NewRenderer() Renderer {
	r := &DefaultRenderer{
		funcMap: template.FuncMap{
			"os":       func() *osInfo { return newOSInfo() },
			"hardware": func() *HardwareInfo { return NewHardwareInfo() },
			"network":  func() *NetworkInfo { return NewNetworkInfo() },
			"user":     func() *UserInfo { return NewUserInfo() },
		},
	}
	return r
}

// Render applies all registered functions and renders the template.
func (r *DefaultRenderer) Render(content []byte) ([]byte, error) {
	// Render the final template after any replacements
	finalResult, err := r.renderTemplate(content)
	if err != nil {
		return nil, err
	}

	return finalResult, nil
}

// renderTemplate renders the final template using the registered function map.
func (r *DefaultRenderer) renderTemplate(content []byte) ([]byte, error) {
	// Create a template with the provided function map
	tmpl, err := template.New("render").Funcs(r.funcMap).Parse(string(content))
	if err != nil {
		return nil, err
	}

	// Render the template
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, nil); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
