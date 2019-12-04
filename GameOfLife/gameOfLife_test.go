package GameOfLife

import (
	"reflect"
	"testing"
)

func Test_emptyUniverse_staysEmpty(t *testing.T) {
	// Given
	emptyUniverse := EmptyUniverse()

	// When
	actualUniverse := emptyUniverse.next()

	// Then
	if !reflect.DeepEqual(emptyUniverse, actualUniverse) {
		t.Errorf("Expected universe to be empty.")
	}
}

func Test_singleCell_dies(t *testing.T) {
	// Given
	singleCell := NewUniverse([]Point{{0, 0}})
	emptyUniverse := EmptyUniverse()

	// When
	actualUniverse := singleCell.next()

	// Then
	if !reflect.DeepEqual(emptyUniverse, actualUniverse) {
		t.Errorf("Expected single cell universe to die")
	}
}

func Test_block(t *testing.T) {
	// Given
	block := NewUniverse([]Point{{0, 0}, {0, 1}, {1, 0}, {1, 1}})

	// When
	actualUniverse := block.next()

	// Then
	if !reflect.DeepEqual(block, actualUniverse) {
		t.Errorf("Expected block to remain the same")
	}
}

func Test_blinker(t *testing.T) {
	// Given
	verticalBlinker := NewUniverse([]Point{{1, 0}, {1, 1}, {1, 2}})
	horizontalBlinker := NewUniverse([]Point{{0, 1}, {1, 1}, {2, 1}})

	// When
	actualHorizontalBlinker := verticalBlinker.next()
	actualVerticalBlinker := horizontalBlinker.next()

	if !reflect.DeepEqual(horizontalBlinker, actualHorizontalBlinker) {
		t.Errorf("Expected vertical blinker to turn horizontal")
	}
	if !reflect.DeepEqual(verticalBlinker, actualVerticalBlinker) {
		t.Errorf("Expected horizontal blinker to turn vertical")
	}
}

func Test_neighbors(t *testing.T) {
	point := Point{5, 5}
	expectedNeighbors := []Point{{4, 4}, {4, 5}, {4, 6}, {5, 4}, {5, 6}, {6, 4}, {6, 5}, {6, 6}}
	actualNeighbors := point.getNeighbors()
	if !reflect.DeepEqual(expectedNeighbors, actualNeighbors) {
		t.Errorf("Expected block to remain the same")
	}
}
