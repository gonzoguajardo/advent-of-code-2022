package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	monkeyInitilizeSlize := make([]string, 0)
	monkeys := make([]*monkey, 0)
	for scanner.Scan() {
		currentString := scanner.Text()
		if currentString == "" {
			monkey := newMonkey(monkeyInitilizeSlize)
			monkeys = append(monkeys, monkey)
			monkeyInitilizeSlize = make([]string, 0)
		} else {
			monkeyInitilizeSlize = append(monkeyInitilizeSlize, currentString)
		}
	}
	monkey := newMonkey(monkeyInitilizeSlize)
	monkeys = append(monkeys, monkey)
	n := 20
	for i := 1; i <= n; i++ {
		for _, m := range(monkeys) {
			m.throw(monkeys)
		}
	}
	sort.Slice(monkeys[:], func(i, j int) bool {
		return monkeys[i].InspectedItems > monkeys[j].InspectedItems
	})
	log.Println(monkeys[0].InspectedItems * monkeys[1].InspectedItems)
}

func print(ms []*monkey) {
	for _, m := range(ms) {
		log.Println(m)
	}
}

func (m *monkey) throw(ms []*monkey) int {
	if len(m.CurrentItems) == 0 {
		return -1
	}
	for _, item := range(m.CurrentItems) {
		currentWorryLevel := item
		switch m.OperationKind {
		case Multiply:
			currentWorryLevel *= m.OperationValue
		case Plus:
			currentWorryLevel += m.OperationValue
		case Square:
			currentWorryLevel = currentWorryLevel * currentWorryLevel
		}
		currentWorryLevel = currentWorryLevel / 3
		if currentWorryLevel % m.Test == 0 {
			ms[m.TrueMonkey].CurrentItems = append(ms[m.TrueMonkey].CurrentItems, currentWorryLevel)
		} else {
			ms[m.FalseMonkey].CurrentItems = append(ms[m.FalseMonkey].CurrentItems, currentWorryLevel)
		}
		m.InspectedItems++
	}
	m.CurrentItems = []int{}
	return 0
}

func newMonkey(initSlice []string) *monkey {
	itemsString := strings.Split(initSlice[1], "Starting items: ")
	itemsStringSplit := strings.Split(itemsString[1], ", ")
	items := make([]int, 0)
	for _, itemsString := range(itemsStringSplit) {
		item, _ := strconv.Atoi(itemsString)
		items = append(items, item)
	}
	operationString := strings.Split(initSlice[2], "Operation: new = old ")[1]
	var operationKind Operation
	var operationValue int
	if strings.Contains(operationString, "old") {
		operationKind = Square
	} else {
		operationStringSplit := strings.Split(operationString, " ")
		operationKind = Operation(operationStringSplit[0])
		operationValue, _ = strconv.Atoi(operationStringSplit[1])
	}
	testString := strings.Split(initSlice[3], "Test: divisible by ")[1]
	test, _ := strconv.Atoi(testString)
	trueMonkeyString := strings.Split(initSlice[4], " throw to monkey ")[1]
	trueMonkey, _ := strconv.Atoi(trueMonkeyString)
	falseMonkeyString := strings.Split(initSlice[5], " throw to monkey ")[1]
	falseMonkey, _ := strconv.Atoi(falseMonkeyString)
	monkey := &monkey{
		CurrentItems: items,
		OperationKind: operationKind,
		OperationValue: operationValue,
		Test: test,
		TrueMonkey: trueMonkey,
		FalseMonkey: falseMonkey,
	}
	return monkey
}

type monkey struct {
	CurrentItems []int
	OperationKind Operation
	OperationValue int
	Test int
	TrueMonkey int
	FalseMonkey int
	InspectedItems int
}

type Operation string

const (
	Multiply Operation = "*"
	Plus Operation = "+"
	Square Operation = "square"
)
