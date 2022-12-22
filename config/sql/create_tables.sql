CREATE DATABASE IF NOT EXISTS QAcommunity;

create table users(
    id bigint not null auto_increment,
    user_id bigint not null ,
    username varchar(30) collate utf8mb4_general_ci not null,
    password varchar(50) collate utf8mb4_general_ci not null,
    createAt timestamp null default current_timestamp,
    updateAt timestamp null default current_timestamp on update current_timestamp,
    primary key (id),
    unique key `idx_username` (username) using btree ,
    unique key `idx_user_id` (user_id)  using btree
)engine =InnoDB character set =utf8mb4 collate =utf8mb4_general_ci;

create table posts(
    id bigint not null auto_increment,
    author_id bigint not null,
    post_id bigint not null,
    title varchar(128) not null ,
    content varchar(8192)  not null,
    createAt timestamp null default current_timestamp ,
    updateAt timestamp null default current_timestamp  on update current_timestamp,
    primary key (id),
    unique key `idx_post_id` (post_id)  using btree ,
    key `idx_author_id` (author_id) using btree
)engine = InnoDB default character set =utf8mb4 collate =utf8mb4_general_ci;

create table comments(
    id bigint not null auto_increment,
    comment_id bigint not null ,
    reply_id bigint default 0 not null ,
    content varchar(800) not null ,
    author_id bigint not null ,
    post_id bigint not null ,
    root_id bigint default 0 not null ,
    parent_id bigint default 0 not null ,
    createAt timestamp null default current_timestamp ,
    updateAt timestamp null default current_timestamp  on update current_timestamp,
    primary key (id),
    unique `idx_comment_id` (comment_id) using btree ,
    key `idx_author_id` (author_id) using btree ,
    key `idx_post_id` (post_id) using btree ,
    key `idx_root_id` (root_id) using btree ,
    key `idx_parent_id` (parent_id)using btree
)engine = InnoDB default character set =utf8mb4 collate =utf8mb4_general_ci;