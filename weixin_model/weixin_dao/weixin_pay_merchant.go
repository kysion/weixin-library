// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package weixin_dao

import (
	"github.com/kysion/base-library/utility/daoctl/dao_interface"
	"github.com/kysion/weixin-library/weixin_model/weixin_dao/internal"
)

type WeixinPayMerchantDao = dao_interface.TIDao[internal.WeixinPayMerchantColumns]

func NewWeixinPayMerchant(dao ...dao_interface.IDao) WeixinPayMerchantDao {
	return (WeixinPayMerchantDao)(internal.NewWeixinPayMerchantDao(dao...))
}

var (
	WeixinPayMerchant = NewWeixinPayMerchant()
)
