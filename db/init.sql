CREATE USER 'rd'@'%' IDENTIFIED BY 'test1234';
GRANT ALL PRIVILEGES ON * . * TO 'rd'@'%' ;

CREATE DATABASE member;

use member;

/*DROP TABLE IF EXISTS member CASCADE;*/

CREATE TABLE `users` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `username` varchar(50) COLLATE utf8mb4_bin NOT NULL COMMENT '用戶',
  `password` varchar(255) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '密碼',
  `phone` varchar(20) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '手機',
  `email` varchar(255) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '信箱',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '建立時間',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

INSERT INTO `users` (`id`, `username`, `password`, `phone`, `email`, `create_time`) VALUES ('1', 'xiao', '123', '12345678910', '123@mail.com', '2020-04-10 01:22:07');
INSERT INTO `users` (`id`, `username`, `password`, `phone`, `email`, `create_time`) VALUES ('2', 'w', '456', '0909090909090909', '123@mail.com', '2020-04-10 01:22:07');
INSERT INTO `users` (`id`, `username`, `password`, `phone`, `email`, `create_time`) VALUES ('3', 'lee', '789', '12345678910', '123@mail.com', '2020-04-10 01:22:07');
INSERT INTO `users` (`id`, `username`, `password`, `phone`, `email`, `create_time`) VALUES ('4', 'lee', '123', '12345678910', '123@mail.com', '2020-04-10 01:22:07');
INSERT INTO `users` (`id`, `username`, `password`, `phone`, `email`, `create_time`) VALUES ('5', 'xiao', '123', 'QQQQQQQQQ', '123@mail.com', '2020-04-10 01:22:07');



