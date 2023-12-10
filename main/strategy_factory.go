package main

func NewStrategyFromArgs(args []string, game Game[State, Action]) Strategy[State, Action] {
	// args 0 = exe path
	// args 1 = first argument
	// etc...
	strategyName := ""
	if len(args) > 1 {
		strategyName = args[1]
	}
	switch strategyName {
	case "error":
		panic("Error arg test!")
	case "mcts":
		return NewMctsStrategy(game)
	case "simple":
		return NewSimpleStrategy(game)
	default:
		return NewMctsStrategy(game)
	}
}
