package entities

type TileNode struct {
	X           uint      `gorm:"unique_index:node_idx_x_y"`
	Y           uint      `gorm:"unique_index:node_idx_x_y"`
	Continent   Continent `gorm:"ForeignKey:ContinentID"`
	ContinentID uint      `gorm:"index"`
	Type        uint
}
