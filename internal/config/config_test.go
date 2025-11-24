package config

import "testing"

func TestConfigDefaults(t *testing.T) {
	cfg, err := New()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.Env != "dev" {
		t.Errorf("expected dev env default, got %s", cfg.Env)
	}
	if cfg.CdnURL == "" {
		t.Errorf("expected default CDN_URL, got empty")
	}
	if cfg.PostBasePath == "" {
		t.Errorf("expected default PostsBasePath, got empty")
	}
}

func TestConfigProdMissingVarsErrors(t *testing.T) {
	t.Setenv("ENV", "prod")
	_, err := New()
	if err == nil {
		t.Fatalf("expected error for missing CDN_URL + GCS_BUCKET in prod")
	}
}

func TestConfigProdOnlyOneMissing(t *testing.T) {
	t.Setenv("ENV", "prod")
	t.Setenv("CDN_URL", "https://cdn.example")
	_, err := New()
	if err == nil {
		t.Fatalf("expected error for missing GCS_BUCKET")
	}
}

func TestConfigProdAllPresent(t *testing.T) {
	t.Setenv("ENV", "prod")
	t.Setenv("CDN_URL", "https://cdn.example")
	t.Setenv("GCS_BUCKET", "bucket-name")
	cfg, err := New()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.CdnURL != "https://cdn.example" {
		t.Errorf("cdn mismatch")
	}
	if cfg.GCSBucket != "bucket-name" {
		t.Errorf("bucket mismatch")
	}
}
