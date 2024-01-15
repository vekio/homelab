# Traefik

Traefik - HTTP reverse proxy and load balancer

## Acceso

Traefik no cuenta con autenticación por lo que uso [Authelia](/docs/services/authelia.md)
para controlar el acceso al dashboard de Traefik.

El único usuario que tiene acceso al dashboard de Traefik es `admin` ya que
pertenece al grupo `admins` del [LLDAP](/docs/services/lldap.md).

## Referencias

- [Docker Hub](https://hub.docker.com/_/traefik)
- [Docs](https://doc.traefik.io/traefik/)