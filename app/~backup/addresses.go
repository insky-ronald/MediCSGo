package main

import (
	"net/http"
	"ibsi/dbase"
	"ibsi/crud"
	"ibsi/utils"
)

func init() {

	crud.Handler(crud.CrudHandler {
		Name: "addresses",
		Action: "addresses",
		KeyName: "id",
		ListDataSource: "DBApp.GetAddresses",
		EditDataSource: "DBApp.GetAddresses",
		UpdateDataSource: "DBApp.AddAddress",
		OnNewRecord: func(mode string, row map[string]interface{}, w http.ResponseWriter, r *http.Request) {
			// row["ewt"] = 0
			row["name_id"] = utils.StrToInt(r.Form.Get("name_id"))
		},
	})
	
	dbase.Connections["DBApp"].NewCommand("GetAddresses", "GetAddresses", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "in", 0, 0)
		cmd.NewParameter("name_id", "int", "in", 0, 0)
		cmd.NewParameter("name", "string", "in", 100, "")
		cmd.NewParameter("action", "int", "in", 0, 0)
		cmd.NewParameter("page", "int", "in", 0, 1)
		cmd.NewParameter("pagesize", "int", "in", 0, 25)
		cmd.NewParameter("row_count", "int", "inout", 0, 0)
		cmd.NewParameter("page_count", "int", "inout", 0, 0)
		cmd.NewParameter("sort", "string", "in", 200, "name")
		cmd.NewParameter("order", "string", "in", 10, "asc")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	}) 

	dbase.Connections["DBApp"].NewCommand("lookup_addresses", "GetAddresses", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "in", 0, 0)
		cmd.NewParameter("name_id", "int", "in", 0, 0)
		cmd.NewParameter("name", "string", "in", 100, "")
		cmd.NewParameter("action", "int", "in", 0, 1)
		cmd.NewParameter("page", "int", "in", 0, 1)
		cmd.NewParameter("pagesize", "int", "in", 0, 1000000)
		cmd.NewParameter("row_count", "int", "inout", 0, 0)
		cmd.NewParameter("page_count", "int", "inout", 0, 0)
		cmd.NewParameter("sort", "string", "in", 200, "street")
		cmd.NewParameter("order", "string", "in", 10, "asc")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	}) 

	dbase.Connections["DBApp"].NewCommand("AddAddress", "AddAddress", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "inout", 0, 0)	
		cmd.NewParameter("name_id", "int", "in", 0, 0)	
		cmd.NewParameter("street", "string", "in", 200, "")
		cmd.NewParameter("city", "string", "in", 60, "")
		cmd.NewParameter("province", "string", "in", 60, "")
		cmd.NewParameter("zip_code", "string", "in", 20, "")
		cmd.NewParameter("country_code", "string", "in", 3, "")
		cmd.NewParameter("action", "int", "in", 0, 10)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
		cmd.NewParameter("action_status_id", "int", "inout", 0, 0)
		cmd.NewParameter("action_msg", "string", "inout", 200, "")
	}) 
}
