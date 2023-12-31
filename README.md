# codingame-tic-tac-toe
Ultimate tic tac toe bot https://www.codingame.com/ide/puzzle/tic-tac-toe

## Build unique file
Install the bundle tool (once)

`go install golang.org/x/tools/cmd/bundle@latest`

Bundle the main package

```PowerShell
& ($env:USERPROFILE + "/go/bin/bundle") -o dist/main.go -dst ./main -prefix '""'  github.com/loicpetit/codingame-tic-tac-toe/main
```

## Get lastest go framework for codingame
Generate the dependency in one file in sources

```PowerShell
go get github.com/loicpetit/codingame-go@main
& ($env:USERPROFILE + "/go/bin/bundle") -o main/runner.go -dst ./main -prefix '""' github.com/loicpetit/codingame-go/runner
& ($env:USERPROFILE + "/go/bin/bundle") -o main/mcts.go -dst ./main -prefix '""' github.com/loicpetit/codingame-go/mcts
& ($env:USERPROFILE + "/go/bin/bundle") -o main/test_helper.go -dst ./main -prefix '""' github.com/loicpetit/codingame-go/test
# previous go command update cache and add the mod file but we dont need it
go mod tidy
```

## Test
In project root `go test -v ./main`

## Benchmark
Debug, execute `go test -v -run nothing -benchtime 1000x -bench Debug ./main`

## Generate documention
In project root `go doc -cmd -u -all main > dist/main.txt`

## Generate main executbale
In project root `go build -o ./dist/main.exe ./main`

## Generate match executbale
In project root `go build -o ./dist/match.exe ./match`

## Run match (from root)
```PowerShell
go build -o ./dist/main.exe ./main
go build -o ./dist/match.exe ./match
.\dist\match.exe
```

## VS code tasks

```JSON
{
    "label": "bundle main",
    "type": "shell",
    "options": {
        "cwd": "${fileWorkspaceFolder}"
    },
    "presentation": {
        "clear": true
    },
    "command": "\"& ($env:USERPROFILE + '/go/bin/bundle') -o dist/main.go -dst ./main -prefix '\\\"\\\"' \"\"$(go list -m)/main\"\"\"",
    "problemMatcher": []
},
{
    "label": "build main",
    "type": "shell",
    "options": {
        "cwd": "${fileWorkspaceFolder}"
    },
    "presentation": {
        "clear": true
    },
    "command": "go build -o ./dist ./main",
    "problemMatcher": []
},
{
    "label": "test main",
    "type": "shell",
    "options": {
        "cwd": "${fileWorkspaceFolder}"
    },
    "presentation": {
        "clear": true
    },
    "command": "go test -v ./main",
    "problemMatcher": []
},
{
    "label": "doc main",
    "type": "shell",
    "options": {
        "cwd": "${fileWorkspaceFolder}"
    },
    "presentation": {
        "clear": true
    },
    "command": "go doc -cmd -u -all main > dist/main.txt",
    "problemMatcher": []
}
```
