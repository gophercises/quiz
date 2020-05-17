package manager

import "testing"

func Test_clean(t *testing.T) {
	type args struct {
		answer string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Remove beginning space",
			args: args{" test"},
			want: "test",
		},
		{
			name: "Remove final space",
			args: args{"test "},
			want: "test",
		},
		{
			name: "Remove both spaces",
			args: args{" test "},
			want: "test",
		},
		{
			name: "Remove endline",
			args: args{"\ntest\n"},
			want: "test",
		},
		{
			name: "Remove tabs",
			args: args{"\ttest\t"},
			want: "test",
		},
		{
			name: "Remove weird windows",
			args: args{"\rtest\r"},
			want: "test",
		},
		{
			name: "Ignore case",
			args: args{"TeSt"},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := clean(tt.args.answer); got != tt.want {
				t.Errorf("clean() = %v, want %v", got, tt.want)
			}
		})
	}
}
