package user

import (
	"io"
	"log"
)

var (
	Logger = log.New(io.Discard, "[user] ", log.LstdFlags)
)
