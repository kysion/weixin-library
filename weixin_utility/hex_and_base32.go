package weixin_utility

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
)

// HexToBase32 十六进制转为32进制
func HexToBase32(hexStr string) (base32 string) {
	//hexStr := "caf4b7b8d6620f00"

	// 解码为字节数组
	bytes, err := hex.DecodeString(hexStr)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 转换为 uint64 数值
	var num uint64
	for i := range bytes {
		num = (num << 8) | uint64(bytes[i])
	}

	// 将 uint64 格式化为 32 进制字符串
	base32Str := strconv.FormatUint(num, 32)

	fmt.Println("32 进制字符串：", base32Str)

	return base32Str
}

// Base32ToHex 32进制转为十六进制
func Base32ToHex(base32Str string) (hexStr string) {
	// 解析为 uint64 数值
	n, _ := new(big.Int).SetString(base32Str, 32)

	bytes := n.Bytes()

	// 将数值格式化为16进制字符串

	// 转换为 16 进制字符串
	hexStr = hex.EncodeToString(bytes)

	fmt.Println("16 进制字符串：", hexStr)

	return hexStr
}

// 正常获取订单，
