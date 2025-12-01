package day1

import "fmt"

type Direction int

const (
	Left  Direction = -1
	Right Direction = 1
)

func DecodeDirection(s string) Direction {
	switch s {
	case "L":
		return Left
	case "R":
		return Right
	default:
		panic("invalid direction")
	}
}

func (d Direction) String() string {
	switch d {
	case Left:
		return "L"
	case Right:
		return "R"
	default:
		return "?"
	}
}

type Rotation struct {
	Direction Direction
	Amount    int
}

func (r Rotation) String() string {
	return fmt.Sprintf("%s%03d", r.Direction.String(), r.Amount)
}

const dialSize = 100

func RotatePosition(current int, rotation Rotation) int {
	newPosition := (current + int(rotation.Direction)*rotation.Amount)

	for newPosition >= dialSize {
		newPosition = newPosition - dialSize
	}

	for newPosition < 0 {
		newPosition = dialSize + newPosition
	}

	return newPosition
}

func Positions(current int, rotations []Rotation) []int {
	positions := make([]int, 0, len(rotations))

	for _, rotation := range rotations {
		current = RotatePosition(current, rotation)
		positions = append(positions, current)
	}

	return positions
}

func Click(current int, rotation Rotation) int {
	if rotation.Amount > dialSize {
		return 1 + Click(current, Rotation{
			Direction: rotation.Direction,
			Amount:    rotation.Amount - dialSize,
		})
	}

	if current != 0 && current <= rotation.Amount && rotation.Direction == Left {
		return 1
	}

	if current+rotation.Amount >= dialSize && rotation.Direction == Right {
		return 1
	}

	return 0
}

func Clicks(current int, rotations []Rotation) int {
	totalClicks := 0

	for _, rotation := range rotations {
		click := Click(current, rotation)
		to := RotatePosition(current, rotation)

		fmt.Printf("From %02d to %02d with rotation %+v: clicks %d\n", current, to, rotation, click)

		totalClicks += click
		current = to
	}

	return totalClicks
}
