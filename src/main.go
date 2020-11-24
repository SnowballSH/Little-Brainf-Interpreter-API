package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"

	"./brainf"

	"encoding/json"
	"io/ioutil"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Brainf Interpreter API!\nUse /run to print result using POST\nUse /test to run the test program")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/run", runCode).Methods("POST")
	myRouter.HandleFunc("/test", testCode)
	log.Fatal(http.ListenAndServe(":8888", myRouter))
}

func runCode(w http.ResponseWriter, r *http.Request) {
	cd, _ := ioutil.ReadAll(r.Body)

	var code string
	json.Unmarshal(cd, &code)

	brainf.RunCode(string(code), w)
}

func testCode(w http.ResponseWriter, r *http.Request) {
	val := `Hello World:
+[-->-[>>+>-----<<]<--<---]>-.>>>+.>>..+++[.>]<<<<.+++.------.<<-.>>>>+.
`

	jsonValue, _ := json.Marshal(val)

	resp, err := http.Post("http://localhost:8888/run", "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		fmt.Fprintf(w, fmt.Sprintf("%v", err))
	}

	fmt.Fprintf(w, "Result:\n")
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Fprintf(w, string(body))
}

func main() {
	handleRequests()
}
