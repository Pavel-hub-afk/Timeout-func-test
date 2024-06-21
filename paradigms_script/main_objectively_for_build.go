// Тест времени выполнения функций в разной последовательности.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/etnz/permute"
)

// ITERATION_LOOP - переменная итераций цикла. Задается с помощью ключа "-i".
var ITERATION_LOOP int = 10
var START_PROPERTY_FLAG int = 1

func init() {
	flag.IntVar(&ITERATION_LOOP, "i", ITERATION_LOOP, "кол-во итераций цикла")
	flag.IntVar(&START_PROPERTY_FLAG, "sp", START_PROPERTY_FLAG, "св-во старта")
}

// FunctionTester - структура, содержащая методы для выполнения различных функций
type FunctionTester struct {
	iterationLoop int
	startProperty int
}

func NewFunctionTester(iterationLoop, startProperty int) *FunctionTester {
	return &FunctionTester{
		iterationLoop: iterationLoop,
		startProperty: startProperty,
	}
}

func (ft *FunctionTester) printScreen() time.Duration {
	t0 := time.Now()

	for i := 0; i < ft.iterationLoop; i++ {
		fmt.Println("Hello TEST")
	}

	t1 := time.Now()
	fmt.Println("Function print | time: ", t1.Sub(t0))
	return t1.Sub(t0)
}

func (ft *FunctionTester) writeFile() time.Duration {
	t0 := time.Now()

	for i := 0; i < ft.iterationLoop; i++ {
		file, err := os.OpenFile("files/test.txt", os.O_WRONLY, 0744)
		if err != nil {
			log.Fatal(err)
		}

		if _, err := file.WriteString("Hello TEST"); err != nil {
			log.Fatal(err)
		}

		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}

	t1 := time.Now()
	fmt.Println("Function write | time: ", t1.Sub(t0))
	return t1.Sub(t0)
}

func (ft *FunctionTester) readFile() time.Duration {
	t0 := time.Now()

	for i := 0; i < ft.iterationLoop; i++ {
		file, err := os.OpenFile("files/test.txt", os.O_RDONLY, 0744)
		if err != nil {
			log.Fatal(err)
		}

		scan := bufio.NewScanner(file)
		for scan.Scan() {
			str := scan.Text()
			if str == "" {
				// Do nothing for empty lines
			}
		}

		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}

	t1 := time.Now()
	fmt.Println("Function read | time: ", t1.Sub(t0))
	return t1.Sub(t0)
}

func (ft *FunctionTester) openDifferentFiles() time.Duration {
	t0 := time.Now()

	fileTypes := []string{"test.doc", "test.html", "test.txt", "test.jpg"}
	for i := 0; i < ft.iterationLoop; i++ {
		for _, fileType := range fileTypes {
			file, err := os.OpenFile("files/"+fileType, os.O_RDONLY, 0744)
			if err != nil {
				log.Fatal(err)
			}
			if err := file.Close(); err != nil {
				log.Fatal(err)
			}
		}
	}

	t1 := time.Now()
	fmt.Println("Function open | time: ", t1.Sub(t0))
	return t1.Sub(t0)
}

func (ft *FunctionTester) outputFuncDifferentTurn() {
	var allTime time.Duration

	arrayTurn := []string{"P", "W", "R", "O"}
	funcMap := map[string]func() time.Duration{
		"P": ft.printScreen,
		"W": ft.writeFile,
		"R": ft.readFile,
		"O": ft.openDifferentFiles,
	}
	fmt.Println(arrayTurn)

	// Начало последовательного выполнения функций
	for _, funcKey := range arrayTurn {
		funcTime := funcMap[funcKey]()
		allTime += funcTime
	}

	fmt.Println("All time:", allTime)
	fmt.Println("------------------------")
	allTime = 0

	p := permute.NewPlainChangeGen(len(arrayTurn))

	var sw [2]int
	for p.Next(&sw) {

		permute.SwapStrings(sw, arrayTurn)
		fmt.Println(arrayTurn)

		for _, funcKey := range arrayTurn {
			funcTime := funcMap[funcKey]()
			allTime += funcTime
		}

		fmt.Println("All time:", allTime)
		fmt.Println("------------------------")
		allTime = 0
	}
}

func main() {
	// Считывание ключа i и определение ключа sp
	flag.Parse()

	// Создание экземпляра FunctionTester
	ft := NewFunctionTester(ITERATION_LOOP, START_PROPERTY_FLAG)

	// Запуск основной функции
	ft.outputFuncDifferentTurn()
}
