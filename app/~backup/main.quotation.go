package main

import (
	"io/ioutil"
	"path/filepath"
	"fmt"
	"time"
	"strconv"
	"bytes"
	"net/http"
	"github.com/gorilla/mux"
	"ibsi/dbase"
	"ibsi/utils"
	"ibsi/template"
	"ibsi/session"
	"github.com/SebastiaanKlippert/wkhtmltopdf"
)

type Quotation struct {
	Pdf string
	Quotation dbase.TDataTableRow
	Owner dbase.TDataTableRow
	Sales dbase.TDataTableRow
	Items dbase.TDataTableRows
}

func (s *Quotation) QuotationDate() string {
	return s.Quotation["date"].(time.Time).Format("January 2, 2006")
}

func QuotationPdf(w http.ResponseWriter, r *http.Request) {
	
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		fmt.Println("QuotationPdf", err.Error())
	}

	vars := mux.Vars(r) 
	id := utils.StrToInt(vars["id"])
	
	// Add one page from an URL
	// page := wkhtmltopdf.NewPage(fmt.Sprintf("http://localhost/app/print/quotation/%d?pdf=1", id))
	// page.Zoom.Set(1.1)
	
	// pdfg.AddPage(page)
	// pdfg.PageSize.Set("Letter")
	// pdfg.MarginTop.Set(127)
	// pdfg.MarginLeft.Set(127)
	// pdfg.MarginRight.Set(127)
	// pdfg.MarginBottom.Set(0)
	// pdfg.Zoom.Set(1.1)
	
	pdfg.AddPage(wkhtmltopdf.NewPage(fmt.Sprintf("http://localhost:%d/app/print/quotation/%d?pdf=1", config.Port, id)))
	
	
	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		fmt.Println("QuotationPdf", err.Error())
	}
	
	var vid int64
	sn := session.GetSession(r)
	if (sn == nil) {
		vid = 1
	} else {
		vid = sn.VisitorId
	}
			
	quotation := dbase.Connections["DBApp"].OpenDataTable("GetQuotations", dbase.TParameters{
		"id": id,
		"action": 10,
		"visit_id": vid,
	})
	
	// fileName := fmt.Sprintf("./documents/%s.pdf", quotation.Get("reference_no"))
	// fileName := fmt.Sprintf("documents/%s.pdf", quotation.Get("reference_no"))
	// fileName := fmt.Sprintf("E:/JCJZ/documents/%s.pdf", quotation.Get("reference_no"))
	// fileName := fmt.Sprintf("./documents/%s.pdf", quotation.Get("reference_no"))
	// fileName := fmt.Sprintf("./%s.pdf", quotation.Get("reference_no"))
	fileName := filepath.Join(config.DocPath, quotation.Get("reference_no").(string) + ".pdf")
	err = pdfg.WriteFile(fileName)
	if err != nil {
		fmt.Println("QuotationPdf", fileName, err.Error())
	}
	
	downloadBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}

	// set the default MIME type to send
	mime := http.DetectContentType(downloadBytes)

	fileSize := len(string(downloadBytes))

	// Generate the server headers
	w.Header().Set("Content-Type", mime)
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=Quotation #%s.pdf", quotation.Get("reference_no")))
	w.Header().Set("Expires", "0")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Content-Length", strconv.Itoa(fileSize))
	w.Header().Set("Content-Control", "private, no-transform, no-store, must-revalidate")

	// force it down the client's.....
	http.ServeContent(w, r, fileName, time.Now(), bytes.NewReader(downloadBytes))
	
	// ServePdf(w, r, fileName)
	// ServePdf(w, r, fmt.Sprintf("%s.pdf", quotation.Get("reference_no")))
	// ServePdf(w, r, fmt.Sprintf("%s.pdf", fileName))
	
	// w.Header().Set("Content-Disposition", "attachment; filename=test.pdf")
	// w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	// w.Header().Set("Content-Length", r.Header.Get("Content-Length"))
	// io.Copy(w, r.Body)
	// http.ServeFile(w, r, fileName)
}

func init() {
	
	Router.HandleFunc("/app/getpdf/quotation/{id}", QuotationPdf)//.Host(config.Domain)
	
	template.NewController(template.Controller {
		Pid: "sales",
		Template: "quotation",
		OnInitHandlers: func(ts *template.Controller) {
			ts.Add("/app/print/quotation/{keyid:[0-9]+\\/?}")
		},
		OnInitPageData: func(r *http.Request, p *template.Page) {
			var vid int64
			sn := session.GetSession(r)
			if (sn == nil) {
				vid = 1
			} else {
				vid = sn.VisitorId
			}
			vars := mux.Vars(r) 
			id := utils.StrToInt(vars["keyid"])
			
			quotation := dbase.Connections["DBApp"].OpenDataTable("GetQuotations", dbase.TParameters{
				"id": id,
				"action": 10,
				// "visit_id": sn.VisitorId,
				"visit_id": vid,
			})
			
			owner := dbase.Connections["DBApp"].OpenDataTable("GetOwner", dbase.TParameters{
				// "visit_id": sn.VisitorId,
				"visit_id": vid,
			})
			
			sales := dbase.Connections["DBApp"].OpenDataTable("GetSales", dbase.TParameters{
				"id": quotation.Get("sale_id"),
				"action": 10,
				// "visit_id": sn.VisitorId,
				"visit_id": vid,
			})
			
			items := dbase.Connections["DBApp"].OpenDataTable("GetSaleItems", dbase.TParameters{
				"sale_id": quotation.Get("sale_id"),
				// "visit_id": sn.VisitorId,
				"visit_id": vid,
			})
			
			// p.Title = fmt.Sprintf("Quotation: %v", quotation.Get("reference_no"))
			
			// Convert.ToDateTime(DBQuotation.Rows[0]["date"]).ToString("MMMM d, yyyy")
			
			p.Custom = &Quotation{
				Pdf: utils.Ifs(r.URL.Query().Get("pdf") == "", "view", "pdf"),
				Quotation: quotation.GetRows()[0],
				Owner: owner.GetRows()[0],
				Sales: sales.GetRows()[0],
				Items: items.GetRows(),
			}
		},
	})
}
