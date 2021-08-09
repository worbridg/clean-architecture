package context

import (
	"io"
	"log"
)

var (
	Logger = log.New(io.Discard, "[user:context] ", log.LstdFlags)
)
