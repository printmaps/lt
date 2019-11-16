# lt

Package lt offers functions to format log lines and to trace the program flow.
The package has no dependencies.

A good logging concept is to provide human and machine readable logs. In order to achieve this,
it's recommended to log everything as key/value pairs. That is what this package does. It can be
used with the standard Go logger.

## Documentation

[![GoDoc](https://godoc.org/github.com/Klaus-Tockloth/lt?status.svg)](https://godoc.org/github.com/Klaus-Tockloth/lt)

## Installing

```bash
go get github.com/Klaus-Tockloth/lt
```

## Format function

The Format() function formats a series of key-value pair into a string.

``` txt
e.g. name="format", version="v1.0.0"

- key has no surrounding '"'
- value has surrounding '"'
- key and value are separated by '='
- key-value pairs are separated by ', '
```

Format() usage sample:

``` golang
package main

import (
	"errors"
	"log"
	"os"

	"github.com/Klaus-Tockloth/lt"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
	log.SetPrefix("FORMAT ")

	progName := os.Args[0]
	progVersion := "v1.0.0"
	log.Println(lt.Format(lt.OPERATIONS, lt.START, lt.MESSAGE, "lt package Format() sample", "name", progName, "version", progVersion))

	file := "rogue.txt"
	err := errors.New("file not found")
	log.Println(lt.Format(lt.ERROR, lt.FATAL, lt.MESSAGE, "error opening file", "error", err, "file", file))

	id := 42
	fruits := []string{"apple", "banana"}
	log.Println(lt.Format(lt.TRANSACTION, lt.INFO, lt.MESSAGE, "map build ok", "id", id, "fruits", fruits))

	person := struct {
		Name string
		Age  int
	}{"Sean", 50}
	hash := map[string]int{"foo": 11, "bar": 22}
	log.Println(lt.Format(lt.TRANSACTION, lt.INFO, lt.MESSAGE, "processing startet", "person", person, "hash", hash))

	// stringify complex objects e.g. with spew
	log.Println(lt.Format(lt.TRANSACTION, lt.INFO, lt.MESSAGE, "processing startet", "person", spew.Sdump(person), "hash", spew.Sdump(hash)))

	user := 23.69
	sys := 33.64
	idle := 42.65
	log.Println(lt.Format(lt.OPERATIONS, lt.METRIC, "cpu-user", user, "cpu-sys", sys, "cpu-idle", idle))
}
}
```

Format output:

``` txt
FORMAT 2019/11/12 08:15:49.496756 main.go:55: operations="start", message="lt package Format() sample", name="./format", version="v1.0.0"
FORMAT 2019/11/12 08:15:49.496898 main.go:59: error="fatal", message="error opening file", error="file not found", file="rogue.txt"
FORMAT 2019/11/12 08:15:49.496932 main.go:63: transaction="information", message="map build ok", id="42", fruits="[apple banana]"
FORMAT 2019/11/12 08:15:49.496971 main.go:70: transaction="information", message="processing startet", person="{Sean 50}", hash="map[bar:22 foo:11]"
FORMAT 2019/11/12 08:15:49.497138 main.go:73: transaction="information", message="processing startet", person="(struct { Name string; Age int }) {\n Name: (string) (len=4) \"Sean\",\n Age: (int) 50\n}\n", hash="(map[string]int) (len=2) {\n (string) (len=3) \"foo\": (int) 11,\n (string) (len=3) \"bar\": (int) 22\n}\n"
FORMAT 2019/11/12 08:15:49.497168 main.go:78: operations="metric", cpu-user="23.69", cpu-sys="33.64", cpu-idle="42.65"
```

## Trace function

The Trace() function traces entry and exit of a function or method. Simply add

``` golang
defer lt.Trace(logger)(logger)
```

at the beginning of each function you want to trace. The Trace() function writes
for each entry and exit a line into the specified logger. The output includes:

``` txt
trace-entry, function name, source file, line number in source file
trace-exit, duration of execution in milliseconds
```

Trace() usage sample:

``` golang
package main

import (
	"log"
	"os"
	"time"

	"github.com/Klaus-Tockloth/lt"
)

func main() {
	logger := log.New(os.Stderr, "TRACE ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	rogueOne(logger)
	rogueTwo(logger)
}

func rogueOne(logger *log.Logger) {
	defer lt.Trace(logger)(logger)
	time.Sleep(100 * time.Millisecond)
}

func rogueTwo(logger *log.Logger) {
	defer lt.Trace(logger)(logger)
	time.Sleep(200 * time.Millisecond)
	rogueOne(logger)
}
```

Trace output:

``` txt
TRACE 2019/11/12 08:16:31.331325 lt.go:183: trace-entry="main.rogueOne", file="/Users/klaustockloth/go/src/github.com/Klaus-Tockloth/lt/cmd/trace/main.go", line="55"
TRACE 2019/11/12 08:16:31.431774 lt.go:186: trace-exit="main.rogueOne", duration="100.430077ms"
TRACE 2019/11/12 08:16:31.431846 lt.go:183: trace-entry="main.rogueTwo", file="/Users/klaustockloth/go/src/github.com/Klaus-Tockloth/lt/cmd/trace/main.go", line="60"
TRACE 2019/11/12 08:16:31.637097 lt.go:183: trace-entry="main.rogueOne", file="/Users/klaustockloth/go/src/github.com/Klaus-Tockloth/lt/cmd/trace/main.go", line="55"
TRACE 2019/11/12 08:16:31.742361 lt.go:186: trace-exit="main.rogueOne", duration="105.266626ms"
TRACE 2019/11/12 08:16:31.742409 lt.go:186: trace-exit="main.rogueTwo", duration="310.600652ms"
```
