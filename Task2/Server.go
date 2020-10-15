package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"log"
	"strconv"
	"strings"
)

type Article struct {
	ID int `json:"ID"`
	Title string `json:"Title"`
	SubTitle string `json:"SubTitle"`
	Content string `json:"Content"`
	Timestamp  string `json:"Timestamp"`
}

// This array stores all the articles
var Articles []Article

// simple homepage
func homePage(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET"{
	    fmt.Fprintf(w, "Welcome to the HomePage!\n")
	    fmt.Fprintf(w, "use url '/articles' to view all articles or send POST request to add new article\n")
	} else {
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

// handles get and post requests to the homepage from client
func get_or_post_articles(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET"{
	    fmt.Fprintf(w, "All articles are :\n")
	    // fmt.Println("Endpoint Hit: homePage")
	    jsonReq, err := json.Marshal(Articles)
		if err != nil {
	        log.Fatalln(err)
	    }
		fmt.Fprintf(w, string(jsonReq))
	} else if r.Method == "POST"{
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
        	log.Fatalln(err)
    	}
    	var articleStruct Article
    	json.Unmarshal(body, &articleStruct)
    	Articles = append(Articles, articleStruct)
    	jsonReq, err := json.Marshal(Articles)
		if err != nil {
	        log.Fatalln(err)
	    }
	    fmt.Fprintf(w, "All articles are :\n")
		fmt.Fprintf(w, string(jsonReq))
	} else {
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

// get the article with given id
func get_article_by_id(w http.ResponseWriter, r *http.Request){

	if r.Method == "GET"{
		id := r.URL.Path[len("/articles/"):]
		fmt.Fprintf(w, "ID of article requested = %q", id)
		fmt.Fprintf(w, "\n")
		//search the article with id
		for i := range Articles {
			if strconv.Itoa(Articles[i].ID) == id {
				jsonReq, err := json.Marshal(Articles[i])
				if err != nil {
			        log.Fatalln(err)
			    }
			    fmt.Fprintf(w, "Articles is :\n")
				fmt.Fprintf(w, string(jsonReq))
				return
			}
		}
		fmt.Fprintf(w, "Article Not Found!")

	} else {
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

// search for article with given key and return
func search_article(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET"{
	    keys, err := r.URL.Query()["q"]
	    
	    if !err || len(keys[0]) < 1 {
	        fmt.Fprintf(w, "Url Param 'q' is missing")
	        log.Fatalln(err)
	        return
	    }

	    key := string(keys[0])

	    // fmt.Fprintf(w, "Url Param 'q' is: " + key + "\n")
	    fmt.Fprintf(w, "Search result for key 'q' = " + key + "\n")
	    // now search for key
	    var search_result []Article
	    var found int = 0
	    for i := range Articles {
			if (strings.Contains(Articles[i].Title, key) || strings.Contains(Articles[i].SubTitle, key) || strings.Contains(Articles[i].Content, key)) {
				search_result = append(search_result, Articles[i])
				found = 1
			}
		}
		if found == 1 {
			jsonReq, err := json.Marshal(search_result)
			if err != nil {
		        log.Fatalln(err)
		    }
			fmt.Fprintf(w, string(jsonReq))
		} else {
			fmt.Fprintf(w, "No results Found")
		}

	} else {
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func handleRequests(){

	http.HandleFunc("/",homePage)
	http.HandleFunc("/articles", get_or_post_articles)
	http.HandleFunc("/articles/", get_article_by_id)
	http.HandleFunc("/articles/search", search_article)
	
	log.Fatal(http.ListenAndServe(":8000",nil))
}


func main(){
	fmt.Println("Server is waiting for requests from client...\n")
	// Initially we have only one article
	Articles = append(Articles, Article{ID: 1, Title: "hello", SubTitle: "hello world", Content: "Hello Welcom to world", Timestamp: "12:20" })
	
	handleRequests()
}