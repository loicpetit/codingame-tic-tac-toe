package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Cell struct {
	x int
	y int
}

type State struct {
	cells      [][]int
	lastAction Cell
}

func (state *State) GetAvailableCells() []Cell {
	cells := make([]Cell, 0)
	for x, column := range state.cells {
		for y, value := range column {
			if value == 0 {
				cells = append(cells, Cell{x: x, y: y})
			}
		}
	}
	return cells
}

func (state *State) Play(action Cell, bot int) {
	state.lastAction = action
	state.cells[action.x][action.y] = bot
}

func (state *State) Draw() bool {
	for _, column := range state.cells {
		for _, value := range column {
			if value == 0 {
				return false
			}
		}
	}
	return true
}

func (state *State) Winner() int {
	size := len(state.cells)
	winner := 0
	var checkRow = func(y int) int {
		if winner != 0 {
			return winner
		}
		if size == 0 {
			return 0
		}
		player := state.cells[0][y]
		for x := 1; x < size; x++ {
			if player != state.cells[x][y] {
				return 0
			}
		}
		return player
	}
	var checkColumn = func(x int) int {
		if winner != 0 {
			return winner
		}
		if size == 0 {
			return 0
		}
		player := state.cells[x][0]
		for y := 1; y < size; y++ {
			if player != state.cells[x][y] {
				return 0
			}
		}
		return player
	}
	var checkDiagonale1 = func() int {
		if winner != 0 {
			return winner
		}
		if size == 0 {
			return 0
		}
		player := state.cells[0][0]
		for i := 1; i < size; i++ {
			if player != state.cells[i][i] {
				return 0
			}
		}
		return player
	}
	var checkDiagonale2 = func() int {
		if winner != 0 {
			return winner
		}
		if size == 0 {
			return 0
		}
		maxIndex := size - 1
		player := state.cells[maxIndex][0]
		for i := 1; i < size; i++ {
			if player != state.cells[maxIndex-i][0+i] {
				return 0
			}
		}
		return player
	}
	for i := 0; i < size; i++ {
		winner = checkRow(i)
		winner = checkColumn(i)
	}
	winner = checkDiagonale1()
	winner = checkDiagonale2()
	return winner
}

func NewState() *State {
	cells := make([][]int, 3)
	for i := range cells {
		cells[i] = make([]int, 3)
	}
	return &State{cells: cells, lastAction: Cell{x: -1, y: -1}}
}

type Bot struct {
	actions chan Cell
	closed  bool
	cmd     *exec.Cmd
	debugs  chan string
	errors  io.ReadCloser
	inputs  io.WriteCloser
	name    string
	outputs io.ReadCloser
}

func (bot *Bot) Close() {
	bot.closed = true
	bot.inputs.Close()
	bot.outputs.Close()
	bot.errors.Close()
	close(bot.actions)
	close(bot.debugs)
}

func (bot *Bot) Start() {
	go func() {
		fmt.Println(bot.name, "start...")
		botError := bot.cmd.Start()
		if botError != nil {
			fmt.Println(bot.name, "error:", botError)
		}
	}()
	go func() {
		scanner := bufio.NewScanner(bot.errors)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			if bot.closed {
				return
			}
			bot.debugs <- scanner.Text()
		}
	}()
	go func() {
		for {
			action := Cell{}
			fmt.Fscan(bot.outputs, &action.y, &action.x)
			if bot.closed {
				return
			}
			bot.actions <- action
		}
	}()
}

func NewBot(name string, strategy string) *Bot {
	exePath, _ := os.Executable()
	dir := filepath.Dir(exePath)
	cmd := exec.Command("./main.exe", strategy)
	cmd.Dir = dir
	inputWriter, _ := cmd.StdinPipe()
	outputReader, _ := cmd.StdoutPipe()
	errorReader, _ := cmd.StderrPipe()
	return &Bot{
		actions: make(chan Cell, 1),
		closed:  false,
		cmd:     cmd,
		debugs:  make(chan string, 3),
		errors:  errorReader,
		inputs:  inputWriter,
		name:    name,
		outputs: outputReader,
	}
}

// run match between bots
func main() {
	fmt.Println("Run match")
	quit := make(chan bool)
	defer close(quit)
	go startGame(quit)
	<-quit
}

func startGame(quit chan bool) {
	// bots
	bot1 := NewBot("bot1", "")
	bot1.Start()
	defer bot1.Close()
	bot2 := NewBot("bot2", "")
	bot2.Start()
	defer bot2.Close()
	// debugs
	var readDebug = func(bot *Bot, msg string) {
		// fmt.Println(bot.name, "debug:", msg)
		if strings.Contains(msg, "panic") {
			quit <- true
		}
	}
	// winner
	var endGame = func(state *State) {
		if state.Draw() {
			fmt.Println("DRAW")
			quit <- true
		}
		winner := state.Winner()
		if winner != 0 {
			fmt.Println("Winner", winner)
			quit <- true
		}
	}
	// play
	state := NewState()
	go turn(bot1, state.lastAction, state.GetAvailableCells())
	for {
		select {
		case action := <-bot1.actions:
			fmt.Println(bot1.name, "action:", "row", action.y, "col", action.x)
			state.Play(action, 1)
			endGame(state)
			go turn(bot2, state.lastAction, state.GetAvailableCells())
		case action := <-bot2.actions:
			fmt.Println(bot2.name, "action:", "row", action.y, "col", action.x)
			state.Play(action, 2)
			endGame(state)
			go turn(bot1, state.lastAction, state.GetAvailableCells())
		case debug := <-bot1.debugs:
			readDebug(bot1, debug)
		case debug := <-bot2.debugs:
			readDebug(bot2, debug)
		}
	}
}

func turn(bot *Bot, opponentAction Cell, availableCells []Cell) {
	opponentString := fmt.Sprintf("%d %d\n", opponentAction.y, opponentAction.x) // opponent move row col (y, x)
	// fmt.Print(bot.name, " ", opponentString)
	io.WriteString(bot.inputs, opponentString)
	nbAvailableCellsString := fmt.Sprintf("%d\n", len(availableCells)) // nb available actions
	// fmt.Print(bot.name, " ", nbAvailableCellsString)
	io.WriteString(bot.inputs, nbAvailableCellsString)
	for _, cell := range availableCells {
		availableActionString := fmt.Sprintf("%d %d\n", cell.y, cell.x) // available action, row col (y, x)
		// fmt.Print(bot.name, " ", availableActionString)
		io.WriteString(bot.inputs, availableActionString)
	}
}
