package main

import (
	"fmt"
	"net/http"
	// "ibsi/dbase"
	"ibsi/crud"
)

func init() {

	ch := crud.CrudHandler {
		Name: "lookup",
		OnInitCrud: func(crud map[string]bool) {
			crud["add"] = false
			crud["edit"] = false
			crud["delete"] = false
		},
	}
	
	Router.HandleFunc(fmt.Sprintf("/app/get/list/%s", ch.Name), func(w http.ResponseWriter, r *http.Request) {
		// list handler
		ch.ListDataSource = "DBApp." + r.URL.Query().Get("name")
		// ch.ListDataSource = "DBMedics." + r.URL.Query().Get("name")
		crud.HandleList(w, r, ch)
	})
}
