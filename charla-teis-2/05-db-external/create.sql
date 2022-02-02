CREATE DATABASE IF NOT EXISTS `asir`;

USE `asir`;

CREATE TABLE IF NOT EXISTS `alumno` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `login` varchar(255),
  `nombre` varchar(255),
  `email` varchar(255)
) ENGINE=InnoDB;

INSERT INTO `alumno` (`id`, `login`, `nombre`, `email`)
    VALUES (NULL, 'jorge.teixeira', 'Jorge Teixeira', 'jorge.teixeira@bizaway.com');

INSERT INTO `alumno` (`id`, `login`, `nombre`, `email`)
    VALUES (NULL, 'rodrigo.souto', 'Rodrigo Souto', 'rodrigo.souto@bizaway.com');

INSERT INTO `alumno` (`id`, `login`, `nombre`, `email`)
    VALUES (NULL, 'elon.musk', 'Elon Musk', 'elon@tesla.com');

INSERT INTO `alumno` (`id`, `login`, `nombre`, `email`)
    VALUES (NULL, 'alan.turing', 'Alan Turing', 'alan.turing@turingcomplete.com');

INSERT INTO `alumno` (`id`, `login`, `nombre`, `email`)
    VALUES (NULL, 'rafa.nadal', 'Rafa Nadal', 'rafa.nadal@21.com');

