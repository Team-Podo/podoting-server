package auth

import (
	"context"
	"fmt"
	"testing"
)

func TestSearchUser(t *testing.T) {
	user, err := Firebase.SearchUser("85HdQy2YDfb0FHDDhD5VnOc2xzI2")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(user)
}

func TestCreateToken(t *testing.T) {
	token, err := Firebase.CreateCustomToken("85HdQy2YDfb0FHDDhD5VnOc2xzI2")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(token)
}

func TestVerifyToken(t *testing.T) {
	verifiedToken, err := Firebase.VerifyToken("eyJhbGciOiJSUzI1NiIsImtpZCI6Ijk5NjJmMDRmZWVkOTU0NWNlMjEzNGFiNTRjZWVmNTgxYWYyNGJhZmYiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJodHRwczovL3NlY3VyZXRva2VuLmdvb2dsZS5jb20vcG9kb3RpbmctMmMzYjAiLCJhdWQiOiJwb2RvdGluZy0yYzNiMCIsImF1dGhfdGltZSI6MTY2NTMwMzA4NywidXNlcl9pZCI6InVUOUhyNWVUWnNTdjVRNmxCNXg4OG52aWd3ZjIiLCJzdWIiOiJ1VDlIcjVlVFpzU3Y1UTZsQjV4ODhudmlnd2YyIiwiaWF0IjoxNjY1MzAzMDg3LCJleHAiOjE2NjUzMDY2ODcsImVtYWlsIjoicG9kb0B0ZXN0LmNvbSIsImVtYWlsX3ZlcmlmaWVkIjpmYWxzZSwiZmlyZWJhc2UiOnsiaWRlbnRpdGllcyI6eyJlbWFpbCI6WyJwb2RvQHRlc3QuY29tIl19LCJzaWduX2luX3Byb3ZpZGVyIjoicGFzc3dvcmQifX0.LQDcF5fpa5ZXvnygFxfPLJIsef_fXvb3HF7E0R4THh0BAAXJO8BZ3yxEyZ5QU1eAjPaLut4a_MT8xa2yhq2A1_LIJy5oR-OvfctnsYTLzG6ZPavz6-MfLMT1Nl1FtoMkfP2nmCoy6-8yXTV7NeyCunMeXuInnU5LilHLaGmz5h-CTJ2HV5atovhvNuilCByWcbPxjKzq2LnPZJmWJXd1Zz3cZJa0LqnjcPvvF9K_FDIoTU30iXuyYJ28GSpy0UOoM57m5UATe-ZDat67rp19fLvWir68HSdy1u4z8_lef5ksyRFBWNCzk97G2AH48GAb1d-r-3uzRrxBAAaPhywjXw", context.Background())
	if err != nil {
		t.Error(err)
	}

	fmt.Println(verifiedToken.UID)
}
