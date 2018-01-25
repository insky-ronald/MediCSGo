package main

import (
	"net/http"
	"ibsi/utils"
	"ibsi/template"
	"ibsi/session"
)

type LoginData struct {
	UserName string
	LoginAttempts int
}

func init() {
	template.NewController(template.Controller {
		Pid: "login",
		// Root: "",
		Template: "login",
		OnInitHandlers: func(ts *template.Controller) {
			ts.Add("/{login:login\\/?}")
		},
		OnInitTemplateCustomData: func(r *http.Request, p *template.Page) interface{} {		
			return session.GetSession(r)
		},
		OnInitCustomData: func(r *http.Request, nav *utils.Navigator) interface{} {
			return nil
		},
		OnInitCallbackData: func(nav *utils.Navigator) {
		},
	})
}
