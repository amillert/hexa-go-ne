package arithmetic

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddition(t *testing.T) {
	core := NewAdapter()

	answer, err := core.Addition(1, 1)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}

	require.Equal(t, answer, int32(2))
}

func TestSubtraction(t *testing.T) {
	core := NewAdapter()

	answer, err := core.Subtraction(2, 1)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}

	require.Equal(t, answer, int32(1))
}

func TestMultiplication(t *testing.T) {
	core := NewAdapter()

	answer, err := core.Multiplication(2, 1)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}

	require.Equal(t, answer, int32(2))
}

func TestDivision(t *testing.T) {
	core := NewAdapter()

	answer, err := core.Division(6, 2)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}

	require.Equal(t, answer, int32(3))
}
