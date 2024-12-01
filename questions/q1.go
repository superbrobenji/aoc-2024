package questions

import (
	"fmt"
	"os"

	"github.com/superbrobenji/advent_of_code/utils"
)

func Q1() {
	data, err := utils.FetchQuestionData(1)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	fmt.Println(*data)
}
