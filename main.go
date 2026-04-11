package main

import (
	"encoding/json"
	"fmt"
	"game-app/repository/mysql"
	"game-app/service/userservice"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

const (
	JwtSignKey = "jwt_secret_key"
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
	userSvc := userservice.New(mysqlRepo, JwtSignKey)
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
	userSvc := userservice.New(mysqlRepo, JwtSignKey)
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

	auth := req.Header.Get("Authorization")
	claims, err := parseJWT(auth)

	if err != nil {
		writer.Write([]byte(`{"message": "token is not valid"}`))
	}

	mysqlRepo := mysql.New()
	userSvc := userservice.New(mysqlRepo, JwtSignKey)
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

func parseJWT(tokenStr string) (*userservice.Claims, error) {
	strings.Replace(tokenStr, "Bearer ", "", 1)

	token, err := jwt.ParseWithClaims(tokenStr, &userservice.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JwtSignKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*userservice.Claims); ok && token.Valid {
		fmt.Printf("userID: %v, ExpiresAt: %v\n", claims.UserID, claims.RegisteredClaims.ExpiresAt)
		return claims, nil
	} else {
		return nil, err
	}
}
