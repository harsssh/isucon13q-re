
ALTER TABLE `themes` auto_increment = 1;
ALTER TABLE `icons` auto_increment = 1;
ALTER TABLE `reservation_slots` auto_increment = 1;
ALTER TABLE `livestream_tags` auto_increment = 1;
ALTER TABLE `livestream_viewers_history` auto_increment = 1;
ALTER TABLE `livecomment_reports` auto_increment = 1;
ALTER TABLE `ng_words` auto_increment = 1;
ALTER TABLE `reactions` auto_increment = 1;
ALTER TABLE `tags` auto_increment = 1;
ALTER TABLE `livecomments` auto_increment = 1;
ALTER TABLE `livestreams` auto_increment = 1;
ALTER TABLE `users` auto_increment = 1;

ALTER TABLE `livestreams` ADD INDEX user_id_idx(user_id);
ALTER TABLE `livestream_tags` ADD INDEX livestream_id_idx(livestream_id);
ALTER TABLE `livecomments` ADD INDEX livestream_id_created_at_idx(livestream_id, created_at desc);
ALTER TABLE `themes` ADD INDEX user_id_idx(user_id);
ALTER TABLE `ng_words` ADD INDEX user_id_livestream_id_idx(user_id, livestream_id);
ALTER TABLE `reservation_slots` ADD INDEX start_at_idx(`start_at`);
ALTER TABLE `icons` ADD INDEX user_id_idx(user_id);
