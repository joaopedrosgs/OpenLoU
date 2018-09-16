package modules

type TroopType struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CanAttack bool   `json:"canAttack"`
	Image     string `json:"image"`
	Attack    int    `json:"attack"`
	Defense   int    `json:"defense"`
	Carry     int    `json:"loot"`
	Speed     int    `json:"speed"`
	Requires  []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"requires"`
	Cost   [5]int
	Upkeep [2]int
}

func (t TroopType) GetType() string {
	return "troop"
}

func (t TroopType) GetId() int {
	return t.ID
}
