package hello

import (
	"fmt"

	"github.com/parthw/go-create-binary-template/internal/logger"
)

func Hello() {
	logger.Log.Info("logging is working fine!")
	fmt.Print("Hello World!")
}
