package main

import (
	"os"

	"github.com/hasssanezzz/gologger/gologger"
)

func main() {
	logger := gologger.New(os.Stderr, "")

	logger.Info("I wanted you to know")
	logger.Debug("This is some debug stuff")
	logger.Warning("I warn you")
	logger.Critical("Critical things are happening")
	logger.Fatal("System shutdown")
}
