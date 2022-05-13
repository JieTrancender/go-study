package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/xuri/excelize/v2"
)

func main() {
	testWriteFile()
	testReadFile()
	testDir()
}

func testWriteFile() {
	f := excelize.NewFile()
	index := f.NewSheet("Sheet2")
	f.SetCellValue("Sheet2", "A2", "Hello world!")
	f.SetCellValue("Sheet1", "B2", 100)
	f.SetActiveSheet(index)
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
}

func testReadFile() {
	f, err := excelize.OpenFile("Book1.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	cell, err := f.GetCellValue("Sheet1", "B2")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(cell)
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}

func testDir() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}
