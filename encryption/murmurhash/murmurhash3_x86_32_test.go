package murmurhash

import "testing"

func TestMurmurHash3_x86_32(t *testing.T) {
	cases := []struct {
		input  string
		output uint32
	}{
		{input: "hello", output: 613153351},
		{input: "world", output: 4220927227},
	}

	for _, c := range cases {
		if MurmurHash3_x86_32([]byte(c.input), 0) != c.output {
			t.FailNow()
		}
	}
}
