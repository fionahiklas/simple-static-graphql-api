package main

import (
	"github.com/sirupsen/logrus"
)

// These values are set by "-X" options to the ldflags for Go, see the Makefile
var (
	commitHash  = "null"
	codeVersion = "null"
)

func main() {
	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)
}
