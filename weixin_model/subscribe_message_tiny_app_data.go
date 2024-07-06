package weixin_model

/*
例如：模板的内容为：
	姓名: {{name01.DATA}}
	金额: {{amount01.DATA}}
	行程: {{thing01.DATA}}
	日期: {{date01.DATA}}

则发送订阅消息对应的json为：
{
  "touser": "OPENID",
  "template_id": "TEMPLATE_ID",
  "page": "index",
  "data": {
      "name01": {
          "value": "某某"
      },
      "amount01": {
          "value": "￥100"
      },
      "thing01": {
          "value": "广州至北京"
      } ,
      "date01": {
          "value": "2018-01-01"
      }
  }
}
*/

type DataValue struct {
	Value string `json:"value"`
}

// CouponUseMessage 优惠券使用提醒
//type CouponUseMessage struct {
//	Thing1 struct {
//		Value string `json:"value"`
//	} `json:"thing1"`
//
//	Thing4 struct {
//		Value string `json:"value"`
//	} `json:"thing4"`
//
//	Time3 struct {
//		Value string `json:"value"`
//	} `json:"time3"`
//}

// ExampleMessageData 订阅消息模板实例内容
type ExampleMessageData struct {
	Number01 DataValue `json:"number01"`
	Date01   DataValue `json:"date01"`
	Site01   DataValue `json:"site01"`
	Site02   DataValue `json:"site02"`
}

type CouponUseMessage struct {
	Thing1 DataValue `json:"thing1"`

	Thing4 DataValue `json:"thing4"`

	Time3 DataValue `json:"time3"`
}
