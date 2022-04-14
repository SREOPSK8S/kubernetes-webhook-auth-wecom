package stores

import (
	"context"
	"testing"
)

func TestEtcdImpl_SetSoreAccessToken(t *testing.T) {
	type args struct {
		ctx   context.Context
		token string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "SetSoreAccessToken",
			args: args{
				ctx:   context.Background(),
				token: "123456",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			et := EtcdImpl{}
			if got := et.SetSoreAccessToken(tt.args.ctx, tt.args.token); got != tt.want {
				t.Errorf("SetSoreAccessToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEtcdImpl_GetSoreAccessToken(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 bool
	}{
		{
			name: "GetSoreAccessToken",
			args: args{
				ctx: context.Background(),
			},
			want: "123456",
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			et := EtcdImpl{}
			got, got1 := et.GetSoreAccessToken(tt.args.ctx)
			if got != tt.want {
				t.Errorf("GetSoreAccessToken() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetSoreAccessToken() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestEtcdImpl_DeleteAccessToken(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "DeleteAccessToken",
			args: args{
				ctx: context.Background(),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			et := EtcdImpl{}
			if got := et.DeleteAccessToken(tt.args.ctx); got != tt.want {
				t.Errorf("DeleteAccessToken() = %v, want %v", got, tt.want)
			}
		})
	}
}