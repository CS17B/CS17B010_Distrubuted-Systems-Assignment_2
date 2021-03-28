package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/gorilla/mux"
	"github.com/thedevsaddam/renderer"
)

var rnd *renderer.Render

var mutex = &sync.Mutex{}
var Empreqs []Empreq

var recover bool = false
var timepos = [24]string{"00:00", "01:00", "02:00", "03:00", "04:00", "05:00", "06:00", "07:00",
	"08:00", "09:00", "10:00", "11:00", "12:00", "13:00", "14:00", "15:00", "16:00",
	"17:00", "18:00", "19:00", "20:00", "21:00", "22:00", "23:00"}

type Empreq struct {
	Name    string `json:"Name"`
	Time    string `json:"Time"`
	Date    string `json:"Date"`
	Members string `json:"Members"`
	Body    string `json:"Body"`
}

var x *Empreq

func returnAllEmpreqs(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	json.NewEncoder(w).Encode(Empreqs)
	mutex.Unlock()
}
func contains(timepos [24]string, e string) bool {
	for _, a := range timepos {
		if a == e {
			return true
		}
	}
	return false
}

func returnSingleEmpreq(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["Name"]

	for _, Empreq := range Empreqs {
		mutex.Lock()
		if Empreq.Name == key {
			json.NewEncoder(w).Encode(Empreq)
		}
		mutex.Unlock()
	}

}

func createNewEmpreq(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	var empreq Empreq
	n := r.FormValue("n")
	t := r.FormValue("t")
	d := r.FormValue("d")
	o := r.FormValue("o")
	m := r.FormValue("m")
	body := r.FormValue("body")

	time1 := t
	date := d

	for _, Empreq := range Empreqs {
		if Empreq.Time == time1 && Empreq.Date == date && Empreq.Name == n {
			rnd.HTML(w, http.StatusOK, "message", nil)
			return
		}
	}

	if !contains(timepos, t) {
		rnd.HTML(w, http.StatusOK, "time", nil)
		return
	}

	if o == "1" {
		x := &Empreq{
			Name:    n,
			Time:    t,
			Date:    d,
			Members: n,
			Body:    body,
		}

		johnJSON, _ := json.Marshal(x)
		json.Unmarshal(johnJSON, &empreq)
		mutex.Lock()
		Empreqs = append(Empreqs, empreq)
		json.NewEncoder(w).Encode(empreq)
		mutex.Unlock()
	} else {
		v := m
		p := n + "," + m
		fmt.Printf("%s \n", p)
		for _, Empreq := range Empreqs {
			if Empreq.Time == time1 && Empreq.Date == date && strings.Contains(p, Empreq.Name) {
				rnd.HTML(w, http.StatusOK, "message", nil)
				recover = true
				return
			}
		}

		x := &Empreq{
			Name:    n,
			Time:    t,
			Date:    d,
			Members: m,
			Body:    body,
		}

		johnJSON, _ := json.Marshal(x)
		json.Unmarshal(johnJSON, &empreq)
		mutex.Lock()
		Empreqs = append(Empreqs, empreq)
		json.NewEncoder(w).Encode(empreq)
		mutex.Unlock()

		m := n + " is the Instructor" + "," + m

		keyVal := strings.Split(v, ",")
		for i := 0; i < len(keyVal); i++ {
			x := &Empreq{
				Name:    keyVal[i],
				Time:    t,
				Date:    d,
				Members: m,
				Body:    body,
			}

			johnJSON, _ := json.Marshal(x)
			json.Unmarshal(johnJSON, &empreq)
			mutex.Lock()
			Empreqs = append(Empreqs, empreq)
			json.NewEncoder(w).Encode(empreq)
			mutex.Unlock()
			fmt.Printf("%s  ", keyVal[i])
		}
	}

}

func deleteEmpreq(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id1 := vars["Name"]
	id2 := vars["Date"]
	id3 := vars["Time"]
	for {

		for index, Empreq := range Empreqs {
			mutex.Lock()
			if Empreq.Name == id1 && Empreq.Date == id2 && Empreq.Time == id3 {
				Empreqs = append(Empreqs[:index], Empreqs[index+1:]...)
				mutex.Unlock()
				return
			}
		}
	}

}
func init() {
	opts := renderer.Options{
		ParseGlobPattern: "./tpl/*.html",
	}
	rnd = renderer.New(opts)
}

