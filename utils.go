package main

import "log"

func checkErr(err error) {
	if err != nil {
		log.Println("Have an error: ", err)
	}
}
