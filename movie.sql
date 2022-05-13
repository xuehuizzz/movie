/*
 Navicat Premium Data Transfer

 Source Server         : mysql
 Source Server Type    : MySQL
 Source Server Version : 80027
 Source Host           : localhost:3306
 Source Schema         : fyouku

 Target Server Type    : MySQL
 Target Server Version : 80027
 File Encoding         : 65001

 Date: 13/05/2022 19:55:54
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for advert
-- ----------------------------
DROP TABLE IF EXISTS `advert`;
CREATE TABLE `advert`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '广告标题',
  `sub_title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '广告副标题',
  `channel_id` int NULL DEFAULT 0 COMMENT '所属频道 0=首页',
  `img` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '广告图片',
  `sort` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '排序',
  `add_time` datetime NULL DEFAULT '0000-00-00 00:00:00' COMMENT '添加时间',
  `url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '跳转地址',
  `status` tinyint NULL DEFAULT 0 COMMENT '0=关闭 1=开启',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of advert
-- ----------------------------
INSERT INTO `advert` VALUES (1, '为中华之崛起而读书', '哈哈哈', 1, '/static/data/new/000.jpg', '1', '2022-05-02 19:12:28', '/show?id=9', 1);

-- ----------------------------
-- Table structure for aliyun_video
-- ----------------------------
DROP TABLE IF EXISTS `aliyun_video`;
CREATE TABLE `aliyun_video`  (
  `id` int NULL DEFAULT NULL,
  `video_id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `add_time` datetime NULL DEFAULT '0000-00-00 00:00:00',
  `log` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of aliyun_video
-- ----------------------------

-- ----------------------------
-- Table structure for barrage
-- ----------------------------
DROP TABLE IF EXISTS `barrage`;
CREATE TABLE `barrage`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `content` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '评论内容',
  `add_time` datetime NULL DEFAULT '0000-00-00 00:00:00' COMMENT '评论时间',
  `user_id` int NULL DEFAULT 0 COMMENT '评论用户',
  `status` tinyint NULL DEFAULT 0 COMMENT '状态0=未审核 1=审核通过',
  `current_time` int NULL DEFAULT 1 COMMENT '视频当前播放时间',
  `episodes_id` int NULL DEFAULT 0 COMMENT '评论视频',
  `video_id` int NULL DEFAULT 0 COMMENT '所属视频',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 55 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of barrage
-- ----------------------------
INSERT INTO `barrage` VALUES (31, '这个面看着真好吃', '2022-05-04 13:07:44', 1, 1, 196, 1, 1);
INSERT INTO `barrage` VALUES (32, '厉害', '2022-05-04 13:07:44', 1, 1, 15, 1, 1);
INSERT INTO `barrage` VALUES (33, '佐助貌似没有准备时间', '2022-05-04 13:07:44', 1, 1, 28, 1, 1);
INSERT INTO `barrage` VALUES (34, '鸣人加油', '2022-05-04 13:07:44', 1, 1, 50, 1, 1);
INSERT INTO `barrage` VALUES (35, '这个太牛了', '2022-05-04 13:07:44', 1, 1, 70, 1, 1);
INSERT INTO `barrage` VALUES (36, '有机会一定吃一次', '2022-05-04 13:07:44', 1, 1, 138, 1, 1);
INSERT INTO `barrage` VALUES (37, '满满的回忆呀', '2022-05-04 13:07:44', 1, 1, 15, 1, 1);
INSERT INTO `barrage` VALUES (38, '满满的回忆', '2022-05-04 13:07:44', 1, 1, 10, 1, 1);
INSERT INTO `barrage` VALUES (39, '好厉害呀', '2022-05-04 13:07:44', 1, 1, 25, 1, 1);
INSERT INTO `barrage` VALUES (40, '佐助', '2022-05-04 13:07:44', 1, 1, 36, 1, 1);
INSERT INTO `barrage` VALUES (54, '1', '2022-05-05 16:18:21', 1, 1, 6, 1, 1);

-- ----------------------------
-- Table structure for channel
-- ----------------------------
DROP TABLE IF EXISTS `channel`;
CREATE TABLE `channel`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '频道名',
  `add_time` datetime NULL DEFAULT NULL COMMENT '添加时间',
  `status` tinyint NULL DEFAULT NULL COMMENT '状态',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of channel
-- ----------------------------
INSERT INTO `channel` VALUES (1, '动漫', '2022-05-02 19:12:28', 1);
INSERT INTO `channel` VALUES (2, '电影', '2022-05-02 19:12:28', 1);
INSERT INTO `channel` VALUES (3, '体育', '2022-05-02 19:12:28', 1);
INSERT INTO `channel` VALUES (4, '音乐', '2022-05-05 15:28:38', 1);

-- ----------------------------
-- Table structure for channel_region
-- ----------------------------
DROP TABLE IF EXISTS `channel_region`;
CREATE TABLE `channel_region`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '分类名称',
  `status` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '状态',
  `add_time` datetime NULL DEFAULT NULL COMMENT '添加时间',
  `channel_id` int NULL DEFAULT NULL COMMENT '所属频道',
  `sort` int NULL DEFAULT 0 COMMENT '排序',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of channel_region
-- ----------------------------
INSERT INTO `channel_region` VALUES (1, '日本', '1', '2022-05-02 19:12:28', 1, 2);
INSERT INTO `channel_region` VALUES (2, '中国大陆', '1', '2022-05-02 19:12:28', 1, 3);
INSERT INTO `channel_region` VALUES (3, '美国', '1', '2022-05-02 19:12:28', 1, 1);
INSERT INTO `channel_region` VALUES (4, '韩国', '1', '2022-05-05 15:30:12', 1, 4);
INSERT INTO `channel_region` VALUES (5, '中国香港', '1', '2022-05-05 15:30:58', 1, 5);

-- ----------------------------
-- Table structure for channel_type
-- ----------------------------
DROP TABLE IF EXISTS `channel_type`;
CREATE TABLE `channel_type`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '分类名称',
  `status` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '状态',
  `add_time` datetime NULL DEFAULT '0000-00-00 00:00:00' COMMENT '添加时间',
  `channel_id` int NULL DEFAULT NULL COMMENT '所属频道',
  `sort` int NULL DEFAULT NULL COMMENT '排序',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of channel_type
-- ----------------------------
INSERT INTO `channel_type` VALUES (1, '少女', '1', '2022-05-02 19:12:28', 1, 6);
INSERT INTO `channel_type` VALUES (2, '热血', '1', '2022-05-02 19:12:28', 1, 5);
INSERT INTO `channel_type` VALUES (3, '青春', '1', '2022-05-02 19:12:28', 1, 4);
INSERT INTO `channel_type` VALUES (4, '社会', '1', '2022-05-02 19:12:28', 1, 3);
INSERT INTO `channel_type` VALUES (5, '科幻', '1', '2022-05-02 19:12:28', 1, 2);
INSERT INTO `channel_type` VALUES (6, '搞笑', '1', '2022-05-02 19:12:28', 1, 1);

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `content` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '评论内容',
  `add_time` datetime NULL DEFAULT '0000-00-00 00:00:00' COMMENT '评论时间',
  `user_id` int NULL DEFAULT 0 COMMENT '评论用户',
  `status` tinyint NULL DEFAULT 0 COMMENT '状态0=未审核 1=审核通过',
  `stamp` tinyint NULL DEFAULT 0 COMMENT '盖章1=热评2=公告',
  `praise_count` int NULL DEFAULT 0 COMMENT '点赞数',
  `episodes_id` int NULL DEFAULT NULL COMMENT '评论视频',
  `video_id` int NULL DEFAULT 0 COMMENT '所属视频',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 48 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of comment
-- ----------------------------
INSERT INTO `comment` VALUES (1, '好期待下一集，净眼马上又要看到了，浦式看到了肯定会说出它的秘密', '2022-05-04 19:12:24', 1, 1, 0, 1, 1, 1);
INSERT INTO `comment` VALUES (2, '博人智商很高，在忍者学校成绩都满分。悟性也都是很高螺旋丸很快就能掌握，但不知道为什么博人在战斗的时候总分不清情况。什么时候该打，什么时候该跑总是表现得一股脑。无论敌人什么实力总是要向前冲，一点也表现不出博人的战术分析。按理说博人比鸣人更有理智，但表现得总是差强人意。我偶尔就看的十分尴尬。', '2022-05-02 19:12:28', 1, 1, 1, 186, 1, 1);
INSERT INTO `comment` VALUES (3, '考试来啦！轻松看剧，认真答题', '2022-05-02 19:12:28', 1, 1, 2, 3459, 1, 1);
INSERT INTO `comment` VALUES (4, '太精彩了', '2022-05-02 19:12:28', 1, 1, 0, 0, 1, 1);
INSERT INTO `comment` VALUES (13, '火影忍者太好看了', '2022-05-02 19:12:28', 1, 1, 0, 0, 1, 1);
INSERT INTO `comment` VALUES (14, '满满的回忆呀', '2022-05-02 19:12:28', 1, 1, 0, 0, 1, 1);

-- ----------------------------
-- Table structure for message
-- ----------------------------
DROP TABLE IF EXISTS `message`;
CREATE TABLE `message`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `content` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '消息内容',
  `add_time` datetime NULL DEFAULT '0000-00-00 00:00:00' COMMENT '添加时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 43 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of message
-- ----------------------------
INSERT INTO `message` VALUES (29, 'aaa', '2022-05-02 19:12:28');
INSERT INTO `message` VALUES (30, 'aaa', '2022-05-02 19:12:28');
INSERT INTO `message` VALUES (31, 'haha', '2022-05-02 19:12:28');
INSERT INTO `message` VALUES (32, 'cvcv', '2022-05-02 19:12:28');
INSERT INTO `message` VALUES (33, 'bnbn', '2022-05-02 19:12:28');
INSERT INTO `message` VALUES (43, 'hello everyone', '2022-05-03 09:53:30');

-- ----------------------------
-- Table structure for message_user
-- ----------------------------
DROP TABLE IF EXISTS `message_user`;
CREATE TABLE `message_user`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `message_id` int NULL DEFAULT 0 COMMENT '所属消息',
  `user_id` int NULL DEFAULT 0 COMMENT '所属用户',
  `add_time` datetime NULL DEFAULT '0000-00-00 00:00:00' COMMENT '添加时间',
  `status` tinyint NULL DEFAULT 0 COMMENT '状态 1=展示',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 49 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of message_user
-- ----------------------------
INSERT INTO `message_user` VALUES (29, 30, 1, '2022-05-02 19:12:28', 1);
INSERT INTO `message_user` VALUES (30, 30, 1, '2022-05-02 19:12:28', 1);
INSERT INTO `message_user` VALUES (31, 31, 1, '2022-05-02 19:12:28', 1);
INSERT INTO `message_user` VALUES (32, 32, 1, '2022-05-02 19:12:28', 1);
INSERT INTO `message_user` VALUES (33, 33, 1, '2022-05-02 19:12:28', 1);
INSERT INTO `message_user` VALUES (34, 33, 1, '2022-05-02 19:12:28', 1);
INSERT INTO `message_user` VALUES (35, 33, 1, '2022-05-02 19:12:28', 1);
INSERT INTO `message_user` VALUES (36, 34, 1, '2022-05-02 19:12:28', 1);
INSERT INTO `message_user` VALUES (37, 35, 1, '2022-05-02 19:12:28', 1);
INSERT INTO `message_user` VALUES (38, 35, 1, '2022-05-02 19:12:28', 1);
INSERT INTO `message_user` VALUES (39, 36, 1, '2022-05-02 19:12:28', 1);
INSERT INTO `message_user` VALUES (40, 36, 1, '2022-05-02 19:12:28', 1);
INSERT INTO `message_user` VALUES (41, 38, 1, '2022-05-02 19:12:28', 1);
INSERT INTO `message_user` VALUES (42, 38, 1, '2022-05-02 19:12:28', 1);
INSERT INTO `message_user` VALUES (43, 39, 1, '2022-05-02 19:12:28', 1);
INSERT INTO `message_user` VALUES (44, 39, 1, '2022-05-02 19:12:28', 1);
INSERT INTO `message_user` VALUES (45, 39, 1, '2022-05-02 19:12:28', 1);
INSERT INTO `message_user` VALUES (46, 39, 1, '2022-05-02 19:12:28', 1);
INSERT INTO `message_user` VALUES (47, 43, 1, '2022-05-03 09:53:30', 1);
INSERT INTO `message_user` VALUES (48, 43, 1, '2022-05-03 09:53:30', 1);
INSERT INTO `message_user` VALUES (49, 43, 1, '2022-05-03 09:53:30', 1);

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户名',
  `password` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '密码',
  `add_time` datetime NULL DEFAULT '0000-00-00 00:00:00' COMMENT '注册时间',
  `status` tinyint NULL DEFAULT 1 COMMENT '状态',
  `mobile` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '手机号',
  `avatar` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '头像',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 21 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, 'xuehui', 'e10adc3949ba59abbe56e057f20f883e', '2022-04-26 19:12:28', 1, '15856236710', '/static/data/user/1.jpg');

-- ----------------------------
-- Table structure for video
-- ----------------------------
DROP TABLE IF EXISTS `video`;
CREATE TABLE `video`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '视频ID',
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '标题',
  `sub_title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '副标题',
  `status` tinyint NULL DEFAULT 0 COMMENT '状态,-1=下架,0=审核中,1=正常',
  `add_time` datetime NULL DEFAULT '0000-00-00 00:00:00' COMMENT '添加时间',
  `img` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '横版缩略图',
  `img1` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '竖版缩略图',
  `channel_id` int NULL DEFAULT 0 COMMENT '所属频道',
  `type_id` int NULL DEFAULT 0 COMMENT '频道类型ID',
  `region_id` int NULL DEFAULT 0 COMMENT '地区ID',
  `user_id` int NULL DEFAULT 0 COMMENT '所属用户',
  `episodes_count` int NULL DEFAULT 0 COMMENT '集数',
  `episodes_update_time` datetime NULL DEFAULT '0000-00-00 00:00:00' COMMENT '集数更新时间',
  `is_end` tinyint NULL DEFAULT 0 COMMENT '是否完结0=未完结 1=已完结',
  `is_hot` tinyint NULL DEFAULT 0 COMMENT '1=正在热播',
  `is_recommend` tinyint NULL DEFAULT 0 COMMENT '1=推荐 0=否',
  `comment` int NULL DEFAULT 0 COMMENT '评论数',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 39 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of video
-- ----------------------------
INSERT INTO `video` VALUES (1, 'Taylor Swift', '泰勒·斯威夫特', 1, '2022-05-02 19:12:28', '/static/data/new/meimei.png', '/static/data/new/mei.png', 1, 2, 1, 1, 10, '2022-05-02 19:12:28', 1, 1, 1, 10077218);
INSERT INTO `video` VALUES (2, '阿衰 第二季', '怕踢学校搞笑故事再次归来', 1, '2022-05-02 19:12:28', '/static/data/new/ashuai.jfif', '/static/data/new/ashuai1.jfif', 1, 2, 2, 1, 15, '2022-05-02 19:12:28', 0, 1, 1, 17860);
INSERT INTO `video` VALUES (3, '名侦探柯南', '万年小学生的推理生涯', 1, '2022-05-02 19:12:28', '/static/data/new/kenan.jfif', '/static/data/new/3a.png', 1, 3, 1, 1, 1021, '2022-05-02 19:12:28', 0, 1, 0, 29876);
INSERT INTO `video` VALUES (4, '博人传 火影忍者新时代', '鸣人之子续写忍者传奇', 1, '2022-05-02 19:12:28', '/static/data/new/borenzhuan.jfif', '/static/data/new/borenzhuan1.jfif', 1, 2, 1, 1, 137, '2022-05-02 19:12:28', 0, 1, 0, 18976);
INSERT INTO `video` VALUES (5, '一拳超人 第二季', '一拳埼玉热血回归', 1, '2022-05-02 19:12:28', '/static/data/new/yiquanchaoren.jfif', '/static/data/new/5a.png', 1, 2, 1, 1, 13, '2022-05-02 19:12:28', 0, 1, 0, 98761);
INSERT INTO `video` VALUES (6, '排球少年 第四季', '排球部迎接更艰巨的挑战', 1, '2022-05-02 19:12:28', '/static/data/new/paiqiu4.jfif', '/static/data/new/paiqiu4-1.jfif', 1, 3, 1, 1, 5, '2022-05-02 19:12:28', 0, 1, 0, 23564);
INSERT INTO `video` VALUES (7, '秦时明月之沧海横流', '十年国漫十年秦时', 1, '2022-05-02 19:12:28', '/static/data/new/canghaihengliu.jfif', '/static/data/new/canghaihengliu1.jfif', 1, 2, 2, 1, 1, '2022-05-02 19:12:28', 0, 1, 0, 0);
INSERT INTO `video` VALUES (8, '猫和老鼠', '汤姆和杰瑞的搞笑日常', 1, '2022-05-02 19:12:28', '/static/data/new/maohelaoshu.jfif', '/static/data/new/maohelaoshu1.jfif', 1, 1, 3, 1, 140, '2022-05-02 19:12:28', 1, 1, 0, 0);
INSERT INTO `video` VALUES (10, '蜡笔小新 第二季', '人小鬼大的野原新之助', 1, '2022-05-02 19:12:28', '/static/data/new/labixiaoxin.jfif', '/static/data/new/labixiaoxin.jpg', 1, 1, 1, 1, 873, '2022-05-02 19:12:28', 1, 1, 1, 0);
INSERT INTO `video` VALUES (14, '网球王子', '运动美少年热血故事', 1, '2022-05-02 19:12:28', '/static/data/new/wangqiuwangzi.jfif', '/static/data/new/wangqiuwangzi1.jfif', 1, 1, 1, 1, 178, '2022-05-02 19:12:28', 1, 1, 1, 0);
INSERT INTO `video` VALUES (15, '守护甜心 第一季', '能实现愿望的魔法蛋', 1, '2022-05-02 19:12:28', '/static/data/new/shouhutianxin.jfif', '/static/data/new/shouhutianxin.jfif', 1, 1, 1, 1, 51, '2022-05-02 19:12:28', 1, 1, 1, 0);
INSERT INTO `video` VALUES (16, '精灵宝可梦 第四季', '星球上不可思议的生物', 1, '2022-05-02 19:12:28', '/static/data/new/jinglingbaokemeng.jfif', '/static/data/new/jinglingbaokemeng1.jfif', 1, 1, 1, 1, 144, '2022-05-02 19:12:28', 1, 1, 1, 0);
INSERT INTO `video` VALUES (17, '中华小子', '中华小子', 1, '2022-05-02 19:12:28', '/static/data/new/zhonghuaxiaozi.jfif', '/static/data/new/zhonghuaxiaozi1.jfif', 1, 2, 2, 1, 366, '2022-05-02 19:12:28', 1, 1, 1, 0);
INSERT INTO `video` VALUES (18, '犬夜叉', '寻找四魂碎片的旅程', 1, '2022-05-02 19:12:28', '/static/data/new/quanyecha.jfif', '/static/data/new/quanyecha1.jpg', 1, 1, 1, 1, 167, '2022-05-02 19:12:28', 1, 1, 1, 0);
INSERT INTO `video` VALUES (19, '伍六七之最强发型师', '强着对手不断现身', 1, '2022-05-02 19:12:28', '/static/data/new/567.jpg', '/static/data/new/567-1.jfif', 1, 3, 3, 1, 10, '2022-05-02 19:12:28', 1, 1, 0, 0);
INSERT INTO `video` VALUES (21, '海贼王', '路飞再踏热血征程', 1, '2022-05-02 19:12:28', '/static/data/new/haizeiwang.jfif', '/static/data/new/haizeiwang1.jfif', 1, 1, 1, 1, 920, '2022-05-02 19:12:28', 0, 0, 0, 0);
INSERT INTO `video` VALUES (22, '元气少女缘结神 第一季', '软萌少女结缘狐神', 1, '2022-05-02 19:12:28', '/static/data/new/22.png', '/static/data/new/22a.png', 1, 1, 1, 1, 13, '2022-05-02 19:12:28', 1, 0, 0, 9988);
INSERT INTO `video` VALUES (23, '爱神巧克力', '高中生成后宫之王', 1, '2022-05-02 19:12:28', '/static/data/new/23.png', '/static/data/new/23a.png', 1, 1, 1, 1, 15, '2022-05-02 19:12:28', 1, 0, 0, 9678);
INSERT INTO `video` VALUES (24, '魔法使的新娘', '骨叔新娘养成计划', 1, '2022-05-02 19:12:28', '/static/data/new/24.png', '/static/data/new/24a.png', 1, 1, 1, 1, 24, '2022-05-02 19:12:28', 1, 0, 0, 8080);
INSERT INTO `video` VALUES (25, '悬崖上的金鱼姬', '宫崎骏版小美人鱼', 1, '2022-05-02 19:12:28', '/static/data/new/jinyuji.jfif', '/static/data/new/jinyuji1.jpg', 1, 1, 1, 1, 1, '2022-05-02 19:12:28', 1, 0, 0, 5674);
INSERT INTO `video` VALUES (26, '飞天小女警', '超能力少女卫东京', 1, '2022-05-02 19:12:28', '/static/data/new/26.png', '/static/data/new/26a.png', 1, 1, 1, 1, 52, '2022-05-02 19:12:28', 1, 0, 0, 8735);
INSERT INTO `video` VALUES (27, '一起来看流星雨', '灰姑娘的校园爱情', 1, '2022-05-02 19:12:28', '/static/data/new/27.png', '/static/data/new/27a.png', 1, 1, 1, 1, 52, '2022-05-02 19:12:28', 1, 0, 0, 3476);
INSERT INTO `video` VALUES (28, '罗小黑战记', '同名大电影热映中', 1, '2022-05-02 19:12:28', '/static/data/new/28.png', '/static/data/new/28a.png', 1, 2, 2, 1, 28, '2022-05-02 19:12:28', 1, 0, 1, 0);
INSERT INTO `video` VALUES (29, '秦时明月之君临天下', '沧海横流 江湖再见', 1, '2022-05-02 19:12:28', '/static/data/new/29.png', '/static/data/new/29a.png', 1, 2, 2, 1, 75, '2022-05-02 19:12:28', 1, 0, 1, 0);
INSERT INTO `video` VALUES (30, '镇魂街', '吸纳亡灵镇压恶灵之地', 1, '2022-05-02 19:12:28', '/static/data/new/zhenhunjie.jfif', '/static/data/new/30a.png', 1, 2, 2, 1, 24, '2022-05-02 19:12:28', 1, 0, 1, 0);
INSERT INTO `video` VALUES (31, '秦时明月之诸子百家', '出机关城伪装入桑海', 1, '2022-05-02 19:12:28', '/static/data/new/31.png', '/static/data/new/31a.png', 1, 2, 2, 1, 34, '2022-05-02 19:12:28', 1, 0, 1, 0);
INSERT INTO `video` VALUES (32, '少年锦衣卫', '三盗大闹皇宫劫公主', 1, '2022-05-02 19:12:28', '/static/data/new/32.png', '/static/data/new/32a.png', 1, 2, 2, 1, 28, '2022-05-02 19:12:28', 1, 0, 1, 0);
INSERT INTO `video` VALUES (33, '十万个冷笑话 第三季', '经典角色悉数回归爆笑玩梗', 1, '2022-05-02 19:12:28', '/static/data/new/33.png', '/static/data/new/33a.png', 1, 2, 2, 1, 52, '2022-05-02 19:12:28', 1, 0, 1, 0);
INSERT INTO `video` VALUES (34, '超神学院', '英雄联盟改编人气动画', 1, '2022-05-02 19:12:28', '/static/data/new/chaoshenxueyuan.jfif', '/static/data/new/chaoshenxueyuan1.jfif', 1, 2, 2, 1, 30, '2022-05-02 19:12:28', 1, 0, 1, 0);
INSERT INTO `video` VALUES (35, '秦时明月之万里长城', '铸造帝国的万里长城', 1, '2022-05-02 19:12:28', '/static/data/new/wanlichangcheng.jfif', '/static/data/new/wanlichangcheng1.jfif', 1, 2, 2, 1, 37, '2022-05-02 19:12:28', 1, 0, 1, 0);
INSERT INTO `video` VALUES (36, '天行九歌', '战国末年一局权力游戏', 1, '2022-05-02 19:12:28', '/static/data/new/tianxingjiuge.jfif', '/static/data/new/tianxingjiuge1.jfif', 1, 2, 2, 1, 60, '2022-05-02 19:12:28', 1, 0, 1, 0);
INSERT INTO `video` VALUES (37, '星际宝贝', '能实现愿望的魔法蛋', 1, '2022-05-02 19:12:28', '/static/data/new/OIP-C.jfif', '/static/data/new/xingjibaobei1.jfif', 1, 1, 1, 1, 51, '2022-05-02 19:12:28', 1, 1, 1, 0);
INSERT INTO `video` VALUES (38, '李克勤', '李克勤', 1, '2022-05-02 19:12:28', '/static/data/new/likeqin.jfif', '/static/data/new/mei.png', 1, 2, 1, 1, 10, '2022-05-02 19:12:28', 1, 1, 1, 10077218);

-- ----------------------------
-- Table structure for video_episodes
-- ----------------------------
DROP TABLE IF EXISTS `video_episodes`;
CREATE TABLE `video_episodes`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '标题',
  `add_time` datetime NULL DEFAULT '0000-00-00 00:00:00' COMMENT '添加时间',
  `num` int NULL DEFAULT 0 COMMENT '第几集',
  `video_id` int NULL DEFAULT 0 COMMENT '所属时间',
  `play_url` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '播放地址',
  `status` tinyint NULL DEFAULT 0 COMMENT '0=不展示 1= 展示',
  `comment` int NULL DEFAULT 0 COMMENT '评论数',
  `aliyun_video_id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '阿里云ID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 14 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of video_episodes
-- ----------------------------
INSERT INTO `video_episodes` VALUES (1, '《Welcome To New York》', '2022-05-02 19:12:28', 1, 1, '/static/video/《Welcome To New York》.mp4', 1, 10058217, '');
INSERT INTO `video_episodes` VALUES (2, '《Love Story》', '2022-05-04 15:51:30', 2, 1, '/static/video/《Love Story》.mp4', 1, 13504189, NULL);
INSERT INTO `video_episodes` VALUES (3, '《Look What You Made Me Do》', '2022-05-02 19:12:28', 3, 1, '/static/video/《Look What You Made Me Do》.mp4', 1, 9973976, '');
INSERT INTO `video_episodes` VALUES (4, '月半小夜曲', '2022-05-06 16:42:17', 1, 38, '/static/video/月半小夜曲.mp4', 1, 8632565, NULL);
INSERT INTO `video_episodes` VALUES (5, '星际宝贝', '2022-05-06 16:42:14', 1, 37, 'https://www.libvio.com/play/100815-1-1.html', 1, 0, NULL);

SET FOREIGN_KEY_CHECKS = 1;
