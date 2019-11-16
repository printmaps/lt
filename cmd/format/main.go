/*
Purpose:
- lt package Format() sample

Description:
- This program shows the usage of the lt.Format() function.

Releases:
- v1.0.0 - 2019/11/12 : initial release

Author:
- Klaus Tockloth

Copyright and license:
- Copyright (c) 2019 Klaus Tockloth
- MIT license

Permission is hereby granted, free of charge, to any person obtaining a copy of this software
and associated documentation files (the Software), to deal in the Software without restriction,
including without limitation the rights to use, copy, modify, merge, publish, distribute,
sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or
substantial portions of the Software.

The software is provided 'as is', without warranty of any kind, express or implied, including
but not limited to the warranties of merchantability, fitness for a particular purpose and
noninfringement. In no event shall the authors or copyright holders be liable for any claim,
damages or other liability, whether in an action of contract, tort or otherwise, arising from,
out of or in connection with the software or the use or other dealings in the software.

Contact (eMail):
- freizeitkarte@googlemail.com
*/

// Shows the usage of the lt.Format() function.
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
