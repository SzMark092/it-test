package query

import (
	"context"
	"it-test/adapters/psql"
	"it-test/pkg/logs"
)

type UpdateUserRepository interface {
	UpdateUser(ctx context.Context) (*psql.User, error)
}

type UpdateUserHandler struct {
	repo UpdateUserRepository
}

type UpdateUser struct {
}

func NewUpdateUserHandler(repo UpdateUserRepository) *UpdateUserHandler {
	return &UpdateUserHandler{repo: repo}
}

func (h *UpdateUserHandler) Handle(ctx context.Context, cmd *UpdateUser) (count int, err error) {
	defer func() {
		logs.LogCommandExecution("UpdateUserHandler", cmd, err)
	}()

	return h.repo.UpdateUser(ctx)
}
