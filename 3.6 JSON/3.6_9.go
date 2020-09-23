package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, err := os.Open("./3.6 JSON/data-20190514T0100.json")
	if err != nil {
		os.Stdout.WriteString(err.Error())
		os.Exit(1)
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		os.Stdout.WriteString(err.Error())
		os.Exit(1)
	}

	var itemsList []okvedItem

	if err := json.Unmarshal(data, &itemsList); err != nil {
		os.Stdout.WriteString(err.Error())
		os.Exit(1)
	}

	var idSum int64
	for _, item := range itemsList {
		idSum += item.Id
	}

	fmt.Println(idSum)
}

type okvedItem struct {
	Id int64 `json:"global_id"`
}
