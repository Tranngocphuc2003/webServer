package main

import(
	"fmt"
	"log"
	"net/http"
	"strconv"
    "sync"		
)
var counter int
var mutex = &sync.Mutex{}
func incrementalCounter(w http.ResponseWriter, r *http.Request){
	mutex.Lock()
    counter++
    fmt.Fprintf(w, strconv.Itoa(counter))
    mutex.Unlock()
}
func formHandler(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err!=nil{
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request succesful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w,"name = %s\n",name)
	fmt.Fprintf(w,"address = %s\n", address)

}

func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello"{
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET"{
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprint(w, "hello!")
}

func main(){
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/increment", incrementalCounter)
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err!=nil{
		log.Fatal(err)
	}	
}