package main

import (
	"fmt"
	"time"
)

func main() {

	for {
		parse := Parse2()

		db := DbConnect2()

		err := Sport(parse, db)
		if err != nil {
			fmt.Println(err)

		}
		err = Events(parse, db)
		if err != nil {
			fmt.Println(err)

		}
		err = Factor(parse, db)
		if err != nil {
			fmt.Println(err)

		}

		err = db.Close()
		if err != nil {
			fmt.Println(err)

		}
		time.Sleep(30 * time.Minute)

	}
}
