# Mini Tutorial de WebAssembly

---

Referencia:
- https://www.youtube.com/watch?v=txC8vKYmHIw

---

## Creación y ejecución de contenedor con volumen vinculado a la carpeta actual:

  ```
  docker run --rm -v .:/webassembly -it -p 8080:8080 ubuntu:latest bash
  ```
  > Se inicia un contendor con Ubuntu que tiene bash
	y dentro una carpeta llamada "webassembly" que va a estar mapeada
	con la carpeta actual (desde donde ejecuté este comando de docker),
	entonces los archivos de la carpeta actual aparecerán en el contenedor
  y los que cree en el contenedor aparecen en la carpeta actual.

## Instalación de dependencias:
```
apt update && apt install -y nano nginx golang
```

## Configuración de Nginx:
```
nano /etc/nginx/nginx.conf
```
Al principio de la sección "http" agregar:
```
  server {
      listen  8080;
      location / {
          root '/webassembly/ejemplos';
          index index.html;
      }
      error_page  500 502 503 504 /50x.html;
      location = /50x.html {
          root    html;
      }
  }

```
  - Guardar y salir presionando:
    - ```Ctrl + x```
    - ```Y``` (la tecla "y")
    - ```Enter```


> Cuando llegue una petición a la URL de la ruta raíz "/", Nginx servirá el contenido de la carpeta '/webassembly/ejemplos'.

## Inicio de Nginx:
```
nginx -g 'daemon off;'
```

## Acceso a los ejemplos:
- http://localhost:8080
