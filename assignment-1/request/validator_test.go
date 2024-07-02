package request

import (
	"reflect"
	"testing"

	"github.com/an-halim/golang-advanced/assignment-1/constant"
)

func TestValidateProfile(t *testing.T) {
	type args struct {
		answers []Answers
	}
	tests := []struct {
		name           string
		args           args
		wantResult     constant.ProfileRisk
		wantTotalScore int
	}{
		{
			name: "Test case 1",
			args: args{
				answers: []Answers{
					{
						QuestionID: 1,
						Answer:     "Pertumbuhan kekayaan untuk jangka panjang",
					},
					{
						QuestionID: 2,
						Answer:     "≥ 10 tahun",
					},
					{
						QuestionID: 3,
						Answer:     "> 10 tahun",
					},
					{
						QuestionID: 4,
						Answer:     "Saham, Reksa Dana terbuka, equity linked structure product",
					},
					{
						QuestionID: 5,
						Answer:     "> 50%",
					},
					{
						QuestionID: 6,
						Answer:     "< -20% - > +20%",
					},
					{
						QuestionID: 7,
						Answer:     "Tidak bergantung pada hasil investasi",
					},
					{
						QuestionID: 8,
						Answer:     "> 50%",
					},
				},
			},
			wantResult: constant.ProfileRisk{
				MinScore:   36,
				MaxScore:   40,
				Category:   "Aggresive",
				Definition: "Anda sangat berpengalaman terhadap produk investasi dan memiliki toleransi yang sangat tinggi atasproduk-produk investasi. Anda bahkan dapat menerima perubahan signifikan pada modal/nilai investasi.Pada umumnya portfolio Anda sebagian besar dialokasikan pada produk investasi.",
			},
			wantTotalScore: 40,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, gotTotalScore := ValidateProfile(tt.args.answers)
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("ValidateProfile() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
			if gotTotalScore != tt.wantTotalScore {
				t.Errorf("ValidateProfile() gotTotalScore = %v, want %v", gotTotalScore, tt.wantTotalScore)
			}
		})
	}
}
