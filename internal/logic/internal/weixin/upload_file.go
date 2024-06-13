package weixin

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/kysion/weixin-library/weixin_model"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

const (
	WechatUploadMediaAPI = "https://api.weixin.qq.com/wxa/uploadmedia"
)

// UploadMedia 上传审核所需素材
func UploadMedia(ctx context.Context, token string, imagePath string) (*weixin_model.UploadAppMediaToAuditRes, error) {
	file, err := os.Open(imagePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("image", filepath.Base(imagePath))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", WechatUploadMediaAPI, body)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	urlQuery := req.URL.Query()

	if err != nil {
		return nil, err
	}
	urlQuery.Add("access_token", token)

	req.URL.RawQuery = urlQuery.Encode()
	fmt.Println(req.URL)

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	jsonbody, _ := ioutil.ReadAll(res.Body)
	media := weixin_model.UploadAppMediaToAuditRes{}
	err = json.Unmarshal(jsonbody, &media)
	if err != nil {
		return nil, err
	}

	if media.Mediaid == "" {
		err = sys_service.SysLogs().ErrorSimple(ctx, err, "素材上传失败！", "WeiXin-App-Version-Manager")
	}

	return &media, err
}
