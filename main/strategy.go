package main

type Strategy interface {
	findAction(state *State) *Cell
}
