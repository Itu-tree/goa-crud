FROM mysql:8.0
#MySQL設定ファイルをイメージ内にコピー
ADD ./my.cnf /etc/mysql/conf.d/my.cnf
ADD ./docker.cnf /etc/mysql/conf.d/docker.cnf
#docker runに実行される
CMD ["mysqld"]