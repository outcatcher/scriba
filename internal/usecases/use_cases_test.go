package usecases

import (
	"testing"

	"github.com/outcatcher/scriba/internal/usecases/mocks"
	"github.com/stretchr/testify/require"
)

func TestUseCasesWithRepo(t *testing.T) {
	t.Parallel()

	repo := mocks.NewMockrepository(t)
	cases := new(UseCases)

	cases.WithRepo(repo)

	require.Equal(t, repo, cases.repo)
}
