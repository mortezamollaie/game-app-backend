package main

import (
	"game-app/config"
	"game-app/delivery/httpserver"
	"game-app/repository/mysql"
	authservice "game-app/service/authService"
	userservice "game-app/service/userservice"

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

	cfg := config.Config{
		HTTPServer: config.HTTPServer{Port: 8080},
		Auth: authservice.Config{
			SignKey:               JwtSignKey,
			AccessExpirationTime:  AccessTokenExpireDuration,
			RefreshExpirationTime: RefreshTokenExpireDuration,
			AccessSubject:         AccessTokenSubject,
			RefreshSubject:        RefreshTokenSubject,
		},
		Mysql: mysql.Config{
			Host:     "127.0.0.1",
			Port:     3306,
			Username: "root",
			Password: "",
			DBName:   "gameapp_db",
		},
	}

	authSvc, userSvc := setupServices(cfg)
	server := httpserver.New(cfg, authSvc, userSvc)
	server.Serve()
}

//func userLoginHandler(writer http.ResponseWriter, req *http.Request) {
//	writer.Header().Set("Content-Type", "application/json")
//
//	if req.Method != http.MethodPost {
//		writer.WriteHeader(http.StatusMethodNotAllowed)
//		fmt.Fprintf(writer, `{"error":"method not allowed"}`)
//
//		return
//	}
//
//	data, err := io.ReadAll(req.Body)
//	if err != nil {
//		writer.WriteHeader(http.StatusInternalServerError)
//		writer.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err.Error())))
//
//		return
//	}
//
//	var uReq userservice.LoginRequest
//	err = json.Unmarshal(data, &uReq)
//	if err != nil {
//		writer.WriteHeader(http.StatusInternalServerError)
//		writer.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err.Error())))
//
//		return
//	}
//
//	authScv := authservice.New(JwtSignKey, AccessTokenSubject, RefreshTokenSubject, AccessTokenExpireDuration, RefreshTokenExpireDuration)
//
//	mysqlRepo := mysql.New()
//	userSvc := userservice.New(authScv, mysqlRepo)
//	resp, err := userSvc.Login(uReq)
//	if err != nil {
//		writer.WriteHeader(http.StatusInternalServerError)
//		writer.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err.Error())))
//
//		return
//	}
//
//	respData, err := json.Marshal(resp)
//	if err != nil {
//		writer.WriteHeader(http.StatusInternalServerError)
//		writer.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err.Error())))
//	}
//
//	writer.Write([]byte(`{"message": "user logged in successfully", "data": ` + string(respData) + `}`))
//}
//
//func userProfile(writer http.ResponseWriter, req *http.Request) {
//	writer.Header().Set("Content-Type", "application/json")
//
//	if req.Method != http.MethodGet {
//		writer.WriteHeader(http.StatusMethodNotAllowed)
//		fmt.Fprintf(writer, `{"error":"method not allowed"}`)
//
//		return
//	}
//
//	authScv := authservice.New(JwtSignKey, AccessTokenSubject, RefreshTokenSubject, AccessTokenExpireDuration, RefreshTokenExpireDuration)
//
//	authToken := req.Header.Get("Authorization")
//	claims, err := authScv.ParseToken(authToken)
//
//	if err != nil {
//		writer.WriteHeader(http.StatusUnauthorized)
//		writer.Write([]byte(`{"message": "token is not valid"}`))
//	}
//
//	mysqlRepo := mysql.New()
//	userSvc := userservice.New(authScv, mysqlRepo)
//	resp, err := userSvc.GetProfile(userservice.ProfileRequest{UserID: claims.UserID})
//	if err != nil {
//		writer.WriteHeader(http.StatusInternalServerError)
//		writer.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err.Error())))
//
//		return
//	}
//
//	respData, err := json.Marshal(resp)
//	if err != nil {
//		writer.WriteHeader(http.StatusInternalServerError)
//		writer.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err.Error())))
//	}
//
//	writer.Write([]byte(`{"message": "user find successfully", "data": ` + string(respData) + `}`))
//}

func setupServices(cfg config.Config) (authservice.Service, userservice.Service) {
	authSvc := authservice.New(cfg.Auth)

	MysqlRepo := mysql.New(cfg.Mysql)
	userSvc := userservice.New(authSvc, MysqlRepo)
	return authSvc, userSvc
}
