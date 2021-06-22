// Package cogtest provides utilities to test cog
package cogtest

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/olekukonko/tablewriter"
	"github.com/thehungry-dev/cog/color"
	"github.com/thehungry-dev/cog/tag/filter"
)

var failureTxt = "✗"
var successTxt = "✓"

type ResultMatrix struct {
	Headers   []string
	RowTitles []string
	Results   [][]bool
}

func toBool(text string) bool {
	switch text {
	case "TRUE":
		return true
	case "FALSE":
		return false
	default:
		panic("Unrecognized boolean text")
	}
}

func AssertMatrix(t *testing.T, matrix string) {
	var headers []string
	var tagsList [][]string
	matrixSuccess := true
	r := csv.NewReader(strings.NewReader(matrix))
	var buf strings.Builder
	twriter := tablewriter.NewWriter(&buf)
	twriter.SetAutoFormatHeaders(false)
	twriter.SetAlignment(tablewriter.ALIGN_CENTER)

	line := 0
	for {
		csvRow, err := r.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatal(err)
		}

		if line == 0 {
			tagsList = make([][]string, len(csvRow))

			for headerIndex, header := range csvRow {
				if headerIndex == 0 {
					continue
				}

				if header == "" {
					tagsList[headerIndex] = []string{}
					continue
				}

				tagsList[headerIndex] = strings.Split(header, filter.TagSeparator)
			}

			headers = csvRow
			twriter.SetHeader(headers)
			line += 1

			continue
		}

		filterText := csvRow[0]
		filter := filter.Parse(filterText)

		outputRow := make([]string, len(csvRow))
		outputRow[0] = filterText

		for testResultIndex, testResult := range csvRow {
			if testResultIndex == 0 {
				continue
			}

			tags := tagsList[testResultIndex]

			actualResult := filter.Select(tags)
			expectedResult := toBool(testResult)

			outputRow[testResultIndex] = successTxt

			if actualResult != expectedResult {
				outputRow[testResultIndex] = failureTxt
				matrixSuccess = false
			}
		}
		twriter.Append(outputRow)

		line += 1
	}

	twriter.Render()

	t.Helper()
	t.Logf("\n%s\n", buf.String())

	if !matrixSuccess {
		t.Error("Matrix test failed")
	}
}

func init() {
	failureTxt = fmt.Sprintf("%s%s%s", color.Red, failureTxt, color.Reset)
	successTxt = fmt.Sprintf("%s%s%s", color.Green, successTxt, color.Reset)
}
