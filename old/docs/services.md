# Servicios

## Iniciar

*Si quieres tener el servicio smtp de protonmail ejecutar `homelab protonmail-bridge init`*

1. Ejecutar `homelab init`
2. Ejecutar `homelab allup`

## Servicios disponibles

- [Authelia](/docs/services/authelia.md) - Authentication and authorization server
    - https://auth.{{.DOMAIN}}
- [Gitea](/docs/services/gitea.md) - Self-hosted Git service
- [Immich](/docs/services/immich.md) - Self-hosted photo and video backup
- [Jellyfin](/docs/services/jellyfin.md) - Media System
- [LLDAP](/docs/services/lldap.md) - Light LDAP implementation for authentication
    - https://lldap.{{.DOMAIN}}
- [Traefik](/docs/services/traefik.md) - HTTP reverse proxy and load balancer.
    - https://traefik.{{.DOMAIN}}
