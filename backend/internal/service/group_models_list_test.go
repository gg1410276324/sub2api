package service

import "testing"

func TestNormalizeProviderBrand(t *testing.T) {
	if got := normalizeProviderBrand(" DeepSeek "); got != "deepseek" {
		t.Fatalf("expected normalized brand, got %q", got)
	}
	if got := normalizeProviderBrand("openai"); got != "" {
		t.Fatalf("unexpected unsupported brand %q", got)
	}
}
