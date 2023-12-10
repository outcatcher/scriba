package usecases

import (
	"testing"

	"github.com/outcatcher/scriba/internal/usecases/mocks"
	"github.com/stretchr/testify/require"
)

func TestUseCasesWithRepo(t *testing.T) {
	t.Parallel()

	repo := new(mocks.Mockrepository)
	cases := new(UseCases)

	cases.WithRepo(repo)

	require.Equal(t, repo, cases.repo)
}
