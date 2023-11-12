traefik-config:
	bash traefik.sh

traefik-up:
	@docker compose -f ./traefik/compose.yml --env-file .env up -d
traefik-down:
	@docker compose -f ./traefik/compose.yml --env-file .env down -v
traefik-logs:
	@docker compose -f ./traefik/compose.yml --env-file .env logs -f
traefik-pull:
	@docker compose -f ./traefik/compose.yml --env-file .env pull
traefik-update: traefik-pull traefik-up
