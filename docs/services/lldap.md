# LLDAP

LLDAP - Light LDAP implementation for authentication

## Instalación

1. Iniciar sesión con `admin`
2. Crear grupos:
    - `admins`
    - `jellyfin_users`
    - `gitea_users`
3. Crear usuarios:
    - `homelab`, grupos: `lldap_admin`, `admins`, `jellyfin_users`, `gitea_users`
    - `lldap`, groups: `lldap_strict_readonly` y `lldap_password_manager`
      La contraseña debe coincidir con el secreto LLDAP_LDAP_USER_PASS_FILE
4. Eliminar usuario `admin`
5. Crear usuarios personales

⚠️ Al crear usuarios hay que añadir le campo name (Display Name) que sino falla
la autenticación de jellyfin

## Acceso

Unicamente el usuario `homelab` tiene acceso al panel de administración.

Temporalmente el acceso al portal es público pero se deberá restringir a para
que solo sea accesible por un equipo o de manera local.

## Reference

- [Docker Hub](https://hub.docker.com/r/lldap/lldap)
- [Docs](https://github.com/lldap/lldap/blob/main/README.md)
