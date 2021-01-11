package grokking_algorithms

type (
	Station  string
	State    string
	States   map[State]struct{}
	Stations map[Station]States
)

func (s States) copy() States {
	statesCopy := make(States)
	for state := range s {
		statesCopy[state] = struct{}{}
	}
	return statesCopy
}

func (s States) intersection(states States) States {
	statesCopy := make(States)
	for state := range states {
		if s.contains(state) {
			statesCopy[state] = struct{}{}
		}
	}
	return statesCopy
}

func (s States) contains(search State) bool {
	for state := range s {
		if state == search {
			return true
		}
	}
	return false
}

func (s States) remove(states States) States {
	statesCopy := make(States)
	for state := range s {
		if !states.contains(state) {
			statesCopy[state] = struct{}{}
		}
	}
	return statesCopy
}

func GreedyStations(stations Stations, statesNeeded States) []Station {
	result := make([]Station, 0, len(stations))
	needed := statesNeeded.copy()
	for len(needed) > 0 {
		var bestStation Station
		bestStationCoverage := make(map[State]struct{})
		for station, states := range stations {
			coveredStates := states.intersection(needed)
			if len(coveredStates) > len(bestStationCoverage) {
				bestStationCoverage = coveredStates
				bestStation = station
			}
		}
		needed = needed.remove(bestStationCoverage)
		result = append(result, bestStation)
	}
	return result
}
