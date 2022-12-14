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