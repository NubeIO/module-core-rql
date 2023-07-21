package apirules

import (
	"fmt"
	"github.com/mandolyte/mdtopdf"
)

type PDFResponse struct {
	Result []pingResult
	Error  string
}

type PdfBody struct {
	Input          string `json:"input" binding:"required"`
	WriteToHomeDir bool   `json:"write_to_home_dir"`
}

func (inst *Client) PDF(pdfBody *PdfBody) *PingResponse {
	// 	"github.com/mandolyte/mdtopdf"
	content := []byte(pdfBody.Input)
	output := "test.pdf"

	fmt.Println(pdfBody.WriteToHomeDir)

	pf := mdtopdf.NewPdfRenderer("", "", output, "")
	err := pf.Process(content)
	if err != nil {
		//die(err)
	}

	r := &PingResponse{
		Result: "wrote PDF ok",
		Error:  errorString(err),
	}
	//pprint.PrintJSON(r)
	return r
}
