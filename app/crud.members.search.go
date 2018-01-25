package main

import (
	// "net/http"
	"ibsi/dbase"
	"ibsi/crud"
)

func init() {

	crud.Handler(crud.CrudHandler {
		Name: "members-search",
		Action: "members-search",
		KeyName: "id",
		ListDataSource: "DBApp.GetMembersEnquiry",
		OnInitCrud: func(crud map[string]bool) {
			crud["delete"] = false
		},
		// OnNewRecord: func(mode string, row map[string]interface{}, w http.ResponseWriter, r *http.Request) {
			// row["ewt"] = 0
			// row["status_code_id"] = 10
		// },
	})

	dbase.Connections["DBApp"].NewCommand("GetMembersEnquiry", "GetMembersEnquiry", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "in", 0, 1)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
		cmd.NewParameter("page", "int", "in", 0, 1)
		cmd.NewParameter("pagesize", "int", "in", 0, 0)
		cmd.NewParameter("row_count", "int", "inout", 0, 0)
		cmd.NewParameter("page_count", "int", "inout", 0, 0)
		cmd.NewParameter("sort", "string", "in", 200, "")
		cmd.NewParameter("order", "string", "in", 10, "")
	})
	
}
