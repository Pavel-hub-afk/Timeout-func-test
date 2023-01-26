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

func init() {
	flag.IntVar(&ITERATION_LOOP, "i", ITERATION_LOOP, "кол-во итераций цикла")
}

// printScreen - функция вывода на экран
func printScreen(l int, channel chan time.Duration) {
	t0 := time.Now()

	for i := 0; i < l; i++ {
		fmt.Println("Hello TEST")
	}

	t1 := time.Now()
	fmt.Println("Function print | time: ", t1.Sub(t0))
	channel <- t1.Sub(t0)
}

// writeFile - функция записи файла
func writeFile(l int, channel chan time.Duration) {
	t0 := time.Now()

	for i := 0; i < l; i++ {
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
	channel <- t1.Sub(t0)
}

// readFile - функция чтения файла
func readFile(l int, channel chan time.Duration) {
	t0 := time.Now()

	for i := 0; i < l; i++ {
		file, err := os.OpenFile("files/test.txt", os.O_RDONLY, 0744)
		if err != nil {
			log.Fatal(err)
		}

		scan := bufio.NewScanner(file)
		for scan.Scan() {
			str := scan.Text()
			if str == "" {

			}
		}

		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}

	t1 := time.Now()
	fmt.Println("Function read | time: ", t1.Sub(t0))
	channel <- t1.Sub(t0)
}

// openDifferentFiles - функция открытия файлов разных типов
func openDifferentFiles(l int, channel chan time.Duration) {
	t0 := time.Now()

	for i := 0; i < l; i++ {
		file, err := os.OpenFile("files/test.doc", os.O_RDONLY, 0744)
		if err != nil {
			log.Fatal(err)
		}
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}

		file, err = os.OpenFile("files/test.html", os.O_RDONLY, 0744)
		if err != nil {
			log.Fatal(err)
		}
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}

		file, err = os.OpenFile("files/test.txt", os.O_RDONLY, 0744)
		if err != nil {
			log.Fatal(err)
		}
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}

		file, err = os.OpenFile("files/test.jpg", os.O_RDONLY, 0744)
		if err != nil {
			log.Fatal(err)
		}
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}

	t1 := time.Now()
	fmt.Println("Function open | time: ", t1.Sub(t0))
	channel <- t1.Sub(t0)
}

// outputFuncDifferentTurn - функция вывода времени выполнения функций в разной последовательности
func outputFuncDifferentTurn() {

	var allTime time.Duration
	channelPrint := make(chan time.Duration)
	channelWrite := make(chan time.Duration)
	channelRead := make(chan time.Duration)
	channelOpen := make(chan time.Duration)

	arrayTurn := []string{"P", "W", "R", "O"}
	fmt.Println(arrayTurn)

	go printScreen(ITERATION_LOOP, channelPrint)
	go writeFile(ITERATION_LOOP, channelWrite)
	go readFile(ITERATION_LOOP, channelRead)
	go openDifferentFiles(ITERATION_LOOP, channelOpen)

	allTime += <-channelPrint
	allTime += <-channelWrite
	allTime += <-channelRead
	allTime += <-channelOpen

	fmt.Println("All time:", allTime)
	fmt.Println("------------------------")
	allTime = 0

	p := permute.NewPlainChangeGen(len(arrayTurn))

	var sw [2]int
	for p.Next(&sw) {

		go printScreen(ITERATION_LOOP, channelPrint)
		go writeFile(ITERATION_LOOP, channelWrite)
		go readFile(ITERATION_LOOP, channelRead)
		go openDifferentFiles(ITERATION_LOOP, channelOpen)

		permute.SwapStrings(sw, arrayTurn)
		fmt.Println(arrayTurn)

		for _, value := range arrayTurn {
			switch value {

			case "P":
				allTime += <-channelPrint

			case "W":
				allTime += <-channelWrite

			case "R":
				allTime += <-channelRead

			case "O":
				allTime += <-channelOpen
			}
		}

		fmt.Println("All time:", allTime)
		fmt.Println("------------------------")
		allTime = 0
	}
}

func main() {
	t0 := time.Now()

	flag.Parse()

	outputFuncDifferentTurn()

	t1 := time.Now()
	fmt.Println("----[", t1.Sub(t0), "]----")
}
