package numwords

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNumbers_Pattern(t *testing.T) {
	t.Parallel()

	ns := numbers{
		{typ: numAnd},
		{typ: numDirect},
		{typ: numSingle},
		{typ: numTens},
		{typ: numBig},
		{typ: numFraction},
		{typ: numDirectOrdinal},
		{typ: numSingleOrdinal},
		{typ: numTensOrdinal},
		{typ: numBigOrdinal},
		{typ: numDone},
		{typ: numDone + 1},
	}

	assert.Equal(t, "&dstbfDSTB__", ns.pattern())
}

func TestNumbers_Strings(t *testing.T) {
	t.Parallel()

	ns := numbers{
		number{numerator: 1000, denominator: 1, typ: numBig},
		number{numerator: 2, denominator: 1, typ: numSingleOrdinal, ordinal: true},
		number{numerator: 1, denominator: 2, typ: numFraction},
	}

	assert.Equal(t, []string{"1000", "2nd", "0.5"}, ns.strings())
}

func TestNumbers_String(t *testing.T) {
	t.Parallel()

	ns := numbers{
		number{numerator: 1000, denominator: 1, typ: numBig},
		number{numerator: 2, denominator: 1, typ: numSingleOrdinal, ordinal: true},
		number{numerator: 1, denominator: 2, typ: numFraction},
	}

	assert.Equal(t, "1000 2nd 0.5", ns.String())
}

func TestNumbers_Reduce(t *testing.T) {
	t.Parallel()

	ns := numbers{
		number{numerator: 1000, denominator: 1, typ: numBig},
		number{numerator: 2, denominator: 1, typ: numSingle},
		number{numerator: 100, denominator: 1, typ: numBig},
	}

	out := reduce(ns)
	assert.Len(t, out, 1)
	assert.InDelta(t, float64(1200), out[0].Value(), 1e-9)
	assert.Equal(t, numBig, out[0].typ)
}

func TestNumbers_Float(t *testing.T) {
	t.Parallel()

	ns := numbers{
		number{numerator: 3, denominator: 2},
	}

	out, err := ns.Float()
	require.NoError(t, err)
	assert.InDelta(t, float64(1.5), out, 1e-9)

	ns = append(ns, number{})

	_, err = ns.Float()
	assert.Equal(t, ErrManyNumbers, err)

	ns = ns[:0]
	_, err = ns.Float()
	assert.Equal(t, ErrNoNumbers, err)
}

func TestNumbers_Int(t *testing.T) {
	t.Parallel()

	ns := numbers{
		number{numerator: 3, denominator: 2},
	}

	out, err := ns.Int()
	require.NoError(t, err)
	assert.Equal(t, 1, out)

	ns = append(ns, number{})

	_, err = ns.Int()
	assert.Equal(t, ErrManyNumbers, err)

	ns = ns[:0]
	_, err = ns.Int()
	assert.Equal(t, ErrNoNumbers, err)
}
