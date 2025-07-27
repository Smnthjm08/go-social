package main

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/smnthjm08/go-social/internal/store"
)

type postKey string

const PostCtx postKey = "post"

type CreatePostPayload struct {
	Title   string `json:"title" validate:"required,max=100"`
	Content string `json:"content" validate:"required,max=1000"`
	// UserId  int64    `json:"user_id" validate:"required,max=100"`
	Tags []string `json:"tags,omitempty"`
}

type UpdatePostPayload struct {
	Title   *string   `json:"title" validate:"omitempty,max=100"`
	Content *string   `json:"content" validate:"omitempty,max=1000"`
	Tags    *[]string `json:"tags"`
}

func getPostFromCtx(r *http.Request) *store.Post {
	post, _ := r.Context().Value(PostCtx).(*store.Post)
	return post
}

func (app *application) postsContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

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
			log.Printf("wrapped error: %v", err)

			switch {
			case errors.Is(err, store.ErrNotFound):
				app.notFoundResponse(w, r, err)

			case errors.Is(err, sql.ErrNoRows):
				// log.Printf("post not found: %v", err)
				app.notFoundResponse(w, r, err)
			default:
				// log.Printf("internal server error: %v", err)
				app.internalServerError(w, r, err)
			}
			return
		}

		ctx = context.WithValue(ctx, PostCtx, post)
		next.ServeHTTP(w, r.WithContext(ctx))
	})

}

// CreatePost godoc
//
//	@Summary		Creates a post
//	@Description	Creates a post
//	@Tags			posts
//	@Accept			json
//	@Produce		json
//	@Param			payload	body		CreatePostPayload	true	"Post payload"
//	@Success		201		{object}	store.Post
//	@Failure		400		{object}	error
//	@Failure		401		{object}	error
//	@Failure		500		{object}	error
//	@Security		ApiKeyAuth
//	@Router			/posts [post]
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

	if err := app.jsonResponse(w, http.StatusCreated, post); err != nil {
		// writeJSONError(w, http.StatusInternalServerError, err.Error())
		app.internalServerError(w, r, err)

		return
	}
}

// GetPost godoc
//
//	@Summary		Fetches a post
//	@Description	Fetches a post by ID
//	@Tags			posts
//	@Accept			json
//	@Produce		json
//	@Param			postID	path		int	true	"Post ID"
//	@Success		200		{object}	store.Post
//	@Failure		404		{object}	error
//	@Failure		500		{object}	error
//	@Security		ApiKeyAuth
//	@Router			/posts/{postID}/ [get]
func (app *application) getPostByIDHandler(w http.ResponseWriter, r *http.Request) {
	post := getPostFromCtx(r)

	comments, err := app.store.Comments.GetByPostID(r.Context(), post.ID)

	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	post.Comments = comments

	if err := app.jsonResponse(w, http.StatusOK, post); err != nil {
		app.internalServerError(w, r, err)
	}

}

// UpdatePost godoc
//
//	@Summary		Updates a post
//	@Description	Updates a post by ID
//	@Tags			posts
//	@Accept			json
//	@Produce		json
//	@Param			postID	path		int					true	"Post ID"
//	@Param			payload	body		UpdatePostPayload	true	"Post payload"
//	@Success		200		{object}	store.Post
//	@Failure		400		{object}	error
//	@Failure		401		{object}	error
//	@Failure		404		{object}	error
//	@Failure		500		{object}	error
//	@Security		ApiKeyAuth
//	@Router			/posts/{postID}/ [put]
func (app *application) updatePostByIDHandler(w http.ResponseWriter, r *http.Request) {
	post := getPostFromCtx(r)

	var payload UpdatePostPayload

	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Vaidate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if payload.Content != nil {
		post.Content = *payload.Content
	}
	if payload.Title != nil {
		post.Title = *payload.Title
	}

	if payload.Tags != nil {
		post.Tags = *payload.Tags
	}

	if err := app.store.Posts.UpdatePostByID(r.Context(), post); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, post); err != nil {
		app.internalServerError(w, r, err)
	}

}

// DeletePost godoc
//
//	@Summary		Deletes a post
//	@Description	Delete a post by ID
//	@Tags			posts
//	@Accept			json
//	@Produce		json
//	@Param			postID	path		int	true	"Post ID"
//	@Success		204		{object}	string
//	@Failure		404		{object}	error
//	@Failure		500		{object}	error
//	@Security		ApiKeyAuth
//	@Router			/posts/{postID}/ [delete]
func (app *application) deletePostByIDHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "postID")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	ctx := r.Context()

	if err := app.store.Posts.DeleteByID(ctx, id); err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			app.notFoundResponse(w, r, err)
		default:
			app.internalServerError(w, r, err)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
