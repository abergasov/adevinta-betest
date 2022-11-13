package parser

import (
	"BeTest-AlexanderBergasov/internal/utils"
	"bufio"
	"io"
	"sort"
	"strings"
)

type columnPosition struct {
	name  string
	index int
}

type PRNDecoder struct {
	columns []string
}

func NewPRNDecoder() *PRNDecoder {
	return &PRNDecoder{
		columns: []string{
			"name",
			"address",
			"postcode",
			"phone",
			"credit limit",
			"birthday",
			"sales",
		},
	}
}

func (c *PRNDecoder) Decode(reader io.Reader) (htmlBody string, err error) {
	scanner := bufio.NewScanner(reader)
	counter := 0
	var positions []columnPosition
	var builder strings.Builder
	for scanner.Scan() {
		tag := "td"
		if counter == 0 {
			tag = "th"
			positions = c.parseHeaderPositions(scanner.Text())
			counter++
		}
		htmlTableRow := utils.GenerateTableRow(c.parseRow(scanner.Text(), positions), tag)
		builder.WriteString(htmlTableRow)
	}
	return builder.String(), nil
}

// parseHeaderPositions parses the header line and returns the positions of the columns.
func (c *PRNDecoder) parseHeaderPositions(headerString string) []columnPosition {
	headerString = strings.ToLower(headerString)
	positions := make([]columnPosition, 0, len(c.columns))
	for _, col := range c.columns {
		index := strings.Index(headerString, col)
		if index < 0 {
			continue // not found in source
		}
		positions = append(positions, columnPosition{
			name:  col,
			index: index,
		})
	}
	sort.Slice(positions, func(i, j int) bool {
		return positions[i].index < positions[j].index
	})
	return positions
}

func (c *PRNDecoder) parseRow(headerString string, positions []columnPosition) []string {
	result := make([]string, 0, len(positions))
	for i, position := range positions {
		var content string
		if i < len(positions)-1 {
			content = headerString[position.index:positions[i+1].index]
		} else {
			content = headerString[position.index:]
		}
		result = append(result, strings.TrimSpace(content))
	}
	return result
}
