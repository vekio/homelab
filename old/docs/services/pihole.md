# Pihole

Pihole - Network-wide Ad Blocking

## Acceso

Deshabilito la autenticación al dashboard de Pihole dejando la password en
blanco y uso [Authelia](/services/authelia.md) para controlar el acceso.

El único usuario que tiene acceso al dashboard de Pihole es `admin` ya que
pertenece al grupo `admins` del [LLDAP](/services/lldap.md).

## Referencias

- [Docker Hub](https://hub.docker.com/r/pihole/pihole)
- [Docs](https://docs.pi-hole.net/)
