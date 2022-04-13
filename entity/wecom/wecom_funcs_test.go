package wecom

import (
	"reflect"
	"testing"
)

func TestAccessTokenResponse_GetAccessToken(t *testing.T) {
	type fields struct {
		BaseResponse BaseResponse
		AccessToken  string
		ExpiresIn    int
	}
	tests := []struct {
		name            string
		fields          fields
		wantAccessToken string
	}{
		{
			name: "GetAccessToken",
			fields: fields{
				BaseResponse: BaseResponse{
					ErrorCode:    0,
					ErrorMessage: "",
				},
				AccessToken:  "123456",
				ExpiresIn:    7200,
			},
			wantAccessToken: "123456",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ac := &AccessTokenResponse{
				BaseResponse: tt.fields.BaseResponse,
				AccessToken:  tt.fields.AccessToken,
				ExpiresIn:    tt.fields.ExpiresIn,
			}
			if gotAccessToken := ac.GetAccessToken(); gotAccessToken != tt.wantAccessToken {
				t.Errorf("GetAccessToken() = %v, want %v", gotAccessToken, tt.wantAccessToken)
			}
		})
	}
}

func TestBaseResponse_GetErrorCode(t *testing.T) {
	type fields struct {
		ErrorCode    int
		ErrorMessage string
	}
	tests := []struct {
		name          string
		fields        fields
		wantErrorCode int
	}{
		{
			name: "GetErrorCode",
			fields: fields{
				ErrorCode:    1,
				ErrorMessage: "",
			},
			wantErrorCode: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			base := &BaseResponse{
				ErrorCode:    tt.fields.ErrorCode,
				ErrorMessage: tt.fields.ErrorMessage,
			}
			if gotErrorCode := base.GetErrorCode(); gotErrorCode != tt.wantErrorCode {
				t.Errorf("GetErrorCode() = %v, want %v", gotErrorCode, tt.wantErrorCode)
			}
		})
	}
}

func TestCorpIDAndSecret_GetCorpIDAndSecret(t *testing.T) {
	type fields struct {
		CorpID     string
		CorpSecret string
	}
	tests := []struct {
		name   string
		fields fields
		want   *CorpIDAndSecret
	}{
		{
			name: "GetCorpIDAndSecret",
			fields: fields{
				CorpID:     "123",
				CorpSecret: "123456",
			},
			want: &CorpIDAndSecret{
				CorpID:     "123",
				CorpSecret: "123456",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			IDS := &CorpIDAndSecret{
				CorpID:     tt.fields.CorpID,
				CorpSecret: tt.fields.CorpSecret,
			}
			if got := IDS.GetCorpIDAndSecret(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCorpIDAndSecret() = %v, want %v", got, tt.want)
			}
		})
	}
}