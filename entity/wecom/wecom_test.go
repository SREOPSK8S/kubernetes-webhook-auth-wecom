package wecom

import "testing"

func TestAccessTokenResponse_GetAccessToken(t *testing.T) {
	type fields struct {
		ErrorCode    int
		ErrorMessage string
		AccessToken  string
		ExpiresIn    int
	}
	tests := []struct {
		name            string
		fields          fields
		wantAccessToken string
	}{
		{
			name: "accessToken",
			fields: fields{
				ErrorCode: 0,
				ErrorMessage: "",
				AccessToken: "Token",
				ExpiresIn: 7200,
			},
			wantAccessToken: "Token",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ac := &AccessTokenResponse{
				ErrorCode:    tt.fields.ErrorCode,
				ErrorMessage: tt.fields.ErrorMessage,
				AccessToken:  tt.fields.AccessToken,
				ExpiresIn:    tt.fields.ExpiresIn,
			}
			if gotAccessToken := ac.GetAccessToken(); gotAccessToken != tt.wantAccessToken {
				t.Errorf("GetAccessToken() = %v, want %v", gotAccessToken, tt.wantAccessToken)
			}
		})
	}
}
