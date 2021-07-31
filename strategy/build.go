package strategy

import (
	"sort"
	"strings"
)

const MinWins = 3
const DefeatResult = "There is no chance of winning"
const InputError = "Error wrong data input"

var canWinOverPlatoons = make(map[Platoon][]Platoon, 0)

func BuildStrategy(rawData string) string {
	armyInfo := strings.Split(rawData, "\n")
	// count how many times an opponent platoon can be defeated
	// used to set the order for platoon wins
	defeatMap := make(map[Platoon]int, 0)
	if len(armyInfo) == 2 {
		kingArmy := CreatePlatoons(armyInfo[0])
		oppArmy := CreatePlatoons(armyInfo[1])
		for _, platoon := range kingArmy {
			canWinOverPlatoons[platoon] = platoon.CanWinOver(oppArmy)
			for _, dp := range canWinOverPlatoons[platoon] {
				defeatMap[dp]++
			}
		}
		// sort win over platoons based on defeat map
		// lesser value in defeat map should be defeated first
		updateDefeatOrder(kingArmy, defeatMap)

		// sort kings army based on number of platoons it can handle
		sort.SliceStable(kingArmy, func(i, j int) bool {
			return len(canWinOverPlatoons[kingArmy[i]]) < len(canWinOverPlatoons[kingArmy[j]])
		})
		// battle
		isWin := battle(kingArmy, oppArmy)
		if isWin {
			iout := make([]string, 0)
			for _, platoon := range kingArmy {
				iout = append(iout, platoon.String())
			}
			return strings.Join(iout, ";")
		} else {
			return DefeatResult
		}
	} else {
		return InputError
	}
}

func updateDefeatOrder(kingArmy []Platoon, defeatMap map[Platoon]int) {
	for _, platoon := range kingArmy {
		defeatedPlatoons := canWinOverPlatoons[platoon]
		sort.SliceStable(defeatedPlatoons, func(i, j int) bool {
			return defeatMap[defeatedPlatoons[i]] < defeatMap[defeatedPlatoons[j]]
		})
		canWinOverPlatoons[platoon] = defeatedPlatoons
	}
}

func battle(kingArmy []Platoon, oppArmy []Platoon) bool {
	wins := 0
	oppArmyStatus := make(map[Platoon]bool, 0)
	for _, op := range oppArmy {
		oppArmyStatus[op] = true // initially alive
	}
	for _, kp := range kingArmy {
		canHandlePlatoons := canWinOverPlatoons[kp]
		for _, op := range canHandlePlatoons {
			if oppArmyStatus[op] {
				// is alive
				oppArmyStatus[op] = false // defeated
				wins++
				break
			}
		}
	}
	return wins > MinWins
}
