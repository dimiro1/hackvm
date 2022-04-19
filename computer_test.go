package hackvm

import "testing"

func TestComputer(t *testing.T) {
	cases := []struct {
		Name          string
		Path          string
		Expected      int
		ResultAddress int
		Steps         int
		Ram           []int
	}{
		{
			Name:          "addition",
			Path:          "testdata/add.hack",
			Expected:      5,
			ResultAddress: 0,
			Steps:         6,
		},
		{
			Name:          "multiplication",
			Path:          "testdata/mul.hack",
			Expected:      25,
			ResultAddress: 2,
			Steps:         100,
			Ram:           []int{5, 5},
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			c := NewComputer()
			err := c.LoadRomFromFile(tc.Path)
			if err != nil {
				t.Fatalf("unexpected error: %s", err.Error())
			}

			for address, data := range tc.Ram {
				if err := c.WriteRam(address, data); err != nil {
					t.Fatalf("unexpected error: %s", err.Error())
				}
			}

			for i := 0; i < tc.Steps; i++ {
				if err := c.Step(); err != nil {
					t.Fatalf("unexpected error: %s", err.Error())
				}
			}

			got, err := c.ReadRam(tc.ResultAddress)
			if err != nil {
				t.Fatalf("unexpected error: %s", err.Error())
			}

			if got != tc.Expected {
				t.Errorf("expected = %d, got = %d", tc.Expected, got)
			}
		})
	}
}
