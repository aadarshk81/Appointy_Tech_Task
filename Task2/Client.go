package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"log"
    "bytes"
	"strconv"

)

type Article struct {
	ID int `json:"ID"`
	Title string `json:"Title"`
	SubTitle string `json:"SubTitle"`
	Content string `json:"Content"`
	Timestamp  string `json:"Timestamp"`
}

// simple get request to homepage
func get_home(){
    fmt.Println("-----Performing Http Get...")
    url := "http://localhost:8000/"
    fmt.Println("URL:>", url)

    resp, err := http.Get(url)
    if err != nil {
        log.Fatalln(err)
    }

    defer resp.Body.Close()
    bodyBytes, _ := ioutil.ReadAll(resp.Body)

    // Convert response body to string
    bodyString := string(bodyBytes)
    fmt.Println("API Response as String:\n" + bodyString)
}

// making a GET request to '/articles' to get all the articles
func get_all_articles(){
    fmt.Println("-----Performing Http Get...")
    url := "http://localhost:8000/articles"
    fmt.Println("URL:>", url)

    resp, err := http.Get(url)
    if err != nil {
        log.Fatalln(err)
    }

    defer resp.Body.Close()
    bodyBytes, _ := ioutil.ReadAll(resp.Body)

    // Convert response body to string
    bodyString := string(bodyBytes)
    fmt.Println("API Response as String:\n" + bodyString)

}

// send POST requests to server to add new artcle
func add_article_post(){
    fmt.Println("-----Performing Http Post...")
    url := "http://localhost:8000/articles"
    fmt.Println("URL:>", url)

    jsonData := Article{2,"IPL 2020","Today's Matches","DC vs RR", "4:35"}
    jsonValue, _ := json.Marshal(jsonData)
    resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        bodyBytes, _ := ioutil.ReadAll(resp.Body)
        bodyString := string(bodyBytes)
        fmt.Println("API Response as String:\n" + bodyString)
    }
}

// GET request to get the article by its ID
func get_artcile_by_id(id int){
    fmt.Println("-----Performing Http Get...")
    url := "http://localhost:8000/articles/" + strconv.Itoa(id)
    fmt.Println("URL:>", url)

    resp, err := http.Get(url)
    if err != nil {
        log.Fatalln(err)
    }

    defer resp.Body.Close()
    bodyBytes, _ := ioutil.ReadAll(resp.Body)

    // Convert response body to string
    bodyString := string(bodyBytes)
    fmt.Println("API Response as String:\n" + bodyString)
}

// GET request to search for article
func search_article(s string){
    fmt.Println("-----Performing Http Get...")
    url := "http://localhost:8000/articles/search?q=" + s
    fmt.Println("URL:>", url)

    resp, err := http.Get(url)
    if err != nil {
        log.Fatalln(err)
    }

    defer resp.Body.Close()
    bodyBytes, _ := ioutil.ReadAll(resp.Body)

    // Convert response body to string
    bodyString := string(bodyBytes)
    fmt.Println("API Response as String:\n" + bodyString)
}


func main(){
    // simple home GET request
    get_home()

    fmt.Println("\nNow sending GET request to '/artciles'...\n\n")

    // GET request to '/artciles' to get all the articles
	get_all_articles()

    fmt.Println("\nNow sending POST request to add new article...\n\n")

    // POST request to '/articles' to add a new article
    add_article_post()

    fmt.Println("\nNow sending GET request to get article with id...\n\n")

    // GET request to '/articles/1' to get article with ID
    get_artcile_by_id(1)

    fmt.Println("\nNow sending GET request to search article...\n\n")

    // GET request to '/articles/1' to get article with ID
    search_article("hello")
    search_article("Aadarsh")
}