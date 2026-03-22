package main

import (
	"error-group/database"
	"error-group/query"
	"fmt"
)

func main() {
	d1 := database.NewDatabase1()
	d2 := database.NewDatabase2()
	shards := []database.Database{d1, d2}

	data, err := query.DistributedQuery(shards, "select * from ...")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(data)
	}
}
