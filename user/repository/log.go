package repository

import (
	"io"
	"log"
)

var (
	Logger = log.New(io.Discard, "[user:repository] ", log.LstdFlags)
)
