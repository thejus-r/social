package main

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/thejus-r/social/package/store"
)

type UserKey string

const userCTX UserKey = "user"

func (app *application) getUserHandler(w http.ResponseWriter, r *http.Request) {

	user := getUserFromCtx(r)

	if err := app.jsonResponse(w, http.StatusOK, user); err != nil {
		app.internalServerError(w, r, err)
		return
	}

}

type FollowUser struct {
	UserID int64 `json:"user_id"`
}

func (app *application) followUserHandler(w http.ResponseWriter, r *http.Request) {
	user := getUserFromCtx(r)

	ctx := r.Context()

	var payload FollowUser
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	app.store.Followers.Follow(ctx, user.ID, payload.UserID)

	if err := app.jsonResponse(w, http.StatusNoContent, nil); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func (app *application) unfollowUserHandler(w http.ResponseWriter, r *http.Request) {
	user := getUserFromCtx(r)

	ctx := r.Context()

	var payload FollowUser
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	app.store.Followers.UnFollow(ctx, user.ID, payload.UserID)

	if err := app.jsonResponse(w, http.StatusNoContent, nil); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func (app *application) userCtxMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "userID")

		id, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			app.internalServerError(w, r, err)
		}

		ctx := r.Context()

		user, err := app.store.Users.GetByID(ctx, id)
		if err != nil {
			switch {
			case errors.Is(err, store.ErrNotFound):
				app.notFoundResponse(w, r, err)
			default:
				app.internalServerError(w, r, err)
			}
			return
		}
		ctx = context.WithValue(ctx, userCTX, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getUserFromCtx(r *http.Request) *store.User {
	user, _ := r.Context().Value(userCTX).(*store.User)
	return user
}
