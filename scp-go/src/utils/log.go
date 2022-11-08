package utils

import (
	"log"
	"os"
)

var LERR *log.Logger = log.New(os.Stderr, log.Prefix(), 0)
var LOUT *log.Logger = log.New(os.Stdout, log.Prefix(), 0)
