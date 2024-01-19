package secrets

import (
	"fmt"

	_rsa "github.com/vekio/crypto/rsa"
	_sgen "github.com/vekio/rand/secretgen"
)

// autheliaSecrets generates various secrets required for Authelia.
// It returns an Authelia struct containing the secrets and an error if any.
func autheliaSecrets() (Authelia, error) {
	jwtSecret, err := _sgen.GenerateRandomAlphaNumeric(64)
	if err != nil {
		return Authelia{}, fmt.Errorf("autheliaSecrets: failed to generate JWT secret: %w", err)
	}

	identityProviderOIDCHMACSecret, err := _sgen.GenerateRandomAlphaNumeric(64)
	if err != nil {
		return Authelia{}, fmt.Errorf("autheliaSecrets: failed to generate OIDC HMAC secret: %w", err)
	}

	keyPair, err := _rsa.GenerateKeyPair(4096)
	if err != nil {
		return Authelia{}, fmt.Errorf("autheliaSecrets: failed to generate RSA key pair: %w", err)
	}
	identityProviderIssuerPrivateKey, err := _rsa.PrivateKeyData(keyPair)
	if err != nil {
		return Authelia{}, fmt.Errorf("autheliaSecrets: failed to extract private key data: %w", err)
	}

	sessionSecret, err := _sgen.GenerateRandomAlphaNumeric(64)
	if err != nil {
		return Authelia{}, fmt.Errorf("autheliaSecrets: failed to generate session secret: %w", err)
	}

	storageEncryptionKey, err := _sgen.GenerateRandomAlphaNumeric(64)
	if err != nil {
		return Authelia{}, fmt.Errorf("autheliaSecrets: failed to generate storage encryption key: %w", err)
	}

	autheliaSecrets := Authelia{
		JWTSecret:                        jwtSecret,
		IdentityProviderOIDCHMACSecret:   identityProviderOIDCHMACSecret,
		IdentityProviderIssuerPrivateKey: identityProviderIssuerPrivateKey,
		SessionSecret:                    sessionSecret,
		StorageEncryptionKey:             storageEncryptionKey,
	}

	return autheliaSecrets, nil
}

// lldapSecrets generates various secrets required for LLDAP.
func lldapSecrets() (Lldap, error) {
	jwtSecret, err := _sgen.GenerateRandomAlphaNumeric(64)
	if err != nil {
		return Lldap{}, fmt.Errorf("lldapSecrets: failed to generate JWT secret: %w", err)
	}

	ldapUserPass, err := _sgen.GenerateRandomAlphaNumeric(16)
	if err != nil {
		return Lldap{}, fmt.Errorf("lldapSecrets: failed to generate LDAP user password: %w", err)
	}

	lldapSecrets := Lldap{
		JWTSecret:    jwtSecret,
		LDAPUserPass: ldapUserPass,
	}

	return lldapSecrets, nil
}

// giteaSecrets generates various secrets required for Gitea.
func giteaSecrets() (Gitea, error) {
	oidcSecret, err := _sgen.GenerateRandomAlphaNumeric(64)
	if err != nil {
		return Gitea{}, fmt.Errorf("giteaSecrets: failed to generate OIDC secret: %w", err)
	}

	giteaSecrets := Gitea{
		OIDCSecret: oidcSecret,
	}

	return giteaSecrets, nil
}

// immichSecrets generates secrets required for Immich.
func immichSecrets() (Immich, error) {
	dbPass, err := _sgen.GenerateRandomAlphaNumeric(16)
	if err != nil {
		return Immich{}, fmt.Errorf("immichSecrets: failed to generate db pass: %w", err)
	}

	oidcSecret, err := _sgen.GenerateRandomAlphaNumeric(64)
	if err != nil {
		return Immich{}, fmt.Errorf("immichSecrets: failed to generate OIDC secret: %w", err)
	}

	immichSecrets := Immich{
		DBPass:     dbPass,
		OIDCSecret: oidcSecret,
	}

	return immichSecrets, nil
}
