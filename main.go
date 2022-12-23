// Тест времени выплнения функций в разной последовательности.

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
func printScreen(l int) time.Duration {
	t0 := time.Now()

	for i := 0; i < l; i++ {
		fmt.Println("Hello TEST")
	}

	t1 := time.Now()
	fmt.Println("Function print | time: ", t1.Sub(t0))
	return t1.Sub(t0)
}

// writeFile - функция записи файла
func writeFile(l int) time.Duration {
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
	return t1.Sub(t0)
}

// readFile - функция чтения файла
func readFile(l int) time.Duration {
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
	return t1.Sub(t0)
}

// openDifferentFiles - функция открытия файлов разных типов
func openDifferentFiles(l int) time.Duration {
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
	return t1.Sub(t0)
}

// outputFuncDifferentTurn - функция вывода времени выполнения функций в разной последовательностии генерации Excel файла
func outputFuncDifferentTurn() {

	var allTime time.Duration

	arrayTurn := []string{"P", "W", "R", "O"}
	fmt.Println(arrayTurn)

	allTime += printScreen(ITERATION_LOOP)
	allTime += writeFile(ITERATION_LOOP)
	allTime += readFile(ITERATION_LOOP)
	allTime += openDifferentFiles(ITERATION_LOOP)

	fmt.Println("All time:", allTime)
	fmt.Println("------------------------")
	allTime = 0

	p := permute.NewPlainChangeGen(len(arrayTurn))

	var sw [2]int
	for p.Next(&sw) {

		permute.SwapStrings(sw, arrayTurn)
		fmt.Println(arrayTurn)

		for _, value := range arrayTurn {
			switch value {

			case "P":
				allTime += printScreen(ITERATION_LOOP)

			case "W":
				allTime += writeFile(ITERATION_LOOP)

			case "R":
				allTime += readFile(ITERATION_LOOP)

			case "O":
				allTime += openDifferentFiles(ITERATION_LOOP)
			}
		}

		fmt.Println("All time:", allTime)
		fmt.Println("------------------------")
		allTime = 0
	}
}

func main() {
	flag.Parse()

	outputFuncDifferentTurn()
}
