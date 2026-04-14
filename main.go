package main

import (
	"encoding/json"
	"fmt"
	"game-app/repository/mysql"
	authservice "game-app/service/authService"
	"game-app/service/userservice"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	JwtSignKey                 = "jwt_secret_key"
	AccessTokenSubject         = "ac"
	RefreshTokenSubject        = "rf"
	AccessTokenExpireDuration  = time.Hour * 24
	RefreshTokenExpireDuration = time.Hour * 24 * 7
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/health-check", healthCheckHandler)
	mux.HandleFunc("/users/register", userRegisterHandler)
	mux.HandleFunc("/users/login", userLoginHandler)
	mux.HandleFunc("/users/profile", userProfile)

	//err := http.ListenAndServe(":8080", mux)

	server := http.Server{Addr: ":8080", Handler: mux}
	fmt.Println("Server is running on port 8080")
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

	authScv := authservice.New(JwtSignKey, AccessTokenSubject, RefreshTokenSubject, AccessTokenExpireDuration, RefreshTokenExpireDuration)

	mysqlRepo := mysql.New()
	userSvc := userservice.New(authScv, mysqlRepo)
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

	authScv := authservice.New(JwtSignKey, AccessTokenSubject, RefreshTokenSubject, AccessTokenExpireDuration, RefreshTokenExpireDuration)

	mysqlRepo := mysql.New()
	userSvc := userservice.New(authScv, mysqlRepo)
	resp, err := userSvc.Login(uReq)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err.Error())))

		return
	}

	respData, err := json.Marshal(resp)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err.Error())))
	}

	writer.Write([]byte(`{"message": "user logged in successfully", "data": ` + string(respData) + `}`))
}

func userProfile(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	if req.Method != http.MethodGet {
		fmt.Fprintf(writer, `{"error":"method not allowed"}`)

		return
	}

	authScv := authservice.New(JwtSignKey, AccessTokenSubject, RefreshTokenSubject, AccessTokenExpireDuration, RefreshTokenExpireDuration)

	authToken := req.Header.Get("Authorization")
	claims, err := authScv.ParseToken(authToken)

	if err != nil {
		writer.Write([]byte(`{"message": "token is not valid"}`))
	}

	mysqlRepo := mysql.New()
	userSvc := userservice.New(authScv, mysqlRepo)
	resp, err := userSvc.GetProfile(userservice.ProfileRequest{UserID: claims.UserID})
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
