package questions

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/superbrobenji/advent_of_code/utils"
)

type Q1Data struct {
	list1      []int
	list2      []int
	diff       int
	similarity int
}
type List struct {
	smallest int
	pos      int
}

func Q1() {
	lists := createQ1Lists()

	lists.sortSlices()
	lists.getDiff()
	lists.getSimilarity()

}

func (D *Q1Data) getSimilarity() {

	for _, num := range D.list1 {
		occurances := 0
		for _, num2 := range D.list2 {
			if num2 == num {
				occurances++
			}
		}
		D.similarity += num * occurances
	}
	fmt.Println("The similarity is: ", D.similarity)
}

func (D *Q1Data) getDiff() {
	for i, num := range D.list1 {
		diff := num - D.list2[i]
		if diff < 0 {
			diff = -diff
		}
		D.diff += diff
	}
	fmt.Println("The difference is: ", D.diff)
}

func (D *Q1Data) sortSlices() {
	sort.Slice(D.list1, func(i, j int) bool {
		return D.list1[i] < D.list1[j]
	})
	sort.Slice(D.list2, func(i, j int) bool {
		return D.list2[i] < D.list2[j]
	})
}

func createQ1Lists() (returnData *Q1Data) {
	data, err := utils.FetchQuestionData(1)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	lines := strings.Split(*data, "\n")

	list1 := make([]int, len(lines))
	list2 := make([]int, len(lines))

	for i, line := range lines {
		set := strings.Split(line, " ")
		if len(set) == 1 {
			break
		}

		value1, err := strconv.Atoi(set[0])
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}
		value2, err := strconv.Atoi(set[len(set)-1])
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}
		list1[i] = value1
		list2[i] = value2
	}

	returnData = &Q1Data{
		list1: list1,
		list2: list2,
	}
	return
}
