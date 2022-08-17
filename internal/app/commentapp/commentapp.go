package commentapp

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/app/userapp"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/entities/comment"
)

const (
	storageTimeout = 10*time.Second
	pageDefault = 1
	perPageDefault = 5
)

var (
	ErrCommentNotExists = errors.New("comment does not exist")
	ErrCommentWrongUserId = errors.New("comment has another user id")
)

type Storage interface {
	Add(context.Context, comment.Comment) error
	Update(context.Context, comment.Comment) error
	Get(context.Context, comment.CommentId) (*comment.Comment, error)
	List(context.Context, uint64, uint64) ([]comment.Comment, error)
	Delete(context.Context, comment.CommentId) error
}

type App struct {
	commentStorage Storage
	userApp userapp.App
}

func New(commentStorage Storage, userApp userapp.App) *App {
	return &App{
		commentStorage: commentStorage,
		userApp: userApp,
	}
}

func (a *App) Add(ctx context.Context, c comment.Comment) error {
	ctx, cancel := context.WithTimeout(ctx, storageTimeout)
	defer cancel()

	if _, err := a.userApp.Get(ctx, c.GetUserId()); err != nil {
		return err
	} 

	return a.commentStorage.Add(ctx, c)
}

func (a *App) Update(ctx context.Context, c comment.Comment) error {
	ctx, cancel := context.WithTimeout(ctx, storageTimeout)
	defer cancel()

	oldComment, err := a.Get(ctx, c.GetId())
	if err != nil {
		return err
	}
	if c.GetUserId() != oldComment.GetUserId() {
		return ErrCommentWrongUserId
	}

	return a.commentStorage.Update(ctx, c)
}

func (a App) Get(ctx context.Context, id comment.CommentId)  (*comment.Comment, error) {
	ctx, cancel := context.WithTimeout(ctx, storageTimeout)
	defer cancel()

	return a.commentStorage.Get(ctx, id)
}

func (a App) List(ctx context.Context, page, perPage uint64) ([]comment.Comment, error) {
	ctx, cancel := context.WithTimeout(ctx, storageTimeout)
	defer cancel()

	if page == 0 {
		page = pageDefault
	}
	if perPage == 0 {
		perPage = perPageDefault
	}

	offset := (page - 1) * perPage
	limit := perPage
	return a.commentStorage.List(ctx, offset, limit)
}

func (a *App)Delete(ctx context.Context, id comment.CommentId) error {
	ctx, cancel := context.WithTimeout(ctx, storageTimeout)
	defer cancel()
	
	return a.commentStorage.Delete(ctx, id)
}
