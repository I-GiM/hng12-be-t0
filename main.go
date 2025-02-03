package main

import (
	"encoding/json"
	"fmt"
	"time"

	// "i/ioutil";
	"log"
	"net/http"

	// "strconv";
	"sync"
)

type Details struct {
  Email string `json:"email"`
  CurrentDateTime string `json:"current_datetime"`
  GithubUrl string `json:"github_url"`
}

type Response struct {
  Status int `json:"status"`
  Message string `json:"message"`
  Data Details `json:"data"`
}

var detMu sync.Mutex


func enableCors(w *http.ResponseWriter) {
  (*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func detailsHandler(w http.ResponseWriter, r *http.Request) {
  enableCors(&w)
  
  switch r.Method {
  case "GET":
    getDetails(w, r)
  default:
    http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
  }
}

func main()  {
  http.HandleFunc("/", detailsHandler)

  fmt.Println("Server is running at http://localhost:2700")
  log.Fatal(http.ListenAndServe(":2700", nil))
}


func getDetails(w http.ResponseWriter, _ *http.Request)  {
  
  detMu.Lock()
  defer detMu.Unlock()

  currentTime := time.Now()

  dt := Response{
    Status: 200,
    Message: "User details fetched successfully",
    Data: Details{
      Email:"ikennaoyiih@gmail.com",
      CurrentDateTime: currentTime.Format(time.RFC3339),
      GithubUrl:"https://github.com/I-GiM/hng12-be-t0",
    },
  }
  details, err := json.Marshal(dt)
  if err != nil {
		http.Error(w, "Failed to marshal data", http.StatusInternalServerError)
		return
	}

  w.Header().Set("Content-Type", "application/json")
  // json.NewEncoder(w).Encode(bytes)
  w.Write(details)
}
