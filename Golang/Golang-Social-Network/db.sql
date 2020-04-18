DROP USER 'netadmin_s96lu'@'%';
CREATE USER 'netadmin_s96lu'@'%' IDENTIFIED BY 'netadmin_s96lu';
GRANT ALL PRIVILEGES ON socialnet.* TO 'netadmin_s96lu'@'%';

DROP TABLE IF EXISTS `Users`;
CREATE TABLE `Users` (
    `user_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `username` VARCHAR(32) NOT NULL,
    `password` VARCHAR(255) NOT NULL,
    `email` VARCHAR(255) NOT NULL,
    PRIMARY KEY(user_id),
    UNIQUE(username),
    UNIQUE(password),
    UNIQUE(email)
);

DROP TABLE IF EXISTS `Profile`;
CREATE TABLE `Profile`(
    `profile_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id` INT UNSIGNED NOT NULL,
    `allow_unfollowed_views` BOOLEAN DEFAULT true,
    `job` VARCHAR(32) NOT NULL,
    `quote` VARCHAR(255) NOT NULL,
    `followers_num` INT UNSIGNED NOT NULL DEFAULT 0,
    `following_num` INT UNSIGNED NOT NULL DEFAULT 0,
    `views` INT UNSIGNED NOT NULL DEFAULT 0,
    `avatar` VARCHAR(255) NOT NULL,
    `created_date` DATETIME NOT NULL,
    PRIMARY KEY(profile_id),
    UNIQUE(user_id),
    UNIQUE(avatar),
    FOREIGN KEY(user_id) REFERENCES Users(user_id) ON DELETE CASCADE ON UPDATE CASCADE
);

DROP TABLE IF EXISTS `Follow`;
CREATE TABLE `Follow` (
    `follow_by` INT UNSIGNED NOT NULL,
    `follow_to` INT UNSIGNED NOT NULL,
    `created_date` DATETIME NOT NULL,
    PRIMARY KEY(follow_by, follow_to),
    FOREIGN KEY(follow_by) REFERENCES Users(user_id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY(follow_to) REFERENCES Users(user_id) ON DELETE CASCADE ON UPDATE CASCADE
);

DROP TABLE IF EXISTS `Blacklist`;
CREATE TABLE `Blacklist` (
    `black_by` INT UNSIGNED NOT NULL,
    `black_to` INT UNSIGNED NOT NULL,
    `created_date` DATETIME NOT NULL,
    PRIMARY KEY(black_by, black_to),
    FOREIGN KEY(black_by) REFERENCES Users(user_id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY(black_to) REFERENCES Users(user_id) ON DELETE CASCADE ON UPDATE CASCADE
);

DROP TABLE IF EXISTS `Posts`;
CREATE TABLE `Posts` (
    `post_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `likes` INT UNSIGNED NOT NULL DEFAULT 0,
    `created_by` INT UNSIGNED NOT NULL NOT NULL,
    `created_date` DATETIME NOT NULL,
    `allow_comments` BOOLEAN DEFAULT true,
    `comments_num` INT UNSIGNED NOT NULL DEFAULT 0,
    `title` VARCHAR(255) NOT NULL,
    `content` TEXT NOT NULL,
    PRIMARY KEY(post_id),
    FOREIGN KEY(created_by) REFERENCES Users(user_id) ON DELETE CASCADE ON UPDATE CASCADE
);

DROP TABLE IF EXISTS `Comments`;
CREATE TABLE `Comments`(
    `comment_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `post_id` INT UNSIGNED NOT NULL,
    `user_id` INT UNSIGNED NOT NULL,
    `content` TEXT NOT NULL,
    `likes` INT UNSIGNED NOT NULL DEFAULT 0,
    `created_date` DATETIME NOT NULL,
    PRIMARY KEY(comment_id),
    FOREIGN KEY(post_id) REFERENCES Posts(post_id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY(user_id) REFERENCES Users(user_id) ON DELETE CASCADE ON UPDATE CASCADE
);

DROP TABLE IF EXISTS `Likes`;
CREATE TABLE `Likes` (
    `post_id` INT UNSIGNED NOT NULL,
    `like_by` INT UNSIGNED NOT NULL NOT NULL,
    `created_date` DATETIME NOT NULL,
    PRIMARY KEY(post_id, like_by),
    FOREIGN KEY(like_by) REFERENCES Users(user_id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY(post_id) REFERENCES Posts(post_id) ON DELETE CASCADE ON UPDATE CASCADE
);

DROP TABLE IF EXISTS `Hashtags`;
CREATE TABLE `Hashtags`(
    `hashtag_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `hashtag_name` VARCHAR(255) NOT NULL,
    `followers_num` INT UNSIGNED NOT NULL DEFAULT 0,
    `posts_num` INT UNSIGNED NOT NULL DEFAULT 0,
    `created_date` DATETIME NOT NULL,
    PRIMARY KEY(hashtag_id),
    UNIQUE(hashtag_name)
);

