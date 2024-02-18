package tblprinter

import (
	"fmt"
	"os"
	"reflect"

	"github.com/olekukonko/tablewriter"
)

func RenderTable(data interface{}) {
	// Reflect the slice to work with its elements
	sliceVal := reflect.ValueOf(data)
	if sliceVal.Kind() != reflect.Slice {
		fmt.Println("renderTable expects a slice as input.")
		return
	}

	table := tablewriter.NewWriter(os.Stdout)

	// Check if the slice is empty to avoid panic
	if sliceVal.Len() == 0 {
		fmt.Println("Empty slice provided.")
		return
	}

	// Use reflection on the first element to set headers dynamically based on struct fields
	firstElem := sliceVal.Index(0)
	var headers []string
	for i := 0; i < firstElem.NumField(); i++ {
		headers = append(headers, firstElem.Type().Field(i).Name)
	}
	table.SetHeader(headers)

	// Populate the table dynamically based on struct fields
	for i := 0; i < sliceVal.Len(); i++ {
		elem := sliceVal.Index(i)
		var row []string
		for j := 0; j < elem.NumField(); j++ {
			field := elem.Field(j)
			row = append(row, fmt.Sprintf("%v", field.Interface()))
		}
		table.Append(row)
	}

	// Render the table
	table.Render()
}