package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/eaburns/johaus/parser"
	_ "github.com/eaburns/johaus/parser/camxes"
	_ "github.com/joho/godotenv/autoload"
)

const dialect = "camxes"

func main() {
	p := os.Getenv("PORT")
	if p == "" {
		p = "9001"
	}

	log.Printf("Listening on http://0.0.0.0:%s", p)
	http.ListenAndServe(":"+p, http.HandlerFunc(handle))
}

func handle(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "can't parse: "+err.Error(), http.StatusBadRequest)
		return
	}

	tree, err := parser.Parse(dialect, string(data))
	if err != nil {
		http.Error(w, "can't parse: "+err.Error(), http.StatusBadRequest)
		return
	}

	parser.RemoveMorphology(tree)
	parser.AddElidedTerminators(tree)
	parser.RemoveSpace(tree)
	parser.CollapseLists(tree)

	json.NewEncoder(w).Encode(tree)
}
