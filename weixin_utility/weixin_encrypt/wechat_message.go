package weixin_encrypt

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/kysion/weixin-library/weixin_model"
)

type msgCrypt struct {
	token  string
	aesKey string
	appid  string
}

// EventEncryptRequest 微信事件推送结构体
type EventEncryptRequest struct {
	XMLName xml.Name `xml:"xml"`
	Encrypt string   `xml:"Encrypt"`
	Appid   string   `xml:"Appid"`
}

// MessageEncryptRequest 微信消息加密结构体
type MessageEncryptRequest struct {
	XMLName      xml.Name `xml:"xml"`
	Encrypt      string   `xml:"Encrypt"`
	MsgSignature string   `xml:"MsgSignature"`
	TimeStamp    string   `xml:"TimeStamp"`
	Nonce        string   `xml:"Nonce"`
}

// NewWechatMsgCrypt 实例化微信加解密
func NewWechatMsgCrypt(token string, aesKey string, appid string) *msgCrypt {
	instance := new(msgCrypt)
	instance.token = token
	instance.aesKey = aesKey
	instance.appid = appid
	return instance
}

// WechatEventDecrypt 微信事件推送解密
func (w *msgCrypt) WechatEventDecrypt(eventRequest weixin_model.EventEncryptMsgReq, msgSignature string, timestamp, nonce string) interface{} {
	errCode, data := w.decryptMsg(msgSignature, timestamp, nonce, eventRequest.Encrypt)
	if errCode != SuccessCode {
		panic(fmt.Sprintf("消息解密失败，code：%d", errCode))
	}
	message := EventMessageBody{}
	_, err := FormatMessage(data, &message)
	if err != nil {
		panic(fmt.Sprintf("消息格式化失败，%s", err.Error()))
	}
	return message
}

// WechatMessageDecrypt 微信消息解密
func (w *msgCrypt) WechatMessageDecrypt(messageEncryptRequest MessageEncryptRequest) interface{} {
	errCode, data := w.decryptMsg(messageEncryptRequest.MsgSignature, messageEncryptRequest.TimeStamp, messageEncryptRequest.Nonce, messageEncryptRequest.Encrypt)
	if errCode != SuccessCode {
		panic(fmt.Sprintf("消息解密失败，code：%d", errCode))
	}
	message := MessageBodyDecrypt{}
	_, err := FormatMessage(data, &message)
	if err != nil {
		panic(fmt.Sprintf("消息格式化失败，%s", err.Error()))
	}
	return message
}

// WechatTextMessage 微信文本消息加密
func (w *msgCrypt) WechatTextMessage(fromUserName string, toUserName string, content string) MessageEncryptBody {
	message := FormatTextMessage(fromUserName, toUserName, content)
	encrypted, err := w.encryptMsg(message)
	if err != nil {
		panic("消息加密失败：" + err.Error())
	}
	data := FormatEncryptData(encrypted, w.token)
	return data
}

// WechatImageMessage 微信图片消息加密
func (w *msgCrypt) WechatImageMessage(fromUserName string, toUserName string, mediaId string) MessageEncryptBody {
	message := FormatImageMessage(fromUserName, toUserName, mediaId)
	encrypted, err := w.encryptMsg(message)
	if err != nil {
		panic("消息加密失败：" + err.Error())
	}
	data := FormatEncryptData(encrypted, w.token)
	return data
}

// WechatVoiceMessage 微信语音消息加密
func (w *msgCrypt) WechatVoiceMessage(fromUserName string, toUserName string, mediaId string) MessageEncryptBody {
	message := FormatVoiceMessage(fromUserName, toUserName, mediaId)
	encrypted, err := w.encryptMsg(message)
	if err != nil {
		panic("消息加密失败：" + err.Error())
	}
	data := FormatEncryptData(encrypted, w.token)
	return data
}

// WechatVideoMessage 微信视频消息加密
func (w *msgCrypt) WechatVideoMessage(fromUserName string, toUserName string, mediaId string, title string, description string) MessageEncryptBody {
	message := FormatVideoMessage(fromUserName, toUserName, mediaId, title, description)
	encrypted, err := w.encryptMsg(message)
	if err != nil {
		panic("消息加密失败：" + err.Error())
	}
	data := FormatEncryptData(encrypted, w.token)
	return data
}

// WechatMusicMessage 微信音乐消息加密
func (w *msgCrypt) WechatMusicMessage(fromUserName string, toUserName string, thumbMediaId string, title string, description string, musicUrl string, hQMusicUrl string) MessageEncryptBody {
	message := FormatMusicMessage(fromUserName, toUserName, thumbMediaId, title, description, musicUrl, hQMusicUrl)
	encrypted, err := w.encryptMsg(message)
	if err != nil {
		panic("消息加密失败：" + err.Error())
	}
	data := FormatEncryptData(encrypted, w.token)
	return data
}

// WechatArticlesMessage 微信图文消息加密
func (w *msgCrypt) WechatArticlesMessage(fromUserName string, toUserName string, items []ArticleItem) MessageEncryptBody {
	message := FormatArticlesMessage(fromUserName, toUserName, items)
	encrypted, err := w.encryptMsg(message)
	if err != nil {
		panic("消息加密失败：" + err.Error())
	}
	data := FormatEncryptData(encrypted, w.token)
	return data
}

// decryptMsg aes消息解密
func (w *msgCrypt) decryptMsg(msgSignature string, timestamp, nonce string, encrypted string) (int, []byte) {
	//验证aes
	if len(w.aesKey) != 43 {
		return IllegalAesKey, nil
	}
	//判断签名是否一致
	if err := w.validSignature(msgSignature, timestamp, nonce, encrypted); err != nil {
		return ValidateSignatureError, nil
	}
	//解密
	prp := NewPrpCrypt(w.aesKey)
	plainText, err := prp.decrypt(encrypted)
	if err != nil {
		return DecryptAESError, nil
	}
	//验证appid是否一致（消息来源是否一致）
	if err := w.validMessageSource(plainText); err != nil {
		return ValidateAppidError, nil
	}
	return SuccessCode, plainText
}

// encryptMsg aes消息加密
func (w *msgCrypt) encryptMsg(message []byte) (string, error) {
	//计算消息长度
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, int32(len(message)))
	if err != nil {
		return "", err
	}
	messageLength := buf.Bytes()
	//生成随机字符串
	randBytes := []byte(makeRandomString(16))
	plainData := bytes.Join([][]byte{randBytes, messageLength, message, []byte(w.appid)}, nil)
	prp := NewPrpCrypt(w.aesKey)
	//消息加密
	encrypted, err := prp.encrypt(plainData)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(encrypted), nil
}

// validSignature 验证签名是否一致
func (w *msgCrypt) validSignature(msgSignature string, timestamp, nonce string, encrypted string) error {
	validSignature := GetSignature(timestamp, nonce, encrypted, w.token)
	if validSignature != msgSignature {
		return errors.New("签名不一致：valid sign error")
	}
	return nil
}

// validMessageSource 验证消息来源
func (w *msgCrypt) validMessageSource(plainText []byte) error {
	messageLength := GetMessageLength(plainText)
	//获取appid位置
	appIdStartPos := 20 + messageLength
	id := plainText[appIdStartPos : int(appIdStartPos)+len(w.appid)]
	if string(id) != w.appid {
		return errors.New("消息来源不一致：Appid is invalid")
	}
	return nil
}
