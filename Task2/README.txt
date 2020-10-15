1. In terminal run "go run server.go"
2. In another terminal run "go run client.go" to run the get/post method, searching, etc.
3. URL are as follows:


http://localhost:8000/
Homepage (Just returns a welcome message)


http://localhost:8000/articles
GET
Returns all the articles

http://localhost:8000/articles
POST
Create a new article
(POST request is sent in client.go "add_article_post" function)


http://localhost:8000/articles/<id>
Returns article with the id
If not found then returns "Article Not Found!" 


http://localhost:8000/articles/search?q="search_text"
Returns ALL articles with search_text in either title or subtitle or content
If not found then returns  "No results Found"