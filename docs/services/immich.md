# Immich

Immich - Self-hosted photo and video backup

## Intalación

1. Acceder al servicio y dar de alta el usuario admin
2. Ir a administración y configurar el OAuth
    - Issuer URL: https://auth.{{DOMAIN}}/.well-known/openid-configuration
    - Client ID: immich
    - Client Secret: coger secreto del configuration.yml
    - Scope: openid email profile
    - Auto Register: Check
    - Auto Launch: Check

## Referencias

- [Docs](https://immich.app/docs/overview/introduction)
