-- phpMyAdmin SQL Dump
-- version 5.2.0-dev+20211224.ca91eaeb38
-- https://www.phpmyadmin.net/
--
-- Хост: 127.0.0.1:3307
-- Время создания: Фев 04 2023 г., 21:28
-- Версия сервера: 8.0.24
-- Версия PHP: 7.4.27

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- База данных: `cto`
--

-- --------------------------------------------------------

--
-- Структура таблицы `test`
--

CREATE TABLE `test` (
  `id` int NOT NULL,
  `mn` decimal(10,2) NOT NULL,
  `mn2` tinyint NOT NULL,
  `mn3` varchar(43) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

--
-- Дамп данных таблицы `test`
--

INSERT INTO `test` (`id`, `mn`, `mn2`, `mn3`) VALUES
(1, '1103.40', 43, '44334sdsa'),
(0, '10.30', 32, '545');
COMMIT;

-- --------------------------------------------------------

--
-- Структура таблицы `category`
--

CREATE TABLE `category` (
  `Id` int NOT NULL,
  `naming` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `categoryId` int NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

--
-- Дамп данных таблицы `category`
--

INSERT INTO `category` (`Id`, `naming`, `categoryId`) VALUES
(1, 'hoodie', 1),
(2, 'Tshirt', 2),
(3, 'Штани', 8),
(4, 'світер', 5);

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
