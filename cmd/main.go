package main

import (
	"github.com/buraksekili/selog"
	"log"
	"os"
)

func main() {
	l := log.New(os.Stdout, "cmd", log.LstdFlags)
	slog := selog.NewLogger(l)
	slog.Success("successfully worked!\n")
	slog.Info("successfully worked!\n")

	// exit status 1 like log.Fatal
	slog.Fatal("successfully worked!\n")
}