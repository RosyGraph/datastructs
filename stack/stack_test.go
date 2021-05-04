package stack

import "testing"

func TestStack(t *testing.T) {
	tests := []struct {
		name string
		proc func() string
		want string
	}{
		{
			name: "init stack",
			proc: func() string {
				s := Stack{1, nil}
				return s.String()
			},
			want: "1",
		},
		{
			name: "add one to stack",
			proc: func() string {
				s := Stack{1, nil}
				s = s.Add(2)
				return s.String()
			},
			want: "2 -> 1",
		},
		{
			name: "add two, pop one to stack",
			proc: func() string {
				s := Stack{1, nil}
				s = s.Add(2)
				s, _ = s.Pop()
				return s.String()
			},
			want: "1",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.proc()
			if got != test.want {
				t.Errorf("test %s: got %s want %s", test.name, got, test.want)
			}
		})
	}
}
