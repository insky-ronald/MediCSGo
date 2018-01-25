package main

import (
	// "net/http"
	"ibsi/dbase"
	"ibsi/crud"
)

func init() {

	crud.Handler(crud.CrudHandler {
		Name: "claim-audit-logs",
		Action: "claim",
		KeyName: "id",
		ListDataSource: "DBApp.GetAuditLogs",
		OnInitCrud: func(crud map[string]bool) {
			crud["add"] = false
			crud["edit"] = false
			crud["delete"] = false
			// DatabaseUtils.GetActionPermission("admin", Crud)
		},
	})

	dbase.Connections["DBApp"].NewCommand("GetAuditLogs", "GetAuditLogs", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("claim_id", "int", "in", 0, 0)
		cmd.NewParameter("service_id", "int", "in", 0, 0)
		cmd.NewParameter("page", "int", "in", 0, 1)
		cmd.NewParameter("pagesize", "int", "in", 0, 0)
		cmd.NewParameter("row_count", "int", "inout", 0, 0)
		cmd.NewParameter("page_count", "int", "inout", 0, 0)
		cmd.NewParameter("sort", "string", "in", 200, "")
		cmd.NewParameter("order", "string", "in", 10, "")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})
}
