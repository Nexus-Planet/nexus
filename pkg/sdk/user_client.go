package auth

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"os/user"
	"time"

	"github.com/bytedance/sonic"
	"github.com/nexus-planet/nexus/internal/auth"
)

var (
	BaseUrl       = "http://localhost:3001"
	CustomBaseUrl = ""
)

type UserClient struct {
	Client *http.Client
}

type User struct {
	ID           string `json:"id"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

func NewUserClient(client *http.Client) *UserClient {
	if client == nil {
		return &UserClient{Client: &http.Client{Timeout: time.Second * 30}}
	}

	return &UserClient{Client: client}
}

func (uc *UserClient) CreateUser(ctx context.Context, userCredentials *User) (*user.User, error) {
	body, err := sonic.Marshal(userCredentials)
	if err != nil {
		return nil, err
	}

	var req *http.Request

	if CustomBaseUrl == "" {
		req, err = http.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf("%s/user/create", BaseUrl), bytes.NewBuffer(body))
		if err != nil {
			return nil, err
		}
	} else {
		req, err = http.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf("%s/user/create", CustomBaseUrl), bytes.NewBuffer(body))
		if err != nil {
			return nil, err
		}
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := uc.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var userResp user.User
	err = sonic.ConfigDefault.NewDecoder(resp.Body).Decode(&userResp)
	if err != nil {
		return nil, err
	}

	return &userResp, nil
}

func (uc *UserClient) FindOneByEmail(ctx context.Context, email string) (*user.User, error) {
	var req *http.Request
	var err error

	if CustomBaseUrl == "" {
		req, err = http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/user/%s", BaseUrl, email), nil)
		if err != nil {
			return nil, err
		}
	} else {
		req, err = http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/user/%s", CustomBaseUrl, email), nil)
		if err != nil {
			return nil, err
		}
	}

	token := ctx.Value(auth.TokenKey).(string)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	q := req.URL.Query()
	q.Set("email", email)
	req.URL.RawQuery = q.Encode()

	resp, err := uc.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var userResp user.User
	err = sonic.ConfigDefault.NewDecoder(resp.Body).Decode(&userResp)
	if err != nil {
		return nil, err
	}

	return &userResp, nil
}
