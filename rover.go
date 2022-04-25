package rover

type Orientation string

func (o Orientation) Turn(m Turn) Orientation {
	if m == LEFT {
		return o.turnLeft()
	}
	return o.turnRight()
}

func (o Orientation) turnLeft() Orientation {
	switch o {
	case NORTH:
		return WEST
	case WEST:
		return SOUTH
	case SOUTH:
		return EAST
	case EAST:
		return NORTH
	}
	return o
}

func (o Orientation) turnRight() Orientation {
	switch o {
	case NORTH:
		return EAST
	case EAST:
		return SOUTH
	case SOUTH:
		return WEST
	case WEST:
		return NORTH
	}
	return o
}

type Turn string
type Movement string

const (
	LEFT  Turn = "L"
	RIGHT Turn = "R"
)

const (
	NORTH Orientation = "N"
	EAST  Orientation = "E"
	SOUTH Orientation = "S"
	WEST  Orientation = "W"
)

type Rover struct {
	orientation Orientation
}

func NewRover(orientation Orientation) *Rover {
	return &Rover{
		orientation: orientation,
	}
}

func (r *Rover) SetOrientation(o Orientation) {
	r.orientation = o
}

func (r *Rover) Turn(m Turn) {
	r.orientation = r.orientation.Turn(m)
}
