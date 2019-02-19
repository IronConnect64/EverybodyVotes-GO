--		Copyright (c) 2019 IronConnect64
--
--		EverybodyVotes-GO, a custom server for the Everybody Votes channel on the PSP.
--
--	Permission is hereby granted, free of charge, to any person obtaining a copy
--	of this software and associated documentation files (the "Software"), to deal
--	in the Software without restriction, including without limitation the rights
--	to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
--	copies of the Software, and to permit persons to whom the Software is
--	furnished to do so, subject to the following conditions:
--
--	The above copyright notice and this permission notice shall be included in all
--	copies or substantial portions of the Software.
--
--	THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
--	IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
--	FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
--	AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
--	LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
--	OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
--	SOFTWARE.
--
-- --------------------------------------------------------
-- 10.3.12-MariaDB               
-- --------------------------------------------------------

create database if not exists `everybodyvotes-go`;
use `everybodyvotes-go`;

create table if not exists `polls` (
  `ID` int(5) unsigned,
  `Question` varchar(200),
  `A` varchar(40),
  `B` varchar(40),
  `Latest` boolean
) engine=InnoDB default charset=utf8;

create table if not exists `users` (
  `Data` binary(200),
  `MAC` varchar(17),
  `Token` varchar(50) unique
) engine=InnoDB default charset=utf8;

create table if not exists `votes` (
  `ID` int(5) unsigned,
  `AVotes` int(3) unsigned zerofill,
  `BVotes` int(3) unsigned zerofill
) engine=InnoDB default charset=utf8;
