package main

import (
	"net/http"
	"ibsi/dbase"
	"ibsi/crud"
)

func init() {

	crud.Handler(crud.CrudHandler {
		Name: "sales",
		Action: "sales",
		KeyName: "id",
		ListDataSource: "DBApp.GetSales",
		EditDataSource: "DBApp.GetSales",
		UpdateDataSource: "DBApp.AddSale",
		OnInitCrud: func(crud map[string]bool) {
			crud["delete"] = false
		},
		OnNewRecord: func(mode string, row map[string]interface{}, w http.ResponseWriter, r *http.Request) {
			// row["ewt"] = 0
			// row["name_id"] = StrToInt(r.Form.Get("name_id"))
		},
	})

	dbase.Connections["DBApp"].NewCommand("GetSales", "GetSales", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "in", 0, 0)
		cmd.NewParameter("customer_id", "int", "in", 0, 0)
		cmd.NewParameter("filter", "string", "in", 100, "")
		cmd.NewParameter("action", "int", "in", 0, 0)
		cmd.NewParameter("page", "int", "in", 0, 1)
		cmd.NewParameter("pagesize", "int", "in", 0, 25)
		cmd.NewParameter("row_count", "int", "inout", 0, 0)
		cmd.NewParameter("page_count", "int", "inout", 0, 0)
		cmd.NewParameter("sort", "string", "in", 200, "reference_no")
		cmd.NewParameter("order", "string", "in", 10, "asc")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	}) 

	dbase.Connections["DBApp"].NewCommand("AddSale", "AddSale", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "inout", 0, 0)	
		cmd.NewParameter("customer_id", "int", "in", 0, 0)	
		cmd.NewParameter("request_date", "datetime", "in", 0, nil)	
		cmd.NewParameter("reference_no", "string", "in", 20, "")
		cmd.NewParameter("terms_id", "int", "in", 0, 0)	
		cmd.NewParameter("lead_time_id", "int", "in", 0, 0)	
		cmd.NewParameter("price_validity_id", "int", "in", 0, 0)	
		cmd.NewParameter("has_vat", "int", "in", 0, 1)	
		cmd.NewParameter("invoice_no", "string", "in", 10, "")
		cmd.NewParameter("invoice_date", "datetime", "in", 0, nil)
		cmd.NewParameter("receipt_no", "string", "in", 10, "")
		cmd.NewParameter("receipt_date", "datetime", "in", 0, nil)
		cmd.NewParameter("delivery_no", "string", "in", 10, "")
		cmd.NewParameter("delivery_date", "datetime", "in", 0, nil)
		cmd.NewParameter("notes", "string", "in", 2000, "")
		cmd.NewParameter("action", "int", "in", 0, 10)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
		cmd.NewParameter("action_status_id", "int", "inout", 0, 0)
		cmd.NewParameter("action_msg", "string", "inout", 200, "")
	}) 

	dbase.Connections["DBApp"].NewQuery("GetTerms", nil, func() string {
		return "SELECT * FROM dbo.f_terms() ORDER by id"
	}) 

	dbase.Connections["DBApp"].NewQuery("GetLeadTimes", nil, func() string {
		return "SELECT * FROM dbo.f_lead_times() ORDER by id"
	}) 

	dbase.Connections["DBApp"].NewQuery("GetPriceValidities", nil, func() string {
		return "SELECT * FROM dbo.f_price_validities() ORDER by id"
	}) 
}
