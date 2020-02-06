package forum

import (
	"context"
	"dbProject/models"
)

type UseCase interface {
	CreateForum(ctx context.Context, forum *models.Forum) (*models.Forum, error)
	CreateForumThread(ctx context.Context, slug string, thread *models.Thread) (*models.Thread, error)
	GetForumDetails(ctx context.Context, slug string) (*models.Forum, error)
	GetForumThreads(ctx context.Context, slug, since string, limit int, sort bool) ([]*models.Thread, error)
	GetForumUsers(ctx context.Context, slug, since string, limit int, sort bool) ([]*models.User, error)

	CreateThreadPost(ctx context.Context, slug string, posts []*models.Post) ([]*models.Post, error)
	GetThread(ctx context.Context, slug string) (*models.Thread, error)
	ChangeThread(ctx context.Context, slug, title, message string) (*models.Thread, error)
	GetThreadPosts(ctx context.Context, slug string, limit, since int, desc bool, sort models.PostSortType) ([]*models.Post, error)
	VoteForThread(ctx context.Context, slug string, vote models.Vote) (*models.Thread, error)
}