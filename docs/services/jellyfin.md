# Jellyfin

Jellyfin - Media System

## Instalación

1. Idioma: Español
2. Admin account (mas adelante será eliminada):
    - `install`
    - random
3. Idioma: Español; Castellano. País: España.
4. Permitir conexiones remotas OK
5. Plugins:
    - Instalar: "LDAP Authentication". Reiniciar
6. LDAP Server Settings:
    - LDAP Server: lldap
    - LDAP Port: 3890
    - Secure LDAP: No Check
    - StartTLS: No Check
    - Skip SSL/TLS Verification: No Check
    - Allow Password Change: No Check
    - Password Reset Url: Empty
    - LDAP Bind User: uid=admin,ou=people,dc={{.SLD}},dc={{.TLD}}
    - User Password: LLDAP_LDAP_USER_PASS_FILE
    - LDAP Base DN: dc={{.SLD}},dc={{.TLD}}
    - Probar haciendo click en "Save and Test LDAP Server Settings"
7. LDAP User Settings:
    - User:
        - LDAP Search Filter: (memberOf=cn=jellyfin_users,ou=groups,dc={{.SLD}},dc={{.TLD}})
        - El resto dejarlo por defecto
    - Administrators:
        - LDAP Admin Base DN: dc={{.SLD}},dc={{.TLD}}
        - (memberOf=cn=admins,ou=groups,dc={{.SLD}},dc={{.TLD}})
    - Probar haciendo click en "Save and Test LDAP Filter Settings"
8. Jellyfin User Settings:
    - Enable User Creation: Check
    - Enable access to all libraries: Check
9. Entrar con el usuario `homelab` y darle permisos de admin
10. Eliminar al usuario `install`

## Referencias

- [Docker Hub](https://hub.docker.com/r/jellyfin/jellyfin)
- [Docs](https://jellyfin.org/docs/)
