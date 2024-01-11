package day02

import (
	"reflect"
	"testing"
)

func Test_parseGame(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want Game
	}{
		{
			name: "happy",
			args: args{"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue"},
			want: Game{
				id:        2,
				highRed:   1,
				highGreen: 3,
				highBlue:  4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseGame(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseGame() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseRounds(t *testing.T) {
	type args struct {
		game         *Game
		roundsString string
	}
	tests := []struct {
		name string
		args args
		want Game
	}{
		{
			name: "happy",
			args: args{
				game:         newGame(),
				roundsString: "1 red, 1 green, 1 blue; 3 blue, 2 green, 1 red",
			},
			want: Game{
				highRed:   1,
				highGreen: 2,
				highBlue:  3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parseRounds(tt.args.game, tt.args.roundsString)
			if tt.args.game.highRed != tt.want.highRed ||
				tt.args.game.highGreen != tt.want.highGreen ||
				tt.args.game.highBlue != tt.want.highBlue {
				t.Errorf("parseRounds() = %v, want %v", tt.args.game, tt.want)
			}
		})
	}
}

func Test_parseRound(t *testing.T) {
	type args struct {
		game  *Game
		round string
	}
	tests := []struct {
		name string
		args args
		want Game
	}{
		{
			name: "happy",
			args: args{
				game:  newGame(),
				round: "1 red, 1 green, 1 blue",
			},
			want: Game{
				highRed:   1,
				highGreen: 1,
				highBlue:  1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parseRound(tt.args.game, tt.args.round)
			if tt.args.game.highRed != tt.want.highRed ||
				tt.args.game.highGreen != tt.want.highGreen ||
				tt.args.game.highBlue != tt.want.highBlue {
				t.Errorf("parseRound() = %v, want %v", tt.args.game, tt.want)
			}
		})
	}
}
