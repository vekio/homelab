# LLDAP

LLDAP - Light LDAP implementation for authentication

## Instalación

1. Crear el grupo `admins` e incluir en él al usuario `admin`.

## Acceso

Unicamente el usuarios `admin` tiene acceso al panel de administración. La
constraseña del usuario `admin` se establece en un secreto que se pasa al
contenedor la primera vez que se arranca, LLDAP_LDAP_USER_PASS_FILE.

<!--
Si se cambia la contraseña al usuario `admin` será necesario:

- informar la nueva en el fichero `secrets.yml` y generar los nuevos secrets
- eliminar los contenedores y volver a levantar los servicios
 -->

Temporalmente el acceso al portal es público pero se deberá restringir a para
que solo sea accesible por un equipo o de manera local.

## Reference

- [Docker Hub](https://hub.docker.com/r/lldap/lldap)
- [Docs](https://github.com/lldap/lldap/blob/main/README.md)
