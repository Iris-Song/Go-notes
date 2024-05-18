package main

import (
	"bytes"
	"fmt"
	"log"
	"log/slog"
	"os"
)

// the log package for free-form output and the log/slog package for structured output.

func main() {
	// Simply invoking functions like Println from the log package uses the standard logger,
	// which is already pre-configured for reasonable logging output to os.Stderr.
	// Additional methods like Fatal* or Panic* will exit the program after logging.
	log.Println("standard loger")

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("with micro")

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("with short file")

	// It may be useful to create a custom logger and pass it around. When creating a new logger,
	// we can set a prefix to distinguish its output from other loggers.
	mylog := log.New(os.Stderr, "MyLog: ", log.LstdFlags)
	mylog.Println("from mylog")

	mylog.SetPrefix("ohmy:")
	mylog.Println("from mylog")

	// Loggers can have custom output targets; any io.Writer works.
	var buf bytes.Buffer
	buflog := log.New(&buf, "buflog: ", log.LstdFlags)
	buflog.Println("hello")
	fmt.Println("from buflog:", buf.String())

	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	myslog := slog.New(jsonHandler)
	myslog.Info("hi there")
	myslog.Info("hello again", "key", "val", "age", 25)
}
