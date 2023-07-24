package apirules

import (
	"github.com/mandolyte/mdtopdf"
	"github.com/wcharczuk/go-chart/v2"
	"os"
)

type PDFResponse struct {
	Result []pingResult
	Error  string
}

type PdfBody struct {
	Text string `json:"text"`
}

func (inst *Client) PDF(pdfBody *PdfBody) *PingResponse {
	// 	"github.com/mandolyte/mdtopdf"
	content := []byte(pdfBody.Text)
	output := "test.pdf"

	pie := chart.DonutChart{
		Width:  512,
		Height: 512,
		Values: []chart.Value{
			{Value: 5, Label: "Blue"},
			{Value: 5, Label: "Green"},
			{Value: 4, Label: "Gray"},
			{Value: 4, Label: "Orange"},
			{Value: 3, Label: "Deep Blue"},
			{Value: 3, Label: "test"},
		},
	}

	f, _ := os.Create("output.png")
	defer f.Close()
	pie.Render(chart.PNG, f)

	pf := mdtopdf.NewPdfRenderer("", "", output, "")
	err := pf.Process(content)
	r := &PingResponse{
		Result: pdfBody.Text,
		Error:  errorString(err),
	}
	return r
}
