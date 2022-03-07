package wsstore

import (
	"fmt"
	"os"
	"testing"
)

var MG *MapGrid

func TestMain(m *testing.M) {

	MG = NewMapGrid()
	fmt.Println(MG.BottomLeft.BottomLeft.BottomLeft.BottomLeft.BottomLeft.BottomLeft.BottomLeft.BottomLeft.BottomLeft.BottomLeft.GridRange)
	MG.Insert(Driver{ID: 1, Cordinate: Cordinate{X: 44.04, Y: 25.0780}})
	fmt.Println(MG.BottomLeft.BottomLeft.BottomLeft.BottomLeft.BottomLeft.BottomLeft.BottomLeft.BottomLeft.BottomLeft.BottomLeft)
	fmt.Println(connections[1])
	fmt.Println(MG.Get(1))
	MG.Delete(1)
	fmt.Println(MG.BottomLeft.BottomLeft.BottomLeft.BottomLeft.BottomLeft.BottomLeft.BottomLeft.BottomLeft.BottomLeft.BottomLeft)
	fmt.Println(connections[1])
	os.Exit(m.Run())
}
