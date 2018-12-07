-- 创建书籍
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;
-- ----------------------------
-- Table structure for history_book_25
-- ----------------------------
DROP TABLE IF EXISTS `history_book_25`;
CREATE TABLE `history_book_25`  (
  `id` int(11)  NOT NULL AUTO_INCREMENT,
  `book_name` varchar(100) NULL DEFAULT '' ,
  `book_index` varchar(255) NULL DEFAULT '',
  `book_content` varchar(255) NULL default '',
  `book_key` varchar(100) NULL default '',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for history_book_25_detail
-- ----------------------------
DROP TABLE IF EXISTS `history_book_25_detail`;
CREATE TABLE `history_book_25_detail`  (
  `id` int(11)  NOT NULL AUTO_INCREMENT,
  `book_name` varchar(100) NULL DEFAULT '' ,
  `book_author` varchar(255) NULL DEFAULT '',
  `book_desc` varchar(1000) NULL default '',
  `book_key` varchar(100) NULL default '',
  PRIMARY KEY (`id`)
)  ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;