package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"text/template"
)

type Role struct {
	Name       string `json:"name"`
	Definition string `json:"definition"`
}

func main() {
	jsonFile, _ := os.Open("./credit_roles.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var roles map[string]Role
	json.Unmarshal(byteValue, &roles)

	tmpl, _ := template.ParseFiles("templates/template_1.tmpl")
	tmpl.Execute(os.Stdout, roles)
}
