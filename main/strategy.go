package main

type Strategy interface {
	findAction(state *State, player int) *Action
}
