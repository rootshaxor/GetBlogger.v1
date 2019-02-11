package main

/*

   Name : Erwin Kurniawan
   Mail : rootshaor@gmail.com
   Web  : https://rootshaxor.github.io
   Fb   : https://facebook.com/rootshaxor
*/

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

var (
	afterLink string = "/atom.xml?redirect=false&start-index=1&max-results="
)

func isError(err error) bool {
	if err != nil {
		log.Println(err.Error())
	}

	return (err != nil)
}

func buat(file string) {
	_, err := os.Stat(file)

	if os.IsNotExist(err) {
		var files, err = os.Create(file)

		if isError(err) {
			return
		}
		defer files.Close()
	}
	log.Println("Succes create " + file)
}

func tulis(file, input string) {
	files, err := os.OpenFile(file, os.O_WRONLY, 0666)
	if isError(err) {
		return
	}
	defer files.Close()

	_, err = files.WriteString(input + "\n")
	if isError(err) {
		return
	}

	err = files.Sync()
	if isError(err) {
		return
	}
	log.Println("Inserting on ", file)

}

func Usage() {
	fmt.Printf("\nUsage : %s --help\n", os.Args[0])
}

func main() {

	link := flag.String("url", "", "URL for target Blogger !!!")
	result := flag.Int("result", 500, "Result post from link blogger")
	output := flag.String("out", "", "Output File !!!, format output is .xml")
	flag.Parse()

	URL := *link
	RES := strconv.Itoa(*result)
	OUT := *output

	if URL == "" && OUT == "" {
		log.Println("Link & Output file is empety!!!")
		Usage()
		return
	} else if URL == "" {
		log.Println("Link is empety !!!")
		Usage()
		return
	} else if OUT == "" {
		log.Println("Output file is empety!!!")
		Usage()
		return
	}

	response, err := http.Get(URL + afterLink + RES)
	if err != nil {
		log.Printf("%s\n", err)
		defer response.Body.Close()
		os.Exit(1)
	} else {
		content, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		//tls := []byte(content)
		buat(OUT)
		tulis(OUT, string(content))
	}
}
