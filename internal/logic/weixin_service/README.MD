# WeixinService

- 该文件夹中主要存放 weixin-library库 的相关SQL数据表的逻辑实现。。。
- 此文件夹主要是针对SQL数据表的实现逻辑，其他的文件，不要放到该目录下


## 一、SQL模型

### 注意事项
- 该文档主要是 weixin-library库 的相关SQL数据表的记录。。。
- 注意：id 全部是主键

### 微信用户表
表名：weixin_consumer_config 

| 字段 | 类型 | 约束       | 描述 |
| ---- | ---- | ---------- | ---- |
| id |  int8  | 主键、非空 | id |
| open_id |  varchar(255)  | 非空| 微信用户openId，不同应用下的用户具备不同的openId |
| sys_user_id |  int8  | | 用户id |
| avatar |  text  | | 头像 |
| province |  varchar(64)  | | 省份 |
| city |  varchar(64)  | | 城市 |
| nick_name |  varchar(64)  | | 昵称 |
| is_student_certified |  int2 | | 是否学生认证 |
| user_type |  int2 | | 用户账号类型，和sysUserType保持一致 |
| user_state |  int2 | | 状态：0未激活、1正常、-1封号、-2异常、-3已注销 |
| is_certified |  int2 | | 是否实名认证 |
| sex |  int2 | | 性别：0未知、1男、2女 |
| access_token |  text  | | 授权token |
| ext_json |  json | | 拓展字段 |
| created_at |  timestamptz(6) | | 微信用户union_id，同一个开放平台帐号下的用户只有一个unionId |
| updated_at |  timestamptz(6) | | 微信用户会话key |
| deleted_at |  timestamptz(6) | | 用户授权刷新令牌 |
| union_id |  varchar(255)  | | 令牌过期时间 |
| session_key |  varchar(255)  | | 微信用户授权状态：1已授权、2未授权 |
| refresh_token |  text  | | 应用类型：1公众号 2小程序 4网站应用H5  8移动应用  16视频小店 |
| expires_in |  timestamptz(6) | | 是否关注公众号：1关注、2未关注 |
| auth_state |  int2 | | 商家应用Id |
| app_type |  int2 | |  |
| is_follow_public |  int2 | |  |
| app_id|  varchar(64)  | |  |

### 微信第三方平台应用配置表
表名： weixin_third_app_config

| 字段 | 类型 | 约束       | 描述 |
| ---- | ---- | ---------- | ---- | 
| id | int8 |  NOT NULL | 服务商id  |
| name | varchar(64) |  NOT NULL | 服务商name |  
| app_id | varchar(64) |  NOT NULL | 服务商应用Id |  
| app_name | varchar(64) |  NOT NULL | 服务商应用名称 |  
| app_type | int2 | | 服务商应用类型 |   
| app_auth_token | varchar(255) | | 服务商应用授权token |   
| expires_in | timestamptz(6) | | Token过期时间 |   
| re_expires_in | timestamptz(6) | | Token限期刷新时间 |   
| union_main_id | int8 | | 关联主体id |   
| sys_user_id | int8 | | 用户id |   
| ext_json | json | | 拓展字段 |   
| app_gateway_url | text |  NOT NULL | 网关地址 | 
| app_callback_url | text |  NOT NULL | 回调地址 | 
| app_secret | varchar(255) |  NOT NULL | 服务商应用密钥 | 
| msg_verfiy_token | text | | 消息校验Token |   
| msg_encrypt_key | text | | 消息加密解密密钥 |   
| auth_init_url | text | | 授权发起页域名 |   
| server_domain | text | | 服务器域名 |   
| business_domain | text | | 业务域名 |   
| auth_test_appIds | text | | 授权测试应用列表 |   
| platform_site | text | | 平台官方 |   
| logo | text | | 服务商logo |   
| state | int2 | | 状态：0禁用 1启用 |   
| release_state | int2 | | 发布状态：0未发布 1已发布 |   
| https_cert | text | | 域名证书 |   
| https_key | text | | 域名私钥 |   
| updated_at | timestamptz(6) | | |    
| app_id_md5 | varchar(32) | | 应用id加密md5后的结果 |   
| user_id | int8 | | 应用所属账号 |   
| refresh_token | varchar(255) | | 刷新应用Token |    



