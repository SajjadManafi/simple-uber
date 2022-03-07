package wsstore

import (
	"github.com/gorilla/websocket"
)

var connections map[int32]*MapGrid

type Cordinate struct {
	X float64
	Y float64
}

type cordinateRange struct {
	xLeft   float64
	xRight  float64
	yTop    float64
	yBottom float64
}

type MapGrid struct {
	GridRange   *cordinateRange
	TopLeft     *MapGrid
	TopRight    *MapGrid
	BottomLeft  *MapGrid
	BottomRight *MapGrid
	Data        map[int32]Driver
}

type Driver struct {
	ID         int32 `json:"id"`
	Cordinate  Cordinate
	Connection *websocket.Conn
}

func NewMapGrid() *MapGrid {
	iran := &MapGrid{
		GridRange: &cordinateRange{
			xLeft:   44.03313,
			xRight:  63.33354,
			yTop:    39.78686,
			yBottom: 25.07480,
		},
		TopLeft:     &MapGrid{},
		TopRight:    &MapGrid{},
		BottomLeft:  &MapGrid{},
		BottomRight: &MapGrid{},
		Data:        make(map[int32]Driver),
	}

	calculateIranGrid(iran, 1)
	connections = make(map[int32]*MapGrid)

	return iran
}

func (mp *MapGrid) Insert(value Driver) {
	if mp.BottomLeft == nil && mp.BottomRight == nil && mp.TopLeft == nil && mp.TopRight == nil {
		mp.Data[value.ID] = value
		connections[value.ID] = mp
		return
	}

	if value.Cordinate.X >= mp.TopLeft.GridRange.xLeft && value.Cordinate.X <= mp.TopLeft.GridRange.xRight && value.Cordinate.Y <= mp.TopLeft.GridRange.yTop && value.Cordinate.Y >= mp.TopLeft.GridRange.yBottom {
		mp.TopLeft.Insert(value)
		return
	}
	if value.Cordinate.X >= mp.TopRight.GridRange.xLeft && value.Cordinate.X <= mp.TopRight.GridRange.xRight && value.Cordinate.Y <= mp.TopRight.GridRange.yTop && value.Cordinate.Y >= mp.TopRight.GridRange.yBottom {
		mp.TopRight.Insert(value)
		return
	}
	if value.Cordinate.X >= mp.BottomLeft.GridRange.xLeft && value.Cordinate.X <= mp.BottomLeft.GridRange.xRight && value.Cordinate.Y <= mp.BottomLeft.GridRange.yTop && value.Cordinate.Y >= mp.BottomLeft.GridRange.yBottom {
		mp.BottomLeft.Insert(value)
		return
	}
	if value.Cordinate.X >= mp.BottomRight.GridRange.xLeft && value.Cordinate.X <= mp.BottomRight.GridRange.xRight && value.Cordinate.Y <= mp.BottomRight.GridRange.yTop && value.Cordinate.Y >= mp.BottomRight.GridRange.yBottom {
		mp.BottomRight.Insert(value)
		return
	}

}

func (mp *MapGrid) Delete(id int32) {
	delete(connections[id].Data, id)
	delete(connections, id)
}

func (mp *MapGrid) Get(id int32) Driver {
	return connections[id].Data[id]
}

func (mp *MapGrid) Update(driver Driver) {
	mp.Delete(driver.ID)
	mp.Insert(driver)
}

func calculateIranGrid(mapGrid *MapGrid, level int) {
	calculate4Grid(mapGrid)
	if level == 10 {
		return
	}
	calculateIranGrid(mapGrid.TopLeft, level+1)
	calculateIranGrid(mapGrid.TopRight, level+1)
	calculateIranGrid(mapGrid.BottomLeft, level+1)
	calculateIranGrid(mapGrid.BottomRight, level+1)

}

func calculate4Grid(mp *MapGrid) {
	mp.TopLeft = &MapGrid{
		GridRange: &cordinateRange{
			xLeft:   mp.GridRange.xLeft,
			xRight:  (mp.GridRange.xLeft + mp.GridRange.xRight) / 2,
			yTop:    mp.GridRange.yTop,
			yBottom: (mp.GridRange.yTop + mp.GridRange.yBottom) / 2,
		},
		Data: make(map[int32]Driver),
	}
	mp.TopRight = &MapGrid{
		GridRange: &cordinateRange{
			xLeft:   (mp.GridRange.xLeft + mp.GridRange.xRight) / 2,
			xRight:  mp.GridRange.xRight,
			yTop:    mp.GridRange.yTop,
			yBottom: (mp.GridRange.yTop + mp.GridRange.yBottom) / 2,
		},
		Data: make(map[int32]Driver),
	}
	mp.BottomLeft = &MapGrid{
		GridRange: &cordinateRange{
			xLeft:   mp.GridRange.xLeft,
			xRight:  (mp.GridRange.xLeft + mp.GridRange.xRight) / 2,
			yTop:    (mp.GridRange.yTop + mp.GridRange.yBottom) / 2,
			yBottom: mp.GridRange.yBottom,
		},
		Data: make(map[int32]Driver),
	}
	mp.BottomRight = &MapGrid{
		GridRange: &cordinateRange{
			xLeft:   (mp.GridRange.xLeft + mp.GridRange.xRight) / 2,
			xRight:  mp.GridRange.xRight,
			yTop:    (mp.GridRange.yTop + mp.GridRange.yBottom) / 2,
			yBottom: mp.GridRange.yBottom,
		},
		Data: make(map[int32]Driver),
	}
}
