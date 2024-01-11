# Immich

Immich - Self-hosted photo and video backup

## Intalación

1. Acceder al servicio y dar de alta el usuario `homelab` que será el admin
2. Ir a administración y configurar el OAuth
    - Issuer URL: https://auth.{{DOMAIN}}/.well-known/openid-configuration
    - Client ID: immich
    - Client Secret: coger secreto del configuration.yml
    - Scope: openid email profile
    - Auto Register: Check
    - Auto Launch: Check
3. Deshabilitar Password Authentication
4. Iniciar sesión con `homelab` en Authelia y se vinculará con la cuenta admin
5. Iniciar sesión con usuario personal
6. Intalar aplicación movil

## Referencias

- [Docs](https://immich.app/docs/overview/introduction)
