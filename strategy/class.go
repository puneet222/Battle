package strategy

type PlatoonClass int

const (
	Militia PlatoonClass = iota
	Spearmen
	LightCavalry
	HeavyCavalry
	FootArcher
	CavalryArcher
)

var advantageMap = map[PlatoonClass][]PlatoonClass{
	Militia: {Spearmen, LightCavalry},
	Spearmen: {LightCavalry, HeavyCavalry},
	LightCavalry: {FootArcher, CavalryArcher},
	HeavyCavalry: {Militia, FootArcher, LightCavalry},
	CavalryArcher: {Spearmen, HeavyCavalry},
	FootArcher: {Militia, CavalryArcher},
}

// HasAdvantageOver function returns whether a particular platoon class
// has an advantage over other platoon class based on advantage map info
func (pc PlatoonClass) HasAdvantageOver(opp PlatoonClass) bool {
	return contains(advantageMap[pc], opp);
}

// contains function returns whether a platoon class exists in a
// slice of platoon classes
func contains(classes []PlatoonClass, pc PlatoonClass) bool {
	for _, class := range classes {
		if class == pc {
			return true
		}
 	}
 	return false
}

func (pc PlatoonClass) String() string {
	return [...]string{"Militia", "Spearmen", "LightCavalry", "HeavyCavalry", "FootArcher", "CavalryArcher"}[pc]
}

func GetPlatoonClass(platoonClass string) PlatoonClass {
	switch platoonClass {
	case "Militia":
		return Militia
	case "Spearmen":
		return Spearmen
	case "LightCavalry":
		return LightCavalry
	case "HeavyCavalry":
		return HeavyCavalry
	case "FootArcher":
		return FootArcher
	case "CavalryArcher":
		return CavalryArcher
	default:
		panic("Platoon Class not supported")
	}
}
