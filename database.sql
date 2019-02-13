-- --------------------------------------------------------
-- 10.3.12-MariaDB               
-- --------------------------------------------------------

create database if not exists `everybodyvotes-go`;
use `everybodyvotes-go`;

create table if not exists `polls` (
  `ID` int(5) unsigned not null,
  `Question` varchar(200) not null,
  `A` varchar(40) not null,
  `B` varchar(40) not null
) engine=InnoDB default charset=utf8;

create table if not exists `users` (
  `Data` binary(200) DEFAULT null,
  `MAC` varchar(17) not null
) engine=InnoDB default charset=utf8;

create table if not exists `votes` (
  `ID` int(5) unsigned not null,
  `AVotes` int(3) unsigned zerofill not null,
  `BVotes` int(3) unsigned zerofill not null
) engine=InnoDB default charset=utf8;
