create database if not exists `everybodyvotes-go`;
use `everybodyvotes-go`;

create table if not exists `polls` (
  `ID` int(5) unsigned zerofill default,
  `Question` varchar(200) default null,
  `A` varchar(40) default null,
  `B` varchar(40) default null
) engine=InnoDB default charset=utf8;

create table if not exists `users` (
  `MAC` varchar(17) NOT null
) engine=InnoDB default charset=utf8;

create table if not exists `votes` (
  `ID` int(5) unsigned zerofill default null,
  `AVotes` int(3) unsigned zerofill default null,
  `BVotes` int(3) unsigned zerofill default null
) engine=InnoDB default charset=utf8;