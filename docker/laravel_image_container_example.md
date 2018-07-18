- docker pull adrianharabula/php7-with-oci8  #decarga imagen
- docker images -a #visualiza las imagenes descargadas
- docker ps -a #visualiza todos los containers
- docker run --name PHP7OCI8 -it php7-with-oci8 bash
#salir del contenedor sin detenerlo CTRL + P + Q
#para volver al contenerdor 'docker attach ID-CONTAINER'
#para commitear una imagen modificaada 'docker commit ID-CONTAINER NEW-REPO-NAME'