package day5

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNumFromStr(t *testing.T) {
  require.Equal(t, 1, numFromStr("2141", 1))
}

func TestIntcodeInterpret(t *testing.T) {
  icode := intcode{3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99}
  out := icode.Interpret(7)
  require.Equal(t, 1, len(out))
  require.Equal(t, 999, out[0])
  out = icode.Interpret(8)
  require.Equal(t, 1, len(out))
  require.Equal(t, 1000, out[0])
  out = icode.Interpret(9)
  require.Equal(t, 1, len(out))
  require.Equal(t, 1001, out[0])
}
