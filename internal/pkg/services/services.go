package services

var composeFile = "/tmp"

var repoConfig = "/tmp"

func Available() map[string]Service {

	services := map[string]Service{
		AUTHELIA:          autheliaSrv,
		GITEA:             giteaSrv,
		IMMICH:            immichSrv,
		JELLYFIN:          jellyfinSrv,
		LLDAP:             lldapSrv,
		PROTONMAIL_BRIDGE: protonmailBridgeSrv,
		TRAEFIK:           traefikSrv,
	}

	return services
}
