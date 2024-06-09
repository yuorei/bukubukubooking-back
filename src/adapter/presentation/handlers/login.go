package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/yuorei/bukubukubooking-back/src/domain"
	"github.com/yuorei/bukubukubooking-back/src/middleware"
)

func SlackLogin(w http.ResponseWriter, r *http.Request) {
	// Slackの認証コードを取得し、必要な処理を行う
	// ここでSlackからアクセストークンを取得します。
	// アクセストークンの取得方法はSlack OAuth 2.0の仕様に従います。
	// https://api.slack.com/authentication/oauth-v2

	// Retrieve query parameters
	queryValues := r.URL.Query()

	// Get the value of a specific query parameter
	paramValue := queryValues.Get("code")

	slackClientID := os.Getenv("SLACK_CLIENT_ID")
	slackClientSecret := os.Getenv("SLACK_SECRET_KEY")
	code := paramValue
	redirect_uri := os.Getenv("SLACK_REDIRECT_URI")

	// 認可コードを使って、アクセストークンをリクエストする
	resp, err := http.PostForm("https://slack.com/api/oauth.access", url.Values{
		"client_id":     {slackClientID},
		"client_secret": {slackClientSecret},
		"code":          {code},
		"redirect_uri":  {redirect_uri},
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var param map[string]interface{}
	if err := json.Unmarshal(body, &param); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// アクセストークンを取得する
	accessToken, ok := param["access_token"].(string)
	if !ok {
		http.Error(w, "access token not found", http.StatusInternalServerError)
		return
	}

	// ユーザIDを取得するためのリクエスト
	authTestResp, err := http.PostForm("https://slack.com/api/auth.test", url.Values{
		"token": {accessToken},
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer authTestResp.Body.Close()

	authTestBody, err := io.ReadAll(authTestResp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var user map[string]interface{}
	if err := json.Unmarshal(authTestBody, &user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// アクセストークンを使ってユーザ情報をリクエスト
	userInfoResp, err := http.PostForm("https://slack.com/api/users.info", url.Values{
		"token": {accessToken},
		"user":  {param["user_id"].(string)},
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer userInfoResp.Body.Close()

	var userInfo domain.SlackUser
	err = json.NewDecoder(userInfoResp.Body).Decode(&userInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	LoginResponse := domain.LoginResponse{
		IsRegistered: true,
		Token:        middleware.GenJwt(userInfo.User.ID),
	}

	jsonData, err := json.Marshal(LoginResponse)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Println(userInfo.User.Name)
	_, err = w.Write(jsonData)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