DROP TABLE IF EXISTS `Posts_Hashtags`;
CREATE TABLE `Posts_Hashtags` (
    `post_id` INT UNSIGNED NOT NULL,
    `hashtag_id` INT UNSIGNED NOT NULL,
    `created_date` DATETIME NOT NULL,
    PRIMARY KEY(post_id, hashtag_id),
    FOREIGN KEY(post_id) REFERENCES Posts(post_id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY(hashtag_id) REFERENCES Hashtags(hashtag_id) ON DELETE CASCADE ON UPDATE CASCADE
);

DROP TABLE IF EXISTS `Users_Hashtags`;
CREATE TABLE `Users_Hashtags` (
    `user_id` INT UNSIGNED NOT NULL,
    `hashtag_id` INT UNSIGNED NOT NULL,
    `created_date` DATETIME NOT NULL,
    PRIMARY KEY(user_id, hashtag_id),
    FOREIGN KEY(user_id) REFERENCES Users(user_id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY(hashtag_id) REFERENCES Hashtags(hashtag_id) ON DELETE CASCADE ON UPDATE CASCADE
);

/*DROP TABLE IF EXISTS `Topics`;
CREATE TABLE `Topics` (
    `topic_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `topic_name` VARCHAR(255) NOT NULL,
    `followers_num` INT UNSIGNED NOT NULL DEFAULT 0,
    `hashtags_num` INT UNSIGNED NOT NULL DEFAULT 0,
    `posts_num` INT UNSIGNED NOT NULL DEFAULT 0,
    `created_date` DATETIME NOT NULL,
    PRIMARY KEY(topic_id),
    UNIQUE(topic_name)
);

DROP TABLE IF EXISTS `Posts_Topics`;
CREATE TABLE `Posts_Topics` (
    `post_id` INT UNSIGNED NOT NULL,
    `topic_id` INT UNSIGNED NOT NULL,
    `created_date` DATETIME NOT NULL,
    PRIMARY KEY(post_id, topic_id),
    FOREIGN KEY(post_id) REFERENCES Posts(post_id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY(topic_id) REFERENCES Topics(topic_id) ON DELETE CASCADE ON UPDATE CASCADE
);

DROP TABLE IF EXISTS `Users_Topics`;
CREATE TABLE `Users_Topics` (
    `user_id` INT UNSIGNED NOT NULL,
    `topic_id` INT UNSIGNED NOT NULL,
    `created_date` DATETIME NOT NULL,
    PRIMARY KEY(user_id, topic_id),
    FOREIGN KEY(user_id) REFERENCES Users(user_id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY(topic_id) REFERENCES Topics(topic_id) ON DELETE CASCADE ON UPDATE CASCADE
);

DROP TABLE IF EXISTS `Hashtags_Topics`;
CREATE TABLE `Hashtags_Topics` (
    `hashtag_id` INT UNSIGNED NOT NULL,
    `topic_id` INT UNSIGNED NOT NULL,
    `created_date` DATETIME NOT NULL,
    PRIMARY KEY(hashtag_id, topic_id),
    FOREIGN KEY(hashtag_id) REFERENCES Hashtags(hashtag_id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY(topic_id) REFERENCES Topics(topic_id) ON DELETE CASCADE ON UPDATE CASCADE
);*/

DROP TABLE IF EXISTS `Mentions`;
CREATE TABLE `Mentions`(
    `user_id` INT UNSIGNED NOT NULL,
    `post_id` INT UNSIGNED NOT NULL,
    `created_date` DATETIME NOT NULL,
    PRIMARY KEY(post_id, user_id),
    FOREIGN KEY(user_id) REFERENCES Users(user_id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY(post_id) REFERENCES Posts(post_id) ON DELETE CASCADE ON UPDATE CASCADE
);

DELIMITER $$
DROP TRIGGER IF EXISTS `create_user_profile`;
CREATE TRIGGER `create_user_profile` AFTER INSERT ON `Users` FOR EACH ROW
BEGIN
    INSERT INTO Profile (user_id, created_date, avatar) VALUES(NEW.user_id, NOW(), CONCAT('users/',NEW.user_id,'/profile/avatar.png'));
END$$

DROP TRIGGER IF EXISTS `new_follows_time`;
CREATE TRIGGER `new_follows_time` BEFORE INSERT ON `Follow` FOR EACH ROW
BEGIN
    SET NEW.created_date = NOW();
END$$

DROP TRIGGER IF EXISTS `new_follows`;
CREATE TRIGGER `new_follows` AFTER INSERT ON `Follow` FOR EACH ROW
BEGIN
    UPDATE Profile SET followers_num = followers_num + 1 WHERE user_id = NEW.follow_to;
    UPDATE Profile SET following_num = following_num + 1 WHERE user_id = NEW.follow_by;
END$$

DROP TRIGGER IF EXISTS `remove_follows`;
CREATE TRIGGER `remove_follows` AFTER DELETE ON `Follow` FOR EACH ROW
BEGIN   
    UPDATE Profile SET followers_num = followers_num - 1 WHERE user_id = OLD.follow_to;
    UPDATE Profile SET following_num = following_num - 1 WHERE user_id = OLD.follow_by;
END$$

DROP TRIGGER IF EXISTS `new_blacks_time`;
CREATE TRIGGER `new_blacks_time` BEFORE INSERT ON `Blacklist` FOR EACH ROW
BEGIN
    SET NEW.created_date = NOW();
END$$
DELIMITER ;


DELIMITER $$
DROP TRIGGER IF EXISTS `new_likes_time`;
CREATE TRIGGER `new_likes_time` BEFORE INSERT ON `Likes` FOR EACH ROW
BEGIN
    SET NEW.created_date = NOW();
END$$

DROP TRIGGER IF EXISTS `new_likes`;
CREATE TRIGGER `new_likes` AFTER INSERT ON `Likes` FOR EACH ROW
BEGIN
    UPDATE Posts SET likes = likes + 1 WHERE post_id = NEW.post_id;
END$$

DROP TRIGGER IF EXISTS `remove_likes`;
CREATE TRIGGER `remove_likes` AFTER DELETE ON `Likes` FOR EACH ROW
BEGIN
    UPDATE Posts SET likes = likes - 1 WHERE post_id = OLD.post_id;
END$$

DROP TRIGGER IF EXISTS `new_posts`;
CREATE TRIGGER `new_posts` BEFORE INSERT ON `Posts` FOR EACH ROW
BEGIN
    SET NEW.created_date = NOW();
END$$

DROP TRIGGER IF EXISTS `new_comments`;
CREATE TRIGGER `new_comments` BEFORE INSERT ON `Comments` FOR EACH ROW
BEGIN
    SET NEW.created_date = NOW();
END$$

DROP TRIGGER IF EXISTS `new_post_comments`;
CREATE TRIGGER `new_post_comments` AFTER INSERT ON `Comments` FOR EACH ROW
BEGIN
    UPDATE Posts SET comments_num = comments_num + 1 WHERE post_id = NEW.post_id;
END$$

DROP TRIGGER IF EXISTS `remove_post_comments`;
CREATE TRIGGER `remove_post_comments` AFTER DELETE ON `Comments` FOR EACH ROW
BEGIN
    UPDATE Posts SET comments_num = comments_num - 1 WHERE post_id = OLD.post_id;
END$$
DELIMITER ;


DELIMITER $$
DROP TRIGGER IF EXISTS `new_hashtags`;
CREATE TRIGGER `new_hashtags` BEFORE INSERT ON `Hashtags` FOR EACH ROW
BEGIN
    SET NEW.created_date = NOW();
END$$

DROP TRIGGER IF EXISTS `new_users_hashtags_time`;
CREATE TRIGGER `new_users_hashtags_time` BEFORE INSERT ON `Users_Hashtags` FOR EACH ROW
BEGIN
    SET NEW.created_date = NOW();
END$$

DROP TRIGGER IF EXISTS `new_users_hashtags`;
CREATE TRIGGER `new_users_hashtags` AFTER INSERT ON `Users_Hashtags` FOR EACH ROW
BEGIN
    UPDATE Hashtags SET followers_num = followers_num + 1 WHERE hashtag_id = NEW.hashtag_id;
END$$

DROP TRIGGER IF EXISTS `remove_users_hashtags`;
CREATE TRIGGER `remove_users_hashtags` AFTER DELETE ON `Users_Hashtags` FOR EACH ROW
BEGIN
    UPDATE Hashtags SET followers_num = followers_num - 1 WHERE hashtag_id = OLD.hashtag_id;
END$$

DROP TRIGGER IF EXISTS `new_posts_hashtags_time`;
CREATE TRIGGER `new_posts_hashtags_time` BEFORE INSERT ON `Posts_Hashtags` FOR EACH ROW
BEGIN
    SET NEW.created_date = NOW();
END$$

DROP TRIGGER IF EXISTS `new_posts_hashtags`;
CREATE TRIGGER `new_posts_hashtags` AFTER INSERT ON `Posts_Hashtags` FOR EACH ROW
BEGIN
    UPDATE Hashtags SET posts_num = posts_num + 1 WHERE hashtag_id = NEW.hashtag_id;
END$$

DROP TRIGGER IF EXISTS `remove_posts_hashtags`;
CREATE TRIGGER `remove_posts_hashtags` AFTER DELETE ON `Posts_Hashtags` FOR EACH ROW
BEGIN
    UPDATE Hashtags SET posts_num = posts_num - 1 WHERE hashtag_id = OLD.hashtag_id;
END$$

/*DROP TRIGGER IF EXISTS `new_posts_hashtags`;
CREATE TRIGGER `new_posts_hashtags` AFTER INSERT ON `Posts_Hashtags` FOR EACH ROW
BEGIN
    UPDATE Hashtags SET posts_num = posts_num + 1 WHERE hashtag_id = NEW.hashtag_id;
END$$

DROP TRIGGER IF EXISTS `remove_posts_hashtags`;
CREATE TRIGGER `remove_posts_hashtags` AFTER DELETE ON `Posts_Hashtags` FOR EACH ROW
BEGIN
    UPDATE Hashtags SET posts_num = posts_num - 1 WHERE hashtag_id = OLD.hashtag_id;
END$$
DELIMITER ;*/

/*
DELIMITER $$
DROP TRIGGER IF EXISTS `new_topics`;
CREATE TRIGGER `new_topics` BEFORE INSERT ON `Topics` FOR EACH ROW
BEGIN
    SET NEW.created_date = NOW();
END$$

DROP TRIGGER IF EXISTS `new_users_topics_time`;
CREATE TRIGGER `new_users_topics_time` BEFORE INSERT ON `Users_Topics` FOR EACH ROW
BEGIN
    SET NEW.created_date = NOW();
END$$

DROP TRIGGER IF EXISTS `new_users_topics`;
CREATE TRIGGER `new_users_topics` AFTER INSERT ON `Users_Topics` FOR EACH ROW
BEGIN
    UPDATE Topics SET followers_num = followers_num + 1 WHERE topic_id = NEW.topic_id;
END$$

DROP TRIGGER IF EXISTS `remove_users_topics`;
CREATE TRIGGER `remove_users_topics` AFTER DELETE ON `Users_Topics` FOR EACH ROW
BEGIN
    UPDATE Topics SET followers_num = followers_num - 1 WHERE topic_id = OLD.topic_id;
END$$

DROP TRIGGER IF EXISTS `new_hashtags_topics_time`;
CREATE TRIGGER `new_hashtags_topics_time` BEFORE INSERT ON `Hashtags_Topics` FOR EACH ROW
BEGIN
    SET NEW.created_date = NOW();
END$$

DROP TRIGGER IF EXISTS `new_hashtags_topics`;
CREATE TRIGGER `new_hashtags_topics` AFTER INSERT ON `Hashtags_Topics` FOR EACH ROW
BEGIN
    UPDATE Topics SET hashtags_num = hashtags_num + 1 WHERE topic_id = NEW.topic_id;
END$$

DROP TRIGGER IF EXISTS `remove_hashtags_topics`;
CREATE TRIGGER `remove_hashtags_topics` AFTER DELETE ON `Hashtags_Topics` FOR EACH ROW
BEGIN
    UPDATE Topics SET hashtags_num = hashtags_num - 1 WHERE topic_id = OLD.topic_id;
END$$

DROP TRIGGER IF EXISTS `new_posts_topics_time`;
CREATE TRIGGER `new_posts_topics_time` BEFORE INSERT ON `Posts_Topics` FOR EACH ROW
BEGIN
    SET NEW.created_date = NOW();
END$$

DROP TRIGGER IF EXISTS `new_posts_topics`;
CREATE TRIGGER `new_posts_topics` AFTER INSERT ON `Posts_Topics` FOR EACH ROW
BEGIN
    UPDATE Topics SET posts_num = posts_num + 1 WHERE topic_id = NEW.topic_id;
END$$

DROP TRIGGER IF EXISTS `remove_posts_topics`;
CREATE TRIGGER `remove_posts_topics` AFTER DELETE ON `Posts_Topics` FOR EACH ROW
BEGIN
    UPDATE Topics SET posts_num = posts_num - 1 WHERE topic_id = OLD.topic_id;
END$$
DELIMITER ;
*/

/*
INSERT INTO `Topics` (`topic_name`) VALUES
('Shoes'),
('Bags'),
('Coats'),
('Computers'),
('Wall Street'),
('Cars');
*/

/*Passwords are encoded by API, but password is the same as username*/
INSERT INTO `Users` (`username`, `password`, `email`) VALUES
('sijia', '$2a$10$eXcpaPHH6tYg8Ie.bhvuZ.PSIykhBdIVpts0BnL0cXl/b3F9XyOKa', 'sijia@gmail.com'),
('takkar', '$2a$10$ttnsVDOPgMlA5vvDE33eneqVO3BHE/zif/axxI5AwNpOuRetkxFk6', 'takkar@gmail.com'),
('faiyaz', '$2a$10$.Wx2jBjYPiMFgWGCW.USze.qFMwrgN1TWOf50CQgqHDBzpcYV2uSa', 'faiyaz@gmail.com'),
('ghalib', '$2a$10$ziw6cqTgpSBIvASZOjTheey8sQYf1iW3HW4N.Xjq4GX6faKqzIrE.', 'ghalib@gmail.com'),
('user1', '$2a$10$PrsYgrp62NJkjOy1FOrL9uaRpzSfuiMv3oL6Xj5Hl90ZQtTlRmfZq', 'user1@gmail.com'),
('user2', '$2a$10$judwDLrzupULLqb8gxRQveGx.knR3LP2qJ/zaPH8YmYzoEdkr.tue', 'user2@gmail.com'),
('user3', '$2a$10$fIt2Gsfntg..wRPgY11yTugAt3HEeJPsbajftyFT4mRKIkJyjuBtS', 'user3@gmail.com'),
('hero1', '$2a$10$vOqEDCYP2ji9MZEp0lg.Jei.uOijw6viV4T5hbmt8/S3.Wi3WpOXS', 'hero1@gmail.com'),
('hero2', '$2a$10$C3XrawSnJIm74IhaVJ7m6upxl8ZWHKp6p.1GtPy6PTV9gMl0qAdr6', 'hero2@gmail.com'),
('hero3', '$2a$10$kLG3iRB1ULBTK.Jnhk.R0.LHuV6sXK1Djcs7X4xI7L2Ap8k9YYMXS', 'hero3@gmail.com'),
('nature', '$2a$10$nBi64BlbJMlzuSJfOhPlXevwdCgHOXKLZQUbJQ1q2Y7Ltbpaf1Woa', 'nature@gmail.com');

INSERT INTO `Follow` (`follow_by`, `follow_to`) VALUES
(2, 1),
(1, 2),
(3, 2),
(4, 2),
(3, 6),
(5, 3),
(7, 6),
(5, 7),
(7, 5),
(2, 3);

INSERT INTO `Blacklist` (`black_by`, `black_to`) VALUES
(2, 1),
(3, 1),
(1, 2),
(5, 2);

INSERT INTO `Posts` (`title`, `content`, `created_by`) VALUES
('Welcome', '#Welcome# Welcome to the community, guys', 1),
('Hello,', 'World!!', 2),
('second', 'second_content', 2),
('third', 'third content..', 2),
('Awesome platform', '#Welcome# I love this platform', 5),
('FirstPost', '#FirstPost# This is my first post !', 4),
('my title..', 'my content...', 1),
('ghalib''s first title..', 'and this is content!!!', 3),
('Wow', '#Welcome# It has been a month now, still loving it', 1),
('Number 8', 'Number 8', 8),
('jkj', 'kj', 4),
('Hey guys', '#FirstPost# First day here', 7),
('Number 9', '#Number9# I am number 9', 9);

INSERT INTO `Likes` (`post_id`, `like_by`) VALUES
(2, 1),
(2, 2),
(3, 1),
(4, 2),
(5, 1),
(6, 2),
(2, 7),
(1, 5),
(6, 3),
(7, 1);

