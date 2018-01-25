package main

import (
	"net/http"
	"ibsi/dbase"
	"ibsi/crud"
)

func init() {

	crud.Handler(crud.CrudHandler {
		Name: "doctors",
		Action: "doctors",
		KeyName: "id",
		ListDataSource: "DBApp.GetDoctors",
		UpdateDataSource: "DBApp.AddDoctors",
		OnInitCrud: func(crud map[string]bool) {
			// crud["delete"] = false
		},
		OnNewRecord: func(mode string, row map[string]interface{}, w http.ResponseWriter, r *http.Request) {
			// row["ewt"] = 0
			row["country"] = ""
			row["specialisation"] = ""
			row["notes"] = ""
			row["status_code"] = "A"
		},
	})
	

	dbase.Connections["DBApp"].NewCommand("GetDoctorSpecialisation", "GetDoctorSpecialisation", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("code", "string", "in", 5, "")
		cmd.NewParameter("lookup", "int", "in", 0, 0)
		cmd.NewParameter("filter", "string", "in", 100, "")
		cmd.NewParameter("page", "int", "in", 0, 1)
		cmd.NewParameter("pagesize", "int", "in", 0, 0)
		cmd.NewParameter("row_count", "int", "inout", 0, 0)
		cmd.NewParameter("page_count", "int", "inout", 0, 0)
		cmd.NewParameter("sort", "string", "in", 200, "specialisation_code")
		cmd.NewParameter("order", "string", "in", 10, "asc")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})

	dbase.Connections["DBApp"].NewCommand("GetDoctors", "GetDoctors", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "in", 0, 0)
		cmd.NewParameter("filter", "string", "in", 100, "")
		cmd.NewParameter("action", "int", "in", 0, 0)
		cmd.NewParameter("page", "int", "in", 0, 1)
		cmd.NewParameter("pagesize", "int", "in", 0, 0)
		cmd.NewParameter("row_count", "int", "inout", 0, 0)
		cmd.NewParameter("page_count", "int", "inout", 0, 0)
		cmd.NewParameter("sort", "string", "in", 200, "name")
		cmd.NewParameter("order", "string", "in", 10, "asc")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})

	dbase.Connections["DBApp"].NewCommand("AddDoctors", "AddDoctors", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "inout", 0, 0)
		cmd.NewParameter("code", "string", "in", 10, "")
		cmd.NewParameter("spin_id", "string", "in", 20, "")
		cmd.NewParameter("name", "string", "in", 100, "")
		cmd.NewParameter("full_name", "string", "in", 100, "")
		cmd.NewParameter("specialisation_code", "string", "in", 5, "")
		cmd.NewParameter("country_code", "string", "in", 3, "")
		cmd.NewParameter("status_code", "string", "in", 1, "")
		cmd.NewParameter("blacklisted", "int", "in", 0, 0)
		cmd.NewParameter("discount_type_id", "int", "in", 0, 0)
		cmd.NewParameter("discount_amount", "float", "in", 0, 0)
		cmd.NewParameter("discount_percent", "float", "in", 0, 0)
		cmd.NewParameter("notes", "string", "in", -1, "")

		cmd.NewParameter("action", "int", "in", 0, 10)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
		cmd.NewParameter("action_status_id", "int", "inout", 0, 0)
		cmd.NewParameter("action_msg", "string", "inout", 200, "")
	})
}
