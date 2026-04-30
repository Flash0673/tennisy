package token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNewJWTService(t *testing.T) {
	t.Parallel()

	type args struct {
		userID, email string
	}

	type want struct {
		userID string
		err    error
	}

	tCases := map[string]struct {
		args args
		want want
	}{
		"base": {
			args: args{
				userID: "user1",
				email:  "email1",
			},
			want: want{
				userID: "user1",
			},
		},
	}
	for name, tCase := range tCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			svc := NewJWTService("", time.Hour)
			pair, err := svc.Generate(tCase.want.userID, tCase.args.email)
			require.NoError(t, err)
			resp, err := svc.Parse(pair.AccessToken)
			require.NoError(t, err)
			require.Equal(t, tCase.want.userID, resp)
		})
	}
}
