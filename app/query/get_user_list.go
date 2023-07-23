package query

import (
	"context"
	"it-test/adapters/psql"
	"it-test/pkg/logs"
)

type GetUserListRepository interface {
	GetUserList(ctx context.Context) ([]psql.User, error)
}

type GetUserListHandler struct {
	repo GetUserListRepository
}

type GetUserList struct {
}

func NewGetUserListHandler(repo GetUserListRepository) *GetUserListHandler {
	return &GetUserListHandler{repo: repo}
}

func (h *GetUserListHandler) Handle(ctx context.Context, cmd *GetUserList) (count int, err error) {
	defer func() {
		logs.LogCommandExecution("GetUserListHandler", cmd, err)
	}()

	return h.repo.GetUserList(ctx)
}
