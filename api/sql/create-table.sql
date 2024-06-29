-- petprojectdb.usuarios definition

CREATE TABLE `usuarios` (
                            `usuario_id` int NOT NULL AUTO_INCREMENT,
                            `cpf` varchar(14) NOT NULL,
                            `nome` varchar(100) NOT NULL,
                            `email` varchar(50) NOT NULL,
                            `senha` varchar(50) NOT NULL,
                            `data_cadastro` datetime default timestamp.now() NOT NULL,
                            `usuario_tipo` varchar(20) DEFAULT NULL,
                            PRIMARY KEY (`usuario_id`),
                            UNIQUE KEY `user_unique` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

----------------------------------------------------------------------
-- petprojectdb.evento definition

CREATE TABLE `evento` (
                          `evento_id` int NOT NULL AUTO_INCREMENT,
                          `nome` varchar(100) NOT NULL,
                          `local` varchar(100) NOT NULL,
                          `capacidade` int NOT NULL,
                          `data` datetime DEFAULT NULL,
                          PRIMARY KEY (`evento_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

----------------------------------------------------------------------
-- petprojectdb.categoria definition

CREATE TABLE `categoria` (
                             `categoria_id` int NOT NULL AUTO_INCREMENT,
                             `nome` varchar(100) NOT NULL,
                             `status` tinyint(1) NOT NULL,
                             PRIMARY KEY (`categoria_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

----------------------------------------------------------------------
-- petprojectdb.bilhete definition

CREATE TABLE `bilhete` (
                           `bilhete_id` int NOT NULL AUTO_INCREMENT,
                           `preco` decimal(10,0) NOT NULL,
                           `status` varchar(30) NOT NULL,
                           `user_id` int DEFAULT NULL,
                           `evento_id` int DEFAULT NULL,
                           `categoria_id` int DEFAULT NULL,
                           PRIMARY KEY (`bilhete_id`),
                           KEY `bilhete_usuarios_FK` (`user_id`),
                           KEY `bilhete_evento_FK` (`evento_id`),
                           KEY `bilhete_categoria_FK` (`categoria_id`),
                           CONSTRAINT `bilhete_categoria_FK` FOREIGN KEY (`categoria_id`) REFERENCES `categoria` (`categoria_id`),
                           CONSTRAINT `bilhete_evento_FK` FOREIGN KEY (`evento_id`) REFERENCES `evento` (`evento_id`),
                           CONSTRAINT `bilhete_usuarios_FK` FOREIGN KEY (`user_id`) REFERENCES `usuarios` (`usuario_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

----------------------------------------------------------------------
-- petprojectdb.pagamento definition

CREATE TABLE `pagamento` (
                             `pagamento_id` int NOT NULL AUTO_INCREMENT,
                             `metodo_pagamento` varchar(50) NOT NULL,
                             `montante` decimal(10,0) NOT NULL,
                             `data_pagamento` datetime NOT NULL,
                             `status` varchar(50) NOT NULL,
                             `bilhete_id` int DEFAULT NULL,
                             PRIMARY KEY (`pagamento_id`),
                             KEY `pagamento_pagamento_id_IDX` (`pagamento_id`) USING BTREE,
                             KEY `pagamento_bilhete_FK` (`bilhete_id`),
                             CONSTRAINT `pagamento_bilhete_FK` FOREIGN KEY (`bilhete_id`) REFERENCES `bilhete` (`bilhete_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;