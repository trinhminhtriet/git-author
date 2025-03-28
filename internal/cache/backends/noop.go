package backends

import (
	"github.com/trinhminhtriet/git-author/internal/cache"
	"github.com/trinhminhtriet/git-author/internal/git"
)

type NoopBackend struct{}

func (b NoopBackend) Name() string {
	return "noop"
}

func (b NoopBackend) Open() error {
	return nil
}

func (b NoopBackend) Close() error {
	return nil
}

func (b NoopBackend) Get(revs []string) (cache.Result, error) {
	return cache.EmptyResult(), nil
}

func (b NoopBackend) Add(commits []git.Commit) error {
	return nil
}

func (b NoopBackend) Clear() error {
	return nil
}
