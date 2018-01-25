package main

import (
	"net/http"
	"ibsi/dbase"
	"ibsi/crud"
)

func init() {

	crud.Handler(crud.CrudHandler {
		Name: "suppliers",
		Action: "suppliers",
		KeyName: "id",
		ListDataSource: "DBApp.GetSuppliers",
		EditDataSource: "DBApp.GetSuppliers",
		UpdateDataSource: "DBApp.AddSupplier",
		OnInitCrud: func(crud map[string]bool) {
			crud["delete"] = false
		},
		OnNewRecord: func(mode string, row map[string]interface{}, w http.ResponseWriter, r *http.Request) {
			// row["ewt"] = 0
			row["status_code_id"] = 10
		},
	})
	
	dbase.Connections["DBApp"].NewCommand("GetSuppliers", "GetSuppliers", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "in", 0, 0)
		cmd.NewParameter("ids", "string", "in", 200, "")
		cmd.NewParameter("name", "string", "in", 100, "")
		cmd.NewParameter("action", "int", "in", 0, 0)
		cmd.NewParameter("page", "int", "in", 0, 1)
		cmd.NewParameter("pagesize", "int", "in", 0, 25)
		cmd.NewParameter("row_count", "int", "inout", 0, 0)
		cmd.NewParameter("page_count", "int", "inout", 0, 0)
		cmd.NewParameter("sort", "string", "in", 200, "full_name")
		cmd.NewParameter("order", "string", "in", 10, "asc")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	}) 

	dbase.Connections["DBApp"].NewCommand("AddSupplier", "AddSupplier", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "inout", 0, 0)	
		cmd.NewParameter("account_code", "string", "in", 20, "")
		cmd.NewParameter("full_name", "string", "in", 100, "")
		cmd.NewParameter("tin", "string", "in", 20, "")
		// cmd.NewParameter("ewt", "float", "in", 0, 0)
		cmd.NewParameter("street", "string", "in", 200, "")
		cmd.NewParameter("city", "string", "in", 60, "")
		cmd.NewParameter("province", "string", "in", 60, "")
		cmd.NewParameter("zip_code", "string", "in", 20, "")
		cmd.NewParameter("country_code", "string", "in", 3, "")
		cmd.NewParameter("status_code_id", "int", "in", 0, 0)
		cmd.NewParameter("action", "int", "in", 0, 10)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
		cmd.NewParameter("action_status_id", "int", "inout", 0, 0)
		cmd.NewParameter("action_msg", "string", "inout", 200, "")
	}) 
}
