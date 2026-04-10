package main

import (
	"encoding/json"
	"fmt"
	"game-app/repository/mysql"
	"game-app/service/userservice"
	"io"
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/health-check", healthCheckHandler)
	mux.HandleFunc("/users/register", userRegisterHandler)
	mux.HandleFunc("/users/login", userLoginHandler)

	//err := http.ListenAndServe(":8080", mux)

	server := http.Server{Addr: ":8080", Handler: mux}
	log.Fatal(server.ListenAndServe())
}

func userRegisterHandler(writer http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		fmt.Fprintf(writer, `{"error":"method not allowed"}`)
	}

	writer.Header().Set("Content-Type", "application/json")

	data, err := io.ReadAll(req.Body)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err.Error())))

		return
	}

	var uReq userservice.RegisterRequest
	err = json.Unmarshal(data, &uReq)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err.Error())))

		return
	}

	mysqlRepo := mysql.New()
	userSvc := userservice.New(mysqlRepo)
	_, err = userSvc.Register(uReq)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err.Error())))

		return
	}

	writer.Write([]byte(`{"message": "user created successfully"}`))
}

func healthCheckHandler(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(writer, `{"alive": true}`)
}

func userLoginHandler(writer http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		fmt.Fprintf(writer, `{"error":"method not allowed"}`)
	}

	writer.Header().Set("Content-Type", "application/json")

	data, err := io.ReadAll(req.Body)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err.Error())))

		return
	}

	var uReq userservice.LoginRequest
	err = json.Unmarshal(data, &uReq)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err.Error())))

		return
	}

	mysqlRepo := mysql.New()
	userSvc := userservice.New(mysqlRepo)
	_, err = userSvc.Login(uReq)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err.Error())))

		return
	}

	writer.Write([]byte(`{"message": "user logged in successfully"}`))
}
