/*
Package asteria is a logging library for go with module support.

	package main

	import(
		"github.com/mylxsw/asteria/log"
	)

	var logger = log.Module("toolkit.process")

	func main() {
		logger.Debugf("xxxx: %s, xxx", "ooo")
		logger.WithFields(log.C{
			"id": 123,
			"name": "lixiaoyao",
		}).Debugf("Hello, %s", "world")
	}

*/
package asteria

