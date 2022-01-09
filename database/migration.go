package database

import (
	"context"
	"log"
	"time"
)

func CreateProductTable() {

	db = ConnectDB()

	query := "CREATE TABLE `products` (\n\t`id` INT(11) NOT NULL AUTO_INCREMENT,\n\t`product_name` VARCHAR(50) NULL DEFAULT NULL COLLATE 'latin1_swedish_ci',\n\t`stock_code` VARCHAR(50) NULL DEFAULT NULL COLLATE 'latin1_swedish_ci',\n\tPRIMARY KEY (`id`) USING BTREE\n)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when creating product table", err)

	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when getting rows affected", err)

	}
	log.Printf("Rows affected when creating table: %d", rows)

}

func CreateShelfTable() {

	db = ConnectDB()

	query := "CREATE TABLE `shelfs` (\n\t`id` INT(11) NOT NULL AUTO_INCREMENT,\n\t`name` VARCHAR(50) NULL DEFAULT NULL COLLATE 'latin1_swedish_ci',\n\t`capacity` SMALLINT(6) NULL DEFAULT NULL,\n\tPRIMARY KEY (`id`) USING BTREE\n)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when creating product table", err)

	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when getting rows affected", err)

	}
	log.Printf("Rows affected when creating table: %d", rows)

}

func CreateUsersTable() {

	db = ConnectDB()

	query := "CREATE TABLE `users` (\n\t`id` INT(11) NOT NULL AUTO_INCREMENT,\n\t`username` VARCHAR(50) NULL DEFAULT NULL COLLATE 'latin1_swedish_ci',\n\t`password` VARCHAR(100) NULL DEFAULT NULL COLLATE 'latin1_swedish_ci',\n\t`role` TINYINT(4) NULL DEFAULT NULL COMMENT '1:İdari Personel, 2: Depo Personeli',\n\tPRIMARY KEY (`id`) USING BTREE\n)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when creating product table", err)

	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when getting rows affected", err)

	}
	log.Printf("Rows affected when creating table: %d", rows)

}

func CreateStocksTable() {

	db = ConnectDB()

	query := "CREATE TABLE `stocks` (\n\t`id` INT(11) NOT NULL AUTO_INCREMENT,\n\t`product_code` INT(11) NULL DEFAULT NULL,\n\t`quantity` INT(11) NULL DEFAULT NULL,\n\t`shelf_id` TINYINT(4) NULL DEFAULT NULL,\n\tPRIMARY KEY (`id`) USING BTREE\n)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when creating product table", err)

	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when getting rows affected", err)

	}
	log.Printf("Rows affected when creating table: %d", rows)

}

func CreateTransactionTable() {

	db = ConnectDB()

	query := "CREATE TABLE `transactions` (\n\t`id` INT(11) NOT NULL AUTO_INCREMENT,\n\t`stock_code` VARCHAR(50) NULL DEFAULT NULL COLLATE 'latin1_swedish_ci',\n\t`quantity` INT(11) NULL DEFAULT NULL,\n\t`type` TINYINT(4) NULL DEFAULT NULL COMMENT '1: Mal Kabul, 2: Sipariş',\n\t`statu` INT(11) NULL DEFAULT '0' COMMENT '0: Beklemede, 1:İşlem Tamamlandı',\n\tPRIMARY KEY (`id`) USING BTREE\n)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when creating product table", err)

	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when getting rows affected", err)

	}
	log.Printf("Rows affected when creating table: %d", rows)

}
