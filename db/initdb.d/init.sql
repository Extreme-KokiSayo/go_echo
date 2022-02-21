### create table

# users
create table if not exists `sns`.`users` (
	`id` int unsigned not null AUTO_INCREMENT,
	`name` varchar(255) not null,
	`password` varchar(255) not null,
	`created_at` timestamp not null default current_timestamp,
  	`updated_at` timestamp not null default current_timestamp on update current_timestamp,
	primary key (`id`)
)
ENGINE = InnoDB;

# posts
create table if not exists `sns`.`posts` (
	`id` int unsigned not null AUTO_INCREMENT,
	`msg` varchar(255) not null,
	`user_id` int unsigned not null,
	`good_cnt` int unsigned default 0,
	`created_at` timestamp not null default current_timestamp,
  	`updated_at` timestamp not null default current_timestamp on update current_timestamp,
	primary key (`id`)
)
ENGINE = InnoDB;

### insert

insert into `sns`.`users` (id, name, password)
values (1, 'test', 'test'),
	(2, 'test1', 'test1'),
	(3, 'test2', 'test2');

insert into `sns`.`posts` (msg, user_id)
values ('Hello, World!', 1),
	('foobar', 2),
	('hogefuga', 3);