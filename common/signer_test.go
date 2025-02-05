package common

import (
	"fmt"
	"github.com/MoyunRz/bitget-sdk/utils"
	"testing"
)

func TestSigner_Sign(t *testing.T) {
	signer := new(Signer)
	result := signer.Sign("GET", "www.bitget.com", "aaaa", utils.TimesStamp())
	fmt.Print(result)
}
