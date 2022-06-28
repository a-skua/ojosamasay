package ojosama

import (
	"testing"
)

func Test_size(t *testing.T) {
	type test struct {
		testcase string
		text     string
		want     int
	}

	do := func(tt *test) {
		t.Run(tt.testcase, func(t *testing.T) {
			got := size(tt.text)
			if tt.want != got {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test{
		{
			text: "hello, サロメイト",
			want: 17,
		},
		{
			text: "おハーブですわ〜",
			want: 16,
		},
		{
			text: "とてもきったねぇですわ～～‼",
			want: 27,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}
