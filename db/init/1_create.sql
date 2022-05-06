-- golang_dbという名前のデータベースを作成
CREATE DATABASE golang_db;
-- golang_dbをアクティブ
use golang_db;
-- usersテーブルを作成。名前とパスワード
CREATE TABLE users (
    id INT(11) AUTO_INCREMENT NOT NULL,
    name VARCHAR(64) NOT NULL,
    password CHAR(30) NOT NULL,
    PRIMARY KEY (id)
);
-- usersテーブルに２つレコードを追加
INSERT INTO users (name, password) VALUES ("gophar", "5555");
INSERT INTO users (name, password) VALUES ("octcat", "0000");

-- golagn-api用のユーザーを作成
CREATE USER 'golang-api'@'%' IDENTIFIED BY 'password';
-- 認証方式を変更
ALTER USER 'golang-api'@'%' IDENTIFIED WITH mysql_native_password BY 'password';
-- 権限を追加(@の直後について、golang-api(docker-composeのサービス名)からのみアクセスできる)
GRANT SELECT,INSERT,UPDATE,DELETE ON golang_db.* TO 'golang-api'@'%';