package main

import "github.com/skycatch/tasque-go/result"

// MessageHandler hello world
type MessageHandler interface {
	id() *string
	body() *string
	initialize()
	receive() bool
	success()
	failure(err result.Result)
	heartbeat()
}
