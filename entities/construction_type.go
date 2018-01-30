package entities

type ConstructionType struct {
	ID    uint   `json:"id"`
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
