# laravel storage folder permissions example
- `sudo chown -R www-data:www-data PROYECT-PATH`
- `sudo find PROYECT-PATH -type f -exec chmod 644 {} \;`
- `sudo find PROYECT-PATH -type d -exec chmod 755 {} \;`
- `sudo chgrp -R www-data PROYECT-PATH/storage PROYECT-PATH/bootstrap/cache`
- `sudo chmod -R ug+rwx PROYECT-PATH/storage PROYECT-PATH/bootstrap/cache`
