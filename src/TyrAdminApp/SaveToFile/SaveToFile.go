package saveToFile

import (
	"fmt"
	"reflect"
	"time"

	"github.com/xuri/excelize/v2"
)

func SaveToExcel(data interface{}, fileName string) error {
	start := time.Now()
	f := excelize.NewFile()
	sheetName := "Sheet1"

	// Reflect the slice to work with its elements
	sliceVal := reflect.ValueOf(data)
	if sliceVal.Kind() != reflect.Slice {
		return fmt.Errorf("SaveToExcel expects a slice as input")
	}

	if sliceVal.Len() == 0 {
		return fmt.Errorf("Empty slice provided")
	}

	// Use reflection on the first element to generate headers based on struct fields
	firstElem := sliceVal.Index(0)
	headers := make([]string, firstElem.NumField())
	for i := 0; i < firstElem.NumField(); i++ {
		headers[i] = firstElem.Type().Field(i).Name
	}

	// Set the headers in the Excel file
	for i, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheetName, cell, header)
	}

	// Populate the Excel file dynamically based on struct fields
	for i := 0; i < sliceVal.Len(); i++ {
		elem := sliceVal.Index(i)
		for j := 0; j < elem.NumField(); j++ {
			cell, _ := excelize.CoordinatesToCellName(j+1, i+2)
			f.SetCellValue(sheetName, cell, elem.Field(j).Interface())
		}
	}

	// Save the Excel file
	err := f.SaveAs(fileName + ".xlsx")
	if err != nil {
		return err
	}

	fmt.Printf("Excel file created in %s\n", time.Since(start))
	return nil
}
