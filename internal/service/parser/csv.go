package parser

import (
	"BeTest-AlexanderBergasov/internal/utils"
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

type CSVDecoder struct {
}

func NewCSVDecoder() *CSVDecoder {
	return &CSVDecoder{}
}

func (c *CSVDecoder) Decode(reader io.Reader) (htmlBody string, err error) {
	csvReader := csv.NewReader(reader)
	data, err := csvReader.ReadAll()
	if err != nil {
		return "", fmt.Errorf("failed to read csv file: %w", err)
	}

	var builder strings.Builder
	for i, row := range data {
		tag := "td"
		if i == 0 {
			tag = "th"
		}
		builder.WriteString(utils.GenerateTableRow(row, tag))
	}
	return builder.String(), nil
}
