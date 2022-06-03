SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

drop table IF EXISTS `roles`;
create TABLE `roles`
(
    `role_id`      int(11)                                               NOT NULL AUTO_INCREMENT,
    `role_name`    varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
    `role_pid`     int(11)                                               NULL DEFAULT 0,
    `role_comment` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
    `tenant_id`    int(11)                                               NULL DEFAULT 0,
    PRIMARY KEY (`role_id`) USING BTREE,
    INDEX `TenantId` (`tenant_id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 11
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_bin
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of roles
-- ----------------------------
insert into `roles`
VALUES (2, 'deptadmin', 0, '部门管理员', 1);
insert into `roles`
VALUES (3, 'deptselecter', 7, '部门查询员', 1);
insert into `roles`
VALUES (7, 'deptupdater', 2, '部门编辑员', 1);
insert into `roles`
VALUES (8, 'deptadmin', 0, '部门管理员', 2);
insert into `roles`
VALUES (9, 'deptselecter', 10, '部门查询员', 2);
insert into `roles`
VALUES (10, 'deptupdater', 8, '部门编辑员', 2);

-- ----------------------------
-- Table structure for router_roles
-- ----------------------------
drop table IF EXISTS `router_roles`;
create TABLE `router_roles`
(
    `id`        int(11) NOT NULL AUTO_INCREMENT,
    `router_id` int(11) NULL DEFAULT NULL,
    `role_id`   int(11) NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE INDEX `router_id` (`router_id`, `role_id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 17
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_bin
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of router_roles
-- ----------------------------
insert into `router_roles`
VALUES (13, 1, 3);
insert into `router_roles`
VALUES (15, 1, 9);
insert into `router_roles`
VALUES (14, 3, 7);
insert into `router_roles`
VALUES (16, 3, 10);

-- ----------------------------
-- Table structure for routers
-- ----------------------------
drop table IF EXISTS `routers`;
create TABLE `routers`
(
    `r_id`     int(11)                                                NOT NULL AUTO_INCREMENT,
    `r_name`   varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin  NULL DEFAULT NULL,
    `r_uri`    varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
    `r_method` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin  NULL DEFAULT NULL,
    `r_status` tinyint(4)                                             NULL DEFAULT NULL,
    PRIMARY KEY (`r_id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 4
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_bin
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of routers
-- ----------------------------
insert into `routers`
VALUES (1, '部门列表', '/dept', 'GET', 1);
insert into `routers`
VALUES (3, '新增部门', '/dept', 'POST', 1);

-- ----------------------------
-- Table structure for tenants
-- ----------------------------
drop table IF EXISTS `tenants`;
create TABLE `tenants`
(
    `tenant_id`   int(11)                                               NOT NULL AUTO_INCREMENT,
    `tenant_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
    PRIMARY KEY (`tenant_id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 3
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_bin
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tenants
-- ----------------------------
insert into `tenants`
VALUES (1, 'domain1');
insert into `tenants`
VALUES (2, 'domain2');

-- ----------------------------
-- Table structure for user_roles
-- ----------------------------
drop table IF EXISTS `user_roles`;
create TABLE `user_roles`
(
    `id`      int(11) NOT NULL AUTO_INCREMENT,
    `user_id` int(11) NOT NULL,
    `role_id` int(11) NOT NULL,
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE INDEX `user_id` (`user_id`, `role_id`) USING BTREE,
    INDEX `role_id` (`role_id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 3
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_bin
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_roles
-- ----------------------------
insert into `user_roles`
VALUES (1, 1, 3);
insert into `user_roles`
VALUES (2, 2, 10);

-- ----------------------------
-- Table structure for users
-- ----------------------------
drop table IF EXISTS `users`;
create TABLE `users`
(
    `user_id`   int(11)                                               NOT NULL AUTO_INCREMENT,
    `user_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
    `tenant_id` int(11)                                               NULL DEFAULT NULL,
    PRIMARY KEY (`user_id`) USING BTREE,
    INDEX `tenant_id` (`tenant_id`) USING BTREE,
    CONSTRAINT `users_ibfk_1` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`tenant_id`) ON delete RESTRICT ON update RESTRICT
) ENGINE = InnoDB
  AUTO_INCREMENT = 3
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_bin
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of users
-- ----------------------------
insert into `users`
VALUES (1, 'dev', 1);
insert into `users`
VALUES (2, 'tester', 2);

SET FOREIGN_KEY_CHECKS = 1;
