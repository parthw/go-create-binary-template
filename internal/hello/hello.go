package hello

import (
	"fmt"

	"github.com/parthw/go-create-binary-template/internal/logger"
)

func Hello() {
	logger.Log.Debug("logging is working fine!")
	fmt.Print("Hello World!")
	// Teesting git hooks
}
