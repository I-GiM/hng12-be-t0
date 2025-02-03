# HNG12 Backend Track - Task 0

## Task: Implement a public API that returns email, current datetime, github url of API

### Description

The project is a simple public API that returns the email of the user, the current datetime of the API response, and the github URL of the API. The prject should be able to handle CORS effectively for security.
The codebase is currently written **#Golang** and is deployed using Dokploy - an open source all in one deployment platform.

### Setup instructions

1. Install Go
   The project is written in Go and would need to be installed on your machine to be able to compile the code and deploy the server locally. Head to [Download Go](https://go.dev/learn/) to download the package that suits your machine. Follow the instructions to install Go on your machine.

2. Download code
   To download/clne the codebase, in your terminal type `git clone https://github.com/I-GiM/hng12-be-t0` and wait for it to complete the download.

3. Run code
   Next enter the folder using `cd hng12-be-t0`
   Next type the command `go run main.go` and the server should respond with `Server is running at http://localhost:2700`. If you see no response, ask your friendly neighbourhood Spider-man when you see him swing by.

4. View result
   Go to your browser or open an API test tool like Postman and enter the url `localhost:2700` and you should see the response.

### Documentatin

Endpoint url - http://localhost:2700<br />
Method: "GET"<br />
Response:

```
{
"status": 200,
"message": "User details fetched successfully",
"data": {
"email": "ikennaoyiih@gmail.com",
"current_datetime": "2025-02-03T12:59:23+01:00",
"github_url": "https://github.com/I-GiM/hng12-be-t0"
}
}

```

To see and hire sound Golang engineers, you can go to [https://hng.tech/hire/golang-developers](https://hng.tech/hire/golang-developers)
