package query

import (
	"context"
	"it-test/pkg/logs"
)

type CreateUserRepository interface {
	CreateUser(ctx context.Context) (int, error)
}

type CreateUserHandler struct {
	repo CreateUserRepository
}

type CreateUser struct {
}

func NewCreateUserHandler(repo CreateUserRepository) *CreateUserHandler {
	return &CreateUserHandler{repo: repo}
}

func (h *CreateUserHandler) Handle(ctx context.Context, cmd *CreateUser) (count int, err error) {
	defer func() {
		logs.LogCommandExecution("CreateUserHandler", cmd, err)
	}()

	return h.repo.CreateUser(ctx)
}
