# Homelab


## TODO

CLI: Conf

- init():
    - comprobar si existe config.yml en la carpeta de configuraciones
        - si no existe leer el fichero plantilla config.yml y crear el fichero config.yml
        - si existe verificar que existe un contexto seleccionado y que el contexto seleccionado es válido
    - comprobar si existen carpetas a servicios que necesitan configuración
        - si no existe crear las carpetas
        - si existe comprobar que en las carpetas estan los enlaces simbólicos a los archivos de configuracion de los servicios
- template config.yml
- el comando init de los servicios tomará las plantillas de los ficheros de
    configuración de la ruta de configuración (~/.config/homelab/servicio)
    - reemplazará los ficheros de configuración en los destinos finales
