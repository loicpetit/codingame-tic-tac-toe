package main

type Strategy[STATE any, ACTION any] interface {
	findAction(state *STATE, player int) *ACTION
}
