package utils

import (
	"fmt"
	"strings"
)

func WrapToHTML(tableContent string) string {
	return fmt.Sprintf("<html><head></head><body><table>%s</table></body></html>", tableContent)
}

func GenerateTableRow(data []string, tableTag string) string {
	var header strings.Builder
	header.WriteString("<tr>")
	for _, col := range data {
		header.WriteString(fmt.Sprintf("<%s>%s</%s>", tableTag, col, tableTag))
	}
	header.WriteString("</tr>")
	return header.String()
}
