package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/superbrobenji/advent_of_code/questions"
)

type AOF struct {
	questions *questions.Questions
}

func main() {
	aof := AOF{
		questions: &questions.Questions{
			QSlice: []func(){
				questions.Q1,
			},
		},
	}
	qToRun := aof.terminalInputHandler()

	aof.questions.Execute(qToRun)

}

func (aof *AOF) terminalInputHandler() (qToRun int) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("which question do you want to run? or type \"0\" to run all of them: ")
	textWithNewLine, err := reader.ReadString('\n')
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	text := strings.TrimSuffix(textWithNewLine, "\n")
	qToRun, err = strconv.Atoi(text)
	if err != nil {
		fmt.Println("please provide a valid input")
		qToRun = aof.terminalInputHandler()
	}

	if text == "" || qToRun < 0 || len(aof.questions.QSlice) < qToRun {
		fmt.Println("please enter a valid question")
		qToRun = aof.terminalInputHandler()
	}
	return
}
