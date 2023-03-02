/*
 Navicat Premium Data Transfer

 Source Server         : 192.168.1.105
 Source Server Type    : PostgreSQL
 Source Server Version : 140005 (140005)
 Source Host           : 192.168.1.105:5432
 Source Catalog        : kuaimk_test
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 140005 (140005)
 File Encoding         : 65001

 Date: 28/02/2023 11:36:38
*/


-- ----------------------------
-- Table structure for third_app_config
-- ----------------------------
DROP TABLE IF EXISTS "public"."third_app_config";
CREATE TABLE "public"."third_app_config" (
                                             "id" int8 NOT NULL,
                                             "name" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
                                             "app_id" int8 NOT NULL,
                                             "app_name" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
                                             "app_type" varchar(32) COLLATE "pg_catalog"."default",
                                             "app_auth_token" varchar(255) COLLATE "pg_catalog"."default",
                                             "is_full_proxy" int2 NOT NULL,
                                             "auth_state" int2,
                                             "expires_in" timestamp(6),
                                             "re_expires_in" timestamp(6),
                                             "user_id" int8 NOT NULL,
                                             "union_main_id" int8,
                                             "sys_user_id" int8,
                                             "tokens" varchar(255) COLLATE "pg_catalog"."default",
                                             "ext_json" json,
                                             "created_at" timestamp(6),
                                             "updated_at" timestamp(6),
                                             "deleted_at" timestamp(6)
)
;
ALTER TABLE "public"."third_app_config" OWNER TO "kuaimk";
COMMENT ON COLUMN "public"."third_app_config"."id" IS '授权商家id';
COMMENT ON COLUMN "public"."third_app_config"."name" IS '授权商家name';
COMMENT ON COLUMN "public"."third_app_config"."app_id" IS '商家授权应用Id';
COMMENT ON COLUMN "public"."third_app_config"."app_name" IS '商家授权应用名称';
COMMENT ON COLUMN "public"."third_app_config"."app_type" IS '应用类型';
COMMENT ON COLUMN "public"."third_app_config"."app_auth_token" IS '授权应用token';
COMMENT ON COLUMN "public"."third_app_config"."is_full_proxy" IS '是否全权委托待开发：0否 1是';
COMMENT ON COLUMN "public"."third_app_config"."auth_state" IS '授权状态';
COMMENT ON COLUMN "public"."third_app_config"."expires_in" IS '生效时间';
COMMENT ON COLUMN "public"."third_app_config"."re_expires_in" IS '失效时间';
COMMENT ON COLUMN "public"."third_app_config"."user_id" IS '用户账号id';
COMMENT ON COLUMN "public"."third_app_config"."union_main_id" IS '关联主体id';
COMMENT ON COLUMN "public"."third_app_config"."sys_user_id" IS '用户id';
COMMENT ON COLUMN "public"."third_app_config"."tokens" IS 'token列表';
COMMENT ON COLUMN "public"."third_app_config"."ext_json" IS '拓展字段';

-- ----------------------------
-- Records of third_app_config
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Primary Key structure for table third_app_config
-- ----------------------------
ALTER TABLE "public"."third_app_config" ADD CONSTRAINT "merchant_config_copy1_pkey" PRIMARY KEY ("id");
/*
 Navicat Premium Data Transfer

 Source Server         : 192.168.1.105
 Source Server Type    : PostgreSQL
 Source Server Version : 140005 (140005)
 Source Host           : 192.168.1.105:5432
 Source Catalog        : kuaimk_test
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 140005 (140005)
 File Encoding         : 65001

 Date: 28/02/2023 11:37:01
*/


