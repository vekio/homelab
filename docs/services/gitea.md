# Gitea

Gitea - Self-hosted Git service

## Instalación

1. Acceder a gitea.{{.DOMAIN}}
2. Registrar usuario `homelab`
3. Ir al panel de administración. Acceder Add Authentication Source
4. Seleccionar OpenID Connect
    - Authentication Name: authelia
    - Client ID: gitea
    - Client Secret: GITEA_OIDC_SECRET
    - Icon URL: https://www.authelia.com/images/branding/logo-cropped.svg
    - Discovery: https://auth.{{.DOMAIN}}/.well-known/openid-configuration
    - Aditional scopes: `email profile groups`
    - Required Claim Name: `groups`
    - Required Claim Value: `gitea_users`
5. Iniciar sesion con `homelab` y linkearlo con el usuario inicial `homelab`

Los usuarios que quieran acceder a gitea deben estar en el grupo `gitea_users`.

## Referencias

- [Docker Hub](https://hub.docker.com/r/gitea/gitea/)
- [Docs](https://docs.gitea.com/)
