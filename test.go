package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/etnz/permute"
)

var ITERATION_LOOP int = 10

func init() {
	flag.IntVar(&ITERATION_LOOP, "i", ITERATION_LOOP, "кол-во итераций цикла")
}

type Executor interface {
	PrintScreen(int) time.Duration
	WriteFile(int) time.Duration
	ReadFile(int) time.Duration
	OpenDifferentFiles(int) time.Duration
}

type ProceduralExecutor struct{}
type FunctionalExecutor struct {
	PrintFunc, WriteFunc, ReadFunc, OpenFunc func(int) time.Duration
}
type OOExecutor struct{}
type ParallelExecutor struct{}

func (pe ProceduralExecutor) PrintScreen(l int) time.Duration {
	return printScreen(l)
}

func (pe ProceduralExecutor) WriteFile(l int) time.Duration {
	return writeFile(l)
}

func (pe ProceduralExecutor) ReadFile(l int) time.Duration {
	return readFile(l)
}

func (pe ProceduralExecutor) OpenDifferentFiles(l int) time.Duration {
	return openDifferentFiles(l)
}

func (fe FunctionalExecutor) PrintScreen(l int) time.Duration {
	return fe.PrintFunc(l)
}

func (fe FunctionalExecutor) WriteFile(l int) time.Duration {
	return fe.WriteFunc(l)
}

func (fe FunctionalExecutor) ReadFile(l int) time.Duration {
	return fe.ReadFunc(l)
}

func (fe FunctionalExecutor) OpenDifferentFiles(l int) time.Duration {
	return fe.OpenFunc(l)
}

func (oe OOExecutor) PrintScreen(l int) time.Duration {
	return printScreen(l)
}

func (oe OOExecutor) WriteFile(l int) time.Duration {
	return writeFile(l)
}

func (oe OOExecutor) ReadFile(l int) time.Duration {
	return readFile(l)
}

func (oe OOExecutor) OpenDifferentFiles(l int) time.Duration {
	return openDifferentFiles(l)
}

func (pe ParallelExecutor) PrintScreen(l int) time.Duration {
	return runInParallel(printScreen, l)
}

func (pe ParallelExecutor) WriteFile(l int) time.Duration {
	return runInParallel(writeFile, l)
}

func (pe ParallelExecutor) ReadFile(l int) time.Duration {
	return runInParallel(readFile, l)
}

func (pe ParallelExecutor) OpenDifferentFiles(l int) time.Duration {
	return runInParallel(openDifferentFiles, l)
}

func printScreen(l int) time.Duration {
	t0 := time.Now()
	for i := 0; i < l; i++ {
		fmt.Println("Hello TEST")
	}
	t1 := time.Now()
	fmt.Println("Function print | time: ", t1.Sub(t0))
	return t1.Sub(t0)
}

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

func runInParallel(fn func(int) time.Duration, l int) time.Duration {
	var wg sync.WaitGroup
	wg.Add(1)
	var duration time.Duration
	go func() {
		defer wg.Done()
		duration = fn(l)
	}()
	wg.Wait()
	return duration
}

func outputFuncDifferentTurn(executor Executor) {
	var funcTime time.Duration
	var allTime time.Duration
	arrayTurn := []string{"P", "W", "R", "O"}
	fmt.Println(arrayTurn)
	funcTime = executor.PrintScreen(ITERATION_LOOP)
	allTime += funcTime
	funcTime = executor.WriteFile(ITERATION_LOOP)
	allTime += funcTime
	funcTime = executor.ReadFile(ITERATION_LOOP)
	allTime += funcTime
	funcTime = executor.OpenDifferentFiles(ITERATION_LOOP)
	allTime += funcTime
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
				funcTime = executor.PrintScreen(ITERATION_LOOP)
				allTime += funcTime
			case "W":
				funcTime = executor.WriteFile(ITERATION_LOOP)
				allTime += funcTime
			case "R":
				funcTime = executor.ReadFile(ITERATION_LOOP)
				allTime += funcTime
			case "O":
				funcTime = executor.OpenDifferentFiles(ITERATION_LOOP)
				allTime += funcTime
			}
		}
		fmt.Println("All time:", allTime)
		fmt.Println("------------------------")
		allTime = 0
	}
}

func main() {
	flag.Parse()
	fmt.Println("=================== Procedural Paradigm ===================")
	outputFuncDifferentTurn(ProceduralExecutor{})
	fmt.Println("=================== Functional Paradigm ===================")
	outputFuncDifferentTurn(FunctionalExecutor{
		PrintFunc: printScreen,
		WriteFunc: writeFile,
		ReadFunc:  readFile,
		OpenFunc:  openDifferentFiles,
	})
	fmt.Println("=================== Object-Oriented Paradigm ===================")
	outputFuncDifferentTurn(OOExecutor{})
	fmt.Println("=================== Parallel Paradigm ===================")
	outputFuncDifferentTurn(ParallelExecutor{})
}
