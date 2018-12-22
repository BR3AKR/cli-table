package clitable

import (
	"fmt"
	"strings"
)

// Print prints the supplied data in a nice table suitable
// for CLIs. If has header is true, it assumes the first
// row is the header.
func Print(data *[][]string, hasHeader bool) {
	rowIdx := 0
	widths := getWidths(data)
	format := getRowFormat(widths)
	totalWidth := getTableWidth(widths)

	fmt.Println(" " + strings.Repeat("_", totalWidth-2))

	if hasHeader {
		fmt.Printf(format+"\n", *convertStringToInterface(&(*data)[rowIdx])...)
		fmt.Printf(format+"\n", *convertStringToInterface(createBreaker(widths))...)
		rowIdx++
	}

	for ; rowIdx < len(*data); rowIdx++ {
		fmt.Printf(format+"\n", *convertStringToInterface(&(*data)[rowIdx])...)
	}

	fmt.Println(" " + strings.Repeat("‾", totalWidth-2))
}

func createBreaker(widths []int) *[]string {
	breaker := make([]string, len(widths))
	for i, width := range widths {
		breaker[i] = strings.Repeat("-", width)
	}
	return &breaker
}

func getWidths(data *[][]string) []int {
	colWidths := make([]int, getColCount(data))
	for _, row := range *data {
		for colIdx, value := range row {
			if len(value) > colWidths[colIdx] {
				colWidths[colIdx] = len(value)
			}
		}
	}
	return colWidths
}

func getColCount(data *[][]string) int {
	count := 0
	for _, cols := range *data {
		if len(cols) > count {
			count = len(cols)
		}
	}
	return count
}

func getRowFormat(widths []int) string {
	var sb strings.Builder

	for _, colWidth := range widths {
		sb.WriteString(fmt.Sprintf("| %%-%vv ", colWidth))
	}
	sb.WriteString("|")
	format := sb.String()
	return format
}

func convertStringToInterface(values *[]string) *[]interface{} {
	converted := make([]interface{}, len(*values))
	for i, v := range *values {
		converted[i] = v
	}
	return &converted
}

func getTableWidth(widths []int) int {
	totalWidth := 3 * len(widths)
	totalWidth++
	for _, width := range widths {
		totalWidth += width
	}
	return totalWidth
}
