package main

import (
	"log"
	"os"
	"./search"
)

//init is called before main to set up things
func init(){
	log.SetOutput(os.Stdout)
}

func main(){
	search.Run("hello")
}