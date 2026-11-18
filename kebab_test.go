package main

import "testing"

func TestToKebabCase(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{"Hello World", "hello-world"},
		{"helloWorld", "hello-world"},
		{"HelloWorld", "hello-world"},
		{"hello_world", "hello-world"},
		{"hello-world", "hello-world"},
		{"  hello   world  ", "hello-world"},
		{"HTTPServer", "http-server"},
		{"getHTTPResponseCode", "get-http-response-code"},
		{"userID", "user-id"},
		{"version2", "version-2"},
		{"v2Beta", "v-2-beta"},
		{"already-kebab-case", "already-kebab-case"},
		{"Hello, World!", "hello-world"},
		{"ÉcoleNormale", "école-normale"},
		{"", ""},
		{"---", ""},
		{"A", "a"},
	}
	for _, tt := range tests {
		if got := ToKebabCase(tt.in); got != tt.want {
			t.Errorf("ToKebabCase(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func BenchmarkToKebabCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ToKebabCase("getHTTPResponseCode for userID 42")
	}
}
