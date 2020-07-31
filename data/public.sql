/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : PostgreSQL
 Source Server Version : 100007
 Source Host           : localhost:5432
 Source Catalog        : cmkit
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 100007
 File Encoding         : 65001

 Date: 31/07/2020 16:15:32
*/


-- ----------------------------
-- Sequence structure for r_auth_role_func_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."r_auth_role_func_id_seq";
CREATE SEQUENCE "public"."r_auth_role_func_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for r_auth_user_role_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."r_auth_user_role_id_seq";
CREATE SEQUENCE "public"."r_auth_user_role_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for t_auth_role_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."t_auth_role_id_seq";
CREATE SEQUENCE "public"."t_auth_role_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for t_auth_user_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."t_auth_user_id_seq";
CREATE SEQUENCE "public"."t_auth_user_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for t_sys_staff_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."t_sys_staff_id_seq";
CREATE SEQUENCE "public"."t_sys_staff_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Table structure for r_auth_role_func
-- ----------------------------
DROP TABLE IF EXISTS "public"."r_auth_role_func";
CREATE TABLE "public"."r_auth_role_func" (
  "id" int4 NOT NULL DEFAULT nextval('r_auth_role_func_id_seq'::regclass),
  "role_id" int4,
  "funcs" text COLLATE "pg_catalog"."default"
)
;

-- ----------------------------
-- Table structure for r_auth_user_role
-- ----------------------------
DROP TABLE IF EXISTS "public"."r_auth_user_role";
CREATE TABLE "public"."r_auth_user_role" (
  "id" int4 NOT NULL DEFAULT nextval('r_auth_user_role_id_seq'::regclass),
  "user_id" int4,
  "role_id" int4
)
;

-- ----------------------------
-- Records of r_auth_user_role
-- ----------------------------
INSERT INTO "public"."r_auth_user_role" VALUES (1, 1, 1);

-- ----------------------------
-- Table structure for t_auth_role
-- ----------------------------
DROP TABLE IF EXISTS "public"."t_auth_role";
CREATE TABLE "public"."t_auth_role" (
  "id" int4 NOT NULL DEFAULT nextval('t_auth_role_id_seq'::regclass),
  "created_at" timestamp(6),
  "updated_at" timestamp(6),
  "deleted_at" timestamp(6),
  "name" varchar(64) COLLATE "pg_catalog"."default",
  "status" int4,
  "remark" text COLLATE "pg_catalog"."default"
)
;

-- ----------------------------
-- Records of t_auth_role
-- ----------------------------
INSERT INTO "public"."t_auth_role" VALUES (1, '2020-07-27 17:49:10.906336', '2020-07-27 17:50:13.979877', NULL, 'admin', 0, '超级管理员');
INSERT INTO "public"."t_auth_role" VALUES (2, '2020-07-28 09:44:01.013753', '2020-07-28 09:44:01.013753', NULL, 'normal', 0, '普通用户');

-- ----------------------------
-- Table structure for t_auth_user
-- ----------------------------
DROP TABLE IF EXISTS "public"."t_auth_user";
CREATE TABLE "public"."t_auth_user" (
  "id" int4 NOT NULL DEFAULT nextval('t_auth_user_id_seq'::regclass),
  "created_at" timestamp(6),
  "updated_at" timestamp(6),
  "deleted_at" timestamp(6),
  "name" varchar(64) COLLATE "pg_catalog"."default",
  "password" varchar(64) COLLATE "pg_catalog"."default",
  "start_time" timestamp(6),
  "end_time" timestamp(6),
  "status" int4,
  "remark" text COLLATE "pg_catalog"."default",
  "staff_id" int4
)
;

-- ----------------------------
-- Records of t_auth_user
-- ----------------------------
INSERT INTO "public"."t_auth_user" VALUES (1, '2020-07-27 15:39:15.298065', '2020-07-27 15:39:15.298065', NULL, 'root', 'ef487843d2b676dc5e4d8de8cbb3bbd3cd21e0622666fae1c911517056f87996', NULL, NULL, 0, 'I''m a root user.', 1);

