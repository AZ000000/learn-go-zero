create table user
(
    id bigint auto_increment,
    username varchar(30) not NULL,
    password varchar(30) default '',
    unique name_index (username),
    primary key (id)
) engine =InnoDB collate utf8mb4_general_ci;

/*goctl model mysql ddl --src user.sql --dir .*/