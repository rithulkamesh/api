package main

import (
	"log"

	"github.com/rithulkamesh/api/util"
	"github.com/rithulkamesh/api/web"
)

func main() {
	db, err := util.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	web.InitServer(db)
}
