package questions

import (
	"github.com/superbrobenji/advent_of_code/utils"
)

type Questions struct {
	QSlice []func()
}

func (questions *Questions) Execute(qToRun int) {
	utils.Assert(qToRun > 0 || qToRun < len(questions.QSlice), "Invalid index provided")

	if qToRun > 0 {
		questions.QSlice[qToRun-1]()
	} else {
		for i := 0; i < len(questions.QSlice); i++ {
			questions.QSlice[i]()
		}
	}

}
