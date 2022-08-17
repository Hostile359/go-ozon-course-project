package comment

import (
	"fmt"

	"gitlab.ozon.dev/Hostile359/homework-1/internal/entities/user"
)

type CommentId uint

type Comment struct {
	Id       CommentId   `db:"id"`
	Comment  string      `db:"comment"`
	UserId   user.UserId `db:"user_id"`
}

func NewComment(id CommentId, comment string, userId user.UserId) Comment {
	c := Comment{
		Id: id,
		Comment: comment,
		UserId: userId,
	}

	return c
}

func (c *Comment) SetComment(comment string) {
	c.Comment = comment
}

func (c *Comment) SetUserId(userId user.UserId) {
	c.UserId = userId
}

func (c Comment) String() string {
	return fmt.Sprintf("%d[%d]: %s", c.Id, c.UserId, c.Comment)
}

func (c Comment) GetComment() string {
	return c.Comment
}

func (c Comment) GetUserId() user.UserId {
	return c.UserId
}

func (c Comment) GetId() CommentId {
	return c.Id
}
