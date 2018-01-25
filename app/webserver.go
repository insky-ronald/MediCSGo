package main

import (
    "fmt"
    "net/http"
    "log"
	"os"
	"bufio"
	"io/ioutil"
	"path/filepath"
	"github.com/gorilla/mux"
	"ibsi/system"
	"ibsi/dbase"
	"ibsi/bundle"
	"ibsi/session"
	"ibsi/crud"
	"ibsi/template"
	_ "ibsi/easy"
)

func init() {
	system.Settings.ShowConnections()
	log.Println(fmt.Sprintf("Web server is listening at port %d", system.Settings.Port))
	if system.Settings.Domain != "" {
		log.Println(fmt.Sprintf("Web server domain is %s", system.Settings.Domain))
	}
}

func ServeVideo(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "video.mp4")
}

func CatchAll(w http.ResponseWriter, r *http.Request) {

}

func GetImage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var image string = vars["image"]

	ext := filepath.Ext(image)
	w.Header().Set("Content-Type", "image/" + ext[1:len(ext)])

	// fileName, err := filepath.Abs("../webserver/images/"+image)
	// if err != nil {
		// log.Fatal(err)
	// }

	// content, err := ioutil.ReadFile(fileName)
	// content, err := ioutil.ReadFile("images/"+image)
	content, err := ioutil.ReadFile("images/"+image)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	w.Write(content)
}

func Api(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	var api, cmd string = mux.Vars(r)["api"], mux.Vars(r)["cmd"]
	
	if api == "session" {
		if cmd == "alive" {
			// jsonText, err := json.MarshalIndent(config, "", "\t") // \t is formatted, remove it to minify
			// if err != nil {			
				// http.Error(w, err.Error(), 500)
			// } else {
				// w.Write(jsonText)			
			// }
		} else if cmd == "login" {
			session.Login(w, r) // main.login.go
		} else if cmd == "logout" {
			session.Logout(w, r) // main.login.go
		}
	}
}

func main2() {
// func main() {
	defer dbase.CloseConnections()
	
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		if input == "x" {
			break
		}
	}

	if scanner.Err() != nil {
		// handle error.
	}
}

func Log(handler http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
		// handler(w, r)
    })
}

func run() {
	// https://www.kaihag.com/https-and-go/
	// FOR HTTPS
	// http.ListenAndServeTLS(":8081", "cert.pem", "key.pem", nil)

	// FinalHandler  is initialized in config.go init()
	// FinalHandler = context.ClearHandler(http.DefaultServeMux)
	err := http.ListenAndServe(fmt.Sprintf(":%d", system.Settings.Port), system.FinalHandler)
	// err := http.ListenAndServe(fmt.Sprintf(":%d", config.Port), Log(FinalHandler))
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

func run2() {
	// https://www.kaihag.com/https-and-go/
	// FOR HTTPS
	// http.ListenAndServeTLS(":8081", "cert.pem", "key.pem", nil)

	// FinalHandler  is initialized in config.go init()
	// FinalHandler = context.ClearHandler(http.DefaultServeMux)
	err := http.ListenAndServe(fmt.Sprintf(":%d", 8080), system.FinalHandler)
	// err := http.ListenAndServe(fmt.Sprintf(":%d", config.Port), FinalHandler)
	// err := http.ListenAndServe(fmt.Sprintf(":%d", config.Port), Log(FinalHandler))
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

// func main0() {
func main() {
	defer dbase.CloseConnections()

	// Router is initialized in config.go init()
	// Router.HandleFunc("/", CatchAll)
	// Router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println("2. here")
		// r.URL.Path = "/login"
		// if sn := session.GetSession(r); sn != nil {
			// if sn.Authenticated {
				// r.URL.Path = sn.Home
			// }
		// }
		
		// Router.ServeHTTP(w,r)
	// })
	system.Router.HandleFunc("/video/{video}", ServeVideo)//.Host(config.Domain)  // Test!
	system.Router.HandleFunc("/images/{image}", GetImage)//.Host(config.Domain)
	system.Router.HandleFunc("/api/{api}/{cmd}", Api)//.Host(config.Domain)
	// Router.HandleFunc("/api/session/alive", KeepAlive).Host(config.Domain)
	
	system.Router.HandleFunc("/loadcss/{path}/css/{css}", bundle.LoadCss)//.Host(config.Domain)
	system.Router.HandleFunc("/loadscript/{path}/scripts/{script}", bundle.LoadScript)//.Host(config.Domain)
	
	system.Router.HandleFunc("/app/{pid}", template.HandleDynamicTemplate)
	
	system.Router.HandleFunc("/app/get/list/{list:lookup\\/?}", crud.HandleLookup)
	system.Router.HandleFunc("/app/get/list/{list}", crud.HandleListGeneric)
	system.Router.HandleFunc("/app/get/edit/{list}", crud.HandleListGeneric)
	system.Router.HandleFunc("/app/get/update/{list}", crud.HandleListGeneric)
	
	// http.Handle("/", Router)
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	go run()
	// go run2()

	fmt.Println("Press x <ENTER> to shutdown server")
	fmt.Print("> ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Print("> ")
		input := scanner.Text()
		if input == "x" {
			break
		}
	}

	if scanner.Err() != nil {
		// handle error.
	}
	// os.Exit(0)
}
