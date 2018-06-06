// Package git provides a few light-weight GIT utilities.
package git

import (
	"bytes"
	"encoding/json"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
)

// Errors.
var (
	ErrDirty  = errors.New("git repo is dirty")
	ErrNoRepo = errors.New("git repo not found")
	ErrLookup = errors.New("git is not installed")
)

// GetRoot returns the git root relative to dir, if present.
func GetRoot(dir string) (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	cmd.Dir = dir
	b, err := output(cmd)
	return string(b), err
}

// Commit is commit meta-data.
type Commit struct {
	AbbreviatedCommit string `json:"abbreviated_commit"`
	AbbreviatedParent string `json:"abbreviated_parent"`
	AbbreviatedTree   string `json:"abbreviated_tree"`
	Author            struct {
		Date  string `json:"date"`
		Email string `json:"email"`
		Name  string `json:"name"`
	} `json:"author"`
	Commit   string `json:"commit"`
	Commiter struct {
		Date  string `json:"date"`
		Email string `json:"email"`
		Name  string `json:"name"`
	} `json:"commiter"`
	Parent               string `json:"parent"`
	Refs                 string `json:"refs"`
	SanitizedSubjectLine string `json:"sanitized_subject_line"`
	Tree                 string `json:"tree"`
}

// Tag returns the tag or empty string.
func (c *Commit) Tag() string {
	parts := strings.Split(c.Refs, ", ")
	for _, p := range parts {
		if strings.HasPrefix(p, "tag: ") {
			return strings.Replace(p, "tag: ", "", 1)
		}
	}
	return ""
}

// Describe returns the tag or sha.
func (c *Commit) Describe() string {
	if t := c.Tag(); t != "" {
		return t
	}

	return c.AbbreviatedCommit
}

// GetCommit returns meta-data for the given commit within a repo.
func GetCommit(dir, commit string) (c *Commit, err error) {
	dir, err = GetRoot(dir)
	if err != nil {
		return nil, err
	}

	cmd := exec.Command("git", "log", "-1", `--pretty=format:{"commit":"%H","abbreviated_commit":"%h","tree":"%T","abbreviated_tree":"%t","parent":"%P","abbreviated_parent":"%p","refs":"%D","author":{"name":"%aN","email":"%aE","date":"%aD"},"commiter":{"name":"%cN","email":"%cE","date":"%cD"}}`, commit)
	cmd.Dir = dir

	b, err := output(cmd)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(b, &c); err != nil {
		return nil, errors.Wrap(err, "unmarshaling")
	}

	return c, nil
}

// output returns GIT command output with error normalization.
func output(cmd *exec.Cmd) ([]byte, error) {
	out, err := cmd.CombinedOutput()

	if e, ok := err.(*exec.Error); ok {
		if e.Err == exec.ErrNotFound {
			return nil, ErrLookup
		}

		return nil, e
	}

	switch {
	case bytes.Contains(out, []byte("Not a git repository")):
		return nil, ErrNoRepo
	case bytes.Contains(out, []byte("ambiguous argument 'HEAD'")):
		return nil, ErrNoRepo
	case bytes.Contains(out, []byte("DIRTY")):
		return nil, ErrDirty
	case err != nil:
		return nil, errors.New(string(out))
	default:
		return bytes.TrimSpace(out), nil
	}
}
