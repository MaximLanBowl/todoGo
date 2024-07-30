package auth

import "github.com/MaximLanBowl/todoGo.git/pkg/models"

type AuthService interface {
	Register(creds *Credentials) (*models.User, error)
	Login(creds *Credentials) (*models.User, error)
}

