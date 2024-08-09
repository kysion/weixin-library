package app_manager

import "github.com/kysion/base-library/utility/enum"

// AppTypeEnum 应用类型：1公众号 2小程序 4网站应用H5  8移动应用  16视频小店
type AppTypeEnum enum.IEnumCode[int]

type appType struct {
	PublicAccount AppTypeEnum
	TinyApp       AppTypeEnum
	H5            AppTypeEnum
	App           AppTypeEnum
	VideoStore    AppTypeEnum
}

var AppType = appType{
	PublicAccount: enum.New[AppTypeEnum](1, "公众号"),
	TinyApp:       enum.New[AppTypeEnum](2, "小程序"),
	H5:            enum.New[AppTypeEnum](4, "网站应用H5"),
	App:           enum.New[AppTypeEnum](8, "移动应用"),
	VideoStore:    enum.New[AppTypeEnum](16, "视频小店"),
}

func (e appType) New(code int, description string) AppTypeEnum {
	if (code & AppType.PublicAccount.Code()) == AppType.PublicAccount.Code() {
		return e.PublicAccount
	}
	if (code & AppType.TinyApp.Code()) == AppType.TinyApp.Code() {
		return e.TinyApp
	}
	if (code & AppType.H5.Code()) == AppType.H5.Code() {
		return e.H5
	}
	if (code & AppType.App.Code()) == AppType.App.Code() {
		return e.App
	}
	if (code & AppType.VideoStore.Code()) == AppType.VideoStore.Code() {
		return e.VideoStore
	}
	return enum.New[AppTypeEnum](code, description)
}
