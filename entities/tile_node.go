package entities

type TileNode struct {
	X           uint      `gorm:"index:idx_x_y"`
	Y           uint      `gorm:"index:idx_x_y"`
	Continent   Continent `gorm:"ForeignKey:ContinentID"`
	ContinentID uint      `gorm:"index"`
	Type        uint
}
