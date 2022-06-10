CREATE DATABASE IF NOT EXISTS notify_sms_xuanwu
    DEFAULT CHARSET utf8mb4
    COLLATE utf8mb4_general_ci;

USE notify_sms_xuanwu;
SET NAMES utf8mb4;
-- ----------------------------
-- Table structure for mail_content
-- ----------------------------
DROP TABLE IF EXISTS `mail_content`;
CREATE TABLE `mail_content`
(
    `id`          int(11)      NOT NULL AUTO_INCREMENT,
    `sender`      varchar(255) NOT NULL,
    `to`          varchar(255) NOT NULL,
    `batch_id`    varchar(255) NOT NULL,
    `send_msg_id` varchar(255) NOT NULL,
    `content`     text         NOT NULL,
    `time`        bigint(20)   NOT NULL,
    PRIMARY KEY (`id`),
    KEY `IDX_msg_id` (`batch_id`, `send_msg_id`) USING BTREE,
    KEY `IDX_time` (`time`) USING BTREE,
    KEY `IDX_sender` (`sender`) USING BTREE,
    KEY `IDX_to` (`to`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;


-- ----------------------------
-- Table structure for mail_state
-- ----------------------------
DROP TABLE IF EXISTS `mail_state`;
CREATE TABLE `mail_state`
(
    `id`            int(11)      NOT NULL AUTO_INCREMENT,
    `to`            varchar(255) NOT NULL,
    `batch_id`      varchar(255) NOT NULL,
    `send_msg_id`   varchar(255) NOT NULL,
    `recv_msg_id`   varchar(255) NOT NULL DEFAULT '',
    `state`         int(11)      NOT NULL DEFAULT -1,
    `submit_time`   bigint(20)   NOT NULL DEFAULT -1,
    `done_time`     bigint(20)   NOT NULL DEFAULT -1,
    `origin_result` varchar(255) NOT NULL DEFAULT '',
    PRIMARY KEY (`id`),
    KEY `IDX_msg_id` (`batch_id`, `send_msg_id`, `recv_msg_id`) USING BTREE,
    KEY `IDX_time` (`submit_time`, `done_time`) USING BTREE,
    KEY `IDX_to` (`to`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

