package utils

import "testing"

func TestGetRiskProfileDefinition(t *testing.T) {
	type args struct {
		score int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{
				score: 39,
			},
			want: "Aggresive",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRiskProfileDefinition(tt.args.score); got != tt.want {
				t.Errorf("GetRiskProfileDefinition() = %v, want %v", got, tt.want)
			}
		})
	}
}
