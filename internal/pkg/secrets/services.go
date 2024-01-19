package secrets

import (
	"fmt"

	_rsa "github.com/vekio/crypto/rsa"
	_sgen "github.com/vekio/rand/secretgen"
)

// autheliaSecrets generates secrets required for Authelia.
func autheliaSecrets() (authelia, error) {
	jwtSecret, err := _sgen.GenerateRandomAlphaNumeric(64)
	if err != nil {
		return authelia{}, fmt.Errorf("autheliaSecrets: failed to generate JWT secret: %w", err)
	}

	identityProviderOIDCHMACSecret, err := _sgen.GenerateRandomAlphaNumeric(64)
	if err != nil {
		return authelia{}, fmt.Errorf("autheliaSecrets: failed to generate OIDC HMAC secret: %w", err)
	}

	keyPair, err := _rsa.GenerateKeyPair(4096)
	if err != nil {
		return authelia{}, fmt.Errorf("autheliaSecrets: failed to generate RSA key pair: %w", err)
	}
	identityProviderIssuerPrivateKey, err := _rsa.PrivateKeyData(keyPair)
	if err != nil {
		return authelia{}, fmt.Errorf("autheliaSecrets: failed to extract private key data: %w", err)
	}

	sessionSecret, err := _sgen.GenerateRandomAlphaNumeric(64)
	if err != nil {
		return authelia{}, fmt.Errorf("autheliaSecrets: failed to generate session secret: %w", err)
	}

	storageEncryptionKey, err := _sgen.GenerateRandomAlphaNumeric(64)
	if err != nil {
		return authelia{}, fmt.Errorf("autheliaSecrets: failed to generate storage encryption key: %w", err)
	}

	autheliaSecrets := authelia{
		JWTSecret:                        jwtSecret,
		IdentityProviderOIDCHMACSecret:   identityProviderOIDCHMACSecret,
		IdentityProviderIssuerPrivateKey: identityProviderIssuerPrivateKey,
		SessionSecret:                    sessionSecret,
		StorageEncryptionKey:             storageEncryptionKey,
	}

	return autheliaSecrets, nil
}

// lldapSecrets generates secrets required for LLDAP.
func lldapSecrets() (lldap, error) {
	jwtSecret, err := _sgen.GenerateRandomAlphaNumeric(64)
	if err != nil {
		return lldap{}, fmt.Errorf("lldapSecrets: failed to generate JWT secret: %w", err)
	}

	ldapUserPass, err := _sgen.GenerateRandomAlphaNumeric(16)
	if err != nil {
		return lldap{}, fmt.Errorf("lldapSecrets: failed to generate LDAP user password: %w", err)
	}

	lldapSecrets := lldap{
		JWTSecret:    jwtSecret,
		LDAPUserPass: ldapUserPass,
	}

	return lldapSecrets, nil
}

// giteaSecrets generates secrets required for Gitea.
func giteaSecrets() (gitea, error) {
	oidcSecret, err := _sgen.GenerateRandomAlphaNumeric(64)
	if err != nil {
		return gitea{}, fmt.Errorf("giteaSecrets: failed to generate OIDC secret: %w", err)
	}

	giteaSecrets := gitea{
		OIDCSecret: oidcSecret,
	}

	return giteaSecrets, nil
}

// immichSecrets generates secrets required for Immich.
func immichSecrets() (immich, error) {
	dbPass, err := _sgen.GenerateRandomAlphaNumeric(16)
	if err != nil {
		return immich{}, fmt.Errorf("immichSecrets: failed to generate db pass: %w", err)
	}

	oidcSecret, err := _sgen.GenerateRandomAlphaNumeric(64)
	if err != nil {
		return immich{}, fmt.Errorf("immichSecrets: failed to generate OIDC secret: %w", err)
	}

	immichSecrets := immich{
		DBPass:     dbPass,
		OIDCSecret: oidcSecret,
	}

	return immichSecrets, nil
}
