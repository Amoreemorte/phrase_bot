package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPhrasesMatch(t *testing.T) {
	t.Run("completely matched phrazes", func(t *testing.T) {
		expectedPhrase, actualPhrase := "sgamga", "sgamga"
		phraser := NewPhraser(expectedPhrase)
		assert.True(t, phraser.IsPhrasesMatch(actualPhrase))
	})
	t.Run("phrases with the same phrases but different registers", func(t *testing.T) {
		expectedPhrase, actualPhrase := "sgamga", "SgAmGa"
		phraser := NewPhraser(expectedPhrase)
		assert.True(t, phraser.IsPhrasesMatch(actualPhrase))
	})
	t.Run("different phrazes", func(t *testing.T) {
		expectedPhrase, actualPhrase := "sgamga", "non-sgamga"
		phraser := NewPhraser(expectedPhrase)
		assert.True(t, !phraser.IsPhrasesMatch(actualPhrase))
	})
}
