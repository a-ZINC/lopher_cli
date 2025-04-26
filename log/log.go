package log

import (
	"log"
	"os"
)

var (
	InfoLogger  = log.New(os.Stdout, "[INFO] ", log.LstdFlags)
	ErrorLogger = log.New(os.Stderr, "[ERROR] ", log.LstdFlags)
	WarnLogger  = log.New(os.Stdout, "[WARB] ", log.LstdFlags)
)
