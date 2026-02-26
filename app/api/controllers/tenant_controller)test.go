package controllers

import (
	"testing"

	"github.com/brunobotter/feature-flag/util/shared"
	"github.com/stretchr/testify/require"
)

func TestParsePaginationParam(t *testing.T) {
	t.Run("returns fallback when empty", func(t *testing.T) {
		parsed := parsePaginationParam("", shared.DefaultPage)
		require.Equal(t, shared.DefaultPage, parsed)
	})

	t.Run("returns parsed value when valid", func(t *testing.T) {
		parsed := parsePaginationParam("3", shared.DefaultPage)
		require.Equal(t, 3, parsed)
	})

	t.Run("returns zero when invalid", func(t *testing.T) {
		parsed := parsePaginationParam("abc", shared.DefaultPage)
		require.Zero(t, parsed)
	})
}
