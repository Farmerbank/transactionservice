package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"
)

var routes = Routes{
	Route{"List", "GET", "/Transactions", List},
	Route{"ListDebit", "GET", "/Transactions/Debit", ListDebit},
	Route{"ListCredit", "GET", "/Transactions/Credit", ListDebit},
}

var transactions = Transactions{
	Transaction{"Debit", "$50", "amazon.com", time.Date(2017, 11, 17, 20, 34, 58, 651387237, time.UTC)},
	Transaction{"Debit", "$154", "ebay.com", time.Date(2017, 11, 17, 20, 34, 58, 651387237, time.UTC)},
	Transaction{"Debit", "$15", "netflix.com", time.Date(2017, 11, 18, 20, 34, 58, 651387237, time.UTC)},
	Transaction{"Debit", "$40", "etsy.com", time.Date(2017, 11, 18, 11, 34, 58, 651387237, time.UTC)},
	Transaction{"Debit", "$500", "homedepot", time.Date(2017, 11, 18, 20, 34, 58, 651387237, time.UTC)},
	Transaction{"Debit", "$50", "wholefoods", time.Date(2017, 11, 18, 20, 34, 58, 651387237, time.UTC)},
	Transaction{"Debit", "$70", "slagerij van kampen", time.Date(2017, 11, 19, 20, 34, 58, 651387237, time.UTC)},
	Transaction{"Debit", "$55", "conrad.nl", time.Date(2017, 11, 19, 20, 34, 58, 651387237, time.UTC)},
	Transaction{"Credit", "$4200", "Carebear Inc.", time.Date(2017, 11, 20, 20, 34, 58, 651387237, time.UTC)},
}

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

type Transaction struct {
	Type         string    `json:"Type"`
	Amount       string    `json:"Amount"`
	CounterParty string    `json:"CounterParty"`
	Date         time.Time `json:"Date"`
}

type Transactions []Transaction

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

func main() {
	router := NewRouter()
	router.HandleFunc("/", logHandler(MessageHandler))
	spew.Dump(router)
	log.Fatal(http.ListenAndServe(":3000", router))

}

func List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if err := json.NewEncoder(w).Encode(transactions); err != nil {
		panic(err)
	}
}

func ListDebit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if err := json.NewEncoder(w).Encode(transactions.getDebit()); err != nil {
		panic(err)
	}
}

func ListCredit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if err := json.NewEncoder(w).Encode(transactions.getCredit()); err != nil {
		panic(err)
	}
}

func (t Transactions) getDebit() Transactions {
	trans := make(Transactions, 0)
	for _, tt := range t {
		if tt.Type == "Debit" {
			trans = append(trans, tt)
		}
	}
	return trans
}

func (t Transactions) getCredit() Transactions {
	trans := make(Transactions, 0)
	for _, tt := range t {
		if tt.Type == "Credit" {
			trans = append(trans, tt)
		}
	}
	return trans
}

func logHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		x, err := httputil.DumpRequest(r, true)
		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
			return
		}
		log.Println(fmt.Sprintf("%q", x))
		rec := httptest.NewRecorder()
		fn(rec, r)
		log.Println(fmt.Sprintf("%q", rec.Body))
	}
}

func MessageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "A message was received")
}

func (t *Transactions) add(ta Transaction) {

}
