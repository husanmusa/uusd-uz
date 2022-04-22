alter table companies alter column name type varchar(16) ;
alter table companies alter column slogan type varchar(128) ;

alter table services alter column name type varchar(64);
alter table services alter column description type varchar(512);

alter table sets alter column name type varchar(64);
alter table sets alter column description type varchar(512);

alter table packages alter column name type varchar(64);
alter table packages alter column description type varchar(512);
