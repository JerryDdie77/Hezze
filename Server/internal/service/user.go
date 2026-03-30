package service

import (
	"context"
	"errors"
	"log"

	"github.com/JerryDdie77/Hezze/Server/internal/domain"
	"github.com/JerryDdie77/Hezze/Server/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return UserService{repo: repo}
}

func (u *UserService) GetUser(ctx context.Context, id int) (*domain.UserResponse, error) {
	user, err := u.repo.GetUserByID(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return &domain.UserResponse{}, domain.ErrNotFound
		}

		log.Printf("GetUserByID failed: id=%d", id)
		return &domain.UserResponse{}, domain.ErrInternal
	}

	if user.IsBlocked == true {
		return &domain.UserResponse{}, domain.ErrForbidden
	}

	return &domain.UserResponse{
		ID:        user.ID,
		UserName:  user.UserName,
		FirstName: user.FirstName,
		Surname:   user.Surname,
		Email:     user.Email,
	}, nil

}
