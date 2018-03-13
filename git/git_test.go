package git_test

import (
	"testing"

	"github.com/tj/assert"
	"github.com/tj/go/git"
)

func TestGetRoot(t *testing.T) {
	t.Run("with a non-git dir", func(t *testing.T) {
		dir, err := git.GetRoot("/tmp")
		assert.EqualError(t, err, `git repo not found`)
		assert.Equal(t, "", dir)
	})

	t.Run("with the root dir", func(t *testing.T) {
		dir, err := git.GetRoot("..")
		assert.NoError(t, err)
		assert.NotEmpty(t, dir)
	})

	t.Run("with the root dir", func(t *testing.T) {
		a, err := git.GetRoot("..")
		assert.NoError(t, err)
		assert.NotEmpty(t, a)

		b, err := git.GetRoot(".")
		assert.NoError(t, err)
		assert.NotEmpty(t, b)

		assert.Equal(t, a, b)
	})
}

func TestGetCommit(t *testing.T) {
	t.Run("direct HEAD", func(t *testing.T) {
		c, err := git.GetCommit("..", "HEAD")
		assert.NoError(t, err)
		assert.Len(t, c.Commit, 40)
		assert.NotEmpty(t, c.Author.Name)
		assert.NotEmpty(t, c.Author.Email)
		assert.NotEmpty(t, c.Author.Date)
	})

	t.Run("relative HEAD", func(t *testing.T) {
		c, err := git.GetCommit(".", "HEAD")
		assert.NoError(t, err)
		assert.Len(t, c.Commit, 40)
		assert.NotEmpty(t, c.Author.Name)
		assert.NotEmpty(t, c.Author.Email)
		assert.NotEmpty(t, c.Author.Date)
	})

	t.Run("relative sha", func(t *testing.T) {
		a, err := git.GetCommit(".", "HEAD")
		assert.NoError(t, err)
		assert.Len(t, a.Commit, 40)

		b, err := git.GetCommit(".", "642d730")
		assert.NoError(t, err)
		assert.Len(t, a.Commit, 40)

		assert.NotEqual(t, a.Commit, b.Commit, "commits")
	})
}
