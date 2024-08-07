/*
 Navicat Premium Data Transfer

 Source Server         : 筷满客-Dev
 Source Server Type    : PostgreSQL
 Source Server Version : 140008 (140008)
 Source Host           : 10.68.74.250:5432
 Source Catalog        : kuaimkdb
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 140008 (140008)
 File Encoding         : 65001

 Date: 06/08/2024 16:59:16
*/

/*
注意：实际执行SQL文件的时候，
需要将 OWNER TO "kuaimk" 中的"kuaimk" 替换成 实际的SQL用户 或者 删除此语句。
*/


-- ----------------------------
-- Table structure for weixin_consumer_config
-- ----------------------------
DROP TABLE IF EXISTS "public"."weixin_consumer_config";
CREATE TABLE "public"."weixin_consumer_config" (
  "id" int8 NOT NULL,
  "open_id" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "sys_user_id" int8,
  "avatar" text COLLATE "pg_catalog"."default",
  "province" varchar(64) COLLATE "pg_catalog"."default",
  "city" varchar(64) COLLATE "pg_catalog"."default",
  "nick_name" varchar(64) COLLATE "pg_catalog"."default",
  "is_student_certified" int2,
  "user_type" int2,
  "user_state" int2,
  "is_certified" int2,
  "sex" int2,
  "access_token" text COLLATE "pg_catalog"."default",
  "ext_json" json,
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6),
  "deleted_at" timestamptz(6),
  "union_id" varchar(255) COLLATE "pg_catalog"."default",
  "session_key" varchar(255) COLLATE "pg_catalog"."default",
  "refresh_token" text COLLATE "pg_catalog"."default",
  "expires_in" timestamptz(6),
  "auth_state" int2,
  "app_type" int2,
  "is_follow_public" int2,
  "app_id" varchar(64) COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "public"."weixin_consumer_config" OWNER TO "kuaimk";
COMMENT ON COLUMN "public"."weixin_consumer_config"."id" IS 'id';
COMMENT ON COLUMN "public"."weixin_consumer_config"."open_id" IS '微信用户openId，不同应用下的用户具备不同的openId';
COMMENT ON COLUMN "public"."weixin_consumer_config"."sys_user_id" IS '用户id';
COMMENT ON COLUMN "public"."weixin_consumer_config"."avatar" IS '头像';
COMMENT ON COLUMN "public"."weixin_consumer_config"."province" IS '省份';
COMMENT ON COLUMN "public"."weixin_consumer_config"."city" IS '城市';
COMMENT ON COLUMN "public"."weixin_consumer_config"."nick_name" IS '昵称';
COMMENT ON COLUMN "public"."weixin_consumer_config"."is_student_certified" IS '是否学生认证';
COMMENT ON COLUMN "public"."weixin_consumer_config"."user_type" IS '用户账号类型，和sysUserType保持一致';
COMMENT ON COLUMN "public"."weixin_consumer_config"."user_state" IS '状态：0未激活、1正常、-1封号、-2异常、-3已注销';
COMMENT ON COLUMN "public"."weixin_consumer_config"."is_certified" IS '是否实名认证';
COMMENT ON COLUMN "public"."weixin_consumer_config"."sex" IS '性别：0未知、1男、2女';
COMMENT ON COLUMN "public"."weixin_consumer_config"."access_token" IS '授权token';
COMMENT ON COLUMN "public"."weixin_consumer_config"."ext_json" IS '拓展字段';
COMMENT ON COLUMN "public"."weixin_consumer_config"."union_id" IS '微信用户union_id，同一个开放平台帐号下的用户只有一个unionId';
COMMENT ON COLUMN "public"."weixin_consumer_config"."session_key" IS '微信用户会话key';
COMMENT ON COLUMN "public"."weixin_consumer_config"."refresh_token" IS '用户授权刷新令牌';
COMMENT ON COLUMN "public"."weixin_consumer_config"."expires_in" IS '令牌过期时间';
COMMENT ON COLUMN "public"."weixin_consumer_config"."auth_state" IS '微信用户授权状态：1已授权、2未授权';
COMMENT ON COLUMN "public"."weixin_consumer_config"."app_type" IS '应用类型：1公众号 2小程序 4网站应用H5  8移动应用  16视频小店';
COMMENT ON COLUMN "public"."weixin_consumer_config"."is_follow_public" IS '是否关注公众号：1关注、2未关注';
COMMENT ON COLUMN "public"."weixin_consumer_config"."app_id" IS '商家应用Id';

