package lt

import (
	"log"
	"os"
	"time"
)

func ExampleFormat() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
	log.SetPrefix("FORMAT ")

	progName := os.Args[0]
	progVersion := "v1.0.0"
	log.Println(Format(OPERATIONS, START, MESSAGE, "lt package Format() sample", "name", progName, "version", progVersion))
}

func ExampleTrace() {
	logger := log.New(os.Stderr, "TRACE ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)

	rogueOne(logger)
	rogueTwo(logger)
}

func rogueOne(logger *log.Logger) {
	defer Trace(logger)(logger)
	time.Sleep(123 * time.Millisecond)
}

func rogueTwo(logger *log.Logger) {
	defer Trace(logger)(logger)
	time.Sleep(456 * time.Millisecond)
	rogueOne(logger)
}
