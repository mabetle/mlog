
-- Create demo db for logger

create database demo default character set utf8 default collate utf8_general_ci;

grant all privileges on demo.* to demo@localhost identified by 'demo';

grant all privileges on demo.* to 'demo'@'%' identified by 'demo';


