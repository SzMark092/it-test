package ports

import (
	"it-test/app/query"
	"it-test/domain"
	"it-test/pkg/server/httperr"
	"net/http"

	"github.com/go-chi/render"
	"github.com/google/uuid"
)

func (h HTTPServer) PostUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	result, err := h.app.Queries.PostUser.Handle(ctx, &query.PostUser{})
	if err != nil {
		httperr.InternalError(domain.ErrorInternalServerErrorLabel, getUserCount, uuid.NewString(), err, w, r)
		return
	}

	response := CreateUser{}
	render.Respond(w, r, response)
}

func (h HTTPServer) GetUserList(w http.ResponseWriter, r *http.Request, params GetUserListParams) {
	ctx := r.Context()
	result, err := h.app.Queries.GetUserList.Handle(ctx, &query.GetUserList{})
	if err != nil {
		httperr.InternalError(domain.ErrorInternalServerErrorLabel, getUserCount, uuid.NewString(), err, w, r)
		return
	}

	response := UserList{}
	render.Respond(w, r, response)
}

func (h HTTPServer) UpdateUserDetails(w http.ResponseWriter, r *http.Request, id string) {
	ctx := r.Context()
	result, err := h.app.Queries.UpdateUserDetails.Handle(ctx, &query.UpdateUserDetails{})
	if err != nil {
		httperr.InternalError(domain.ErrorInternalServerErrorLabel, getUserCount, uuid.NewString(), err, w, r)
		return
	}

	response := UpdateUser{}
	render.Respond(w, r, response)
}

func (h HTTPServer) Count(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	result, err := h.app.Queries.GetUserCount.Handle(ctx, &query.GetUserCount{})
	if err != nil {
		httperr.InternalError(domain.ErrorInternalServerErrorLabel, getUserCount, uuid.NewString(), err, w, r)
		return
	}
	response := Count{Count: &result}

	render.Respond(w, r, response)
}
