package services

var ServicesMap = Services{
	AUTHELIA: &Service{
		Name:     AUTHELIA,
		Priority: 4,
	},
	GITEA: &Service{
		Name:     GITEA,
		Priority: 99,
	},
	IMMICH: &Service{
		Name:     IMMICH,
		Priority: 99,
	},
	JELLYFIN: &Service{
		Name:     JELLYFIN,
		Priority: 99,
	},
	LLDAP: &Service{
		Name:     LLDAP,
		Priority: 2,
	},
	PROTONMAIL_BRIDGE: &Service{
		Name:     PROTONMAIL_BRIDGE,
		Priority: 3,
	},
	TRAEFIK: &Service{
		Name:     TRAEFIK,
		Priority: 1,
	},
}
