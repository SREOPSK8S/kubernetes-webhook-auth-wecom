package config

import (
	"os"
	"testing"
)

func TestGetCorpID(t *testing.T) {
	wantCorpID := "want"
	err := os.Setenv("CORP_ID", "want")
	if err != nil {
		t.Fatalf("set env error %v\n",err)
	}
	outPut := GetCorpID()
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
		t.Fatalf("set env error %v\n",err)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCorpSecret := GetCorpSecret(); gotCorpSecret != tt.wantCorpSecret {
				t.Errorf("GetCorpSecret() = %v, want %v", gotCorpSecret, tt.wantCorpSecret)
			}
		})
	}
}

func TestGetServerPort(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		// TODO: Add test cases.
		{
			name: "GetServerPort",
			want: 8443,
		},
	}
	os.Setenv("SERVICE_PORT","8443")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetServerPort(); got != tt.want {
				t.Errorf("GetServerPort() = %v, want %v", got, tt.want)
			}
		})
	}
}