-- ----------------------------
-- Table structure for t_sys_city
-- ----------------------------
DROP TABLE IF EXISTS "public"."t_sys_city";
CREATE TABLE "public"."t_sys_city" (
  "id" int4 NOT NULL,
  "name" varchar(32) COLLATE "pg_catalog"."default",
  "province_id" int4,
  "status" int4,
  "remark" text COLLATE "pg_catalog"."default"
)
;

-- ----------------------------
-- Records of t_sys_city
-- ----------------------------
INSERT INTO "public"."t_sys_city" VALUES (370100, '济南市', 370000, 0, '');
INSERT INTO "public"."t_sys_city" VALUES (370200, '青岛市', 370000, 0, NULL);

-- ----------------------------
-- Table structure for t_sys_company
-- ----------------------------
DROP TABLE IF EXISTS "public"."t_sys_company";
CREATE TABLE "public"."t_sys_company" (
  "id" int4 NOT NULL,
  "name" varchar(128) COLLATE "pg_catalog"."default",
  "status" int4,
  "remark" text COLLATE "pg_catalog"."default"
)
;

-- ----------------------------
-- Records of t_sys_company
-- ----------------------------
INSERT INTO "public"."t_sys_company" VALUES (1, '默认公司', 0, '默认公司');

-- ----------------------------
-- Table structure for t_sys_department
-- ----------------------------
DROP TABLE IF EXISTS "public"."t_sys_department";
CREATE TABLE "public"."t_sys_department" (
  "id" int4 NOT NULL,
  "name" varchar(128) COLLATE "pg_catalog"."default",
  "company_id" int4,
  "status" int4,
  "remark" text COLLATE "pg_catalog"."default"
)
;

-- ----------------------------
-- Records of t_sys_department
-- ----------------------------
INSERT INTO "public"."t_sys_department" VALUES (1, '默认部门', 1, 0, '默认部门');

-- ----------------------------
-- Table structure for t_sys_province
-- ----------------------------
DROP TABLE IF EXISTS "public"."t_sys_province";
CREATE TABLE "public"."t_sys_province" (
  "id" int4 NOT NULL,
  "name" varchar(32) COLLATE "pg_catalog"."default",
  "status" int4,
  "remark" text COLLATE "pg_catalog"."default"
)
;

-- ----------------------------
-- Records of t_sys_province
-- ----------------------------
INSERT INTO "public"."t_sys_province" VALUES (100000, '全国', 0, '全国');
INSERT INTO "public"."t_sys_province" VALUES (370000, '山东省', 0, '山东省');

-- ----------------------------
-- Table structure for t_sys_staff
-- ----------------------------
DROP TABLE IF EXISTS "public"."t_sys_staff";
CREATE TABLE "public"."t_sys_staff" (
  "id" int4 NOT NULL DEFAULT nextval('t_sys_staff_id_seq'::regclass),
  "created_at" timestamp(6),
  "updated_at" timestamp(6),
  "deleted_at" timestamp(6),
  "name" varchar(64) COLLATE "pg_catalog"."default",
  "company_id" int4 DEFAULT 1,
  "department_id" int4 DEFAULT 1,
  "post_name" text COLLATE "pg_catalog"."default",
  "birthday" timestamp(6),
  "status" int4,
  "remark" text COLLATE "pg_catalog"."default"
)
;

