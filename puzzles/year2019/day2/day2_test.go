package day2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIntCodeInterpret(t *testing.T) {
	icode := intcode{1, 0, 0, 0, 99}
	icode.Interpret()
	require.Equal(t, "[2 0 0 0 99]", fmt.Sprintf("%v", icode))

	icode = intcode{2, 3, 0, 3, 99}
	icode.Interpret()
	require.Equal(t, "[2 3 0 6 99]", fmt.Sprintf("%v", icode))

	icode = intcode{2, 4, 4, 5, 99, 0}
	icode.Interpret()
	require.Equal(t, "[2 4 4 5 99 9801]", fmt.Sprintf("%v", icode))

	icode = intcode{1, 1, 1, 4, 99, 5, 6, 0, 99}
	icode.Interpret()
	require.Equal(t, "[30 1 1 4 2 5 6 0 99]", fmt.Sprintf("%v", icode))
}
