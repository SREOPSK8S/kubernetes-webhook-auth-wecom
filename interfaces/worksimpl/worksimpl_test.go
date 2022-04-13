package worksimpl
//
//import (
//	"github.com/SREOPSK8S/kubernetes-webhook-auth-wecom/entity/auth"
//	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"reflect"
//	"testing"
//)
//
//func TestWorkChatImpl_TokenReviewVerify(t *testing.T) {
//	type args struct {
//		review auth.TokenReview
//	}
//	tests := []struct {
//		name string
//		args args
//		want bool
//	}{
//		{
//			name: "TokenReviewVerifyPass",
//			args: args{
//				review: auth.TokenReview{
//					TypeMeta: metav1.TypeMeta{
//						Kind:       "TokenReview",
//						APIVersion: "authentication.k8s.io/v1beta1",
//					},
//					Spec:     auth.TokenReviewSpec{
//						Token:     "Y2hhb3lhbmcK",
//						Audiences: []string{"dev"},
//					},
//				},
//			},
//			want: true,
//		},
//		{
//			name: "TokenReviewVerifyFaire",
//			args: args{
//				review: auth.TokenReview{
//					TypeMeta: metav1.TypeMeta{
//						Kind:       "TokenReview",
//						APIVersion: "authentication.k8s.io/v1beta1",
//					},
//					Spec:     auth.TokenReviewSpec{
//						Token:     "Y2hhb3lhbmcK123",
//						Audiences: []string{"dev"},
//					},
//				},
//			},
//			want: false,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			wo := WorkChatImpl{}
//			if got := wo.TokenReviewVerify(tt.args.review); got != tt.want {
//				t.Errorf("TokenReviewVerify() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestWorkChatImpl_TokenReviewFailure(t *testing.T) {
//	type args struct {
//		review auth.TokenReview
//	}
//	tests := []struct {
//		name string
//		args args
//		want auth.TokenReviewResponse
//	}{
//		{
//			name: "TokenReviewFailure",
//			args: args{
//				review: auth.TokenReview{
//					TypeMeta: metav1.TypeMeta{
//						Kind:       "TokenReview",
//						APIVersion: "authentication.k8s.io/v1beta1",
//					},
//					Spec:     auth.TokenReviewSpec{
//						Token:     "123456",
//						Audiences: []string{},
//					},
//				},
//			},
//			want: auth.TokenReviewResponse{
//				TypeMeta: metav1.TypeMeta{
//					Kind:       "TokenReview",
//					APIVersion: "authentication.k8s.io/v1beta1",
//				},
//				Status:   auth.TokenReviewStatus{
//					Authenticated: false,
//					User:          auth.UserInfo{},
//					Audiences:     []string{},
//					Error:         "",
//				},
//			},
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			wo := WorkChatImpl{}
//			if got := wo.TokenReviewFailure(tt.args.review); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("TokenReviewFailure() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestWorkChatImpl_TokenReviewSuccess(t *testing.T) {
//	type args struct {
//		review auth.TokenReview
//	}
//	tests := []struct {
//		name                string
//		args                args
//		wantSuccessResponse auth.TokenReviewResponse
//	}{
//		{
//			name: "TokenReviewSuccess",
//			args: args{
//				review: auth.TokenReview{
//					TypeMeta: metav1.TypeMeta{
//						Kind:       "TokenReview",
//						APIVersion: "authentication.k8s.io/v1beta1",
//					},
//					Spec:     auth.TokenReviewSpec{
//						Token:     "123456",
//						Audiences: []string{},
//					},
//				},
//			},
//			wantSuccessResponse: auth.TokenReviewResponse{
//				TypeMeta: metav1.TypeMeta{
//					Kind:       "TokenReview",
//					APIVersion: "authentication.k8s.io/v1beta1",
//				},
//				Status:   auth.TokenReviewStatus{
//					Authenticated: true,
//					User:          auth.UserInfo{
//						Username: "chaoyang",
//						UID:      "0",
//						Groups:   []string{"dev"},
//						Extra:    map[string]auth.ExtraValue{},
//					},
//					Audiences:     []string{},
//					Error:         "",
//				},
//			},
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			wo := WorkChatImpl{}
//			if gotSuccessResponse := wo.TokenReviewSuccess(tt.args.review); !reflect.DeepEqual(gotSuccessResponse, tt.wantSuccessResponse) {
//				t.Errorf("TokenReviewSuccess() = %v, want %v", gotSuccessResponse, tt.wantSuccessResponse)
//			}
//		})
//	}
//}
//
//func TestWorkChatImpl_GetReadMember(t *testing.T) {
//	type fields struct {
//		AccessTokenMap map[string]string
//	}
//	type args struct {
//		token string
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//		want   bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			w := WorkChatImpl{
//				AccessTokenMap: tt.fields.AccessTokenMap,
//			}
//			if got := w.GetReadMember(tt.args.token); got != tt.want {
//				t.Errorf("GetReadMember() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}