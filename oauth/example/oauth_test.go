package oauth

import (
	"errors"
	"testing"

	"github.com/SituChengxiang/WeJH-SDK/oauth"
	"github.com/SituChengxiang/WeJH-SDK/oauth/oauthException"
)

func TestLogin(t *testing.T) {
	for i, tc := range []struct {
		inputUsername string
		inputPassword string
		expect        error
	}{
		{
			inputUsername: "302024315300",
			inputPassword: "i m password",
			expect:        oauthException.NotActivatedError,
		},
		{
			inputUsername: "wrong username",
			inputPassword: "this password is correct",
			expect:        oauthException.WrongAccount,
		},
		{
			inputUsername: "MangoGovo",
			inputPassword: "wrong password",
			expect:        oauthException.WrongPassword,
		},
	} {
		cookies, userInfo, err := oauth.GetUserInfo(tc.inputUsername, tc.inputPassword)
		if err != nil || len(cookies) == 0 {
			t.Errorf("测试点 %d (%s,%s): 遇到了异常: %v", i+1, tc.inputUsername, tc.inputPassword, err)
			continue
		}

		if !errors.Is(err, tc.expect) {
			t.Errorf("测试点 %d (%s,%s): 期望值:=%s 实际值=%s", i+1, tc.inputUsername, tc.inputPassword, tc.expect, err)
			continue
		}
		t.Logf("测试点 %d %s %s", i+1, cookies, userInfo)
	}
}
