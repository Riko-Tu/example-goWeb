create table user(
                     id int primary key ,
                     email varchar(20) unique ,
                     iphone varchar(20) unique ,
                     password varchar(20) ,
                     create_time bigint not null ,
                     update_time bigint not null
)
