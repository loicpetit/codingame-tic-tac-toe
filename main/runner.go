// Code generated by golang.org/x/tools/cmd/bundle. DO NOT EDIT.
//go:generate bundle -o main/runner.go -dst ./main -prefix "" github.com/loicpetit/codingame-go/runner

package main

import (
	"fmt"
	"os"
	"time"
)

type Game[STATE any, ACTION any] interface {
	Start() *STATE
	GetAvailableActions(state *STATE, player int) []*ACTION
	GetLastPlayer(state *STATE) int // 0 if any player played
	GetNextPlayer(state *STATE) int // minimum 1
	Play(state *STATE, action *ACTION) *STATE
	Winner(state *STATE) int // 0 = no winner, 1,2,etc the winner
}

type InputReader[INPUT any, STATE any, ACTION any] interface {
	Read() chan INPUT
	UpdateState(state *STATE, input INPUT) *STATE
	ValidateAction(action *ACTION, input INPUT)
}

type OutputWriter[ACTION any] interface {
	Write(action *ACTION)
}

func WriteOutput(a ...any) {
	fmt.Println(a...)
}

func WriteDebug(a ...any) {
	fmt.Fprintln(os.Stderr, a...)
}

type Runner[INPUT any, STATE any, ACTION any] struct {
	game     Game[STATE, ACTION]
	reader   InputReader[INPUT, STATE, ACTION]
	strategy Strategy[STATE, ACTION]
	writer   OutputWriter[ACTION]
}

func (runner Runner[INPUT, STATE, ACTION]) Run(quit chan bool) *STATE {
	timer := NewTimer()
	inputs := runner.reader.Read()
	state := runner.game.Start()
	WriteDebug("State:", state)
	round := 0
	for {
		select {
		case input := <-inputs:
			round++
			maxTime := startTimer(timer, round)
			// opponent turn
			state = runner.reader.UpdateState(state, input)
			if runner.game.Winner(state) > 0 {
				return state
			}
			// my turn
			nextAction := runner.strategy.FindAction(state, 1, maxTime)
			runner.reader.ValidateAction(nextAction, input)
			state = runner.game.Play(state, nextAction)
			WriteDebug("State:", state)
			runner.writer.Write(nextAction)
			endTimer(timer, round)
			WriteDebug("Timer:", timer)
			if runner.game.Winner(state) > 0 {
				return state
			}
		case <-quit:
			return state
		}
	}
}

func startTimer(timer *Timer, round int) time.Time {
	if round == 1 {
		timer.StartInit()
		return timer.initStart.Add(980 * time.Millisecond)
	} else {
		timer.StartRound()
		return timer.roundStart.Add(80 * time.Millisecond)
	}
}

func endTimer(timer *Timer, round int) {
	if round == 1 {
		timer.EndInit()
	} else {
		timer.EndRound()
	}
}

func NewRunner[INPUT any, STATE any, ACTION any](
	game Game[STATE, ACTION],
	reader InputReader[INPUT, STATE, ACTION],
	strategy Strategy[STATE, ACTION],
	writer OutputWriter[ACTION],
) Runner[INPUT, STATE, ACTION] {
	return Runner[INPUT, STATE, ACTION]{
		game,
		reader,
		strategy,
		writer,
	}
}

type Strategy[STATE any, ACTION any] interface {
	FindAction(state *STATE, player int, maxTime time.Time) *ACTION
}

type SimpleStrategy[STATE any, ACTION any] struct {
	game Game[STATE, ACTION]
}

func (strategy SimpleStrategy[STATE, ACTION]) FindAction(state *STATE, player int, maxTime time.Time) *ACTION {
	if state == nil {
		panic("State cannot be nil")
	}
	availableActions := strategy.game.GetAvailableActions(state, player)
	for maxTime.After(time.Now()) {
		time.Sleep(5 * time.Millisecond)
	}
	if len(availableActions) == 0 {
		return nil
	}
	return availableActions[0]
}

func NewSimpleStrategy[STATE any, ACTION any](game Game[STATE, ACTION]) Strategy[STATE, ACTION] {
	return SimpleStrategy[STATE, ACTION]{game: game}
}

type Timer struct {
	initStart      time.Time
	initDuration   time.Duration
	roundStart     time.Time
	roundDuration  time.Duration
	roundMin       time.Duration
	roundMax       time.Duration
	roundDurations []time.Duration
}

func (timer *Timer) StartInit() {
	if timer == nil {
		return
	}
	timer.initStart = time.Now()
	timer.initDuration = 0
}

func (timer *Timer) EndInit() {
	if timer == nil {
		return
	}
	timer.initDuration = time.Since(timer.initStart)
}

func (timer *Timer) StartRound() {
	if timer == nil {
		return
	}
	timer.roundStart = time.Now()
	timer.roundDuration = 0
}

func (timer *Timer) EndRound() {
	if timer == nil {
		return
	}
	timer.roundDuration = time.Since(timer.roundStart)
	timer.roundDurations = append(timer.roundDurations, timer.roundDuration)
	if timer.roundMin == 0 || timer.roundDuration < timer.roundMin {
		timer.roundMin = timer.roundDuration
	}
	if timer.roundDuration > timer.roundMax {
		timer.roundMax = timer.roundDuration
	}
}

func (timer *Timer) roundAverage() time.Duration {
	if timer == nil || len(timer.roundDurations) == 0 {
		return 0
	}
	durations := time.Duration(0)
	for _, timerDuration := range timer.roundDurations {
		durations += timerDuration
	}
	return durations / time.Duration(len(timer.roundDurations))
}

func (timer *Timer) String() string {
	return fmt.Sprintf("{init: %v, round: %v, min: %v, max: %v, average: %v}", timer.initDuration, timer.roundDuration, timer.roundMin, timer.roundMax, timer.roundAverage())
}

func NewTimer() *Timer {
	return &Timer{}
}
