package cli

import (
	"github.com/gophers-mx/just-go/config"
	"github.com/gophers-mx/just-go/pkg/files"
)

// GeneratorCLI creates a new CLI project.
type GeneratorCLI struct{}

// NewGeneratorCLI retrieves a generator for a CLI project.
func NewGeneratorCLI() *GeneratorCLI {
	return &GeneratorCLI{}
}

// Run creates a CLI project using the embedded assets.
func (lg *GeneratorCLI) Run(cfg *config.Cfg) {
	files.GenFromTemplate("assets", "", "README.md")
	files.GenFromTemplate("assets", "", "go.mod")
	files.GenFromTemplate("assets", "", "main.go")
	files.CopyDirectory("assets/cmd", "cmd")
	files.CopyFile("assets", "", "hello.md")
}
