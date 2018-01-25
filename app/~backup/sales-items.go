package main

import (
	"net/http"
	"ibsi/dbase"
	"ibsi/crud"
	"ibsi/utils"
)

func init() {

	crud.Handler(crud.CrudHandler {
		Name: "sale-items",
		Action: "sales",
		KeyName: "id",
		ListDataSource: "DBApp.GetSaleItems",
		EditDataSource: "DBApp.GetSaleItems",
		UpdateDataSource: "DBApp.AddSaleItem",
		OnInitCrud: func(crud map[string]bool) {
			// crud["delete"] = false
		},
		OnNewRecord: func(mode string, row map[string]interface{}, w http.ResponseWriter, r *http.Request) {
			// row["ewt"] = 0
			row["sale_id"] = utils.StrToInt(r.Form.Get("sale_id"))
		},
	})

	dbase.Connections["DBApp"].NewCommand("GetSaleItems", "GetSaleItems", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "in", 0, 0)
		cmd.NewParameter("sale_id", "int", "in", 0, 0)
		cmd.NewParameter("action", "int", "in", 0, 0)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	}) 

	dbase.Connections["DBApp"].NewCommand("AddSaleItem", "AddSaleItem", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "inout", 0, 0)	
		cmd.NewParameter("sale_id", "int", "in", 0, 0)	
		cmd.NewParameter("sequence_no", "int", "in", 0, 0)	
		cmd.NewParameter("description", "string", "in", 100, "")
		cmd.NewParameter("header", "string", "in", 100, "")
		cmd.NewParameter("footer", "string", "in", 100, "")
		cmd.NewParameter("packaging", "string", "in", 20, "")
		cmd.NewParameter("unit", "string", "in", 20, "")
		cmd.NewParameter("quantity", "float", "in", 0, 0)
		cmd.NewParameter("unit_price", "float", "in", 0, 0)
		cmd.NewParameter("actual_unit_price", "float", "in", 0, 0)
		cmd.NewParameter("notes", "string", "in", 2000, "")
		cmd.NewParameter("status_code_id", "int", "in", 0, 10)
		cmd.NewParameter("action", "int", "in", 0, 10)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
		cmd.NewParameter("action_status_id", "int", "inout", 0, 0)
		cmd.NewParameter("action_msg", "string", "inout", 200, "")
	}) 
}
