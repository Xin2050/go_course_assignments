
using the gin framework for an HTTP server example that demonstrates the use of pkg/errors Wrap errors,
layer by layer back to the controller, and then through the middleware error handling.

according to the error returns user-readable JSON format rest error,
using uber zap logging JSON format, and print stack trace information for debugging.

log middleware: middleware/logger_middleware.go

error start with :/domain/users/user_dao.go:23

request path: http://localhost:3000/user/:user_id
test sample:
    Correct: http://localhost:3000/user/25  return: code:200
    Failed: http://localhost:3000/user/999  return: code:404
    Failed: http://localhost:3000/user/ab  return: code:400

mysql info
host: localhost:3306
username&password: go
db: go_bookstore

sql:
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` mediumint NOT NULL AUTO_INCREMENT,
  `firstName` varchar(255) NOT NULL,
  `lastName` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) DEFAULT NULL,
  `status` varchar(50) DEFAULT NULL,
  `dateUpdated` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `dateCreated` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `emailcheck` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of users
-- ----------------------------
BEGIN;
INSERT INTO `users` VALUES (25, 'Xin3', 'Lee7', 'test19@gmail.com', '$2a$10$KCq4fI9xmYDAGq7NPOo8oOtR1W3X5bpq41alCEclq7kK1EAsF8zkO', 'Active', '2021-11-11 16:09:16', '2021-11-11 14:27:02');
INSERT INTO `users` VALUES (27, 'Xin3', 'Lee2', 'teste9@gmail.com', '$2a$10$dXACGuN8sLhSKeRb2B5qz.3S9wv5h/ukJje72E3GCTYUJOgnnnYOG', 'Active', '2021-11-11 15:40:49', '2021-11-11 15:40:49');
INSERT INTO `users` VALUES (29, 'Xin3', 'Lee2', 'tes3te9@gmail.com', '$2a$10$i.Dom68BXMx6ErArpQZP4eqHBb/C22ktxrLsdWMP8O7Is2gXVGVmy', 'Active', '2021-11-11 15:42:26', '2021-11-11 15:42:26');
INSERT INTO `users` VALUES (31, 'Xin3', 'Lee2', 'te4s3te9@gmail.com', '$2a$10$fKk/cc3ilJulwkUNyxH.Lu.Ip1RMvba2ADWtAsYr7SobGKujdeXwG', 'Active', '2021-11-11 15:42:39', '2021-11-11 15:42:39');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;