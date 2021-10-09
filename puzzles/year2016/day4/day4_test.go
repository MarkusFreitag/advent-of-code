package day4

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	input := `aaaaa-bbb-z-y-x-123[abxyz]
a-b-c-d-e-f-g-h-987[abcde]
not-a-real-room-404[oarel]
totally-real-room-200[decoy]`
	solution, err := Part1(input)
	require.Nil(t, err)
	require.Equal(t, "1514", solution)
}

func TestRoomDecrypt(t *testing.T) {
	r := &room{ID: 343, Name: "qzmt-zixmtkozy-ivhz"}
	require.Equal(
		t,
		"very encrypted name",
		r.Decrypt(),
	)
}
