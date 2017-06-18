package main

import (
	"encoding/json"
	"fmt"
	_ "log"
	"net/http"
	_ "os"
)

type Family struct {
	Father string
}

type Familys struct {
	All []Family
}

//fatherAge int
//mother string `json:"mother_name"`
//motherAge int }

func main() {
	http.HandleFunc("/test", handler)
	http.ListenAndServe("localhost:8080", nil)
}

//	fatherAge : 30,
//mother : "sun huimin",
//motherAge : 18 }

func handler(w http.ResponseWriter, r *http.Request) {

	//obj := family{father: "bizhenchao"}
	//	fatherAge : 30,
	//mother : "sun huimin",
	//motherAge : 18 }
	//var objs = familys{[]family{obj}}

	var objs Familys
	objs.All = append(objs.All, Family{Father: "bizhenchao"})
	objs.All = append(objs.All, Family{Father: "1234"})

	b, err := json.Marshal(objs)

	if err != nil {
		fmt.Fprintf(w, "Error is  %s", err)
	} else {
		fmt.Println(string(b))
		fmt.Fprint(w, string(b))

	}

}
