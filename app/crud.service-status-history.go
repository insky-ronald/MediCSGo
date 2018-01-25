package main

import (
	// "net/http"
	"ibsi/dbase"
	"ibsi/crud"
)

func init() {

	crud.Handler(crud.CrudHandler {
		Name: "service-status-history",
		Action: "service",
		KeyName: "id",
		ListDataSource: "DBApp.GetServiceStatusHistory",
		OnInitCrud: func(crud map[string]bool) {
			crud["delete"] = false
		},
		// OnNewRecord: func(mode string, row map[string]interface{}, w http.ResponseWriter, r *http.Request) {
			// row["ewt"] = 0
			// row["status_code_id"] = 10
		// },
	})

	dbase.Connections["DBApp"].NewCommand("GetServiceStatusHistory", "GetServiceStatusHistory", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "in", 0, 0)
		cmd.NewParameter("sort", "string", "in", 200, "")
		cmd.NewParameter("order", "string", "in", 10, "")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})		
}
