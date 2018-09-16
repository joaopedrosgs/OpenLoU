package modules

type ConstructionType struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Bonus []struct {
		Name  string `json:"name"`
		Value []int  `json:"value"`
	} `json:"bonus"`
	Adjascent []struct {
		Builds []string `json:"builds"`
		Bonus  []int    `json:"bonus"`
	} `json:"adjascent"`
	ResourceCost [][]int `json:"resourceCost"`
	Score        []int   `json:"score"`
}

func (c ConstructionType) GetType() string {
	return "construction"
}
func (c ConstructionType) GetId() int {
	return c.ID
}
