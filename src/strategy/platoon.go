package strategy

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Platoon struct {
	Class PlatoonClass
	Soldiers int
}

func (p Platoon) String() string {
	return fmt.Sprintf("%v#%d", p.Class.String(), p.Soldiers)
}

// CreatePlatoons function converts raw data to slice of platoons for an army
// this function assuming raw data of type class1#num1;class2#num2
// e.g.  Spearmen#10;Militia#30;FootArcher#20;LightCavalry#1000;HeavyCavalry#120
func CreatePlatoons(rawData string) []Platoon {
	platoons := make([]Platoon, 0);
	rawPlatoons := strings.Split(rawData, ";");
	for _, p := range rawPlatoons {
		tokens := strings.Split(p, "#")
		class := GetPlatoonClass(tokens[0]);
		soldiers, err := strconv.Atoi(tokens[1])
		if err != nil {
			log.Printf("Error while converting string to int for soldiers %v", err)
		}
		platoons = append(platoons, Platoon{class, soldiers})
	}
	return platoons
}

// CanWin function returns whether a particular platoon
//	can win over opponent platoon
func (p Platoon) CanWin(opp Platoon) bool {
	if p.Class.HasAdvantageOver(opp.Class) {
		// if platoon has advantage over opponent class it can
		// handle double soldiers of opponent class
		p.Soldiers *= 2
	}
	return p.Soldiers > opp.Soldiers
}

// CanWinOver function returns a list of opp platoons
//	whom a particular platoon can win
func (p Platoon) CanWinOver(oppPlatoons []Platoon) []Platoon {
	winPlatoons := make([]Platoon, 0)
	for _, op := range oppPlatoons {
		if p.CanWin(op) {
			winPlatoons = append(winPlatoons, op)
		}
	}
	return winPlatoons
}

// CanDefeatedBy function returns a list of platoons
//	who can defeat a particular platoon
func (p Platoon) CanDefeatedBy(platoons []Platoon) []Platoon {
	defeatPlatoons := make([]Platoon, 0)
	for _, platoon := range platoons {
		if platoon.CanWin(p) {
			defeatPlatoons = append(defeatPlatoons, platoon)
		}
	}
	return defeatPlatoons
}