-- ----------------------------
-- Table structure for consumer_config
-- ----------------------------
DROP TABLE IF EXISTS "public"."consumer_config";
CREATE TABLE "public"."consumer_config" (
                                            "id" int8 NOT NULL,
                                            "user_id" int8 NOT NULL,
                                            "sys_user_id" int8,
                                            "avatar" text COLLATE "pg_catalog"."default",
                                            "province" varchar(64) COLLATE "pg_catalog"."default",
                                            "city" varchar(64) COLLATE "pg_catalog"."default",
                                            "nick_name" varchar(64) COLLATE "pg_catalog"."default",
                                            "is_student_certified" int2,
                                            "user_type" varchar(64) COLLATE "pg_catalog"."default",
                                            "user_state" int2,
                                            "is_certified" int2,
                                            "sex" int2,
                                            "auth_token" text COLLATE "pg_catalog"."default",
                                            "ext_json" json,
                                            "created_at" timestamp(6),
                                            "updated_at" timestamp(6),
                                            "deleted_at" timestamp(6)
)
;
ALTER TABLE "public"."consumer_config" OWNER TO "kuaimk";
COMMENT ON COLUMN "public"."consumer_config"."id" IS 'id';
COMMENT ON COLUMN "public"."consumer_config"."user_id" IS '用户账号id';
COMMENT ON COLUMN "public"."consumer_config"."sys_user_id" IS '用户id';
COMMENT ON COLUMN "public"."consumer_config"."avatar" IS '头像';
COMMENT ON COLUMN "public"."consumer_config"."province" IS '省份';
COMMENT ON COLUMN "public"."consumer_config"."city" IS '城市';
COMMENT ON COLUMN "public"."consumer_config"."nick_name" IS '昵称';
COMMENT ON COLUMN "public"."consumer_config"."is_student_certified" IS '学生认证';
COMMENT ON COLUMN "public"."consumer_config"."user_type" IS '用户账号类型';
COMMENT ON COLUMN "public"."consumer_config"."user_state" IS '状态：0未激活、1正常、-1封号、-2异常、-3已注销';
COMMENT ON COLUMN "public"."consumer_config"."is_certified" IS '是否实名认证';
COMMENT ON COLUMN "public"."consumer_config"."sex" IS '性别：0女 1男';
COMMENT ON COLUMN "public"."consumer_config"."auth_token" IS '授权token';
COMMENT ON COLUMN "public"."consumer_config"."ext_json" IS '拓展字段';

-- ----------------------------
-- Records of consumer_config
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Primary Key structure for table consumer_config
-- ----------------------------
ALTER TABLE "public"."consumer_config" ADD CONSTRAINT "consumer_config_pkey" PRIMARY KEY ("id");

/*
 Navicat Premium Data Transfer

 Source Server         : 192.168.1.105
 Source Server Type    : PostgreSQL
 Source Server Version : 140005 (140005)
 Source Host           : 192.168.1.105:5432
 Source Catalog        : kuaimk_test
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 140005 (140005)
 File Encoding         : 65001

 Date: 28/02/2023 11:36:52
*/


-- ----------------------------
-- Table structure for merchant_app_config
-- ----------------------------
DROP TABLE IF EXISTS "public"."merchant_app_config";
CREATE TABLE "public"."merchant_app_config" (
                                                "id" int8 NOT NULL,
                                                "name" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
                                                "app_id" int8 NOT NULL,
                                                "app_name" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
                                                "app_type" varchar(32) COLLATE "pg_catalog"."default",
                                                "app_auth_token" varchar(255) COLLATE "pg_catalog"."default",
                                                "is_full_proxy" int2 NOT NULL,
                                                "auth_state" int2,
                                                "expires_in" timestamp(6),
                                                "re_expires_in" timestamp(6),
                                                "user_id" int8 NOT NULL,
                                                "union_main_id" int8,
                                                "sys_user_id" int8,
                                                "tokens" varchar(255) COLLATE "pg_catalog"."default",
                                                "ext_json" json,
                                                "created_at" timestamp(6),
                                                "updated_at" timestamp(6),
                                                "deleted_at" timestamp(6)
)
;
ALTER TABLE "public"."merchant_app_config" OWNER TO "kuaimk";
COMMENT ON COLUMN "public"."merchant_app_config"."id" IS '授权商家id';
COMMENT ON COLUMN "public"."merchant_app_config"."name" IS '授权商家name';
COMMENT ON COLUMN "public"."merchant_app_config"."app_id" IS '商家授权应用Id';
COMMENT ON COLUMN "public"."merchant_app_config"."app_name" IS '商家授权应用名称';
COMMENT ON COLUMN "public"."merchant_app_config"."app_type" IS '应用类型';
COMMENT ON COLUMN "public"."merchant_app_config"."app_auth_token" IS '授权应用token';
COMMENT ON COLUMN "public"."merchant_app_config"."is_full_proxy" IS '是否全权委托待开发：0否 1是';
COMMENT ON COLUMN "public"."merchant_app_config"."auth_state" IS '授权状态';
COMMENT ON COLUMN "public"."merchant_app_config"."expires_in" IS '生效时间';
COMMENT ON COLUMN "public"."merchant_app_config"."re_expires_in" IS '失效时间';
COMMENT ON COLUMN "public"."merchant_app_config"."user_id" IS '用户账号id';
COMMENT ON COLUMN "public"."merchant_app_config"."union_main_id" IS '关联主体id';
COMMENT ON COLUMN "public"."merchant_app_config"."sys_user_id" IS '用户id';
COMMENT ON COLUMN "public"."merchant_app_config"."tokens" IS 'token列表';
COMMENT ON COLUMN "public"."merchant_app_config"."ext_json" IS '拓展字段';

-- ----------------------------
-- Records of merchant_app_config
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Primary Key structure for table merchant_app_config
-- ----------------------------
ALTER TABLE "public"."merchant_app_config" ADD CONSTRAINT "merchant_config_pkey" PRIMARY KEY ("id");

