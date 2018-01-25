package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"ibsi/dbase"
	"ibsi/utils"
	"ibsi/template"
)

type SalesData struct {
	NewRecord int `json:"newRecord"`
	Id int `json:"sale_id"`
	Sales dbase.TDataTableRows `json:"data"`
	Customer dbase.TDataTableRows `json:"customer"`
	Terms dbase.TDataTableRows `json:"terms"`
	LeadTimes dbase.TDataTableRows `json:"lead_times"`
	PriceValidities dbase.TDataTableRows `json:"price_validities"`
}

func init() {
	template.NewController(template.Controller {
		Pid: "sales",
		// Root: "/app",
		// Format: "/app/%v",
		// Params: "keyid:[0-9]+",
		// ValueParam: "[0-9]+", // get this is "keyid"
		Template: "template-2",
		OnInitHandlers: func(ts *template.Controller) {
			ts.Add("/app/sales/{keyid:[0-9]+\\/?}")
			ts.Add("/app/sales/{new}/{keyid:[0-9]+\\/?}")
		},
		OnInitPageData: func(r *http.Request, p *template.Page) {
			vars := mux.Vars(r) 
			id := utils.StrToInt(vars["keyid"])
			
			sales := dbase.Connections["DBApp"].OpenDataTable("GetSales", dbase.TParameters{
				"id": id,
				"action": 10,
			})
			
			if id == 0 {
				p.Title = "New Sales Transaction"
				p.Nav.PageTitle = "New Sales Transaction"
				p.Nav.WindowTitle = "New Sales Transaction"
			} else {
				p.Title = sales.Get("reference_no").(string)
				p.Nav.PageTitle = sales.Get("reference_no").(string)
				p.Nav.WindowTitle = sales.Get("reference_no").(string)
			}
			
			customer := dbase.Connections["DBApp"].OpenDataTable("GetCustomers", dbase.TParameters{
				"id": sales.Get("customer_id"),
				"action": 10,
			})
			
			p.Nav.CustomData = SalesData{
				Id: id,
				NewRecord: 0,
				Sales: sales.GetRows(),
				Customer: customer.GetRows(),
				Terms: dbase.Connections["DBApp"].OpenDataTable("GetTerms", dbase.TParameters{}).GetRows(),
				LeadTimes: dbase.Connections["DBApp"].OpenDataTable("GetLeadTimes", dbase.TParameters{}).GetRows(),
				PriceValidities: dbase.Connections["DBApp"].OpenDataTable("GetPriceValidities", dbase.TParameters{}).GetRows(),
			}
			
			utils.NewNavigatorItem(p.Nav, "main", "Main", func(item *utils.NavigatorItem) {
				utils.NewMenuItem(item, func(s *utils.MenuItem) {
					s.ID = "details"
					s.Action = "sales"
					s.Title = "Details"
					s.Description = "Sales Transaction"
					s.Icon = "table-edit"
					s.Run = "SaleDetailsView"
					s.Css = "*"
				})
			})
		},
	})
}
