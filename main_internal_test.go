package main

import "testing"

func TestGreet(t *testing.T) {
	type testCase struct {
		lang language
		want string
	}

	var tests = map[string]testCase{
		"English":                 {lang: "en", want: "Hello world"},
		"Greek":                   {lang: "el", want: "Γειά σου κόσμε"},
		"French":                  {lang: "fr", want: "Bonjour le monde"},
		"Hebrew":                  {lang: "he", want: "שלום עולם"},
		"Urdu":                    {lang: "ur", want: "ہیلو دنیا"},
		"Vietnamese":              {lang: "vi", want: "Xin chào thế giới"},
		"Akkadian, not supported": {lang: "akk", want: `unsupported language: "akk"`},
		"Indonesian":              {lang: "id", want: "Halo dunia"},
		"Empty":                   {lang: "", want: `unsupported language: ""`},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := greet(tc.lang)

			if got != tc.want {
				t.Errorf("expected: %q, got: %q", tc.want, got)
			}
		})
	}
}