func edithandler(w http.ResponseWriter, r *http.Request) {
	rnd.HTML(w, http.StatusOK, "edit", nil)
}

func updateEmpreq(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id1 := vars["Name"]
	id2 := vars["Date"]
	id3 := vars["Time"]
	mutex.Lock()
	for _, Empreq := range Empreqs {
		if Empreq.Name == id1 && Empreq.Date == id2 && Empreq.Time == id3 {
			x = &Empreq
			break
		}
	}
	rnd.HTML(w, http.StatusOK, "new", x)
	if recover {
		for index, Empreq := range Empreqs {
			if Empreq.Name == id1 && Empreq.Date == id2 && Empreq.Time == id3 {
				Empreqs = append(Empreqs[:index], Empreqs[index+1:]...)
			}
		}
		recover = false
	}
	mutex.Unlock()
}

func returnmeetings(w http.ResponseWriter, r *http.Request) {
	for _, Empreq := range Empreqs {
		y := Empreq.Members
		keyVal := strings.Split(y, ",")
		var keyValsub bool = strings.Contains(keyVal[0], "Instructor")
		z := len(keyVal)
		if z > 1 && !keyValsub {
			json.NewEncoder(w).Encode(Empreq)
		}
	}
}

func returnsingeuserspecificDateTimeEmpreq(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id1 := vars["Name"]
	id2 := vars["Date"]
	id3 := vars["Time"]

	for _, Empreq := range Empreqs {
		if Empreq.Name == id1 && Empreq.Date == id2 && Empreq.Time == id3 {
			mutex.Lock()
			json.NewEncoder(w).Encode(Empreq)
			mutex.Unlock()
			return
		}
	}

}

func returnsingeuserspecificDateEmpreq(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id1 := vars["Name"]
	id2 := vars["Date"]

	for _, Empreq := range Empreqs {
		if Empreq.Name == id1 && Empreq.Date == id2 {
			mutex.Lock()
			json.NewEncoder(w).Encode(Empreq)
			mutex.Unlock()
		}

	}

}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/empreqs", returnAllEmpreqs)
	myRouter.HandleFunc("/edit/", edithandler)
	myRouter.HandleFunc("/empreqs/empreq", createNewEmpreq)
	myRouter.HandleFunc("/empreqs/delete/{Name}/{Date}/{Time}", deleteEmpreq)
	myRouter.HandleFunc("/empreqs/update/{Name}/{Date}/{Time}", updateEmpreq)
	myRouter.HandleFunc("/empreqs/{Name}/{Date}/{Time}", returnsingeuserspecificDateTimeEmpreq)
	myRouter.HandleFunc("/f10/", returnmeetings)
	myRouter.HandleFunc("/empreqs/{Name}/{Date}", returnsingeuserspecificDateEmpreq)
	myRouter.HandleFunc("/empreqs/{Name}", returnSingleEmpreq)
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {

	Empreqs = []Empreq{
		{Name: "f1", Time: "02:00", Date: "22-03-2021", Members: "f1", Body: "I want to attend class A"},
		{Name: "f1", Time: "03:00", Date: "24-03-2021", Members: "f1", Body: "I want to attend DS class"},
		{Name: "f2", Time: "04:00", Date: "22-03-2021", Members: "f2", Body: "I want to meet Mr.Raghu"},
		{Name: "f2", Time: "05:00", Date: "23-03-2021", Members: "f2", Body: "I want to meet Mr.J"},
		{Name: "f2", Time: "06:00", Date: "29-03-2021", Members: "f1,f3,f4", Body: "Discussion about class commitee meeting"},
		{Name: "f1", Time: "06:00", Date: "29-03-2021", Members: "f2 is the Instructor,f1,f3,f4", Body: "Discussion about class commitee meeting"},
		{Name: "f3", Time: "06:00", Date: "29-03-2021", Members: "f2 is the Instructor,f1,f3,f4", Body: "Discussion about class commitee meeting"},
		{Name: "f4", Time: "06:00", Date: "29-03-2021", Members: "f2 is the Instructor,f1,f3,f4", Body: "Discussion about class commitee meeting"},
	}
	handleRequests()
}
