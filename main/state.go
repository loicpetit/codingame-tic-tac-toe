package main

type State struct {
	width            int
	height           int
	opponent         *Cell
	availableActions []*Cell
}

func (state *State) SetWidth(width int) *State {
	if state == nil {
		state = NewState()
	}
	return &State{
		width:            width,
		height:           state.height,
		opponent:         state.opponent,
		availableActions: state.availableActions,
	}
}

func (state *State) SetHeight(height int) *State {
	if state == nil {
		state = NewState()
	}
	return &State{
		width:            state.width,
		height:           height,
		opponent:         state.opponent,
		availableActions: state.availableActions,
	}
}

func (state *State) SetOpponent(opponent *Cell) *State {
	if state == nil {
		state = NewState()
	}
	return &State{
		width:            state.width,
		height:           state.height,
		opponent:         opponent,
		availableActions: state.availableActions,
	}
}

func (state *State) SetAvailableActions(availableActions []*Cell) *State {
	if state == nil {
		state = NewState()
	}
	return &State{
		width:            state.width,
		height:           state.height,
		opponent:         state.opponent,
		availableActions: availableActions,
	}
}

func NewState() *State {
	return &State{
		width:  0,
		height: 0,
	}
}
