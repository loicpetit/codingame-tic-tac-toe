package main

import (
	"time"
)

type Strategy[STATE Hashable, ACTION any] interface {
	findAction(state *STATE, player int, maxTime time.Time) *ACTION
}
