


- USE mysql;
- UPDATE user SET plugin='mysql_native_password' WHERE User='root';
- UPDATE mysql.user set authentication_string=PASSWORD('root') where user='root';
- FLUSH PRIVILEGES;
