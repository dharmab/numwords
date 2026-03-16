package numwords

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDictionary_IncludeSecond(t *testing.T) {
	t.Parallel()

	_, ok := dictionary.m["second"]
	assert.True(t, ok)

	IncludeSecond(false)
	_, ok = dictionary.m["second"]
	assert.False(t, ok)

	IncludeSecond(true)
	n, ok := dictionary.m["second"]
	assert.True(t, ok)
	assert.Equal(t, second, n)
}

func TestDictionary_IncludeIndefiniteArticle(t *testing.T) {
	t.Parallel()

	// Indefinite articles present by default
	for _, word := range []string{"a", "an"} {
		_, ok := lookupNumber(word)
		assert.True(t, ok, "expected %q to be in dictionary by default", word)
	}

	// Disable indefinite articles
	IncludeIndefiniteArticle(false)
	for _, word := range []string{"a", "an"} {
		_, ok := lookupNumber(word)
		assert.False(t, ok, "expected %q to be removed after IncludeIndefiniteArticle(false)", word)
	}

	// Other entries unaffected
	_, ok := lookupNumber("one")
	assert.True(t, ok)

	// Re-enable indefinite articles
	IncludeIndefiniteArticle(true)
	for _, word := range []string{"a", "an"} {
		n, ok := lookupNumber(word)
		assert.True(t, ok)
		assert.Equal(t, indefiniteArticles[word], n)
	}

	// Idempotency
	IncludeIndefiniteArticle(true)
	_, ok = lookupNumber("a")
	assert.True(t, ok)

	IncludeIndefiniteArticle(false)
	IncludeIndefiniteArticle(false)
	_, ok = lookupNumber("a")
	assert.False(t, ok)

	// Restore default state
	IncludeIndefiniteArticle(true)
}

func TestDictionary_IncludeFractions(t *testing.T) {
	t.Parallel()

	// Fractions present by default
	for _, word := range []string{"half", "thirds", "quarter"} {
		_, ok := lookupNumber(word)
		assert.True(t, ok, "expected %q to be in dictionary by default", word)
	}

	// Disable fractions
	IncludeFractions(false)
	for _, word := range []string{"half", "thirds", "quarter"} {
		_, ok := lookupNumber(word)
		assert.False(t, ok, "expected %q to be removed after IncludeFractions(false)", word)
	}

	// Non-fraction entries unaffected
	for _, word := range []string{"one", "hundred", "first"} {
		_, ok := lookupNumber(word)
		assert.True(t, ok, "expected non-fraction %q to remain in dictionary", word)
	}

	// Re-enable fractions
	IncludeFractions(true)
	n, ok := lookupNumber("half")
	assert.True(t, ok)
	assert.Equal(t, fractions["half"], n)

	n, ok = lookupNumber("thirds")
	assert.True(t, ok)
	assert.Equal(t, fractions["thirds"], n)

	// Idempotency
	IncludeFractions(true)
	_, ok = lookupNumber("half")
	assert.True(t, ok)

	IncludeFractions(false)
	IncludeFractions(false)
	_, ok = lookupNumber("half")
	assert.False(t, ok)

	// Restore default state
	IncludeFractions(true)
}

func TestDictionary_LookupNumber(t *testing.T) {
	t.Parallel()

	_, ok := lookupNumber("one")
	assert.True(t, ok)

	_, ok = lookupNumber("ONE")
	assert.True(t, ok)

	_, ok = lookupNumber("foobar")
	assert.False(t, ok)
}
