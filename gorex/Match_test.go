package gorex

import "testing"

type matchArgs struct {
	pattern string
	subject string
}

type matchTestCases struct {
	name string
	args matchArgs
	want bool
}

func TestMatch(t *testing.T) {

	tests := matchCases()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Match(tt.args.pattern, tt.args.subject); got != tt.want {
				t.Errorf("Match() = %v, want %v", got, tt.want)
			}
		})
	}
}

func matchCases() []matchTestCases {

	return []matchTestCases{
		{
			name: "Simple True Test 1",
			args: matchArgs{
				pattern: "^[a-z]+$",
				subject: "abc",
			},
			want: true,
		},
		{
			name: "Simple True Test 2",
			args: matchArgs{
				pattern: "^[a-zA-Z]+$",
				subject: "aBc",
			},
			want: true,
		},
		{
			name: "Simple False Test 1",
			args: matchArgs{
				pattern: "^[a-z]+$",
				subject: "ABc",
			},
			want: false,
		},
		{
			name: "Simple False Test 2",
			args: matchArgs{
				pattern: "^[a-zA-Z]+$",
				subject: "1ABC",
			},
			want: false,
		},

		{
			name: "Digit True Test 1",
			args: matchArgs{
				pattern: "^[[:digit:]]+$",
				subject: "1234",
			},
			want: true,
		},
		{
			name: "Digit True Test 2",
			args: matchArgs{
				pattern: "^[a-zA-Z[:digit:]]+$",
				subject: "ABC123a12",
			},
			want: true,
		},
		{
			name: "Digit True Test 3",
			args: matchArgs{
				pattern: "^[^[:digit:]]+$",
				subject: "abd",
			},
			want: true,
		},
		{
			name: "Digit False Test 1",
			args: matchArgs{
				pattern: "^[[:digit:]]+$",
				subject: "ABc",
			},
			want: false,
		},
		{
			name: "Digit False Test 2",
			args: matchArgs{
				pattern: "^[[:digit:]a-z]+$",
				subject: "ABasdasdC",
			},
			want: false,
		},
		{
			name: "Digit false Test 3",
			args: matchArgs{
				pattern: "^[^[:digit:]]+$",
				subject: "1234",
			},
			want: false,
		},
	}
}
