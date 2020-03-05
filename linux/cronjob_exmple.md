
# OPEN CRONJOB-FILE

```
export VISUAL=nano; crontab -e
```

# EXAMPLE

```
* * * * * cd /var/www/html/YOURPROYECTNAME && php artisan schedule:run >> /dev/null 2>&1
*/10 * * * * cd /var/www/html/YOURPROYECTNAME && yarn && yarn production && php artisan config:clear && php artisan clear-compiled && /usr/bin/git reset --hard && /usr/bin/git pull origin develop && php a$

*/5 * * * * cd /var/www/html/YOURPROYECTNAME && /usr/bin/git pull origin develop

* * * * * cd /var/www/html/YOURPROYECTNAME && php artisan config:clear && php artisan clear-compiled && /usr/bin/git reset --hard && php artisan schedule:run >> /dev/null 2>&1
*/10 * * * * cd /var/www/html/YOURPROYECTNAME && yarn && yarn production && /usr/bin/git pull origin develop && php artisan migrate >> /dev/null 2>&1

* * * * * cd /var/www/html/YOURPROYECTNAME && php artisan config:clear && php artisan clear-compiled && /usr/bin/git reset --hard && php artisan schedule:run >> /dev/null 2>&1
*/10 * * * * cd /var/www/html/YOURPROYECTNAME && yarn && yarn production && /usr/bin/git pull origin develop && php artisan migrate >> /dev/null 2>&1
```
