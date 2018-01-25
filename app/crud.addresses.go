package main

import (
	"net/http"
	// "github.com/gorilla/mux"
	"ibsi/dbase"
	"ibsi/crud"
	"ibsi/utils"
)

func init() {

	crud.Handler(crud.CrudHandler {
		Name: "addresses",
		Action: "addresses",
		Path: "",
		KeyName: "id",
		ListDataSource: "DBApp.GetAddress",
		UpdateDataSource: "DBApp.AddAddress",
		OnInitCrud: func(crud map[string]bool) {
			// crud["delete"] = false
		},
		OnNewRecord: func(mode string, row map[string]interface{}, w http.ResponseWriter, r *http.Request) {
			row["name_id"] = utils.StrToInt(r.Form.Get("name_id"))
			// row["address_type_name"] = ""
			// row["country"] = ""
			// utils.Println(vars)
			// row["notes"] = ""
			// row["status_code"] = "A"
		},
	})
	
	dbase.Connections["DBApp"].NewCommand("GetAddress", "GetAddress", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "in", 0, 0)
		cmd.NewParameter("name_id", "int", "in", 0, 0)
		cmd.NewParameter("action", "int", "in", 0, 0)
		cmd.NewParameter("sort", "string", "in", 200, "")
		cmd.NewParameter("order", "string", "in", 10, "")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})

	dbase.Connections["DBApp"].NewCommand("AddAddress", "AddAddress", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("id", "int", "inout", 0, 0)
		cmd.NewParameter("name_id", "int", "in", 0, 0)
		cmd.NewParameter("address_type", "string", "in", 3, "")
		cmd.NewParameter("street", "string", "in", 100, "")
		cmd.NewParameter("city", "string", "in", 20, "")
		cmd.NewParameter("province", "string", "in", 30, "")
		cmd.NewParameter("zip_code", "string", "in", 20, "")
		cmd.NewParameter("country_code", "string", "in", 3, "")
		cmd.NewParameter("phone_no1", "string", "in", 20, "")
		cmd.NewParameter("phone_no2", "string", "in", 20, "")
		cmd.NewParameter("fax_no1", "string", "in", 20, "")
		cmd.NewParameter("fax_no2", "string", "in", 20, "")
		cmd.NewParameter("email", "string", "in", 60, "")

		cmd.NewParameter("action", "int", "in", 0, 10)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
		cmd.NewParameter("action_status_id", "int", "inout", 0, 0)
		cmd.NewParameter("action_msg", "string", "inout", 200, "")
	})
	
	dbase.Connections["DBApp"].NewQuery("lookup_address_types", func(cmd dbase.ICommand) {
		}, 
		func() string {
			return `
				select
					code,
					address_type
				from address_types
				order by address_type
			`
		},
	) 

	dbase.Connections["DBApp"].NewCommand("GetAddressTypes", "GetAddressTypes", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("code", "string", "in", 3, "")
		cmd.NewParameter("lookup", "int", "in", 0, 0)
		cmd.NewParameter("filter", "string", "in", 100, "")
		cmd.NewParameter("page", "int", "in", 0, 1)
		cmd.NewParameter("pagesize", "int", "in", 0, 0)
		cmd.NewParameter("row_count", "int", "inout", 0, 0)
		cmd.NewParameter("page_count", "int", "inout", 0, 0)
		cmd.NewParameter("sort", "string", "in", 200, "code")
		cmd.NewParameter("order", "string", "in", 10, "asc")
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
	})

	dbase.Connections["DBApp"].NewCommand("AddAddressTypes", "AddAddressTypes", "procedure", func(cmd dbase.ICommand) {
		cmd.NewParameter("code", "string", "inout", 3, "")
		cmd.NewParameter("address_type", "string", "in", 60, "")

		cmd.NewParameter("action", "int", "in", 0, 10)
		cmd.NewParameter("visit_id", "int", "in", 0, 0)
		cmd.NewParameter("action_status_id", "int", "inout", 0, 0)
		cmd.NewParameter("action_msg", "string", "inout", 200, "")
	})
}
