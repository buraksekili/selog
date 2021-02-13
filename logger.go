package selog // ch

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Selog struct {
	l      *log.Logger
	prefix string
	output io.Writer
}

func NewLogger(l *log.Logger) *Selog {
	return &Selog{l, strings.Trim(l.Prefix(), " "), l.Writer()}
}

func (s *Selog) Debug(msg string, i ...interface{}) {

	prefix := GetPrefix("DEBUG", s.prefix)
	color.New(color.FgBlue).Fprintf(s.output, prefix+msg, i...)
}

func (s *Selog) Info(msg string, i ...interface{}) {
	prefix := GetPrefix("INFO", s.prefix)
	color.New(color.FgCyan).Fprintf(s.output, prefix+msg, i...)
}

func (s *Selog) Warn(msg string, i ...interface{}) {
	prefix := GetPrefix("WARN", s.prefix)
	color.New(color.FgYellow).Fprintf(s.output, prefix+msg, i...)
}

func (s *Selog) Error(msg string, i ...interface{}) {
	prefix := GetPrefix("ERROR", s.prefix)
	color.New(color.FgRed).Fprintf(s.output, prefix+msg, i...)
}

func (s *Selog) Fatal(msg string, i ...interface{}) {
	prefix := GetPrefix("FATAL", s.prefix)
	color.New(color.FgRed, color.Bold).Fprintf(s.output, prefix+msg, i...)
	os.Exit(1)
}

func (s *Selog) Success(msg string, i ...interface{}) {
	prefix := GetPrefix("SUCCESS", s.prefix)
	color.New(color.FgGreen, color.Bold).Fprintf(s.output, prefix+msg, i...)
}

func GetPrefix(status, loggerPrefix string) string {
	t := time.Now().Format("2006/01/02 Monday 15:04:05")
	prefix := fmt.Sprintf("%v %s: [%s] ", t, loggerPrefix, status)
	return prefix
}
