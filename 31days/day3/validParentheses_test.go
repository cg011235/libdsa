package day3

import "testing"

type stack struct {
	arr []rune
	top int
}

func (st *stack) push(c rune) {
	st.arr = append(st.arr, c)
}

func (st *stack) pop() rune {
	if len(st.arr) == 0 {
		return 0
	}
	ret := st.arr[len(st.arr)-1]
	st.arr = st.arr[:len(st.arr)-1]
	return ret
}

func (st *stack) isEmpty() bool {
	return len(st.arr) == 0
}

func validate(s string) bool {
	var st stack
	for _, c := range s {
		switch c {
		case '{':
			fallthrough
		case '[':
			fallthrough
		case '(':
			st.push(c)

		case '}':
			r := st.pop()
			if r != '{' {
				return false
			}
		case ']':
			r := st.pop()
			if r != '[' {
				return false
			}
		case ')':
			r := st.pop()
			if r != '(' {
				return false
			}
		}
	}
	if st.isEmpty() {
		return true
	}
	return false
}

func TestValidate(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
		{"", true}, // Edge case: empty string should be considered valid
		{"{", false},
		{"}", false},
		{"{}}", false},
		{"((()))", true},
		{"((())", false},
		{"(()))", false},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			actual := validate(tc.input)
			if actual != tc.expected {
				t.Errorf("For input \"%s\", expected %v but got %v", tc.input, tc.expected, actual)
			}
		})
	}
}
