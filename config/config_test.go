package config

import (
	"os"
	"testing"
)

func TestGetCorpID(t *testing.T) {
	wantCorpID := "want"
	err := os.Setenv("CORP_ID", "want")
	outPut := GetCorpID()
	if err != nil {
		return
	}
	if wantCorpID != outPut {
		t.Fatalf("want %s, got %s", wantCorpID, outPut)
	}
}

func TestGetCorpSecret(t *testing.T) {
	tests := []struct {
		name           string
		wantCorpSecret string
	}{
		{wantCorpSecret: "CorpSecret"},
	}
	err := os.Setenv("CORP_SECRET", "CorpSecret")
	if err != nil {
		return
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCorpSecret := GetCorpSecret(); gotCorpSecret != tt.wantCorpSecret {
				t.Errorf("GetCorpSecret() = %v, want %v", gotCorpSecret, tt.wantCorpSecret)
			}
		})
	}
}
