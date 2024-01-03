package services

var repoConfig string = "/home/alberto/src/homelab/services"

func AvailableServices() []string {
	return []string{AUTHELIA, GITEA, IMMICH, LLDAP, TRAEFIK}
}
