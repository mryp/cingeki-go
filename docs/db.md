create database cingekigo default character set utf8;

create table stories
(
    id int not null unique auto_increment,
    number int not null,
    title varchar(256) not null,
    uptime datetime not null,
    primary key (id)
) engine=innodb;
