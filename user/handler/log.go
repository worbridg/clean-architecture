package handler

import (
	"io"
	"log"
)

var (
	Logger = log.New(io.Discard, "[user:handler] ", log.LstdFlags)
)
