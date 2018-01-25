package main

import (
	"net/http"
	"ibsi/dbase"
	"ibsi/crud"
)

func init() {

	crud.Handler(crud.CrudHandler {
		Name: "quotations",
		Action: "quotations",
		KeyName: "id",
		ListDataSource: "DBApp.GetQuotations",
		EditDataSource: "DBApp.GetQuotations",
		UpdateDataSource: "DBApp.AddQuotation",
		OnInitCrud: func(crud map[string]bool) {
			crud["delete"] = false
		},
		OnNewRecord: func(mode string, row map[string]interface{}, w http.ResponseWriter, r *http.Request) {
			// row["ewt"] = 0
			// row["name_id"] = StrToInt(r.Form.Get("name_id"))
		},
	})

	dbase.Connections["DBApp"].NewCommand("GetQuotations", "GetQuotations", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "in", 0, 0)
		cmd.NewParameter("sale_id", "int", "in", 0, 0)
		cmd.NewParameter("action", "int", "in", 0, 0)
		cmd.NewParameter("sort", "string", "in", 200, "reference_no")
		cmd.NewParameter("order", "string", "in", 10, "asc")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	}) 

	dbase.Connections["DBApp"].NewCommand("AddQuotation", "AddQuotation", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "inout", 0, 0)	
		cmd.NewParameter("sale_id", "int", "in", 0, 0)	
		cmd.NewParameter("date", "datetime", "in", 0, nil)
		cmd.NewParameter("request_date", "datetime", "in", 0, nil)
		cmd.NewParameter("reference_no", "string", "in", 20, "")
		cmd.NewParameter("customer_reference_no", "string", "in", 20, "")
		cmd.NewParameter("address_id", "int", "in", 0, 0)	
		cmd.NewParameter("contact_id", "int", "in", 0, 0)	
		cmd.NewParameter("undersign_contact_id", "int", "in", 0, 0)	
		cmd.NewParameter("status_code_id", "int", "in", 0, 10)	
		cmd.NewParameter("notes", "string", "in", -1, "")
		cmd.NewParameter("action", "int", "in", 0, 10)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
		cmd.NewParameter("action_status_id", "int", "inout", 0, 0)
		cmd.NewParameter("action_msg", "string", "inout", 200, "")
	}) 

	dbase.Connections["DBApp"].NewCommand("GetQuotationContacts", "GetQuotationContacts", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "in", 0, 0)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	}) 
}
