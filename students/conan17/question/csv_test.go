package question

import (
	"testing"
)

func TestSliceQuestion(t *testing.T) {
	type args struct {
		slice []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"success", args{[]string{"1+1", "2"}}, false},
		{"throw error", args{[]string{"1+1", "2", "3"}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := SliceQuestion(tt.args.slice)
			if (err != nil) != tt.wantErr {
				t.Errorf("SliceQuestion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestCSVQuizzes(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantLen int
		wantErr bool
	}{
		{"success", args{"../problem.csv"}, 12, false},
		{"format error", args{"../problem_err.csv"}, 0, true},
		{"not found", args{"../abc.csv"}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CSVQuizzes(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("CSVQuizzes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Size() != tt.wantLen {
				t.Errorf("CSVQuizzes() len = %d, want %d", got.Size(), tt.wantLen)
			}
		})
	}
}
