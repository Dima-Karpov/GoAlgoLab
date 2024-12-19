package countLetters

import "testing"

func TestCountLetters(t *testing.T) {
	type args struct {
		s string
		l byte
	}
	tests := []struct {
		name  string
		args  args
		wantC int
	}{
		{name: "test1", args: args{s: "hello world", l: byte('l')}, wantC: 3},
		{name: "test2", args: args{s: "abcb", l: byte('b')}, wantC: 2},
		{name: "test3", args: args{s: "", l: byte('t')}, wantC: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotC := CountLetters(tt.args.s, tt.args.l); gotC != tt.wantC {
				t.Errorf("CountLetters() = %v, want %v", gotC, tt.wantC)
			}
		})
	}
}
