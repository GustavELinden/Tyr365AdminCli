package saveToFile

import (
	"encoding/json"
	"fmt"

	"os"
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

func SaveDataToJSONFile(data interface{}, filename string) error {
	// Marshal the data into JSON
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return fmt.Errorf("error marshaling data: %w", err)
	}

	// Write the JSON data to file
	err = os.WriteFile(filename, jsonData, 0644)
	
	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}

	fmt.Printf("Data successfully saved to %s\n", filename)
	return nil
}
func ReadDataFromJSONFile(filename string, data interface{}) error {
	// Read the JSON data from file
	fileData, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("error reading from file: %w", err)
	}

	// Unmarshal the JSON data into the provided data structure
	// Note: Ensure that `data` is a pointer to the expected structure
	err = json.Unmarshal(fileData, data)
	if err != nil {
		return fmt.Errorf("error unmarshaling JSON: %w", err)
	}

	// Convert the unmarshaled data back to JSON for pretty printing
	// prettyJSON, err := json.MarshalIndent(data, "", "    ")
	// if err != nil {
	// 	return fmt.Errorf("error marshaling data for print: %w", err)
	// }

	// fmt.Printf("Data successfully read from %s:\n", filename)
	// fmt.Println(string(prettyJSON))
	return nil
}

func DeleteFile(filename string) error {
	
	err := os.Remove(filename)
	if err != nil {
		return fmt.Errorf("error deleting file: %w", err)
	}
	return nil
}