package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

type Role struct {
	Name       string `json:"name"`
	Definition string `json:"definition"`
}

func main() {
	jsonFile, err := ioutil.ReadFile("./credit_roles.json")
	if err != nil {
		log.Fatal("Unable to open credit roles file: ", err)
	}

	var roles map[string]Role
	json.Unmarshal(jsonFile, &roles)

	tmpl, _ := template.ParseFiles("templates/template_1.tmpl")
	tmpl.Execute(os.Stdout, roles)
}
