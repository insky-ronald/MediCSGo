package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"ibsi/utils"
	"ibsi/template"
	"ibsi/session"
	"ibsi/dbase"
)

// AddTemplateController is in ibsi.controller.template.go
func init() {
	
	template.NewController(template.Controller {
		Pid: "service",
		// Root: "/app",
		Template: "template-2",
		OnInitHandlers: func(ts *template.Controller) {
			// ts.Add("/app/claim/{id\\/?}")
			// ts.Add("/app/service/{keyid:new}/{memberid:[0-9]+\\/?}")
			ts.Add("/app/service/{module}/{keyid:[0-9]+\\/?}")
			// ts.Add("/app/claim/{new}/{keyid:[0-9]+\\/?}")
		},
		OnInitPageData: func(r *http.Request, p *template.Page) {
			vid := session.GetVisitorId(r)
			vars := mux.Vars(r)
			
			var id int64 = 0
			serviceType := vars["module"]
			if vars["keyid"] != "new" {
				id = utils.StrToInt64(vars["keyid"])
			}
			
			dbService := dbase.Connections["DBApp"].OpenDataTable("GetService", dbase.TParameters{"id":id, "service_type":serviceType, "visit_id":vid})
			
			var serviceSubType string = ""			
			if id == 0 {
				serviceSubType = "take from url parameter type?"
			} else {
				serviceSubType = dbService.Get("service_sub_type").(string)
			}
			
			dbServiceSubType := dbase.Connections["DBApp"].OpenDataTable("GetServiceSubType", dbase.TParameters{"id":id, "service_type":serviceType, "sub_type":serviceSubType, "visit_id":vid})
			dbGopCalculationDates := dbase.Connections["DBApp"].OpenDataTable("GetGopCalculationDates", dbase.TParameters{"id":id, "visit_id":vid})			
			if dbGopCalculationDates.RowCount() == 0 {
				// ...
			}
			
			dbGopEstimates := dbase.Connections["DBApp"].OpenDataTable("GetGopEstimates", dbase.TParameters{"id":id, "visit_id":vid})
			if dbGopEstimates.RowCount() == 0 {
				dbGopEstimates.Add(dbase.TDataTableRow{
					"id": id,
					"average_cost": 0,
					"average_los": 0,
					"estimated_cost": 0,
					"estimated_los": 0,
					"estimated_provider_cost": 0,
					"estimated_provider_los": 0,
				})
			}
			
			if id == 0 {
				p.Title = "New Service"
				p.Nav.PageTitle = "New Service"
				p.Nav.WindowTitle = "New Service"
			} else {
				p.Title = dbService.Get("service_no").(string)
				p.Nav.PageTitle = dbService.Get("service_no").(string)
				p.Nav.WindowTitle = dbService.Get("service_no").(string)
			}
			
			p.Nav.CustomData = map[string]interface{}{
				"newRecord": utils.Ifx(id == 0, 1, 0),
				"service_id": id,
				"module_type": serviceType,
				"service_sub_type": serviceSubType,
				"service_type_name": dbServiceSubType.Get("service_display_name").(string),
				// "service_type_name": dbServiceSubType.Get("service_name").(string),
				"data": dbService.GetRows(),
				"sub_type_data": dbServiceSubType.GetRows(),
				"estimates": dbGopEstimates.GetRows(),
				"calculation_dates": dbGopCalculationDates.GetRows(),
				"countries": dbase.Connections["DBApp"].OpenDataTable("GetCountries", dbase.TParameters{"action":1, "visit_id":vid}).GetRows(),
			}
			
			description := dbServiceSubType.Get("service_display_name").(string)
			
			utils.NewNavigatorItem(p.Nav, "service", description, func(item *utils.NavigatorItem) {
				
				run := "ServiceDetailsView"
				
				utils.NewMenuItem(item, func(s *utils.MenuItem) {
					s.ID = "details"
					s.Action = "admin"				
					s.Title = "Details"
					s.Description = description
					s.Icon = "table-edit"
					s.Css = "*"
					s.Run = run
					s.Params["service_id"] = id
				})
				
				if serviceType == "inv" || serviceType == "gop" {
					utils.NewMenuItem(item, func(s *utils.MenuItem) {
						s.ID = "breakdown"
						s.Action = "admin"				
						s.Title = "Breakdown"
						s.Icon = "table-edit"
						s.Url = "app/service-breakdown"
						s.Params["service_id"] = id
					})
				}
			
				utils.NewMenuItem(item, func(s *utils.MenuItem) {
					s.ID = "documents"
					s.Title = "Documents"
					s.Icon = "documents"
					s.Action = "admin"
					s.Url = "app/claim-documents"
					s.Params["claim_id"] = 0
					s.Params["service_id"] = id
				})
			})
		},
	})
	
	
}
