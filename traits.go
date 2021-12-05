package opensea

type Trait struct {
	Type        string      `json:"trait_type"`
	Value       interface{} `json:"value"`
	DisplayType string      `json:"display_type"`
	MaxValue    int64       `jsson:"max_value"`
	Count       int64       `json:"trait_count"`
	Order       string      `json:"order"`
}

type BoostTrait struct {
	Value       interface{}
	DisplayType string
}

type CosmeticTrait struct {
	Value string
	Type  string
	Count int
}

func displayTraitValue(traits []Trait, name string) interface{} {
	for _, trait := range traits {
		if trait.Type == name {
			return trait.Value
		}
	}
	return nil
}

func organizeTraits(traits []Trait) (map[string]BoostTrait, []CosmeticTrait) {
	var cosmetics []CosmeticTrait
	boosts := make(map[string]BoostTrait)
	for _, trait := range traits {
		if trait.DisplayType != "" {
			boosts[trait.Type] = BoostTrait{
				Value:       trait.Value,
				DisplayType: trait.DisplayType,
			}
		} else {
			cosmetics = append(cosmetics, CosmeticTrait{
				Type:  trait.Type,
				Value: trait.Value.(string),
				Count: int(trait.Count),
			})
		}
	}
	return boosts, cosmetics
}
