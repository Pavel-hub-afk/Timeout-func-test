// Тест времени выплнения функций в разной последовательности.

package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/etnz/permute"
	"github.com/xuri/excelize/v2"
)

// ITERATION_LOOP - переменная итераций цикла. Задается с помощью ключа "-i".
var ITERATION_LOOP int = 10
var START_PROPERTY_FLAG int = 1

func init() {
	flag.IntVar(&ITERATION_LOOP, "i", ITERATION_LOOP, "кол-во итераций цикла")
	flag.IntVar(&START_PROPERTY_FLAG, "sp", START_PROPERTY_FLAG, "св-во старта")
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
		// TODO: почему-то пустое условие if
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

// outputFuncDifferentTurn - функция вывода времени выполнения функций в разной последовательностии, генерации Excel и CSV файлов
func outputFuncDifferentTurn(file *excelize.File, writerCSV *csv.Writer) {
	var indexExcle int = 1
	var funcTime time.Duration
	var allTime time.Duration

	arrayTurn := []string{"P", "W", "R", "O"}
	fmt.Println(arrayTurn)

	// Запись заколовка csv файла
	if err := writerCSV.Write([]string{"start_combination", "func", "time", "system", "start_property"}); err != nil {
		panic(err)
	}

	file.SetCellValue("Sheet1", fmt.Sprintf("A%d", indexExcle), arrayTurn)
	indexExcle++

	// Начало последовательного выполнения функций
	funcTime = printScreen(ITERATION_LOOP)
	file.SetCellValue("Sheet1", fmt.Sprintf("A%d", indexExcle), "Вывод на экран | время: ")
	file.SetCellValue("Sheet1", fmt.Sprintf("B%d", indexExcle), fmt.Sprint(funcTime))

	if err := writerCSV.Write([]string{strings.Join(arrayTurn, ","), "Вывод на экран", funcTime.String(), runtime.GOOS, strconv.Itoa(START_PROPERTY_FLAG)}); err != nil {
		panic(err)
	}

	allTime += funcTime
	indexExcle++

	funcTime = writeFile(ITERATION_LOOP)
	file.SetCellValue("Sheet1", fmt.Sprintf("A%d", indexExcle), "Запись в файл | время: ")
	file.SetCellValue("Sheet1", fmt.Sprintf("B%d", indexExcle), fmt.Sprint(funcTime))

	if err := writerCSV.Write([]string{strings.Join(arrayTurn, ","), "Запись в файл", funcTime.String(), runtime.GOOS, strconv.Itoa(START_PROPERTY_FLAG)}); err != nil {
		panic(err)
	}

	allTime += funcTime
	indexExcle++

	funcTime = readFile(ITERATION_LOOP)
	file.SetCellValue("Sheet1", fmt.Sprintf("A%d", indexExcle), "Чтение файла | время: ")
	file.SetCellValue("Sheet1", fmt.Sprintf("B%d", indexExcle), fmt.Sprint(funcTime))

	if err := writerCSV.Write([]string{strings.Join(arrayTurn, ","), "Чтение файла", funcTime.String(), runtime.GOOS, strconv.Itoa(START_PROPERTY_FLAG)}); err != nil {
		panic(err)
	}

	allTime += funcTime
	indexExcle++

	funcTime = openDifferentFiles(ITERATION_LOOP)
	file.SetCellValue("Sheet1", fmt.Sprintf("A%d", indexExcle), "Октрытие файлаов разного типа | время: ")
	file.SetCellValue("Sheet1", fmt.Sprintf("B%d", indexExcle), fmt.Sprint(funcTime))

	if err := writerCSV.Write([]string{strings.Join(arrayTurn, ","), "Октрытие файлаов разного типа", funcTime.String(), runtime.GOOS, strconv.Itoa(START_PROPERTY_FLAG)}); err != nil {
		panic(err)
	}

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

				if err := writerCSV.Write([]string{strings.Join(arrayTurn, ","), "Вывод на экран", funcTime.String(), runtime.GOOS, strconv.Itoa(START_PROPERTY_FLAG)}); err != nil {
					panic(err)
				}

				allTime += funcTime
				indexExcle++

			case "W":
				funcTime = writeFile(ITERATION_LOOP)
				file.SetCellValue("Sheet1", fmt.Sprintf("A%d", indexExcle), "Запись в файл | время: ")
				file.SetCellValue("Sheet1", fmt.Sprintf("B%d", indexExcle), fmt.Sprint(funcTime))

				if err := writerCSV.Write([]string{strings.Join(arrayTurn, ","), "Запись в файл", funcTime.String(), runtime.GOOS, strconv.Itoa(START_PROPERTY_FLAG)}); err != nil {
					panic(err)
				}

				allTime += funcTime
				indexExcle++

			case "R":
				funcTime = readFile(ITERATION_LOOP)
				file.SetCellValue("Sheet1", fmt.Sprintf("A%d", indexExcle), "Чтение файла | время: ")
				file.SetCellValue("Sheet1", fmt.Sprintf("B%d", indexExcle), fmt.Sprint(funcTime))

				if err := writerCSV.Write([]string{strings.Join(arrayTurn, ","), "Чтение файла", funcTime.String(), runtime.GOOS, strconv.Itoa(START_PROPERTY_FLAG)}); err != nil {
					panic(err)
				}

				allTime += funcTime
				indexExcle++

			case "O":
				funcTime = openDifferentFiles(ITERATION_LOOP)
				file.SetCellValue("Sheet1", fmt.Sprintf("A%d", indexExcle), "Октрытие файлаов разного типа | время: ")
				file.SetCellValue("Sheet1", fmt.Sprintf("B%d", indexExcle), fmt.Sprint(funcTime))

				if err := writerCSV.Write([]string{strings.Join(arrayTurn, ","), "Октрытие файлаов разного типа", funcTime.String(), runtime.GOOS, strconv.Itoa(START_PROPERTY_FLAG)}); err != nil {
					panic(err)
				}

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
	// Считывание ключа i и определение ключа sp
	flag.Parse()

	// Создание excel файла
	file := excelize.NewFile()

	// Создание CSV файла и его записывателя
	fileName := "secondary_start_windows_data.csv"
	if START_PROPERTY_FLAG == 0 {
		fileName = "primary_start_windows_data.csv"
	}

	fileCSV, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writerCSV := csv.NewWriter(fileCSV)
	defer writerCSV.Flush()

	// Запуск основной функции
	outputFuncDifferentTurn(file, writerCSV)

	if err := file.SaveAs("IntermediaTimeFunc.xlsx"); err != nil {
		fmt.Println(err)
	}
}
