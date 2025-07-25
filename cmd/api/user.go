package main

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/smnthjm08/go-social/internal/store"
)

type userKey string

const userCtx userKey = "user"

func (app *application) userContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId, err := strconv.ParseInt(chi.URLParam(r, "userId"), 10, 64)
		if err != nil {
			app.badRequestResponse(w, r, err)
			return
		}

		ctx := r.Context()

		user, err := app.store.Users.GetByID(ctx, userId)

		if err != nil {
			switch err {
			case store.ErrNotFound:
				app.notFoundResponse(w, r, err)
				return
			default:
				app.internalServerError(w, r, err)
				return
			}
		}

		ctx = context.WithValue(ctx, userCtx, user)

		next.ServeHTTP(w, r.WithContext(ctx))

	})
}

func getUserFromContext(r *http.Request) *store.User {
	user, _ := r.Context().Value(userCtx).(*store.User)
	return user

}

func (app *application) getUserHandler(w http.ResponseWriter, r *http.Request) {
	user := getUserFromContext(r)

	if err := app.jsonResponse(w, http.StatusOK, user); err != nil {
		app.internalServerError(w, r, err)
	}

}

func (app *application) followUserHandler(w http.ResponseWriter, r *http.Request) {
	followerUser := getUserFromContext(r)

	type FollowUser struct {
		UserId int64 `json:"user_id"`
	}

	// pending
	var payload FollowUser

	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()

	followerID, err := strconv.ParseInt(followerUser.ID, 10, 64)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := app.store.Followers.Follow(ctx, followerID, payload.UserId); err != nil {
		switch err {
		case store.ErrConflict:
			app.conflictResponse(w, r, err)
			return
		default:
			app.internalServerError(w, r, err)
			return
		}
	}

	if err := app.jsonResponse(w, http.StatusOK, nil); err != nil {
		app.internalServerError(w, r, err)
	}
}

func (app *application) unfollowUserHandler(w http.ResponseWriter, r *http.Request) {
	unfollowUser := getUserFromContext(r)

	type UnfollowUser struct {
		UserId int64 `json:"user_id"`
	}

	var payload UnfollowUser

	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()

	followerID, err := strconv.ParseInt(unfollowUser.ID, 10, 64)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := app.store.Followers.Unfollow(ctx, followerID, payload.UserId); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, nil); err != nil {
		app.internalServerError(w, r, err)
	}
}

// func (app *application) unfollowUserHandler(w http.ResponseWriter, r *http.Request) {
// 	unfollowUser := getUserFromContext(r)

// 	type UnfollowUser struct {
// 		UserId int64 `json:"user_id"`
// 	}

// 	// pending
// 	var payload UnfollowUser

// 	if err := readJSON(w, r, payload); err != nil {
// 		app.badRequestResponse(w, r, err)
// 		return
// 	}

// 	ctx := r.Context()

// 	followerID, err := strconv.ParseInt(unfollowUser.ID, 10, 64)
// 	if err != nil {
// 		app.badRequestResponse(w, r, err)
// 		return
// 	}

// 	if err := app.store.Followers.Unfollow(ctx, followerID, payload.UserId); err != nil {
// 		app.internalServerError(w, r, err)
// 		return
// 	}

// 	if err := app.jsonResponse(w, http.StatusOK, nil); err != nil {
// 		app.internalServerError(w, r, err)
// 	}
// }
