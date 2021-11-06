CREATE TABLE `user` (
    `id` varchar(50),
    `username` varchar(20),
    `password` varchar(100),
    `email` varchar(100),
    PRIMARY KEY (`id`),
    UNIQUE KEY idx_unique_username (`username`),
    UNIQUE KEY idx_unique_email (`email`),
    INDEX idx_user_username (`username`),
    INDEX idx_user_email (`email`),
);