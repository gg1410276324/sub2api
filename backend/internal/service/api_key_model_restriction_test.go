package service

import (
	"context"
	"errors"
	"testing"

	"github.com/Wei-Shaw/sub2api/internal/pkg/ctxkey"
)

func TestAPIKeyModelRestriction(t *testing.T) {
	group := &Group{ModelsListConfig: GroupModelsListConfig{
		Enabled: true,
		Models:  []string{"deepseek-v4-flash", "deepseek-v4-pro"},
	}}
	if err := validateAPIKeyAllowedModel(group, "deepseek-v4-pro"); err != nil {
		t.Fatalf("configured model rejected: %v", err)
	}
	if err := validateAPIKeyAllowedModel(group, "gpt-5.6-sol"); !errors.Is(err, ErrAPIKeyModelNotAvailable) {
		t.Fatalf("unexpected model error: %v", err)
	}
	if err := validateAPIKeyAllowedModel(group, ""); !errors.Is(err, ErrAPIKeyModelRequired) {
		t.Fatalf("configured groups must require a model: %v", err)
	}

	ctx := WithAPIKeyAllowedModel(context.Background(), "deepseek-v4-pro")
	if err := ValidateAPIKeyAllowedModel(ctx, "deepseek-v4-pro"); err != nil {
		t.Fatalf("bound model rejected: %v", err)
	}
	if err := ValidateAPIKeyAllowedModel(ctx, "deepseek-v4-flash"); !errors.Is(err, ErrAPIKeyModelNotAllowed) {
		t.Fatalf("different model should be blocked: %v", err)
	}
}

func TestAPIKeyModelRestrictionUsesOriginalModelBeforeUpstreamMapping(t *testing.T) {
	ctx := WithAPIKeyAllowedModel(context.Background(), "deepseek-chat")
	ctx = context.WithValue(ctx, ctxkey.Model, "deepseek-chat")
	if err := ValidateAPIKeyAllowedModel(ctx, "upstream-deepseek-v3"); err != nil {
		t.Fatalf("mapped upstream model should use the original request model: %v", err)
	}

	ctx = context.WithValue(ctx, ctxkey.Model, "deepseek-reasoner")
	if err := ValidateAPIKeyAllowedModel(ctx, "upstream-deepseek-r1"); !errors.Is(err, ErrAPIKeyModelNotAllowed) {
		t.Fatalf("wrong original request model should be rejected, got %v", err)
	}
}
