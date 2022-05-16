package main

import(
	"log"
	"net/http"
	"fmt"
)

func sumHandler(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil{
		fmt.Fprintf(w, "ParseForm() err: %v",err)
		return
	}
	fmt.Fprintf(w , "POST request succesful\n")
	num1 := r.FormValue("num1")
	num2 := r.FormValue("num2")
	res,err := sum(&num1, &num2)
	if err == nil{

	}
	fmt.Fprintf(w, "Sum = %s\n",res)
}

func main(){
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/sum", sumHandler)
	fmt.Printf("Starting server at port: 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}