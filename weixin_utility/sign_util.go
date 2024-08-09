package weixin_utility

import (
	"crypto/sha1"
	"fmt"
	"sort"
	"strings"
)

// CheckSignature 微信公众号签名检查
func CheckSignature(signature, timestamp, nonce, token string) bool {
	arr := []string{timestamp, nonce, token}
	// 字典序排序
	sort.Strings(arr)

	n := len(timestamp) + len(nonce) + len(token)
	var b strings.Builder
	b.Grow(n)
	for i := 0; i < len(arr); i++ {
		b.WriteString(arr[i])
	}

	fmt.Println("签名：" + signature)
	fmt.Println("加密后: " + b.String())
	// 进行Sha1编码
	return Sha1(b.String()) == signature
}

// VerifyByteDanceServer 验签
func VerifyByteDanceServer(tpToken string, timestamp string, nonce string, encrypt string, msgSignature string) bool {
	values := []string{tpToken, timestamp, nonce, encrypt}
	sort.Strings(values)
	newMsgSignature := Sha1(strings.Join(values, ""))

	if newMsgSignature == msgSignature {
		return true
	}

	return false
}

// Sha1 进行Sha1编码
func Sha1(str string) string {
	h := sha1.New()
	h.Write([]byte(str))
	encodeStr := fmt.Sprintf("%x", h.Sum(nil))
	return encodeStr
}