### 商户应用配置表
表名：weixin_merchant_app_config

| 字段 | 类型 | 约束       | 描述 |
| ---- | ---- | ---------- | ---- |
| id | int8 | NOT NULL | ID |
| name | varchar(64) | NOT NULL | 商家id |
| app_id | varchar(64) | NOT NULL | 商家name |
| app_name | varchar(64) | NOT NULL | 商家应用Id |
| app_type | int2 | ---------- | 商家应用名称 |
| app_auth_token | varchar(255) | ---------- | 应用类型：1公众号 2小程序 4网站应用H5  8移动应用  16视频小店 |
| is_full_proxy | int2 | NOT NULL | 商家应用token：1、当第三方代开发的时候，这个是商家授权的应用token (authorizer_access_token)；2、当是商家自研模式时，这个是商家的应用token(access_token)。 |
| state | int2 | NOT NULL, | 是否全权委托待开发：0否 1是 |
| expires_in | timestamptz(6) |   | 状态： 0禁用 1启用 |
| re_expires_in | timestamptz(6) |   | Token过期时间 |
| user_id | int8 |   | Token限期刷新时间 |
| union_main_id | int8 |   | 应用所属账号 |
| sys_user_id | int8 |   | 关联主体id |
| ext_json | json |   | 用户id |
| app_gateway_url | text |   | 拓展字段 |
| app_callback_url | text |   | 网关地址 |
| app_secret | varchar(255) |   | 回调地址 |
| msg_verfiy_token | text |   | 商家应用密钥 |
| msg_encrypt_key | text |   | 消息校验Token |
| msg_encrypt_type | int2 |   | 消息加密解密密钥（EncodingAESKey） |
| business_domain | text |   | 消息加密模式：1兼容模式 2明文模式 4安全模式 |
| js_domain | text |   | 业务域名 |
| auth_domain | varchar(255) |   | JS接口安全域名 |
| logo | text |   | 网页授权域名 |
| https_cert | text |   | 商家logo |
| https_key | text |   | 域名证书 |
| server_domain | text |   | 域名私钥 |
| app_id_md5 | varchar(32) |   | 服务器域名 |
| third_app_id | varchar(64) |   | 应用id加密md5后的结果 |
| notify_url | text |   | 服务商appId |
| server_rate | float4 |   | 异步通知地址，允许业务层追加相关参数 |
| union_main_type | varchar(255) |   | 手续费比例，默认0.6% |
| version | varchar(16) |   | 应用关联主体类型，和user_type保持一致 |
| privacy_policy | text |   | 应用版本 |
| user_policy | text |   | 隐私协议 |
| dev_state | int2 |   | 用户协议 |
| updated_at | timestamptz(6) |   | 开发状态：0未上线 1已上线 |
| refresh_token | varchar(255) |   | 刷新商家授权应用Token |
| primitive_id | varchar(64) |   | 应用原始ID |


### 微信支付商户表 
表名：weixin_pay_merchant

| 字段 | 类型 | 约束       | 描述 |
| ---- | ---- | ---------- | ---- |
| id | int8 |  NOT NULL| ID |
| mchid | int4 |  NOT NULL| 微信支付商户号 |
| merchant_name | varchar(128) |  NOT NULL| 商户号公司名称 |
| merchant_short_name | varchar(255) |  | 商户号简称 |
| merchant_type | int2 |  NOT NULL| 商户号类型：1服务商、2商户、4门店商家 |
| api_v3_key | varchar(32) | | 用于ApiV3平台证书解密、回调信息解密   |
| api_v2_key | varchar(32) | | 用于ApiV2平台证书解密、回调信息解密   |
| pay_cert_p12 | text | | 支付证书文件 |
| pay_public_key_pem | text | | 公钥文件 |
| pay_private_key_pem | text | | 私钥文件 |
| cert_serial_number | text | | 证书序列号 |
| jsapi_auth_path | text | | JSAPI支付授权目录 |
| sys_user_id | int8 | | 用户ID |
| union_main_id | int8 | | 用户关联主体 |
| union_main_type | int2 | | 用户类型 |
| bankcard_account | varchar(20) | | 银行结算账户,用于交易和提现 |
| union_appid | json | | 该商户号关联的AppId，微信支付接入模式属于直连模式，限制只能是同一主体下的App列表 |
| updated_at | timestamptz(6) | | |
| app_id | varchar(128) | |  商户号 对应的公众号的服务号APPID |

