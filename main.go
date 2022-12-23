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
	"github.com/xuri/excelize/v2"
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
func outputFuncDifferentTurn(file *excelize.File) {
	var indexExcle int = 1
	var funcTime time.Duration
	var allTime time.Duration

	arrayTurn := []string{"P", "W", "R", "O"}
	fmt.Println(arrayTurn)

	file.SetCellValue("Sheet1", fmt.Sprintf("A%d", indexExcle), arrayTurn)
	indexExcle++

	funcTime = printScreen(ITERATION_LOOP)
	file.SetCellValue("Sheet1", fmt.Sprintf("A%d", indexExcle), "Вывод на экран | время: ")
	file.SetCellValue("Sheet1", fmt.Sprintf("B%d", indexExcle), fmt.Sprint(funcTime))
	allTime += funcTime
	indexExcle++

	funcTime = writeFile(ITERATION_LOOP)
	file.SetCellValue("Sheet1", fmt.Sprintf("A%d", indexExcle), "Запись в файл | время: ")
	file.SetCellValue("Sheet1", fmt.Sprintf("B%d", indexExcle), fmt.Sprint(funcTime))
	allTime += funcTime
	indexExcle++

	funcTime = readFile(ITERATION_LOOP)
	file.SetCellValue("Sheet1", fmt.Sprintf("A%d", indexExcle), "Чтение файла | время: ")
	file.SetCellValue("Sheet1", fmt.Sprintf("B%d", indexExcle), fmt.Sprint(funcTime))
	allTime += funcTime
	indexExcle++

	funcTime = openDifferentFiles(ITERATION_LOOP)
	file.SetCellValue("Sheet1", fmt.Sprintf("A%d", indexExcle), "Октрытие файлаов разного типа | время: ")
	file.SetCellValue("Sheet1", fmt.Sprintf("B%d", indexExcle), fmt.Sprint(funcTime))
	allTime += funcTime
	indexExcle++

	file.SetCellValue("Sheet1", fmt.Sprintf("A%d", indexExcle), "Общее время выполнения | время: ")
	file.SetCellValue("Sheet1", fmt.Sprintf("B%d", indexExcle), fmt.Sprint(allTime))
	indexExcle++
	indexExcle++

	fmt.Println("All time:", allTime)
	fmt.Println("------------------------")
	allTime = 0

	p := permute.NewPlainChangeGen(len(arrayTurn))

	var sw [2]int
	for p.Next(&sw) {

		permute.SwapStrings(sw, arrayTurn)
		fmt.Println(arrayTurn)

		file.SetCellValue("Sheet1", fmt.Sprintf("A%d", indexExcle), arrayTurn)
		indexExcle++

		for _, value := range arrayTurn {
			switch value {

			case "P":
				funcTime = printScreen(ITERATION_LOOP)
				file.SetCellValue("Sheet1", fmt.Sprintf("A%d", indexExcle), "Вывод на экран | время: ")
				file.SetCellValue("Sheet1", fmt.Sprintf("B%d", indexExcle), fmt.Sprint(funcTime))
				allTime += funcTime
				indexExcle++

			case "W":
				funcTime = writeFile(ITERATION_LOOP)
				file.SetCellValue("Sheet1", fmt.Sprintf("A%d", indexExcle), "Запись в файл | время: ")
				file.SetCellValue("Sheet1", fmt.Sprintf("B%d", indexExcle), fmt.Sprint(funcTime))
				allTime += funcTime
				indexExcle++

			case "R":
				funcTime = readFile(ITERATION_LOOP)
				file.SetCellValue("Sheet1", fmt.Sprintf("A%d", indexExcle), "Чтение файла | время: ")
				file.SetCellValue("Sheet1", fmt.Sprintf("B%d", indexExcle), fmt.Sprint(funcTime))
				allTime += funcTime
				indexExcle++

			case "O":
				funcTime = openDifferentFiles(ITERATION_LOOP)
				file.SetCellValue("Sheet1", fmt.Sprintf("A%d", indexExcle), "Октрытие файлаов разного типа | время: ")
				file.SetCellValue("Sheet1", fmt.Sprintf("B%d", indexExcle), fmt.Sprint(funcTime))
				allTime += funcTime
				indexExcle++
			}
		}
		file.SetCellValue("Sheet1", fmt.Sprintf("A%d", indexExcle), "Общее время выполнения | время: ")
		file.SetCellValue("Sheet1", fmt.Sprintf("B%d", indexExcle), fmt.Sprint(allTime))
		indexExcle++
		indexExcle++

		fmt.Println("All time:", allTime)
		fmt.Println("------------------------")
		allTime = 0
	}
}

func main() {
	flag.Parse()

	file := excelize.NewFile()

	outputFuncDifferentTurn(file)

	if err := file.SaveAs("IntermediaTimeFunc.xlsx"); err != nil {
		fmt.Println(err)
	}
}
