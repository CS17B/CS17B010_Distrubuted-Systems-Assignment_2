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

type Empreq struct {
	Name    string `json:"Name"`
	Time    string `json:"Time"`
	Date    string `json:"Date"`
	Members string `json:"Members"`
	Body    string `json:"Body"`
	mutex   bool   `json:"Locked"`
}

var x *Empreq

func returnAllEmpreqs(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	json.NewEncoder(w).Encode(Empreqs)
	mutex.Unlock()
}

func returnSingleEmpreq(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["Name"]
	for {

		for _, Empreq := range Empreqs {
			//	mutex.Lock()
			if Empreq.Name == key && Empreq.Locked == false {
				Empreq.Locked = true
				json.NewEncoder(w).Encode(Empreq)

				Empreq.Locked = false
			}
			mutex.Unlock()
		}

		return
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
		if Empreq.Time == time1 && Empreq.Date == date {
			rnd.HTML(w, http.StatusOK, "message", nil)
			return
		}
	}
	for {
		if o == "1" {
			x := &Empreq{
				Name:    n,
				Time:    t,
				Date:    d,
				Members: n,
				Body:    body,
				Locked:  false,
			}

			johnJSON, _ := json.Marshal(x)
			json.Unmarshal(johnJSON, &empreq)
			mutex.Lock()
			empreq.Locked = true
			Empreqs = append(Empreqs, empreq)
			json.NewEncoder(w).Encode(empreq)
			empreq.Locked = false
			//time.Sleep(2 * time.Second)
			mutex.Unlock()

		} else {
			x := &Empreq{
				Name:    n,
				Time:    t,
				Date:    d,
				Members: m,
				Body:    body,
				Locked:  false,
			}

			johnJSON, _ := json.Marshal(x)
			json.Unmarshal(johnJSON, &empreq)
			mutex.Lock()
			empreq.Locked = true
			Empreqs = append(Empreqs, empreq)
			json.NewEncoder(w).Encode(empreq)
			mutex.Unlock()
			empreq.Locked = false
			v := m
			m := n + " is the Instructor" + "," + m

			keyVal := strings.Split(v, ",")
			for i := 0; i < len(keyVal); i++ {
				x := &Empreq{
					Name:    keyVal[i],
					Time:    t,
					Date:    d,
					Members: m,
					Body:    body,
					Locked:  false,
				}

				johnJSON, _ := json.Marshal(x)
				json.Unmarshal(johnJSON, &empreq)
				mutex.Lock()
				empreq.Locked = true
				Empreqs = append(Empreqs, empreq)
				json.NewEncoder(w).Encode(empreq)
				fmt.Printf("%s  ", keyVal[i])
				mutex.Unlock()
				empreq.Locked = false

			}
		}
		return
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
			if Empreq.Name == id1 && Empreq.Date == id2 && Empreq.Time == id3 && Empreq.Locked == false {
				Empreq.Locked = true
				Empreqs = append(Empreqs[:index], Empreqs[index+1:]...)
				Empreq.Locked = false
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
	for {
		//mutex.Lock()
		for _, Empreq := range Empreqs {
			if Empreq.Name == id1 && Empreq.Date == id2 && Empreq.Time == id3 && Empreq.Locked == false {
				Empreq.Locked = true
				x = &Empreq
				break
			}
		}
		for index, Empreq := range Empreqs {
			if Empreq.Name == id1 && Empreq.Date == id2 && Empreq.Time == id3 {
				Empreqs = append(Empreqs[:index], Empreqs[index+1:]...)
				Empreq.Locked = false
			}
		}

		rnd.HTML(w, http.StatusOK, "new", x)
		//	mutex.Unlock()

		return
	}
}

func returnmeetings(w http.ResponseWriter, r *http.Request) {
	for _, Empreq := range Empreqs {
		y := Empreq.Members
		keyVal := strings.Split(y, ",")
		var keyValsub bool = strings.Contains(keyVal[0], "Instructor")
		z := len(keyVal)
		if z > 1 && !keyValsub == true {
			json.NewEncoder(w).Encode(Empreq)
		}
	}
}

func returnsingeuserspecificDateTimeEmpreq(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id1 := vars["Name"]
	id2 := vars["Date"]
	id3 := vars["Time"]
	for {

		for _, Empreq := range Empreqs {
			mutex.Lock()
			if Empreq.Name == id1 && Empreq.Date == id2 && Empreq.Time == id3 && Empreq.Locked == false {
				Empreq.Locked = true
				json.NewEncoder(w).Encode(Empreq)
				Empreq.Locked = false
				mutex.Unlock()
				return
			}
		}
	}

}

func returnsingeuserspecificDateEmpreq(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id1 := vars["Name"]
	id2 := vars["Date"]
	for {
		//mutex.Lock()
		for _, Empreq := range Empreqs {
			mutex.Lock()
			if Empreq.Name == id1 && Empreq.Date == id2 && Empreq.Locked == false {

				Empreq.Locked = true
				json.NewEncoder(w).Encode(Empreq)
				Empreq.Locked = false
			}
			mutex.Unlock()
		}
		//	mutex.Unlock()
		return
	}
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/Empreqs", returnAllEmpreqs)
	myRouter.HandleFunc("/edit/", edithandler)
	myRouter.HandleFunc("/Empreqs/Empreq", createNewEmpreq)
	myRouter.HandleFunc("/Empreqs/Delete/{Name}/{Date}/{Time}", deleteEmpreq)
	myRouter.HandleFunc("/Empreqs/Update/{Name}/{Date}/{Time}", updateEmpreq)
	myRouter.HandleFunc("/Empreqs/{Name}/{Date}/{Time}", returnsingeuserspecificDateTimeEmpreq)
	myRouter.HandleFunc("/F10/", returnmeetings)
	myRouter.HandleFunc("/Empreqs/{Name}/{Date}", returnsingeuserspecificDateEmpreq)
	myRouter.HandleFunc("/Empreqs/{Name}", returnSingleEmpreq)
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {

	Empreqs = []Empreq{
		{Name: "F1", Time: "02:00", Date: "22-03-2021", Members: "F1", Body: "I want to meet Mr.J"},
		{Name: "F1", Time: "03:00", Date: "22-03-2021", Members: "F1", Body: "I want to attend ethics class"},
		{Name: "F2", Time: "04:00", Date: "22-03-2021", Members: "F2", Body: "I want to meet Mr.X"},
		{Name: "F2", Time: "05:00", Date: "22-03-2021", Members: "F2", Body: "I want to meet Mr.y"},
	}
	handleRequests()
}
