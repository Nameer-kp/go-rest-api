//go:build integration
// +build integration

package database

import (
	"context"
	"fmt"
	"github.com/Nameer-kp/go-rest-api/internal/comment"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCommentDatabase(t *testing.T) {
	t.Run("test create comment", func(t *testing.T) {
		db, err := NewDatabase()
		assert.NoError(t, err)
		cmt, err := db.PostComment(context.Background(), comment.Comment{
			Slug:   "blah",
			Body:   "blahblabha",
			Author: "Dingan",
		})
		assert.NoError(t, err)
		newCmt, err := db.GetComment(context.Background(), cmt.ID)
		assert.NoError(t, err)
		assert.Equal(t, "blah", newCmt.Slug)
		fmt.Println("testing the creation of comments")
	})
	t.Run("test delete comment", func(t *testing.T) {
		db, err := NewDatabase()
		assert.NoError(t, err)
		cmt, err := db.PostComment(context.Background(), comment.Comment{
			Slug:   "bbbb",
			Author: "Dingan v2",
			Body:   "test comment",
		})
		assert.NoError(t, err)
		err = db.DeleteComment(context.Background(), cmt.ID)
		assert.NoError(t, err)
		_, err = db.GetComment(context.Background(), cmt.ID)
		assert.Error(t, err)
	})
}
