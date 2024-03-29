package auth

import (
	"context"
	"log"
	"os"

	"github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
)

// Authenticator is oauth and oidc authentication
type Authenticator struct {
	Provider *oidc.Provider
	Config   oauth2.Config
	Ctx      context.Context
}

// NewAuthenticator is to create new Authenticator struct
func NewAuthenticator() (*Authenticator, error) {
	ctx := context.Background()

	domain := os.Getenv("AUTH0_DOMAIN")
	// oidc
	// oidc で認証を行う
	provider, err := oidc.NewProvider(ctx, "https://"+domain+"/")
	if err != nil {
		log.Printf("failed to get provider: %v", err)
		return nil, err
	}

	// oauth
	// oidc で行った認証の結果を元に oauth の設定を行う
	// privider で定義されているエンドポイント、スコープ（権限）を設定
	conf := oauth2.Config{
		ClientID:     os.Getenv("AUTH0_CLIENT_ID"),
		ClientSecret: os.Getenv("AUTH0_CLIENT_SECRET"),
		RedirectURL:  "http://localhost:3000/callback",
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile"},
	}

	return &Authenticator{
		Provider: provider,
		Config:   conf,
		Ctx:      ctx,
	}, nil
}
