package main

import (
	log "github.com/alecthomas/log4go"
	"testing"
)

func TestName(t *testing.T) {
	s := "string value"
	log.Info("test: %s", s)
}
