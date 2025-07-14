package main

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/smnthjm08/go-social/internal/store"
)

type CreatePostPayload struct {
	Title   string `json:"title" validate:"required,max=100"`
	Content string `json:"content" validate:"required,max=1000"`
	// UserId  int64    `json:"user_id" validate:"required,max=100"`
	Tags []string `json:"tags,omitempty"`
}

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreatePostPayload

	if err := readJSON(w, r, &payload); err != nil {
		// writeJSONError(w, http.StatusBadRequest, err.Error())
		app.badRequestResponse(w, r, err)
		return
	}

	// if payload.Content == "" {
	// 	app.badRequestResponse(w, r, fmt.Errorf("content is required"))
	// 	return
	// }

	if err := Vaidate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	post := &store.Post{
		Title:   payload.Title,
		Content: payload.Content,
		// UserId:  payload.UserId,
		// change later after the auth
		UserId: 1,
		Tags:   payload.Tags,
	}

	ctx := r.Context()

	if err := app.store.Posts.Create(ctx, post); err != nil {
		// writeJSONError(w, http.StatusInternalServerError, err.Error())
		app.internalServerError(w, r, err)

		return
	}

	if err := writeJSON(w, http.StatusCreated, post); err != nil {
		// writeJSONError(w, http.StatusInternalServerError, err.Error())
		app.internalServerError(w, r, err)

		return
	}
}

func (app *application) getPostByIDHandler(w http.ResponseWriter, r *http.Request) {
	// postID := 1
	idParams := chi.URLParam(r, "postID")
	id, err := strconv.ParseInt(idParams, 10, 64)
	if err != nil {
		// writeJSONError(w, http.StatusInternalServerError, err.Error())
		app.internalServerError(w, r, err)
		return
	}

	ctx := r.Context()

	post, err := app.store.Posts.GetByID(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			// writeJSONError(w, http.StatusNotFound, err.Error())
			app.notFoundResponse(w, r, err)
		default:
			// writeJSONError(w, http.StatusInternalServerError, err.Error())
			app.internalServerError(w, r, err)
		}
		return
	}

	comments, err := app.store.Comments.GetByPostID(ctx, id)

	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	post.Comments = comments

	if err := writeJSON(w, http.StatusOK, post); err != nil {
		app.internalServerError(w, r, err)

		// writeJSONError(w, http.StatusInternalServerError, err.Error())
	}

}
