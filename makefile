traefik-config:
	bash traefik.sh

traefik-up:
	@docker compose -f ./services/traefik/compose.yml --env-file .env up -d
traefik-down:
	@docker compose -f ./services/traefik/compose.yml --env-file .env down -v
traefik-logs:
	@docker compose -f ./services/traefik/compose.yml --env-file .env logs -f
traefik-pull:
	@docker compose -f ./services/traefik/compose.yml --env-file .env pull
traefik-update: traefik-pull traefik-up

authelia-up:
	@docker compose -f ./services/authelia/compose.yml --env-file .env up -d
authelia-down:
	@docker compose -f ./services/authelia/compose.yml --env-file .env down -v
authelia-logs:
	@docker compose -f ./services/authelia/compose.yml --env-file .env logs -f
authelia-pull:
	@docker compose -f ./services/authelia/compose.yml --env-file .env pull
authelia-update: authelia-pull authelia-up
