package main

import (
	"net/http"
	"ibsi/utils"
	"ibsi/template"
	"ibsi/session"
	"ibsi/dbase"
	// "ibsi/system"
)

// AddTemplateController is in ibsi.controller.template.go
func init() {
	
	type PageData struct {
		CustomerService dbase.TDataTableRows `json:"data"`
	}
	
	template.NewController(template.Controller {
		Pid: "providers",
		// Root: "/app",
		Template: "template-2",
		OnInitHandlers: func(ts *template.Controller) {
			ts.Add("/app/{pid:providers\\/?}")
		},
		OnInitPageData: func(r *http.Request, p *template.Page) {
			vid := session.GetVisitorId(r)
			
			p.Title = "Providers"
			p.Nav.WindowTitle = "Providers"
			
			p.Nav.CustomData = map[string]interface{}{
				"countries": dbase.Connections["DBApp"].OpenDataTable("GetCountries", dbase.TParameters{"action":1, "visit_id":vid}).GetRows(),
				"specialisation": dbase.Connections["DBApp"].OpenDataTable("GetDoctorSpecialisation", dbase.TParameters{"lookup":1, "visit_id":vid}).GetRows(),
			}
			
			utils.NewNavigatorItem(p.Nav, "medical", "Medical", func(item *utils.NavigatorItem) {
			
				utils.NewMenuItem(item, func(s *utils.MenuItem) {
					s.ID = "doctors"
					s.Action = "admin"				
					s.Title = "Doctors"
					s.Icon = "doctor"
					s.Url = "app/providers"
					s.Run = "ProviderMasterDetailView"
					s.Css = "*"
					s.Params["provider_type"] = "D"
				})
				
				// utils.NewMenuItem(item, func(s *utils.MenuItem) {
					// s.ID = "members"
					// s.Action = "admin"
					// s.Title = "Claims Processing"
					// s.Icon = "users"
					// s.Url = "app/claims-entry"
				// })
			
				// utils.NewMenuItem(item, func(s *utils.MenuItem) {
					// s.ID = "enquiries"
					// s.Action = "admin"
					// s.Title = "Enquiries"
					// s.Icon = "search"
				// })
				
				// utils.NewMenuItem(item, func(s *utils.MenuItem) {
					// s.ID = "reports"
					// s.Action = "admin"
					// s.Title = "Reports"
					// s.Icon = "table"
					// s.Url = "engine/sys-reports"
				// })
			})
			
			utils.NewNavigatorItem(p.Nav, "travel", "Travel", func(item *utils.NavigatorItem) {
			
				// utils.NewMenuItem(item, func(s *utils.MenuItem) {
					// s.ID = "users"
					// s.Action = "sys-users"
					// s.Title = "Users"
					// s.Icon = "user"
					// s.Url = "engine/sys-users"
				// });
			})
		},
	})
}
