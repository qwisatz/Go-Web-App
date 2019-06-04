package main

import (
	"encoding/json"
	"fmt"
    "net/http"
    "log"
	"github.com/gorilla/mux"
	"time"

	"io/ioutil"
	"strings"

	"github.com/030/go-utils"
)

/*------------------------------------------------------------------*/
// STRUCT MODELS
/*------------------------------------------------------------------*/
type Todo struct {
    Name      string    `json:"name"`
    Completed bool      `json:"completed"`
    Due       time.Time `json:"due"`
}
 
type Todos []Todo

/*------------------------------------------------------------------*/
// HANDLER: INDEX
/*------------------------------------------------------------------*/
func indexHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(r) 

	// WRITING THINGS
	// WTIRE OUT TITLE OF HTML
	w.Write([]byte( "<html><head><title>MUX INDEX PAGE</title>"))

	w.Write([]byte( "<body>" + "<h2> HELLO WORLD! </h2>"))
	w.Write([]byte( "<p><a href='/1'>LINK TO JSON</a></p>"))

	w.Write([]byte( "</body></html>"))
}

/*------------------------------------------------------------------*/
// HANDLER: JSON PAGE
/*------------------------------------------------------------------*/
func todoJsonHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(r) 
	w.Write([]byte( "<html><head><title>JSON DATA</title>"))

	// CREATE JSONS FOR Todo STRUCT
    todos := Todos{
        Todo{Name: "Write presentation"},
        Todo{Name: "Host meetup"},
    }
 
	// WRITE JSON
	json.NewEncoder(w).Encode(todos)
	
	w.Write([]byte( "</body></html>"))
}

/*------------------------------------------------------------------*/
// HANDLER: CUSTOM LINK THAT STORES THE SAID LINK AND DISPLAYS
/*------------------------------------------------------------------*/
func customLinkHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	//PARSING NAME
	name := params["name"]
	fmt.Println("USER TYPED = ",name) 
	name = strings.ToLower(name)
	fmt.Println("LOWERCASE = ",name) 
	// -1 MEANS REPLACES ALL OCCURRENCES OF CHARACTER
	name = strings.Replace(name, "/", ".",-1)

	fmt.Println("LOWERCASE = ",name) 

	website := "http://" + name 
	fmt.Println(website) 
	fmt.Println("***********************************")
	fmt.Println(utils.URLExists(website))
	fmt.Println("***********************************")

	// IF WEBSITE FALSE = CUSTOM HANDLER
	// ELSE = SAVE AND DISPLAY WEBSITE
	if utils.URLExists(website) == false {

		w.Write([]byte( "<html><head><title>MUX INDEX PAGE</title>"))
		w.Write([]byte( "<body>" + "<h2>THIS PAGE DOESN'T EXIST!</h2>"))
		w.Write([]byte( "</body></html>"))

	} else {
		resp, err := http.Get(website)
		if err != nil {
			//log.Fatal(err)
		}
	
		// DEFER MEANS EXCUTE THIS LAST IN THIS BLOCK OF CODE
		defer resp.Body.Close()
	
		if resp.StatusCode == http.StatusOK {
			// READ THE HTML CODE
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				//log.Fatal(err)
			}
	
			// WRITE OUT THE WEBPAGE
			w.Write([]byte(bodyBytes))
		} 
	}
}
func customLinkHandler2(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	//PARSING NAME
	name := params["name"]
	fmt.Println("USER TYPED = ",name) 
	name = strings.ToLower(name)
	fmt.Println("LOWERCASE = ",name) 
	// -1 MEANS REPLACES ALL OCCURRENCES OF CHARACTER
	name = strings.Replace(name, "/", ".",-1)

	fmt.Println("LOWERCASE = ",name) 

	website := "http://" + name 
	fmt.Println(website) 

	resp, err := http.Get(website)

	if err != nil {
		w.Write([]byte( "<html><head><title>MUX INDEX PAGE</title>"))
		w.Write([]byte( "<body><h2>THIS PAGE DOESN'T EXIST!</h2>"))
		w.Write([]byte( "</body></html>"))

	} else {
		print(string(resp.StatusCode) + resp.Status)

		resp2, err := http.Get(website)
		if err != nil {
			//log.Fatal(err)
		}
	
		// DEFER MEANS EXCUTE THIS LAST IN THIS BLOCK OF CODE
		defer resp2.Body.Close()
	
		if resp2.StatusCode == http.StatusOK {
			// READ THE HTML CODE
			bodyBytes, err := ioutil.ReadAll(resp2.Body)
			if err != nil {
				//log.Fatal(err)
			}
	
			// WRITE OUT THE WEBPAGE
			w.Write([]byte(bodyBytes))
		} 
	}
}

/*------------------------------------------------------------------*/
// HANDLER FOR 404
/*------------------------------------------------------------------*/
func NotFoundHandler(w http.ResponseWriter, r *http.Request) { 
	w.Write([]byte( "<html><head><title>MUX INDEX PAGE</title>"))

	w.Write([]byte( "<body>" + "<h2>THIS PAGE DOESN'T EXIST!</h2>"))

	w.Write([]byte( "</body></html>"))
} 


/*------------------------------------------------------------------*/
// MUX SERVER
/*------------------------------------------------------------------*/
func main() {
	fmt.Println("/*------------------------------------------------------------------*/") 

	// CREATE MUX ROUTER
	r := mux.NewRouter()
	
	// Routes consist of a path and a handler function.
	// PATH OF THE PAGE
	// HANDLER FUNCTION IS THE METHOD, I.E. CREATES YOUR WEB PAGE
	r.HandleFunc("/", indexHandler)
	// THIS PAGE DISPLAYS JSON
	r.HandleFunc("/1", todoJsonHandler)
	// CUSTOM WEBSITE HANDLER
	r.HandleFunc("/{name:(?i)[A-Za-z0-9_/]{1,}/*}", customLinkHandler2)

	// THIS PAGE ACTIVATES THE METHOD TO STORE AND DISPLAY THE PAGE
	//r.HandleFunc("/2", storeDisplayHtml)

	// 404 HANDLER
	r.NotFoundHandler = http.HandlerFunc(NotFoundHandler)

    log.Fatal(http.ListenAndServe(":8000", r))
}