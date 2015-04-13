package slugify

import "testing"

func TestSlugify(t *testing.T) {
	var tests = []struct{ in, out string }{
		{"simple test", "simple-test"},
		{"I'm go developer", "i-m-go-developer"},
		{"Simples código em go", "simples-codigo-em-go"},
		{"日本語の手紙をテスト", "日本語の手紙をテスト"},
	}

	for _, test := range tests {
		if out := Slugify(test.in); out != test.out {
			t.Errorf("%q: %q != %q", test.in, out, test.out)
		}
	}
}
