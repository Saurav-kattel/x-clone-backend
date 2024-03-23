package test

import (
	"testing"

	"x-clone.com/backend/src/utils/validator"
)

func TestHashValidator(t *testing.T) {
	type args struct {
		hash     string
		password string
	}
	tt := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test for true",
			args: args{
				hash:     "1f33b146637c260aa09787f5b5357ecc976f6700cb55f5bbbac6b7884ccfd86b",
				password: "saurav",
			},
			want: true,
		}, {
			name: "test for false",
			args: args{
				hash:     "1f33b146637c260aa09787f5b5357ecc976f6700cb55f5bbbac6b7884ccfd86b",
				password: "raaam",
			},
			want: false,
		},
	}

	for _, r := range tt {
		t.Run(r.name, func(t *testing.T) {
			if res := validator.HashValidator(r.args.hash, r.args.password); res != r.want {
				t.Errorf("failed for test %v want=%v got%v", r.name, r.want, res)
			}
		})

	}
}
