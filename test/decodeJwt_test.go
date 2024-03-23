package test

import (
	"reflect"
	"testing"

	"x-clone.com/backend/src/utils/validator"
)

func TestDecodeJwt(t *testing.T) {
	type args struct {
		jwtToken string
	}
	tests := []struct {
		name    string
		args    args
		want    *validator.DecodedJwtVal
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				jwtToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImthdHRlbHNhdXJhdjMyQGdtYWlsLmNvbSIsInVzZXJJZCI6Ijk2Mjc0MTEwLWY4NGMtNDZhYi04MGRiLWZlNmZiNDYyM2ViYiJ9.fC0Uc284w2V7I-6K9oo5UvKIr2Phedwm2vRvBN-vFRc",
			},
			want: &validator.DecodedJwtVal{
				Email:  "kattelsaurav32@gmail.com",
				UserId: "96274110-f84c-46ab-80db-fe6fb4623ebb",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := validator.ValidateJwt(tt.args.jwtToken)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeJwt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodeJwt() = %v, want %v", got, tt.want)
			}
		})
	}
}
