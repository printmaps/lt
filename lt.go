/*
Purpose:
- See package description in doc.go.

Releases:
- v1.0.0 - 2019/11/12 : initial release
- v2.0.0 - 2019/11/15 : non essential Pretty() function removed

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

Remarks:
- NN

Links:
- NN
*/

package lt

import (
	"fmt"
	"log"
	"runtime"
	"strings"
	"time"
)

// Logging main categories (for convenience only).
const (
	TRANSACTION string = "transaction"
	OPERATIONS  string = "operations"
	ERROR       string = "error"
)

// Logging subcategory transactions (for convenience only).
const (
	REQUEST  string = "request"
	RESPONSE string = "response"
	INFO     string = "information"
)

// Loggging subcategory operations (for convenience only).
const (
	START  string = "start"
	STOP   string = "stop"
	CONFIG string = "configuration"
	METRIC string = "metric"
	HEALTH string = "health"
)

// Logging subcategory errors (for convenience only).
const (
	NOTCLASSIFIED string = "not-classified"
	FATAL         string = "fatal"
	TECHNICAL     string = "technical"
	FUNCTIONAL    string = "functional"
)

// General log data (for convenience only).
const (
	MESSAGE string = "message"
	DATA    string = "data"
	TEXT    string = "text"
)

/*
Format formats all log elements in key-value form.
*/
func Format(kvPairs ...interface{}) string {

	logElements := []string{}

	if len(kvPairs)%2 != 0 {
		// error: key or value missing
		logElements = append(logElements, "internal-log-format-error=\"key-or-value-missing\"")
	} else {
		for index := 0; index < len(kvPairs); index += 2 {
			logElements = append(logElements, fmt.Sprintf("%v=%q", kvPairs[index], fmt.Sprintf("%v", kvPairs[index+1])))
		}
	}

	return strings.Join(logElements, ", ")
}

/*
Trace traces entry and exit of a function.
*/
func Trace(tracelogger *log.Logger) func(tracelogger *log.Logger) {

	callerName := "NO_CALLER"
	callerFile := "NO_FILE"
	callerLine := -1

	fpcs := make([]uintptr, 1)

	// get caller parameters (skip two levels)
	n := runtime.Callers(2, fpcs)
	if n != 0 {
		caller := runtime.FuncForPC(fpcs[0] - 1)
		if caller == nil {
			callerName = "NIL_CALLER"
		} else {
			callerName = caller.Name()
			callerFile, callerLine = caller.FileLine(fpcs[0] - 1)
		}
	}

	start := time.Now()
	tracelogger.Println(Format("trace-entry", callerName, "file", callerFile, "line", callerLine))

	return func(*log.Logger) {
		tracelogger.Println(Format("trace-exit", callerName, "duration", time.Since(start)))
	}
}