-- ----------------------------
-- Records of t_sys_staff
-- ----------------------------
INSERT INTO "public"."t_sys_staff" VALUES (1, '2020-07-30 16:30:45.004882', '2020-07-30 16:30:45.004882', NULL, '管理员', 1, 1, '', NULL, 0, '管理员');

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."r_auth_role_func_id_seq"
OWNED BY "public"."r_auth_role_func"."id";
SELECT setval('"public"."r_auth_role_func_id_seq"', 1, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."r_auth_user_role_id_seq"
OWNED BY "public"."r_auth_user_role"."id";
SELECT setval('"public"."r_auth_user_role_id_seq"', 2, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."t_auth_role_id_seq"
OWNED BY "public"."t_auth_role"."id";
SELECT setval('"public"."t_auth_role_id_seq"', 3, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."t_auth_user_id_seq"
OWNED BY "public"."t_auth_user"."id";
SELECT setval('"public"."t_auth_user_id_seq"', 2, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."t_sys_staff_id_seq"
OWNED BY "public"."t_sys_staff"."id";
SELECT setval('"public"."t_sys_staff_id_seq"', 2, true);

-- ----------------------------
-- Primary Key structure for table r_auth_role_func
-- ----------------------------
ALTER TABLE "public"."r_auth_role_func" ADD CONSTRAINT "r_auth_role_func_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table r_auth_user_role
-- ----------------------------
ALTER TABLE "public"."r_auth_user_role" ADD CONSTRAINT "r_auth_user_role_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table t_auth_role
-- ----------------------------
ALTER TABLE "public"."t_auth_role" ADD CONSTRAINT "t_auth_role_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table t_auth_user
-- ----------------------------
ALTER TABLE "public"."t_auth_user" ADD CONSTRAINT "t_auth_user_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table t_sys_city
-- ----------------------------
ALTER TABLE "public"."t_sys_city" ADD CONSTRAINT "t_sys_city_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table t_sys_company
-- ----------------------------
ALTER TABLE "public"."t_sys_company" ADD CONSTRAINT "t_sys_company_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table t_sys_department
-- ----------------------------
ALTER TABLE "public"."t_sys_department" ADD CONSTRAINT "t_sys_department_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table t_sys_province
-- ----------------------------
ALTER TABLE "public"."t_sys_province" ADD CONSTRAINT "t_sys_company_copy1_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table t_sys_staff
-- ----------------------------
ALTER TABLE "public"."t_sys_staff" ADD CONSTRAINT "t_sys_staff_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Foreign Keys structure for table r_auth_role_func
-- ----------------------------
ALTER TABLE "public"."r_auth_role_func" ADD CONSTRAINT "fk_role_id" FOREIGN KEY ("role_id") REFERENCES "public"."t_auth_role" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- ----------------------------
-- Foreign Keys structure for table r_auth_user_role
-- ----------------------------
ALTER TABLE "public"."r_auth_user_role" ADD CONSTRAINT "fk_role_id" FOREIGN KEY ("role_id") REFERENCES "public"."t_auth_role" ("id") ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE "public"."r_auth_user_role" ADD CONSTRAINT "fk_user_id" FOREIGN KEY ("user_id") REFERENCES "public"."t_auth_user" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- ----------------------------
-- Foreign Keys structure for table t_auth_user
-- ----------------------------
ALTER TABLE "public"."t_auth_user" ADD CONSTRAINT "kf_staff_id" FOREIGN KEY ("staff_id") REFERENCES "public"."t_sys_staff" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- ----------------------------
-- Foreign Keys structure for table t_sys_city
-- ----------------------------
ALTER TABLE "public"."t_sys_city" ADD CONSTRAINT "fk_province_id" FOREIGN KEY ("province_id") REFERENCES "public"."t_sys_province" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- ----------------------------
-- Foreign Keys structure for table t_sys_department
-- ----------------------------
ALTER TABLE "public"."t_sys_department" ADD CONSTRAINT "fk_company_id" FOREIGN KEY ("company_id") REFERENCES "public"."t_sys_company" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- ----------------------------
-- Foreign Keys structure for table t_sys_staff
-- ----------------------------
ALTER TABLE "public"."t_sys_staff" ADD CONSTRAINT "fk_company_id" FOREIGN KEY ("company_id") REFERENCES "public"."t_sys_company" ("id") ON DELETE SET DEFAULT ON UPDATE SET DEFAULT;
ALTER TABLE "public"."t_sys_staff" ADD CONSTRAINT "fk_depatment_id" FOREIGN KEY ("department_id") REFERENCES "public"."t_sys_department" ("id") ON DELETE SET DEFAULT ON UPDATE SET DEFAULT;
