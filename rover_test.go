package rover

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitialState(t *testing.T) {
	r := NewRover(NORTH)
	assert.Equal(t, r.orientation, NORTH)
}

func TestTurnLeft1(t *testing.T) {
	r := NewRover(NORTH)
	r.Turn(LEFT)

	if r.orientation != WEST {
		t.Logf("orientation unexpected: %T %q", r.orientation, r.orientation)
	}

	assert.Equal(t, r.orientation, WEST)
}

func TestTurnLeft2(t *testing.T) {
	r := NewRover(WEST)
	r.Turn(LEFT)

	assert.Equal(t, r.orientation, SOUTH)
}

func TestTurnLeft3(t *testing.T) {
	r := NewRover(SOUTH)
	r.Turn(LEFT)

	assert.Equal(t, r.orientation, EAST)
}

func TestTurnLeft4(t *testing.T) {
	r := NewRover(EAST)
	r.Turn(LEFT)

	assert.Equal(t, r.orientation, NORTH)
}

func TestTurnRight(t *testing.T) {
	testcases := []struct {
		start  Orientation
		expect Orientation
	}{
		{NORTH, EAST},
		{EAST, SOUTH},
		{SOUTH, WEST},
		{WEST, NORTH},
	}
	for idx, tc := range testcases {
		t.Run(fmt.Sprintf("test case %d", idx), func(t *testing.T) {
			r := NewRover(tc.start)
			r.Turn(RIGHT)
			assert.Equal(t, r.orientation, tc.expect)
		})
	}
}
