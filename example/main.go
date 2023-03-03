package main

import (
	"fmt"
	_ "github.com/SupenBysz/gf-admin-community"
	"github.com/kysion/kys-weixin-library/utility"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/kysion/kys-weixin-library/example/internal/boot"
	_ "github.com/kysion/kys-weixin-library/internal/logic"
)

func main() {
	// 消息加密解密ket， 参数是token
	fmt.Println(utility.Md5Hash("comjditcokuaimk")) // EC68C5852D342DD42A7F64071E6C5EE9  EC68C5852D342DD42A7F64071E6C5EE9E0123456789
	boot.Main.Run(gctx.New())
}

// com.jditco.kuaimk
// douyin.jditco.com/douyin/gateway.services
// comjditcokuaimk
// EC68C5852D342DD42A7F64071E6C5EE901234567890

//-----BEGIN PUBLIC KEY-----
//MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAi66Zj1NBbLP7amr38Q0X
//GNOZ9tYPC2ta5roYTkJmfA6uspVAzV3VE3vKVZ7ROJDCFWKrzly8K2lDX1/CypOe
//lqLDVAUF8BV966CRCZYOBH23+fmxIdP0GNZ+KyeDXXZKRliAi9dWRo8Mhsw9XlWx
//K5jkGv1Mu8fgXdmZPrtC7TkeK2jq0x9NLtO6q9UR3D+wdX7FQdD13fI76n95RSna
//2t797nZzrZBh7FCtXWfp+ObWQKQf/qfPe93v6uPwfHKxTep5QeiLAREb6JtuTaB5
//PwDLE/LAxMQnjpG/fNwTTtaM3cEOvmngXkkYCj0j8XDFX+Gx3n1rirnaw1IrHhnL
//hwIDAQAB
//-----END PUBLIC KEY-----
