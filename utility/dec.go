package utility

//
//import (
//	"bytes"
//	"crypto/aes"
//	"crypto/cipher"
//	"encoding/base64"
//	"encoding/binary"
//	"encoding/xml"
//	"errors"
//	"fmt"
//	"github.com/gogf/gf/v2/util/gconv"
//	"github.com/kysion/merchant-library/weixin_consts"
//	"github.com/kysion/merchant-library/weixin_model"
//	"math/rand"
//)
//
//var SuccessCode = 0
//var ValidateSignatureError = -40001
//var ParseXmlError = -40002
//var ComputeSignatureError = -40003
//var IllegalAesKey = -40004
//var ValidateAppidError = -40005
//var EncryptAESError = -40006
//var DecryptAESError = -40007
//var IllegalBuffer = -40008
//
//// WechatEventDecrypt 微信事件推送解密
//func WechatEventDecrypt(eventRequest weixin_model.EventEncryptMsgReq, msgSignature string, timestamp, nonce string) interface{} {
//	errCode, data := decryptMsg(msgSignature, timestamp, nonce, eventRequest.Encrypt)
//	if errCode != SuccessCode {
//		panic(fmt.Sprintf("消息解密失败，code：%d", errCode))
//	}
//	message := weixin_model.EventMessageBody{}
//	_, err := FormatMessage(data, &message)
//	if err != nil {
//		panic(fmt.Sprintf("消息格式化失败，%s", err.Error()))
//	}
//	return message
//}
//
//func GetMessageLength(plainText []byte) int32 {
//	// Read length
//	buf := bytes.NewBuffer(plainText[16:20])
//	var length int32
//	err := binary.Read(buf, binary.BigEndian, &length)
//	if err != nil {
//		panic("获取消息长度失败：read message length error")
//	}
//	return length
//}
//
//func FormatMessage(plainText []byte, data interface{}) (*interface{}, error) {
//	length := GetMessageLength(plainText)
//	err := xml.Unmarshal(plainText[20:20+length], data)
//	if err != nil {
//		return nil, errors.New("格式化消息失败：format message error")
//	}
//	return &data, nil
//}
//
//// WechatMessageDecrypt 微信消息解密
//func WechatMessageDecrypt(messageEncryptRequest weixin_model.MessageEncryptReq) interface{} {
//	errCode, data := decryptMsg(messageEncryptRequest.MsgSignature, messageEncryptRequest.TimeStamp, messageEncryptRequest.Nonce, messageEncryptRequest.Encrypt)
//	if errCode != SuccessCode {
//		panic(fmt.Sprintf("消息解密失败，code：%d", errCode))
//	}
//	message := weixin_model.EventMessageBody{}
//	_, err := FormatMessage(data, &message)
//	if err != nil {
//		panic(fmt.Sprintf("消息格式化失败，%s", err.Error()))
//	}
//	return message
//}
//
//// decryptMsg aes消息解密
//func decryptMsg(msgSignature string, timestamp, nonce string, encrypted string) (int, []byte) {
//	//验证aes
//	if len(weixin_consts.Global.DecryptKey) != 43 {
//		return IllegalAesKey, nil
//	}
//	//判断签名是否一致
//	if err := validSignature(msgSignature, timestamp, nonce, encrypted); err != nil {
//		return ValidateSignatureError, nil
//	}
//	//解密
//	prp := NewPrpCrypt(weixin_consts.Global.DecryptKey)
//	plainText, err := prp.decrypt(encrypted)
//	if err != nil {
//		return DecryptAESError, nil
//	}
//	//验证appid是否一致（消息来源是否一致）
//	if err := w.validMessageSource(plainText); err != nil {
//		return ValidateAppidError, nil
//	}
//	return SuccessCode, plainText
//}
//
//var prpKey = gconv.Bytes("PKCS#7")
//
//func decrypt(encrypted string) ([]byte, error) {
//	encryptedBytes, _ := base64.StdEncoding.DecodeString(encrypted)
//	k := len(prpKey) //PKCS#7
//	if len(encryptedBytes)%k != 0 {
//		panic("ciphertext size is not multiple of aes key length")
//	}
//	block, err := aes.NewCipher(prpKey)
//	if err != nil {
//		return nil, err
//	}
//	blockMode := cipher.NewCBCDecrypter(block, prp.Iv)
//	plainText := make([]byte, len(encryptedBytes))
//	blockMode.CryptBlocks(plainText, encryptedBytes)
//	return plainText, nil
//}
//
//func NewPrpCrypt(aesKey string) *prpCrypt {
//	instance := new(prpCrypt)
//	//网络字节序
//	instance.Key, _ = base64.StdEncoding.DecodeString(aesKey + "=")
//	instance.Iv = randomIv()
//	return instance
//}
//
//// encryptMsg aes消息加密
//func encryptMsg(message []byte) (string, error) {
//	//计算消息长度
//	buf := new(bytes.Buffer)
//	err := binary.Write(buf, binary.BigEndian, int32(len(message)))
//	if err != nil {
//		return "", err
//	}
//	messageLength := buf.Bytes()
//	//生成随机字符串
//	randBytes := []byte(makeRandomString(16))
//	plainData := bytes.Join([][]byte{randBytes, messageLength, message, []byte(w.appid)}, nil)
//	prp := NewPrpCrypt(w.aesKey)
//	//消息加密
//	encrypted, err := prp.weixin_encrypt(plainData)
//	if err != nil {
//		return "", err
//	}
//	return base64.StdEncoding.EncodeToString(encrypted), nil
//}
//
//func makeRandomString(length int) string {
//	randStr := ""
//	strSource := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyl"
//	maxLength := len(strSource) - 1
//	for i := 0; i < length; i++ {
//		randomNum := rand.Intn(maxLength)
//		randStr += strSource[randomNum : randomNum+1]
//	}
//	return randStr
//}
//
//// validSignature 验证签名是否一致
//func validSignature(msgSignature string, timestamp, nonce string, encrypted string) error {
//	validSignature := GetSignature(timestamp, nonce, encrypted, w.token)
//	if validSignature != msgSignature {
//		return errors.New("签名不一致：valid sign error")
//	}
//	return nil
//}
//
//// validMessageSource 验证消息来源
//func validMessageSource(plainText []byte) error {
//	messageLength := GetMessageLength(plainText)
//	//获取appid位置
//	appIdStartPos := 20 + messageLength
//	id := plainText[appIdStartPos : int(appIdStartPos)+len(w.appid)]
//	if string(id) != w.appid {
//		return errors.New("消息来源不一致：Appid is invalid")
//	}
//	return nil
//}
