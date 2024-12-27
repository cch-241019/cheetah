drop table if exists test_create_table;
create table test_create_table
(
    id            bigint(20)  not null auto_increment,
    test_varchar  varchar(20) not null default '',
    test_default  int(10)              default 1,
    test_unsigned int(10) unsigned,
    primary key (id)
) ;