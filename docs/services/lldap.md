# LLDAP

LLDAP - Light LDAP implementation for authentication

## Instalación

1. Iniciar sesión con `admin`
2. Crear grupos:
    - `traefik_admin`
3. Crear usuarios:
    - `homelab`, grupos: `lldap_admin`, `traefik_admin`
    - `lldap`, groups: `lldap_strict_readonly` y `lldap_password_manager`
      La contraseña debe coincidir con el secreto LLDAP_LDAP_USER_PASS_FILE
4. Eliminar `admin`

## Acceso

Unicamente el usuario `homelab` tiene acceso al panel de administración.

Temporalmente el acceso al portal es público pero se deberá restringir a para
que solo sea accesible por un equipo o de manera local.

## Reference

- [Docker Hub](https://hub.docker.com/r/lldap/lldap)
- [Docs](https://github.com/lldap/lldap/blob/main/README.md)
