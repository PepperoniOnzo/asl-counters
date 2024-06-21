package services

import (
	"errors"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	pdfconstants "github.com/PepperoniOnzo/asl-counters/internal/constants/pdf_constants"
	"github.com/PepperoniOnzo/asl-counters/internal/models"
	"github.com/go-pdf/fpdf"
)

const assetsPdf = "../assets/pdf/"

func GeneratePdf(data *models.GenerateRequest) ([]byte, error) {

	if len(data.Counters) == 0 {
		return nil, errors.New("no counters to generate, specify ids in body")
	}

	sortedCounters := sortCountersBySize(data.Counters)

	pdf := fpdf.New("P", "mm", pdfconstants.PageFormatA4, "")
	for key, value := range sortedCounters {
		generateCountersPage(key, data.Spacing, value, pdf)
	}

	nowTime := time.Now()
	time := nowTime.Format(time.RFC822)
	fileName := assetsPdf + strings.ReplaceAll(time, " ", "") + ".pdf"

	err := pdf.OutputFileAndClose(fileName)

	if err != nil {
		return nil, errors.New("failed to create pdf")
	}

	file, err := os.Open(fileName)

	if err != nil {
		return nil, errors.New("failed to open created pdf")
	}

	defer file.Close()

	var res []byte
	_, err = file.Read(res)

	if err != nil {
		return nil, errors.New("failed to read created pdf")
	}

	return res, nil
}

func sortCountersBySize(counters []*models.CounterRequest) map[float64][]*models.CounterRequest {
	res := map[float64][]*models.CounterRequest{}

	for _, counter := range counters {
		res[counter.Size] = append(res[counter.Size], counter)
	}

	return res
}

func getMaxCountersForPage(sizeWithSpacing float64) (int, int) {
	maxWidthTokens := int(math.Floor(pdfconstants.A4Width / sizeWithSpacing))
	if maxWidthTokens%2 != 0 {
		maxWidthTokens--
	}

	maxHeightTokens := int(math.Floor(pdfconstants.A4Height / sizeWithSpacing))

	return maxWidthTokens, maxHeightTokens
}

func generateCountersPage(size float64, spacing float64, counters []*models.CounterRequest, pdf *fpdf.Fpdf) error {
	maxWidthTokens, maxHeightTokens := getMaxCountersForPage(size + spacing)

	if maxWidthTokens == 0 || maxHeightTokens == 0 {

		pdf.AddPage()
		pdf.SetFont("Arial", "B", 16)
		pdf.Cell(15, 15, "Cant generate tokens with size: "+strconv.FormatFloat(size, 'f', -1, 64))

		return nil
	}

	currentRow := 0
	currentColumn := 0
	counterAmount := 0

	pdf.AddPage()

	for counterIndex, counter := range counters {
		counterAmount = counter.Amount

		if counterAmount == 0 {
			continue
		}

		for row := currentRow; row < maxHeightTokens; row++ {
			for column := currentColumn; column < maxWidthTokens/2; column++ {
				pdf.ImageOptions(assets+counter.FrontPathId, float64(size*float64(column)+spacing*float64(column)), float64(size*float64(row)+spacing*float64(row)), size, size, false, fpdf.ImageOptions{}, 0, "")
				pdf.ImageOptions(assets+counter.BackgroundPathId, float64(pdfconstants.A4Width-size*float64(column+1)-spacing*float64(column)), float64(size*float64(row)+spacing*float64(row)), size, size, false, fpdf.ImageOptions{}, 0, "")

				counterAmount--

				if counterAmount == 0 {
					currentColumn = column
					break
				} else {
					currentColumn = 0
				}
			}

			if row == maxHeightTokens-1 {
				if counterIndex == len(counters)-1 {
					break
				}

				pdf.AddPage()
				currentColumn = 0
				currentRow = 0
				row = -1
			}

			if counterAmount == 0 {
				currentRow = row
				break
			}
		}
	}

	return nil
}
