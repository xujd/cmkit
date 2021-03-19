/*
 Navicat Premium Data Transfer

 Source Server         : local
 Source Server Type    : PostgreSQL
 Source Server Version : 120003
 Source Host           : localhost:5432
 Source Catalog        : cmkit
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 120003
 File Encoding         : 65001

 Date: 02/08/2020 19:01:49
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
-- Sequence structure for t_res_cabinet_grid_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."t_res_cabinet_grid_id_seq";
CREATE SEQUENCE "public"."t_res_cabinet_grid_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for t_res_cabinet_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."t_res_cabinet_id_seq";
CREATE SEQUENCE "public"."t_res_cabinet_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for t_res_sling_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."t_res_sling_id_seq";
CREATE SEQUENCE "public"."t_res_sling_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for t_res_use_log_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."t_res_use_log_id_seq";
CREATE SEQUENCE "public"."t_res_use_log_id_seq" 
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
-- Records of r_auth_role_func
-- ----------------------------

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
INSERT INTO "public"."r_auth_user_role" VALUES (3, 3, 1);

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
INSERT INTO "public"."t_auth_user" VALUES (3, '2020-08-01 07:20:53.072274', '2020-08-01 23:24:47.584485', NULL, 'test', 'ba9d08c52f3b5e81cc4e2df9222946bd37d5a5adcc06e558beffd63245c5e803', '2020-08-01 00:00:00', '2020-08-06 00:00:00', 0, '222', 1);

-- ----------------------------
-- Table structure for t_res_cabinet
-- ----------------------------
DROP TABLE IF EXISTS "public"."t_res_cabinet";
CREATE TABLE "public"."t_res_cabinet" (
  "id" int4 NOT NULL DEFAULT nextval('t_res_cabinet_id_seq'::regclass),
  "created_at" timestamp(6),
  "updated_at" timestamp(6),
  "deleted_at" timestamp(6),
  "name" varchar(64) COLLATE "pg_catalog"."default",
  "grid_count" int4,
  "location" text COLLATE "pg_catalog"."default",
  "status" int4,
  "remark" text COLLATE "pg_catalog"."default"
)
;

-- ----------------------------
-- Records of t_res_cabinet
-- ----------------------------
INSERT INTO "public"."t_res_cabinet" VALUES (1, '2020-08-01 23:09:12.21249', '2020-08-01 23:19:59.416501', NULL, '测试1', 20, '济南市中区1111号111单元楼南10米', 0, '1111');
INSERT INTO "public"."t_res_cabinet" VALUES (2, '2020-08-01 23:22:03.817087', '2020-08-02 16:27:06.238669', NULL, '2号柜', 10, '南门外大街11号', 0, 'www');
INSERT INTO "public"."t_res_cabinet" VALUES (3, '2020-08-02 16:28:43.098206', '2020-08-02 16:28:43.098206', NULL, '3号柜', 15, '125454', 0, '');

-- ----------------------------
-- Table structure for t_res_cabinet_grid
-- ----------------------------
DROP TABLE IF EXISTS "public"."t_res_cabinet_grid";
CREATE TABLE "public"."t_res_cabinet_grid" (
  "id" int4 NOT NULL DEFAULT nextval('t_res_cabinet_grid_id_seq'::regclass),
  "created_at" timestamp(6),
  "updated_at" timestamp(6),
  "deleted_at" timestamp(6),
  "grid_no" int4,
  "cabinet_id" int4,
  "in_res_id" int4,
  "is_out" int4
)
;

-- ----------------------------
-- Records of t_res_cabinet_grid
-- ----------------------------
INSERT INTO "public"."t_res_cabinet_grid" VALUES (5, '2020-08-02 09:04:51.093326', '2020-08-02 09:07:20.793728', '2020-08-02 09:12:15.591338', 1, 1, 5, 0);
INSERT INTO "public"."t_res_cabinet_grid" VALUES (7, '2020-08-02 09:18:22.891876', '2020-08-02 09:18:22.891876', '2020-08-02 15:13:13.948103', 6, 1, 11, 0);
INSERT INTO "public"."t_res_cabinet_grid" VALUES (4, '2020-08-02 08:43:54.416505', '2020-08-02 15:17:38.193199', NULL, 4, 2, 3, 1);
INSERT INTO "public"."t_res_cabinet_grid" VALUES (8, '2020-08-02 17:12:19.124877', '2020-08-02 17:12:19.124877', NULL, 6, 1, 12, 0);
INSERT INTO "public"."t_res_cabinet_grid" VALUES (1, '2020-08-02 08:35:06.737916', '2020-08-02 17:16:07.332119', NULL, 5, 1, 2, 1);
INSERT INTO "public"."t_res_cabinet_grid" VALUES (6, '2020-08-02 09:13:52.303283', '2020-08-02 18:00:43.107746', NULL, 1, 1, 8, 0);
INSERT INTO "public"."t_res_cabinet_grid" VALUES (2, '2020-08-02 08:35:45.35172', '2020-08-02 18:00:49.109844', NULL, 2, 2, 4, 0);
INSERT INTO "public"."t_res_cabinet_grid" VALUES (9, '2020-08-02 18:27:21.086288', '2020-08-02 18:27:21.086288', NULL, 7, 1, 13, 0);

-- ----------------------------
-- Table structure for t_res_sling
-- ----------------------------
DROP TABLE IF EXISTS "public"."t_res_sling";
CREATE TABLE "public"."t_res_sling" (
  "id" int4 NOT NULL DEFAULT nextval('t_res_sling_id_seq'::regclass),
  "created_at" timestamp(6),
  "updated_at" timestamp(6),
  "deleted_at" timestamp(6),
  "rf_id" varchar(64) COLLATE "pg_catalog"."default",
  "name" varchar(64) COLLATE "pg_catalog"."default",
  "sling_type" int4,
  "max_tonnage" int4,
  "use_status" int4,
  "inspect_status" int4,
  "put_time" text COLLATE "pg_catalog"."default",
  "use_permission" text COLLATE "pg_catalog"."default"
)
;

-- ----------------------------
-- Records of t_res_sling
-- ----------------------------
INSERT INTO "public"."t_res_sling" VALUES (1, '2020-08-01 20:14:53.130602', '2020-08-01 20:19:19.960691', '2020-08-01 20:19:29.589598', 'RF01', '测试111', 1, 2, 1, 1, '2020-08-01 00:00:00', '111');
INSERT INTO "public"."t_res_sling" VALUES (5, '2020-08-02 09:03:29.765448', '2020-08-02 09:07:20.778734', '2020-08-02 09:12:15.589339', '222', '333', 0, 0, 0, 0, '', '');
INSERT INTO "public"."t_res_sling" VALUES (9, '2020-08-02 09:16:28.131168', '2020-08-02 09:16:28.131168', '2020-08-02 09:16:28.150155', '111湖广会馆', '222222344', 0, 0, 0, 0, '', '');
INSERT INTO "public"."t_res_sling" VALUES (11, '2020-08-02 09:18:22.879892', '2020-08-02 09:18:22.879892', '2020-08-02 15:13:13.947106', '啊顺丰到付', '似懂非懂', 0, 0, 0, 0, '', '');
INSERT INTO "public"."t_res_sling" VALUES (3, '2020-08-01 20:20:08.476107', '2020-08-02 15:17:38.202216', NULL, 'RF03', '2324', 2, 2, 2, 1, '2020-08-02 00:00:00', '2223341asdsd');
INSERT INTO "public"."t_res_sling" VALUES (12, '2020-08-02 17:12:19.09789', '2020-08-02 17:12:19.09789', NULL, 'aaa', 'bbbb', 3, 3, 1, 1, '2020-08-03 00:00:00', '');
INSERT INTO "public"."t_res_sling" VALUES (2, '2020-08-01 20:19:46.704101', '2020-08-02 17:16:07.337152', NULL, 'RF02', '1123', 1, 1, 2, 1, '2020-08-03 00:00:00', '222');
INSERT INTO "public"."t_res_sling" VALUES (8, '2020-08-02 09:13:52.28028', '2020-08-02 18:00:43.075754', NULL, '1221212', '123232', 0, 1, 1, 1, '', '');
INSERT INTO "public"."t_res_sling" VALUES (4, '2020-08-01 23:49:26.546588', '2020-08-02 18:00:49.102824', NULL, '111', '222', 0, 2, 1, 1, '', '');
INSERT INTO "public"."t_res_sling" VALUES (13, '2020-08-02 18:27:21.057292', '2020-08-02 18:27:21.057292', NULL, 'hhhh', 'hhhh', 0, 2, 1, 1, '2020-08-02 18:27:14', '');

-- ----------------------------
-- Table structure for t_res_use_log
-- ----------------------------
DROP TABLE IF EXISTS "public"."t_res_use_log";
CREATE TABLE "public"."t_res_use_log" (
  "id" int4 NOT NULL DEFAULT nextval('t_res_use_log_id_seq'::regclass),
  "created_at" timestamp(6),
  "updated_at" timestamp(6),
  "deleted_at" timestamp(6),
  "res_id" int4,
  "take_staff_id" int4,
  "take_time" timestamp(6),
  "return_plan_time" timestamp(6),
  "return_staff_id" int4,
  "return_time" timestamp(6),
  "remark" text COLLATE "pg_catalog"."default",
  "rf_id" varchar(64) COLLATE "pg_catalog"."default",
  "res_name" varchar(64) COLLATE "pg_catalog"."default",
  "take_staff_name" varchar(64) COLLATE "pg_catalog"."default",
  "return_staff_name" varchar(64) COLLATE "pg_catalog"."default"
)
;

-- ----------------------------
-- Records of t_res_use_log
-- ----------------------------
INSERT INTO "public"."t_res_use_log" VALUES (11, '2021-03-19 23:19:52.450714', '2021-03-19 23:56:58.666918', NULL, 13, 3, '2021-03-20 00:00:00', '2021-04-01 00:00:00', 3, '2021-03-19 23:56:58', '111', 'hhhh', 'hhhh', '张三', '张三');
INSERT INTO "public"."t_res_use_log" VALUES (10, '2021-03-19 22:17:53.277037', '2021-03-20 00:00:47.946658', NULL, 12, 3, '2021-03-20 00:00:00', '2021-03-31 00:00:00', 3, '2021-03-20 00:00:47', '222', 'aaa', 'bbbb', '张三', '张三');
INSERT INTO "public"."t_res_use_log" VALUES (14, '2021-03-19 23:45:20.471539', '2021-03-20 00:16:43.694007', NULL, 4, 3, NULL, '2021-03-24 00:00:00', 1, '2021-03-20 00:16:43', 'aaaa', '111', '222', '张三', '管理员');
INSERT INTO "public"."t_res_use_log" VALUES (13, '2021-03-19 23:42:44.373859', '2021-03-19 23:42:44.373859', NULL, 8, 1, NULL, '2021-03-31 00:00:00', 0, NULL, '1111', '1221212', '123232', '管理员', '');
INSERT INTO "public"."t_res_use_log" VALUES (9, '2020-08-02 17:16:07.35411', '2020-08-02 17:16:07.35411', NULL, 2, 3, '2020-08-02 00:00:00', '2020-08-04 00:00:00', 0, NULL, '2222', 'RF02', '1123', '张三', NULL);
INSERT INTO "public"."t_res_use_log" VALUES (12, '2021-03-19 23:20:33.576427', '2021-03-19 23:20:33.576427', NULL, 3, 1, '2021-03-19 00:00:00', '2021-03-20 00:00:00', 0, NULL, '1qqw', 'RF03', '2324', '管理员', NULL);
INSERT INTO "public"."t_res_use_log" VALUES (8, '2020-08-02 15:17:38.209195', '2021-03-19 22:16:13.515499', NULL, 3, 1, '2020-08-02 00:00:00', '2020-08-28 00:00:00', 1, '2021-03-19 22:16:13', '111', 'RF03', '2324', '管理员', '管理员');
INSERT INTO "public"."t_res_use_log" VALUES (7, '2020-08-02 15:14:49.646364', '2020-08-02 15:15:01.86944', NULL, 3, 1, '2020-08-02 00:00:00', '2020-08-12 00:00:00', 1, '2020-08-02 15:15:01', '111acdf', 'RF03', '2324', '管理员', '管理员');
INSERT INTO "public"."t_res_use_log" VALUES (5, '2020-08-02 15:00:22.395874', '2020-08-02 15:00:53.202961', NULL, 2, 1, '2020-08-12 00:00:00', '2020-08-28 00:00:00', 1, '2020-08-02 15:00:53', '2154hello--15:00', 'RF02', '1123', '管理员', '管理员');
INSERT INTO "public"."t_res_use_log" VALUES (4, '2020-08-02 14:55:54.468577', '2020-08-02 14:59:59.565194', NULL, 4, 1, '2020-08-02 00:00:00', '2020-08-05 00:00:00', 1, '2020-08-02 14:59:59', '2222qqqqqqqqqqqqqqqqqq', '111', '222', '管理员', '管理员');
INSERT INTO "public"."t_res_use_log" VALUES (3, '2020-08-02 11:38:17.112042', '2020-08-02 14:57:55.732006', NULL, 4, 1, '2020-08-03 00:00:00', '2020-08-06 00:00:00', 1, '2020-08-02 14:57:55', '2222aaaaaaaaaaaaaaaaaaa', '111', '222', '管理员', '管理员');
INSERT INTO "public"."t_res_use_log" VALUES (6, '2020-08-02 15:01:23.991075', '2020-08-02 15:41:57.831611', NULL, 4, 1, '2020-08-02 00:00:00', '2020-08-03 00:00:00', 3, '2020-08-02 15:41:57', '21343', '111', '222', '管理员', '张三');

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
-- Table structure for t_sys_dict
-- ----------------------------
DROP TABLE IF EXISTS "public"."t_sys_dict";
CREATE TABLE "public"."t_sys_dict" (
  "id" int4 NOT NULL,
  "name" varchar(32) COLLATE "pg_catalog"."default",
  "type" varchar(32) COLLATE "pg_catalog"."default",
  "note" varchar(64) COLLATE "pg_catalog"."default",
  "scene" varchar(32) COLLATE "pg_catalog"."default",
  "key" int4
)
;

-- ----------------------------
-- Records of t_sys_dict
-- ----------------------------
INSERT INTO "public"."t_sys_dict" VALUES (1, '吊具类型1', 'SLING_TYPE', '吊具类型', 'Sling', 1);
INSERT INTO "public"."t_sys_dict" VALUES (2, '吊具类型2', 'SLING_TYPE', '吊具类型', 'Sling', 2);
INSERT INTO "public"."t_sys_dict" VALUES (3, '吊具类型3', 'SLING_TYPE', '吊具类型', 'Sling', 3);
INSERT INTO "public"."t_sys_dict" VALUES (4, '1吨', 'TON_TYPE', '吨位', 'Sling', 1);
INSERT INTO "public"."t_sys_dict" VALUES (5, '10吨', 'TON_TYPE', '吨位', 'Sling', 2);
INSERT INTO "public"."t_sys_dict" VALUES (6, '100吨', 'TON_TYPE', '吨位', 'Sling', 3);
INSERT INTO "public"."t_sys_dict" VALUES (9, '正常', 'INSPECT_STATUS_TYPE', '点检状态', 'Sling', 1);
INSERT INTO "public"."t_sys_dict" VALUES (10, '异常', 'INSPECT_STATUS_TYPE', '点检状态', 'Sling', 2);
INSERT INTO "public"."t_sys_dict" VALUES (7, '在库', 'USE_STATUS_TYPE', '使用状态', 'Sling', 1);
INSERT INTO "public"."t_sys_dict" VALUES (11, '不可用', 'USE_STATUS_TYPE', '使用状态', 'Sling', 3);
INSERT INTO "public"."t_sys_dict" VALUES (8, '借出', 'USE_STATUS_TYPE', '使用状态', 'Sling', 2);

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
INSERT INTO "public"."t_sys_staff" VALUES (3, '2020-08-02 15:41:47.817811', '2020-08-02 15:41:47.817811', NULL, '张三', 1, 1, '', '1990-01-01 00:00:00', 0, '');

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."r_auth_role_func_id_seq"
OWNED BY "public"."r_auth_role_func"."id";
SELECT setval('"public"."r_auth_role_func_id_seq"', 2, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."r_auth_user_role_id_seq"
OWNED BY "public"."r_auth_user_role"."id";
SELECT setval('"public"."r_auth_user_role_id_seq"', 4, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."t_auth_role_id_seq"
OWNED BY "public"."t_auth_role"."id";
SELECT setval('"public"."t_auth_role_id_seq"', 4, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."t_auth_user_id_seq"
OWNED BY "public"."t_auth_user"."id";
SELECT setval('"public"."t_auth_user_id_seq"', 4, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."t_res_cabinet_grid_id_seq"
OWNED BY "public"."t_res_cabinet_grid"."id";
SELECT setval('"public"."t_res_cabinet_grid_id_seq"', 10, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."t_res_cabinet_id_seq"
OWNED BY "public"."t_res_cabinet"."id";
SELECT setval('"public"."t_res_cabinet_id_seq"', 4, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."t_res_sling_id_seq"
OWNED BY "public"."t_res_sling"."id";
SELECT setval('"public"."t_res_sling_id_seq"', 14, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."t_res_use_log_id_seq"
OWNED BY "public"."t_res_use_log"."id";
SELECT setval('"public"."t_res_use_log_id_seq"', 10, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."t_sys_staff_id_seq"
OWNED BY "public"."t_sys_staff"."id";
SELECT setval('"public"."t_sys_staff_id_seq"', 4, true);

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
-- Primary Key structure for table t_res_cabinet
-- ----------------------------
ALTER TABLE "public"."t_res_cabinet" ADD CONSTRAINT "t_res_cabinet_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table t_res_cabinet_grid
-- ----------------------------
ALTER TABLE "public"."t_res_cabinet_grid" ADD CONSTRAINT "t_res_cabinet_grid_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table t_res_sling
-- ----------------------------
ALTER TABLE "public"."t_res_sling" ADD CONSTRAINT "t_res_sling_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table t_res_use_log
-- ----------------------------
ALTER TABLE "public"."t_res_use_log" ADD CONSTRAINT "t_res_use_log_pkey" PRIMARY KEY ("id");

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
-- Primary Key structure for table t_sys_dict
-- ----------------------------
ALTER TABLE "public"."t_sys_dict" ADD CONSTRAINT "t_sys_dict_pkey" PRIMARY KEY ("id");

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
