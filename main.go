package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Vehicle struct {
	Name  string  `json:"nom"`
	Speed float32 `json:"vitesse,omitempty"` // won't be displayed if no initial value
}

type Dragster struct {
	v       Vehicle // member, manually initialize in the struct
	Vehicle         // anonymous member
}

func (v Vehicle) Start() {
	fmt.Println("Vehicle", v.Name, " starting...")
}

func (v Vehicle) Run() {
	fmt.Println("Vehicle", v.Name, " running...")
}

func (d Dragster) Run() {
	d.v.Run()
}

func SpeedHandler(w http.ResponseWriter, r *http.Request) {

	v := Vehicle{Name: "monvehicule"}

	jsonBytes, err := json.Marshal(v)
	if err != nil {
		log.Printf("Error JSON marshal: %v", err)
		return
	}

	fmt.Fprintf(w, "%v", string(jsonBytes))

}

func main() {
	//var d Dragster // variable declraration (everything zeroed)
	//d := Dragster{v: Vehicle{name: "toto", speed: 70}} // declaration with init
	//d.Start() // runs even if Dragster.start() doesn"t exist but name is not defined
	//d.Run()

	http.HandleFunc("/speed", SpeedHandler)
	http.ListenAndServe(":8080", nil)
}
