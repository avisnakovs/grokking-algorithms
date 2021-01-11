package grokking_algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreedyStations(t *testing.T) {
	type args struct {
		stations     Stations
		statesNeeded States
	}
	tests := []struct {
		name string
		args args
		want []Station
	}{
		{
			"empty",
			args{
				stations:     nil,
				statesNeeded: States{},
			},
			[]Station{},
		},
		{
			"five stations",
			args{
				stations: map[Station]States{
					"one":   map[State]struct{}{"id": {}, "nv": {}, "ut": {}},
					"two":   map[State]struct{}{"wa": {}, "id": {}, "mt": {}},
					"three": map[State]struct{}{"or": {}, "nv": {}, "ca": {}},
					"four":  map[State]struct{}{"nv": {}, "ut": {}},
					"five":  map[State]struct{}{"ca": {}, "az": {}},
				},
				statesNeeded: map[State]struct{}{"mt": {}, "or": {}, "id": {}, "nv": {}, "ut": {}, "ca": {}, "az": {}},
			},
			[]Station{"two", "three", "one", "five"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GreedyStations(tt.args.stations, tt.args.statesNeeded)
			assert.Len(t, got, len(tt.want))
			for _, station := range tt.want {
				assert.Contains(t, got, station)
			}
		})
	}
}
