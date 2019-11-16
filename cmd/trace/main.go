/*
Purpose:
- lt package Trace() sample

Description:
- This program shows the usage of the lt.Trace() function.

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

// Shows the usage of the lt.Trace() function.
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
