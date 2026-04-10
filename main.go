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
	mux.HandleFunc("/users/profile", userProfile)

	//err := http.ListenAndServe(":8080", mux)

	server := http.Server{Addr: ":8080", Handler: mux}
	log.Fatal(server.ListenAndServe())
}

func healthCheckHandler(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(writer, `{"alive": true}`)
}

func userRegisterHandler(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	if req.Method != http.MethodPost {
		fmt.Fprintf(writer, `{"error":"method not allowed"}`)

		return
	}

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

func userLoginHandler(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	if req.Method != http.MethodPost {
		fmt.Fprintf(writer, `{"error":"method not allowed"}`)

		return
	}

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

func userProfile(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	if req.Method != http.MethodGet {
		fmt.Fprintf(writer, `{"error":"method not allowed"}`)

		return
	}

	data, err := io.ReadAll(req.Body)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err.Error())))

		return
	}

	var pReq userservice.ProfileRequest
	err = json.Unmarshal(data, &pReq)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err.Error())))

		return
	}

	mysqlRepo := mysql.New()
	userSvc := userservice.New(mysqlRepo)
	resp, err := userSvc.GetProfile(pReq)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err.Error())))

		return
	}

	respData, err := json.Marshal(resp)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err.Error())))
	}

	writer.Write([]byte(`{"message": "user find successfully", "data": ` + string(respData) + `}`))
}