-- ----------------------------
-- Table structure for weixin_merchant_app_config
-- ----------------------------
DROP TABLE IF EXISTS "public"."weixin_merchant_app_config";
CREATE TABLE "public"."weixin_merchant_app_config" (
  "id" int8 NOT NULL,
  "name" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "app_id" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "app_name" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "app_type" int2,
  "app_auth_token" varchar(255) COLLATE "pg_catalog"."default",
  "is_full_proxy" int2 NOT NULL,
  "state" int2 NOT NULL,
  "expires_in" timestamptz(6),
  "re_expires_in" timestamptz(6),
  "user_id" int8,
  "union_main_id" int8,
  "sys_user_id" int8,
  "ext_json" json,
  "app_gateway_url" text COLLATE "pg_catalog"."default",
  "app_callback_url" text COLLATE "pg_catalog"."default",
  "app_secret" varchar(255) COLLATE "pg_catalog"."default",
  "msg_verfiy_token" text COLLATE "pg_catalog"."default",
  "msg_encrypt_key" text COLLATE "pg_catalog"."default",
  "msg_encrypt_type" int2,
  "business_domain" text COLLATE "pg_catalog"."default",
  "js_domain" text COLLATE "pg_catalog"."default",
  "auth_domain" varchar(255) COLLATE "pg_catalog"."default",
  "logo" text COLLATE "pg_catalog"."default",
  "https_cert" text COLLATE "pg_catalog"."default",
  "https_key" text COLLATE "pg_catalog"."default",
  "server_domain" text COLLATE "pg_catalog"."default",
  "app_id_md5" varchar(32) COLLATE "pg_catalog"."default",
  "third_app_id" varchar(64) COLLATE "pg_catalog"."default",
  "notify_url" text COLLATE "pg_catalog"."default",
  "server_rate" float4 DEFAULT 0.0006,
  "union_main_type" varchar(255) COLLATE "pg_catalog"."default",
  "version" varchar(16) COLLATE "pg_catalog"."default",
  "privacy_policy" text COLLATE "pg_catalog"."default",
  "user_policy" text COLLATE "pg_catalog"."default",
  "dev_state" int2,
  "updated_at" timestamptz(6),
  "refresh_token" varchar(255) COLLATE "pg_catalog"."default",
  "primitive_id" varchar(64) COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "public"."weixin_merchant_app_config" OWNER TO "kuaimk";
COMMENT ON COLUMN "public"."weixin_merchant_app_config"."id" IS '商家id';
COMMENT ON COLUMN "public"."weixin_merchant_app_config"."name" IS '商家name';
COMMENT ON COLUMN "public"."weixin_merchant_app_config"."app_id" IS '商家应用Id';
COMMENT ON COLUMN "public"."weixin_merchant_app_config"."app_name" IS '商家应用名称';
COMMENT ON COLUMN "public"."weixin_merchant_app_config"."app_type" IS '应用类型：1公众号 2小程序 4网站应用H5  8移动应用  16视频小店';
COMMENT ON COLUMN "public"."weixin_merchant_app_config"."app_auth_token" IS '商家应用token：1、当第三方代开发的时候，这个是商家授权的应用token (authorizer_access_token)；2、当是商家自研模式时，这个是商家的应用token(access_token)。';
COMMENT ON COLUMN "public"."weixin_merchant_app_config"."is_full_proxy" IS '是否全权委托待开发：0否 1是';
COMMENT ON COLUMN "public"."weixin_merchant_app_config"."state" IS '状态： 0禁用 1启用';
COMMENT ON COLUMN "public"."weixin_merchant_app_config"."expires_in" IS 'Token过期时间';
COMMENT ON COLUMN "public"."weixin_merchant_app_config"."re_expires_in" IS 'Token限期刷新时间';
COMMENT ON COLUMN "public"."weixin_merchant_app_config"."user_id" IS '应用所属账号';
COMMENT ON COLUMN "public"."weixin_merchant_app_config"."union_main_id" IS '关联主体id';
COMMENT ON COLUMN "public"."weixin_merchant_app_config"."sys_user_id" IS '用户id';
COMMENT ON COLUMN "public"."weixin_merchant_app_config"."ext_json" IS '拓展字段';
COMMENT ON COLUMN "public"."weixin_merchant_app_config"."app_gateway_url" IS '网关地址';
COMMENT ON COLUMN "public"."weixin_merchant_app_config"."app_callback_url" IS '回调地址';
COMMENT ON COLUMN "public"."weixin_merchant_app_config"."app_secret" IS '商家应用密钥';
COMMENT ON COLUMN "public"."weixin_merchant_app_config"."msg_verfiy_token" IS '消息校验Token';
COMMENT ON COLUMN "public"."weixin_merchant_app_config"."msg_encrypt_key" IS '消息加密解密密钥（EncodingAESKey）';
COMMENT ON COLUMN "public"."weixin_merchant_app_config"."msg_encrypt_type" IS '消息加密模式：1兼容模式 2明文模式 4安全模式';
COMMENT ON COLUMN "public"."weixin_merchant_app_config"."business_domain" IS '业务域名';
COMMENT ON COLUMN "public"."weixin_merchant_app_config"."js_domain" IS 'JS接口安全域名';
COMMENT ON COLUMN "public"."weixin_merchant_app_config"."auth_domain" IS '网页授权域名';
COMMENT ON COLUMN "public"."weixin_merchant_app_config"."logo" IS '商家logo';
COMMENT ON COLUMN "public"."weixin_merchant_app_config"."https_cert" IS '域名证书';
COMMENT ON COLUMN "public"."weixin_merchant_app_config"."https_key" IS '域名私钥';
COMMENT ON COLUMN "public"."weixin_merchant_app_config"."server_domain" IS '服务器域名';
COMMENT ON COLUMN "public"."weixin_merchant_app_config"."app_id_md5" IS '应用id加密md5后的结果';
COMMENT ON COLUMN "public"."weixin_merchant_app_config"."third_app_id" IS '服务商appId';
COMMENT ON COLUMN "public"."weixin_merchant_app_config"."notify_url" IS '异步通知地址，允许业务层追加相关参数';
COMMENT ON COLUMN "public"."weixin_merchant_app_config"."server_rate" IS '手续费比例，默认0.6%';
COMMENT ON COLUMN "public"."weixin_merchant_app_config"."union_main_type" IS '应用关联主体类型，和user_type保持一致';
COMMENT ON COLUMN "public"."weixin_merchant_app_config"."version" IS '应用版本';
COMMENT ON COLUMN "public"."weixin_merchant_app_config"."privacy_policy" IS '隐私协议';
COMMENT ON COLUMN "public"."weixin_merchant_app_config"."user_policy" IS '用户协议';
COMMENT ON COLUMN "public"."weixin_merchant_app_config"."dev_state" IS '开发状态：0未上线 1已上线';
COMMENT ON COLUMN "public"."weixin_merchant_app_config"."refresh_token" IS '刷新商家授权应用Token';
COMMENT ON COLUMN "public"."weixin_merchant_app_config"."primitive_id" IS '应用原始ID';

-- ----------------------------
-- Table structure for weixin_pay_merchant
-- ----------------------------
DROP TABLE IF EXISTS "public"."weixin_pay_merchant";
CREATE TABLE "public"."weixin_pay_merchant" (
  "id" int8 NOT NULL,
  "mchid" int4 NOT NULL,
  "merchant_name" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "merchant_short_name" varchar(255) COLLATE "pg_catalog"."default",
  "merchant_type" int2 NOT NULL,
  "api_v3_key" varchar(32) COLLATE "pg_catalog"."default",
  "api_v2_key" varchar(32) COLLATE "pg_catalog"."default",
  "pay_cert_p12" text COLLATE "pg_catalog"."default",
  "pay_public_key_pem" text COLLATE "pg_catalog"."default",
  "pay_private_key_pem" text COLLATE "pg_catalog"."default",
  "cert_serial_number" text COLLATE "pg_catalog"."default",
  "jsapi_auth_path" text COLLATE "pg_catalog"."default",
  "sys_user_id" int8,
  "union_main_id" int8,
  "union_main_type" int2,
  "bankcard_account" varchar(20) COLLATE "pg_catalog"."default",
  "union_appid" json,
  "updated_at" timestamptz(6),
  "app_id" varchar(128) COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "public"."weixin_pay_merchant" OWNER TO "kuaimk";
COMMENT ON COLUMN "public"."weixin_pay_merchant"."id" IS 'ID';
COMMENT ON COLUMN "public"."weixin_pay_merchant"."mchid" IS '微信支付商户号';
COMMENT ON COLUMN "public"."weixin_pay_merchant"."merchant_name" IS '商户号公司名称';
COMMENT ON COLUMN "public"."weixin_pay_merchant"."merchant_short_name" IS '商户号简称';
COMMENT ON COLUMN "public"."weixin_pay_merchant"."merchant_type" IS '商户号类型：1服务商、2商户、4门店商家';
COMMENT ON COLUMN "public"."weixin_pay_merchant"."api_v3_key" IS '用于ApiV3平台证书解密、回调信息解密  ';
COMMENT ON COLUMN "public"."weixin_pay_merchant"."api_v2_key" IS '用于ApiV2平台证书解密、回调信息解密  ';
COMMENT ON COLUMN "public"."weixin_pay_merchant"."pay_cert_p12" IS '支付证书文件';
COMMENT ON COLUMN "public"."weixin_pay_merchant"."pay_public_key_pem" IS '公钥文件';
COMMENT ON COLUMN "public"."weixin_pay_merchant"."pay_private_key_pem" IS '私钥文件';
COMMENT ON COLUMN "public"."weixin_pay_merchant"."cert_serial_number" IS '证书序列号';
COMMENT ON COLUMN "public"."weixin_pay_merchant"."jsapi_auth_path" IS 'JSAPI支付授权目录';
COMMENT ON COLUMN "public"."weixin_pay_merchant"."sys_user_id" IS '用户ID';
COMMENT ON COLUMN "public"."weixin_pay_merchant"."union_main_id" IS '用户关联主体';
COMMENT ON COLUMN "public"."weixin_pay_merchant"."union_main_type" IS '用户类型';
COMMENT ON COLUMN "public"."weixin_pay_merchant"."bankcard_account" IS '银行结算账户,用于交易和提现';
COMMENT ON COLUMN "public"."weixin_pay_merchant"."union_appid" IS '该商户号关联的AppId，微信支付接入模式属于直连模式，限制只能是同一主体下的App列表';
COMMENT ON COLUMN "public"."weixin_pay_merchant"."app_id" IS '商户号 对应的公众号的服务号APPID';

-- ----------------------------
-- Table structure for weixin_pay_sub_merchant
-- ----------------------------
DROP TABLE IF EXISTS "public"."weixin_pay_sub_merchant";
CREATE TABLE "public"."weixin_pay_sub_merchant" (
  "id" int8 NOT NULL,
  "sub_mchid" int4 NOT NULL,
  "sp_mchid" int4 NOT NULL,
  "sub_appid" varchar(128) COLLATE "pg_catalog"."default",
  "sub_app_name" varchar(128) COLLATE "pg_catalog"."default",
  "sub_app_type" int2,
  "sub_merchant_name" varchar(255) COLLATE "pg_catalog"."default",
  "sub_merchant_short_name" varchar(255) COLLATE "pg_catalog"."default",
  "sys_user_id" int8,
  "union_main_id" int8,
  "union_main_type" int2,
  "jsapi_auth_path" text COLLATE "pg_catalog"."default",
  "h5_auth_path" text COLLATE "pg_catalog"."default",
  "updated_at" timestamptz(6),
  "merchant_type" int2,
  "merchant_union_type" int2,
  "bankcard_account" varchar(128) COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "public"."weixin_pay_sub_merchant" OWNER TO "kuaimk";
COMMENT ON COLUMN "public"."weixin_pay_sub_merchant"."id" IS 'ID';
COMMENT ON COLUMN "public"."weixin_pay_sub_merchant"."sub_mchid" IS '特约商户商户号';
COMMENT ON COLUMN "public"."weixin_pay_sub_merchant"."sp_mchid" IS '服务商商户号';
COMMENT ON COLUMN "public"."weixin_pay_sub_merchant"."sub_appid" IS '特约商户App唯一标识ID';
COMMENT ON COLUMN "public"."weixin_pay_sub_merchant"."sub_app_name" IS '特约商户App名称';
COMMENT ON COLUMN "public"."weixin_pay_sub_merchant"."sub_app_type" IS '特约商户App类型：1公众号 2小程序 4网站应用H5  8移动应用  16视频小店';
COMMENT ON COLUMN "public"."weixin_pay_sub_merchant"."sub_merchant_name" IS '特约商户公司名称';
COMMENT ON COLUMN "public"."weixin_pay_sub_merchant"."sub_merchant_short_name" IS '特约商户商家简称';
COMMENT ON COLUMN "public"."weixin_pay_sub_merchant"."sys_user_id" IS '特约商户用户ID';
COMMENT ON COLUMN "public"."weixin_pay_sub_merchant"."union_main_id" IS '特约商户用户主体';
COMMENT ON COLUMN "public"."weixin_pay_sub_merchant"."union_main_type" IS '特约商户主体类型';
COMMENT ON COLUMN "public"."weixin_pay_sub_merchant"."jsapi_auth_path" IS 'JSAPI支付授权目录';
COMMENT ON COLUMN "public"."weixin_pay_sub_merchant"."h5_auth_path" IS 'H5支付授权目录';
COMMENT ON COLUMN "public"."weixin_pay_sub_merchant"."merchant_type" IS '商户号类型：1服务商、2商户、4门店商家';
COMMENT ON COLUMN "public"."weixin_pay_sub_merchant"."merchant_union_type" IS '特约商户主体类型：1个体工商户、2企业、4事业单位、8社会组织、16政府机关';
COMMENT ON COLUMN "public"."weixin_pay_sub_merchant"."bankcard_account" IS '结算账号，添加特约商户的时候填写的结算银行账户';

-- ----------------------------
-- Table structure for weixin_subscribe_message_template
-- ----------------------------
DROP TABLE IF EXISTS "public"."weixin_subscribe_message_template";
CREATE TABLE "public"."weixin_subscribe_message_template" (
  "id" int8 NOT NULL,
  "template_id" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "type" int2 NOT NULL,
  "title" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "key_words" text COLLATE "pg_catalog"."default",
  "server_category" text COLLATE "pg_catalog"."default",
  "server_category_id" int4,
  "content" text COLLATE "pg_catalog"."default" NOT NULL,
  "content_example" text COLLATE "pg_catalog"."default",
  "content_data_json" json,
  "key_word_enum_value_list_json" json,
  "scene_desc" varchar(128) COLLATE "pg_catalog"."default",
  "scene_type" int2 DEFAULT 0,
  "message_type" int2,
  "merchant_app_id" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "merchant_app_type" int2 NOT NULL DEFAULT 0,
  "third_app_id" varchar(64) COLLATE "pg_catalog"."default",
  "user_id" int8,
  "union_main_id" int8,
  "ext_json" json,
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6)
)
;
ALTER TABLE "public"."weixin_subscribe_message_template" OWNER TO "kuaimk";
COMMENT ON COLUMN "public"."weixin_subscribe_message_template"."id" IS 'ID';
COMMENT ON COLUMN "public"."weixin_subscribe_message_template"."template_id" IS '模板ID';
COMMENT ON COLUMN "public"."weixin_subscribe_message_template"."type" IS '模板类型：2一次性订阅、3长期订阅';
COMMENT ON COLUMN "public"."weixin_subscribe_message_template"."title" IS '模板标题';
COMMENT ON COLUMN "public"."weixin_subscribe_message_template"."key_words" IS '模板主题词/关键词';
COMMENT ON COLUMN "public"."weixin_subscribe_message_template"."server_category" IS '模板服务类目';
COMMENT ON COLUMN "public"."weixin_subscribe_message_template"."server_category_id" IS '模板服务类目ID';
COMMENT ON COLUMN "public"."weixin_subscribe_message_template"."content" IS '模板内容';
COMMENT ON COLUMN "public"."weixin_subscribe_message_template"."content_example" IS '模板内容示例';
COMMENT ON COLUMN "public"."weixin_subscribe_message_template"."content_data_json" IS '模板内容Json';
COMMENT ON COLUMN "public"."weixin_subscribe_message_template"."key_word_enum_value_list_json" IS '模板枚举参数值范围列表';
COMMENT ON COLUMN "public"."weixin_subscribe_message_template"."scene_desc" IS '场景描述';
COMMENT ON COLUMN "public"."weixin_subscribe_message_template"."scene_type" IS '场景类型【业务层自定义】：1活动即将开始提醒、2活动开始提醒、3活动即将结束提醒、4活动结束提醒、5活动获奖提醒、6券即将生效提醒、7券的生效提醒、8券的失效提醒、9券即将失效提醒、10券核销提醒、8192系统通知、';
COMMENT ON COLUMN "public"."weixin_subscribe_message_template"."message_type" IS '消息类型【业务层自定义】：1系统消息、2活动消息、4免啦券消息';
COMMENT ON COLUMN "public"."weixin_subscribe_message_template"."merchant_app_id" IS '商家应用APPID';
COMMENT ON COLUMN "public"."weixin_subscribe_message_template"."merchant_app_type" IS '商家应用类型：1公众号、2小程序、4网站应用H5、8移动应用、16视频小店';
COMMENT ON COLUMN "public"."weixin_subscribe_message_template"."third_app_id" IS '第三方平台应用APPID';
COMMENT ON COLUMN "public"."weixin_subscribe_message_template"."user_id" IS '应用所属账号';
COMMENT ON COLUMN "public"."weixin_subscribe_message_template"."union_main_id" IS '关联主体id';
COMMENT ON COLUMN "public"."weixin_subscribe_message_template"."ext_json" IS '拓展字段Json';

-- ----------------------------
-- Table structure for weixin_third_app_config
-- ----------------------------
DROP TABLE IF EXISTS "public"."weixin_third_app_config";
CREATE TABLE "public"."weixin_third_app_config" (
  "id" int8 NOT NULL,
  "name" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "app_id" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "app_name" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "app_type" int2,
  "app_auth_token" varchar(255) COLLATE "pg_catalog"."default",
  "expires_in" timestamptz(6),
  "re_expires_in" timestamptz(6),
  "union_main_id" int8,
  "sys_user_id" int8,
  "ext_json" json,
  "app_gateway_url" text COLLATE "pg_catalog"."default" NOT NULL,
  "app_callback_url" text COLLATE "pg_catalog"."default" NOT NULL,
  "app_secret" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "msg_verfiy_token" text COLLATE "pg_catalog"."default",
  "msg_encrypt_key" text COLLATE "pg_catalog"."default",
  "auth_init_url" text COLLATE "pg_catalog"."default",
  "server_domain" text COLLATE "pg_catalog"."default",
  "business_domain" text COLLATE "pg_catalog"."default",
  "auth_test_appIds" text COLLATE "pg_catalog"."default",
  "platform_site" text COLLATE "pg_catalog"."default",
  "logo" text COLLATE "pg_catalog"."default",
  "state" int2 NOT NULL DEFAULT 0,
  "release_state" int2 NOT NULL DEFAULT 0,
  "https_cert" text COLLATE "pg_catalog"."default",
  "https_key" text COLLATE "pg_catalog"."default",
  "updated_at" timestamptz(6),
  "app_id_md5" varchar(32) COLLATE "pg_catalog"."default",
  "user_id" int8,
  "refresh_token" varchar(255) COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "public"."weixin_third_app_config" OWNER TO "kuaimk";
COMMENT ON COLUMN "public"."weixin_third_app_config"."id" IS '服务商id';
COMMENT ON COLUMN "public"."weixin_third_app_config"."name" IS '服务商name';
COMMENT ON COLUMN "public"."weixin_third_app_config"."app_id" IS '服务商应用Id';
COMMENT ON COLUMN "public"."weixin_third_app_config"."app_name" IS '服务商应用名称';
COMMENT ON COLUMN "public"."weixin_third_app_config"."app_type" IS '服务商应用类型';
COMMENT ON COLUMN "public"."weixin_third_app_config"."app_auth_token" IS '服务商应用授权token';
COMMENT ON COLUMN "public"."weixin_third_app_config"."expires_in" IS 'Token过期时间';
COMMENT ON COLUMN "public"."weixin_third_app_config"."re_expires_in" IS 'Token限期刷新时间';
COMMENT ON COLUMN "public"."weixin_third_app_config"."union_main_id" IS '关联主体id';
COMMENT ON COLUMN "public"."weixin_third_app_config"."sys_user_id" IS '用户id';
COMMENT ON COLUMN "public"."weixin_third_app_config"."ext_json" IS '拓展字段';
COMMENT ON COLUMN "public"."weixin_third_app_config"."app_gateway_url" IS '网关地址';
COMMENT ON COLUMN "public"."weixin_third_app_config"."app_callback_url" IS '回调地址';
COMMENT ON COLUMN "public"."weixin_third_app_config"."app_secret" IS '服务商应用密钥';
COMMENT ON COLUMN "public"."weixin_third_app_config"."msg_verfiy_token" IS '消息校验Token';
COMMENT ON COLUMN "public"."weixin_third_app_config"."msg_encrypt_key" IS '消息加密解密密钥';
COMMENT ON COLUMN "public"."weixin_third_app_config"."auth_init_url" IS '授权发起页域名';
COMMENT ON COLUMN "public"."weixin_third_app_config"."server_domain" IS '服务器域名';
COMMENT ON COLUMN "public"."weixin_third_app_config"."business_domain" IS '业务域名';
COMMENT ON COLUMN "public"."weixin_third_app_config"."auth_test_appIds" IS '授权测试应用列表';
COMMENT ON COLUMN "public"."weixin_third_app_config"."platform_site" IS '平台官方';
COMMENT ON COLUMN "public"."weixin_third_app_config"."logo" IS '服务商logo';
COMMENT ON COLUMN "public"."weixin_third_app_config"."state" IS '状态：0禁用 1启用';
COMMENT ON COLUMN "public"."weixin_third_app_config"."release_state" IS '发布状态：0未发布 1已发布';
COMMENT ON COLUMN "public"."weixin_third_app_config"."https_cert" IS '域名证书';
COMMENT ON COLUMN "public"."weixin_third_app_config"."https_key" IS '域名私钥';
COMMENT ON COLUMN "public"."weixin_third_app_config"."app_id_md5" IS '应用id加密md5后的结果';
COMMENT ON COLUMN "public"."weixin_third_app_config"."user_id" IS '应用所属账号';
COMMENT ON COLUMN "public"."weixin_third_app_config"."refresh_token" IS '刷新应用Token';

-- ----------------------------
-- Primary Key structure for table weixin_consumer_config
-- ----------------------------
ALTER TABLE "public"."weixin_consumer_config" ADD CONSTRAINT "weixin_consumer_config_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table weixin_merchant_app_config
-- ----------------------------
ALTER TABLE "public"."weixin_merchant_app_config" ADD CONSTRAINT "weixin_merchant_app_config_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table weixin_pay_merchant
-- ----------------------------
ALTER TABLE "public"."weixin_pay_merchant" ADD CONSTRAINT "weixin_pay_merchant_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table weixin_pay_sub_merchant
-- ----------------------------
ALTER TABLE "public"."weixin_pay_sub_merchant" ADD CONSTRAINT "weixin_pay_sub_merchant_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table weixin_subscribe_message_template
-- ----------------------------
ALTER TABLE "public"."weixin_subscribe_message_template" ADD CONSTRAINT "weixin_message_template_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table weixin_third_app_config
-- ----------------------------
ALTER TABLE "public"."weixin_third_app_config" ADD CONSTRAINT "weixin_third_app_config_pkey" PRIMARY KEY ("id");