### 微信支付服务商的特约商户表 
表名：weixin_pay_sub_merchant

| 字段 | 类型 | 约束        | 描述 |
| ---- | ---- |-----------| ---- |
| id | int8 | NOT NULL  | ID|
| sub_mchid | int4 | NOT NULL  | 特约商户商户号|
| sp_mchid | int4 | NOT NULL  | 服务商商户号|
| sub_appid | varchar(128) |           | 特约商户App唯一标识ID|
| sub_app_name | varchar(128) |           | 特约商户App名称|
| sub_app_type | int2 |           | 特约商户App类型：1公众号 2小程序 4网站应用H5  8移动应用  16视频小店|
| sub_merchant_name | varchar(255) |           | 特约商户公司名称|
| sub_merchant_short_name | varchar(255) |           | 特约商户商家简称|
| sys_user_id | int8 |           | 特约商户用户ID|
| union_main_id | int8 |           | 特约商户用户主体|
| union_main_type | int2 |           | 特约商户主体类型|
| jsapi_auth_path | text |           | JSAPI支付授权目录|
| h5_auth_path | text |           | H5支付授权目录|
| updated_at | timestamptz |           | |
| merchant_type | int2 |           | 商户号类型：1服务商、2商户、4门店商家|
| merchant_union_type | int2 |           | 特约商户主体类型：1个体工商户、2企业、4事业单位、8社会组织、16政府机关|
| bankcard_account | varchar(128) |           | 结算账号，添加特约商户的时候填写的结算银行账户|



### 微信订阅消息模板表
表名： weixin_subscribe_message_template

| 字段                             | 类型 | 约束       | 描述 |
|--------------------------------| ---- | ---------- | ---- |
| id                             | int8 |   NOT NULL | ID |
| template_id                    | varchar |   NOT NULL | 模板ID |
| type                           | int2 |   NOT NULL | 模板类型：2一次性订阅、3长期订阅 |
| title                          | varchar |   NOT NULL | 模板标题 |
| key_words                      | text |  | 模板主题词/关键词 |
| server_category                | text |  | 模板服务类目 |
| server_category_id             | int4 |  | 模板服务类目ID |
| content                        | text |  NOT NULL | 模板内容 |
| content_example                | text |  | 模板内容示例 |
| content_data_json              | json |  | 模板内容Json |
| key_word_enum_value_list_json  | json |  | 模板枚举参数值范围列表 |
| scene_desc                     | varchar |  | 场景描述 |
| scene_type                     | int2 |  | 场景类型【业务层自定义】：1活动即将开始提醒、2活动开始提醒、3活动即将结束提醒、4活动结束提醒、5活动获奖提醒、6券即将生效提醒、7券的生效提醒、8券的失效提醒、9券即将失效提醒、10券核销提醒、8192系统通知、 |
| message_type                   | int2 |  | 消息类型【业务层自定义】：1系统消息、2活动消息、4免啦券消息 |
| merchant_app_id                | varchar | NOT NULL   | 商家应用APPID |
| merchant_app_type              | int2 |  NOT NULL | 商家应用类型：1公众号、2小程序、4网站应用H5、8移动应用、16视频小店 |
| third_app_id                   | varchar |  | 第三方平台应用APPID |
| user_id                        | int8 |  | 应用所属账号 |
| union_main_id                  | int8 |  | 关联主体id |
| ext_json                       | json |  | 拓展字段Json |
| created_at                     | timestamptz |  |  |
| updated_at                     | timestamptz |  |  |


## 二、实现逻辑

weixin-library库：https://github.com/kysion/weixin-library.git

逻辑代码：

![image-20240806174820074](https://gitee.com/LinFeiFei0852/typora-images/raw/master/user/imgs/image-20240806174820074.png)



