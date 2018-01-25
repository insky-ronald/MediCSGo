package main

import (
	"net/http"
	"ibsi/dbase"
	"ibsi/crud"
)

func init() {

	crud.Handler(crud.CrudHandler {
		Name: "countries",
		Action: "countries",
		KeyName: "id",
		ListDataSource: "DBApp.GetCountries",
		UpdateDataSource: "DBApp.AddCountries",
		OnInitCrud: func(crud map[string]bool) {
			// crud["delete"] = false
		},
		OnNewRecord: func(mode string, row map[string]interface{}, w http.ResponseWriter, r *http.Request) {
			// row["ewt"] = 0
			// row["status_code_id"] = 10
		},
	})
	

	dbase.Connections["DBApp"].NewCommand("GetCountries", "GetCountries", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("code", "string", "in", 3, "")
		cmd.NewParameter("filter", "string", "in", 100, "")

		cmd.NewParameter("action", "int", "in", 0, 0)
		cmd.NewParameter("page", "int", "in", 0, 1)
		cmd.NewParameter("pagesize", "int", "in", 0, 0)
		cmd.NewParameter("row_count", "int", "inout", 0, 0)
		cmd.NewParameter("page_count", "int", "inout", 0, 0)
		cmd.NewParameter("sort", "string", "in", 200, "country")
		cmd.NewParameter("order", "string", "in", 10, "asc")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})

	dbase.Connections["DBApp"].NewCommand("AddCountries", "AddCountries", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("code", "string", "inout", 3, "")
		cmd.NewParameter("iso_code", "string", "in", 3, "")
		cmd.NewParameter("country", "string", "in", 20, "")

		cmd.NewParameter("action", "int", "in", 0, 10)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
		cmd.NewParameter("action_status_id", "int", "inout", 0, 0)
		cmd.NewParameter("action_msg", "string", "inout", 200, "")
	})
}
