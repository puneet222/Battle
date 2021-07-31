package strategy

import (
	"sort"
	"strings"
)

const MinWins = 3
const DefeatResult = "There is no chance of winning"
const InputError = "Error wrong data input"

var canDefeatedByPlatoons = make(map[Platoon][]Platoon, 0)

func BuildStrategy(rawData string) string {
	armyInfo := strings.Split(rawData, "\n")
	// count how many times a platoon can defeat an opponent platoon
	defeatMap := make(map[Platoon]int, 0)
	if len(armyInfo) == 2 {
		kingArmy := CreatePlatoons(armyInfo[0])
		oppArmy := CreatePlatoons(armyInfo[1])
		for _, op := range oppArmy {
			canDefeatedByPlatoons[op] = op.CanDefeatedBy(kingArmy)
			for _, wp := range canDefeatedByPlatoons[op] {
				defeatMap[wp]++
			}
		}
		// sort win over platoons based on defeat map
		// lesser value in defeat map can defeat opponent first
		updateDefeatOrder(oppArmy, defeatMap)
		// battle
		order, isWin := battle(kingArmy, oppArmy)
		if isWin {
			iout := make([]string, 0)
			for _, platoon := range order {
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

func updateDefeatOrder(oppArmy []Platoon, defeatMap map[Platoon]int) {
	for _, op := range oppArmy {
		winnerPlatoons := canDefeatedByPlatoons[op]
		sort.SliceStable(winnerPlatoons, func(i, j int) bool {
			return defeatMap[winnerPlatoons[i]] < defeatMap[winnerPlatoons[j]]
		})
		canDefeatedByPlatoons[op] = winnerPlatoons
	}
}

func battle(kingArmy []Platoon, oppArmy []Platoon) ([]*Platoon, bool) {
	wins := 0
	kingArmyStatus := make(map[Platoon]bool, 0)
	kingArmyOrder := make([]*Platoon, len(kingArmy))
	for _, kp := range kingArmy {
		kingArmyStatus[kp] = true // initially alive
	}
	for i, op := range oppArmy {
		defeatPlatoons := canDefeatedByPlatoons[op]
		for _, kp := range defeatPlatoons {
			if kingArmyStatus[kp] {
				// is alive can fight

				kingArmyStatus[kp] = false // defeated the opponent
				wins++
				kingArmyOrder[i] = &kp
				break
			}
		}
	}
	// fill all lose battles
	for i, p := range kingArmyOrder {
		if p == nil {
			kingArmyOrder[i] = getAlivePlatoon(kingArmyStatus)
		}
	}
	return kingArmyOrder, wins > MinWins
}

func getAlivePlatoon(kingArmyStatus map[Platoon]bool) *Platoon {
	for p, isAlive := range kingArmyStatus {
		if isAlive {
			kingArmyStatus[p] = false
			return &p
		}
	}
	return nil
}
