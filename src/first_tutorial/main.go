package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"

	"github.com/gorilla/mux"
)

type Article struct {
	TYPE      string `json:"TYPE"`
	BOOTPROTO string `json:"BOOTPROTO"`
	NAME      string `json:"NAME"`
	DEVICE    string `json:"DEVICE"`
	ONBOOT    string `json:"ONBOOT"`
	IPADDR    string `json:"IPADDR"`
	PREFIX    string `json:"PREFIX"`
}

var line_data [7]string
var i int = 0
var x int
var result string
var Articles Article

func fetch() Article { // Reads the config file

	Articles = Article{TYPE: "", BOOTPROTO: "", NAME: "", DEVICE: "", ONBOOT: "", IPADDR: "", PREFIX: ""}
	file, err := os.Open("/Users/mohammedajmalkhan/Desktop/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		line_data[i] = scanner.Text()
		i = i + 1
	}

	for i := 0; i < 7; i++ {
		x := strings.Index(line_data[i], "=")
		if i == 0 {
			Articles.TYPE = line_data[i][x+1:]
		}
		if i == 1 {
			Articles.BOOTPROTO = line_data[i][x+1:]
		}
		if i == 2 {
			Articles.NAME = line_data[i][x+1:]
		}
		if i == 3 {
			Articles.DEVICE = line_data[i][x+1:]
		}
		if i == 4 {
			Articles.ONBOOT = line_data[i][x+1:]
		}
		if i == 5 {
			Articles.IPADDR = line_data[i][x+1:]
		}
		if i == 6 {
			Articles.PREFIX = line_data[i][x+1:]
			break
		}
	}
	return Articles
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func createConfig(w http.ResponseWriter, r *http.Request) { //Post method to create a completely new configuration
	// get the body of our POST request
	// return the string response containing the request body

	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)
	writeToFile(article)
	fmt.Println(article.TYPE)
	fmt.Println(reflect.TypeOf(reqBody))
	fmt.Fprintf(w, "%+v", string(reqBody))
}

func writeToFile(data Article) { //Write the new configuartion into the config file
	//file, err := os.Create("/Users/mohammedajmalkhan/Desktop/test.txt")
	f, err := os.OpenFile("/Users/mohammedajmalkhan/Desktop/test.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	response := fmt.Sprintf("TYPE=%s\nBOOTPROTO=%s\nNAME=%s\nDEVICE=%s\nONBOOT=%s\nIPADDR=%s\nPREFIX=%s\n", data.TYPE, data.BOOTPROTO, data.NAME, data.DEVICE, data.ONBOOT, data.IPADDR, data.PREFIX)
	//response := "Omar is here"
	len, err := f.WriteString(response)

	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}
	fmt.Println(len)
	f.Close()
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) { // Returns the config of file (after reading it using fetch() )
	data := fetch()
	fmt.Println("Endpoint Hit: READ CONFIG")
	json.NewEncoder(w).Encode(data)
}

func deleteConfig(w http.ResponseWriter, r *http.Request) {

	data := fetch()
	vars := mux.Vars(r)
	field := vars["field"]
	if field == "TYPE" {
		data.TYPE = ""
	}
	if field == "BOOTPROTO" {
		data.BOOTPROTO = ""
	}
	if field == "NAME" {
		data.NAME = ""
	}
	if field == "DEVICE" {
		data.DEVICE = ""
	}
	if field == "ONBOOT" {
		data.ONBOOT = ""
	}
	if field == "IPADDR" {
		data.IPADDR = ""
	}
	if field == "PREFIX" {
		data.PREFIX = ""
	}
	writeToFile(data)
	fmt.Println("Successfully deleted!")

}

func updateConfig(w http.ResponseWriter, r *http.Request) {

	// read the field to update
	// read the json to know what the value must be updated to
	// make a new article I guess?
	// write the contents back in to the file

	data := fetch()
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)
	vars := mux.Vars(r)
	field := vars["field"]
	//fmt.Println(field)
	if field == "TYPE" {
		data.TYPE = article.TYPE
	} else if field == "BOOTPROTO" {
		data.BOOTPROTO = article.BOOTPROTO
	} else if field == "NAME" {
		data.NAME = article.NAME
	} else if field == "DEVICE" {
		data.DEVICE = article.DEVICE
	} else if field == "ONBOOT" {
		data.ONBOOT = article.ONBOOT
	} else if field == "IPADDR" {
		data.IPADDR = article.IPADDR
	} else if field == "PREFIX" {
		data.PREFIX = article.PREFIX
	} else {
		fmt.Println("Field Not Found!")
	}
	writeToFile(data)
	json.NewEncoder(w).Encode(data)

}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/read", returnAllArticles)
	myRouter.HandleFunc("/create", createConfig).Methods("POST")
	myRouter.HandleFunc("/delete/{field}", deleteConfig).Methods("DELETE")
	myRouter.HandleFunc("/update/{field}", updateConfig).Methods("POST")
	log.Fatal(http.ListenAndServe(":10002", myRouter))
}

func main() {
	fmt.Println("Started server!")
	handleRequests()
}
