package main

import (
	"bogoDB/backend"
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func showTitle(){
	title := `BogoDb : A toy database management system.`
	fmt.Println(title)
}

func client(){
	showTitle()
	stdin := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(">>")
		stdin.Scan()
		q := stdin.Text()

		var err error
		if q == "exit"{
			_, err = http.Get("http://localhost:32198/exit")
		}else{
			_, err = http.Get("http://localhost:32198/execute?query=" + q)
		}

		if err != nil{
			fmt.Println(err)
		}
	}
}

func server(){
	db, err := backend.NewBogoDb()
	if err != nil{
		log.Fatal(err)
	}
	db.Init()

	err = os.Setenv("BOGO_HOME", "/tmp/bogodb/"); if err != nil{
		log.Fatal(err)
	}

	backend.NewApiServer(db).Host()
}

var (
	serverMode = flag.Bool("server", false, "boot the db server")
)

func main(){
	flag.Parse()

	if *serverMode{
		server()
		return
	}

	client()
}
