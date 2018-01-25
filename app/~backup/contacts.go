package main

import (
	"net/http"
	"ibsi/dbase"
	"ibsi/crud"
	"ibsi/utils"
)

func init() {

	crud.Handler(crud.CrudHandler {
		Name: "contacts",
		Action: "contacts",
		KeyName: "id",
		ListDataSource: "DBApp.GetContacts",
		EditDataSource: "DBApp.GetContacts",
		UpdateDataSource: "DBApp.AddContact",
		OnNewRecord: func(mode string, row map[string]interface{}, w http.ResponseWriter, r *http.Request) {
			// row["ewt"] = 0
			row["name_id"] = utils.StrToInt(r.Form.Get("name_id"))
		},
	})

	dbase.Connections["DBApp"].NewCommand("GetContacts", "GetContacts", "procedure", func(cmd dbase.ICommand) {
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

	dbase.Connections["DBApp"].NewCommand("lookup_contacts", "GetContacts", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "in", 0, 0)
		cmd.NewParameter("name_id", "int", "in", 0, 0)
		cmd.NewParameter("name", "string", "in", 100, "")
		cmd.NewParameter("action", "int", "in", 0, 1)
		cmd.NewParameter("page", "int", "in", 0, 1)
		cmd.NewParameter("pagesize", "int", "in", 0, 1000000)
		cmd.NewParameter("row_count", "int", "inout", 0, 0)
		cmd.NewParameter("page_count", "int", "inout", 0, 0)
		cmd.NewParameter("sort", "string", "in", 200, "name")
		cmd.NewParameter("order", "string", "in", 10, "asc")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})

	dbase.Connections["DBApp"].NewCommand("AddContact", "AddContact", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "inout", 0, 0)	
		cmd.NewParameter("name_id", "int", "in", 0, 0)	
		cmd.NewParameter("name", "string", "in", 100, "")
		cmd.NewParameter("title", "string", "in", 20, "")
		cmd.NewParameter("department", "string", "in", 100, "")
		cmd.NewParameter("position", "string", "in", 60, "")
		cmd.NewParameter("phone_no", "string", "in", 20, "")
		cmd.NewParameter("fax_no", "string", "in", 20, "")
		cmd.NewParameter("mobile_no", "string", "in", 20, "")
		cmd.NewParameter("email", "string", "in", 200, "")
		cmd.NewParameter("action", "int", "in", 0, 10)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
		cmd.NewParameter("action_status_id", "int", "inout", 0, 0)
		cmd.NewParameter("action_msg", "string", "inout", 200, "")
	})
}
