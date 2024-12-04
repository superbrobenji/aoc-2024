package questions

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/superbrobenji/advent_of_code/utils"
)

type Q2Data struct {
	reports [][]int
	safe    int
}

func Q2() {
	data := createQ2List()
	data.checkSafetey()
	data.checkSafeteyWithRetry()
}

func (D *Q2Data) checkSafeteyWithRetry() {
	for i, report := range D.reports {
		fmt.Println("testing report number ", i)
		fmt.Println("testing: ", report)
		safe := testReportSafety(report, true)
		if safe {
			D.safe += 1
		}
		fmt.Println("safe rerports: ", D.safe)
	}
	fmt.Println("number of safe reports with retries: ", D.safe)
}

func (D *Q2Data) checkSafetey() {
	for _, report := range D.reports {
		safe := testReportSafety(report, false)
		if safe {
			D.safe += 1
		}
	}
	fmt.Println("number of safe reports: ", D.safe)
}

func testReportSafety(report []int, withRetry bool) (safe bool) {
	safe = true
	var unsafePairs []int
	retested := !withRetry
	//loop entries
	for y, entry := range report {
		tested := false
		if len(report)-1 > y {
			nextEntry := report[y+1]
			diff := entry - nextEntry

			//if not the first item
			if y != 0 {
				prevEntry := report[y-1]

				//if decending cehck if everything is decending
				if prevEntry > entry {
					if entry < nextEntry {
						if !retested {
							if !contains(unsafePairs, y) {
								unsafePairs = append(unsafePairs, y)
							}
						}
						tested = true
						safe = false
					}
				}

				//if ascending cehck if everything is ascending
				if prevEntry < entry {
					if entry > nextEntry && !tested {
						if !retested {
							if !contains(unsafePairs, y) {
								unsafePairs = append(unsafePairs, y)
							}
						}
						tested = true
						safe = false
					}
				}
			}

			//if negative, make positive
			if diff < 0 {
				diff = -diff
			}

			//if diff bigger than 3 or smaller than 1 and diff = 1 make unsafe
			if (diff < 1 || diff > 3 || diff == 0) && !tested {
				if !retested {
					if !contains(unsafePairs, y) {
						unsafePairs = append(unsafePairs, y)
					}
				}
				tested = true
				safe = false
			}
		}
	}

	//retry system
	if !safe && !retested {
		retested = true

		//loop all positions of errors that occured on report
		for _, position := range unsafePairs {
			fmt.Println("number of errors in report: ", unsafePairs)
			originalReport := report

			//rerun test with 1 removed
			fmt.Println("testing: ", remove(originalReport, position))
			safe = testReportSafety(remove(originalReport, position), false)
			if safe {
				break
			}
		}
	}
	return
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func remove(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice
	}

	newSlice := make([]int, 0, len(slice)-1)
	newSlice = append(newSlice, slice[:index]...)
	newSlice = append(newSlice, slice[index+1:]...)

	return newSlice
}

func createQ2List() (returnData *Q2Data) {
	data, err := utils.FetchQuestionData(2)
	fmt.Println(*data)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	lines := strings.Split(*data, "\n")

	returnData = &Q2Data{
		reports: make([][]int, len(lines)-1),
	}

	for i, line := range lines {
		numbers := strings.Split(line, " ")
		if len(numbers) == 1 {
			break
		}
		lineSlice := make([]int, len(numbers))

		for y, number := range numbers {
			value, err := strconv.Atoi(number)
			if err != nil {
				fmt.Print(err.Error())
				os.Exit(1)
			}
			lineSlice[y] = value
		}
		returnData.reports[i] = lineSlice
	}
	return
}
