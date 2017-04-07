package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"time"
)

const NaturalDate = "January 2, 2006"

type timestamp struct {
	Unix    *int64  `json:"unix"`
	Natural *string `json:"natural"`
}

func (t timestamp) writeJson(w http.ResponseWriter) {
	js, err := json.Marshal(t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func parse(s string) (t time.Time, err error) {
	t, err = time.Parse(NaturalDate, s)
	if err != nil {
		i, e := strconv.ParseInt(s, 10, 32)
		err = e
		t = time.Unix(i, 0)
	}
	return
}

func handler(w http.ResponseWriter, r *http.Request) {
	p := r.URL.Path[1:]
	ts := &timestamp{}
	if t, err := parse(p); err == nil {
		u := t.Unix()
		n := t.Format(NaturalDate)
		ts.Unix = &u
		ts.Natural = &n
	}
	ts.writeJson(w)
}

func main() {
	p := os.Getenv("PORT")
	if p == "" {
		p = "3750"
	}
	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+p, nil)
}
