package gorex

import "testing"

type testArgs struct {
	pattern string
	subject string
}

type testSimpleTestCases struct {
	name string
	args testArgs
	want bool
}

type errorWant struct {
	result  bool
	message string
}

type testErrorTestCases struct {
	name string
	args testArgs
	want errorWant
}

func TestTest(t *testing.T) {

	tests := testSimpleCases()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := Test(tt.args.pattern, tt.args.subject); got != tt.want {
				t.Errorf("Test() = %v, want %v", got, tt.want)
			}
		})
	}

	testsErrors := testErrorCases()

	for _, tt := range testsErrors {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := Test(tt.args.pattern, tt.args.subject); got != tt.want.result || err.Error() != tt.want.message {
				t.Errorf(
					"Test(%s, %s) = {%v, %v}, want %v",
					tt.args.pattern, tt.args.subject,
					got,
					err.Error(),
					tt.want,
				)
			}
		})
	}
}

func testSimpleCases() []testSimpleTestCases {

	return []testSimpleTestCases{
		{
			name: "Simple True Test 1",
			args: testArgs{
				pattern: "^[a-z]+$",
				subject: "abc",
			},
			want: true,
		},
		{
			name: "Simple True Test 2",
			args: testArgs{
				pattern: "^[a-zA-Z]+$",
				subject: "aBc",
			},
			want: true,
		},
		{
			name: "Simple False Test 1",
			args: testArgs{
				pattern: "^[a-z]+$",
				subject: "ABc",
			},
			want: false,
		},
		{
			name: "Simple False Test 2",
			args: testArgs{
				pattern: "^[a-zA-Z]+$",
				subject: "1ABC",
			},
			want: false,
		},

		{
			name: "Digit True Test 1",
			args: testArgs{
				pattern: "^[[:digit:]]+$",
				subject: "1234",
			},
			want: true,
		},
		{
			name: "Digit True Test 2",
			args: testArgs{
				pattern: "^[a-zA-Z[:digit:]]+$",
				subject: "ABC123a12",
			},
			want: true,
		},
		{
			name: "Digit True Test 3",
			args: testArgs{
				pattern: "^[^[:digit:]]+$",
				subject: "abd",
			},
			want: true,
		},
		{
			name: "Digit False Test 1",
			args: testArgs{
				pattern: "^[[:digit:]]+$",
				subject: "ABc",
			},
			want: false,
		},
		{
			name: "Digit False Test 2",
			args: testArgs{
				pattern: "^[[:digit:]a-z]+$",
				subject: "ABasdasdC",
			},
			want: false,
		},
		{
			name: "Digit false Test 3",
			args: testArgs{
				pattern: "^[^[:digit:]]+$",
				subject: "1234",
			},
			want: false,
		},
	}
}

func testErrorCases() []testErrorTestCases {

	return []testErrorTestCases{
		{
			name: "Invalid Regex1",
			args: testArgs{
				pattern: "^{[[:digit:]]+$",
				subject: "1234",
			},
			want: errorWant{
				result:  false,
				message: "Invalid preceding regular expression",
			},
		},
	}
}
