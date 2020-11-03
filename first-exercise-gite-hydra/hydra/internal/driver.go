package internal

import (
	"context"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ory/hydra/x"
	"github.com/ory/x/sqlcon/dockertest"

	"github.com/jmoiron/sqlx"

	"github.com/ory/viper"

	"github.com/ory/hydra/driver"
	"github.com/ory/hydra/driver/configuration"
	"github.com/ory/hydra/jwk"
	"github.com/ory/x/logrusx"
)

func resetConfig() {
	viper.Set(configuration.ViperKeyWellKnownKeys, nil)
	viper.Set(configuration.ViperKeySubjectTypesSupported, nil)
	viper.Set(configuration.ViperKeyDefaultClientScope, nil)
	viper.Set(configuration.ViperKeyDSN, nil)
	viper.Set(configuration.ViperKeyEncryptSessionData, true)
	viper.Set(configuration.ViperKeyBCryptCost, nil)
	viper.Set(configuration.ViperKeyAdminListenOnHost, nil)
	viper.Set(configuration.ViperKeyAdminListenOnPort, nil)
	viper.Set(configuration.ViperKeyPublicListenOnHost, nil)
	viper.Set(configuration.ViperKeyPublicListenOnPort, nil)
	viper.Set(configuration.ViperKeyConsentRequestMaxAge, nil)
	viper.Set(configuration.ViperKeyAccessTokenLifespan, nil)
	viper.Set(configuration.ViperKeyRefreshTokenLifespan, nil)
	viper.Set(configuration.ViperKeyIDTokenLifespan, nil)
	viper.Set(configuration.ViperKeyAuthCodeLifespan, nil)
	viper.Set(configuration.ViperKeyScopeStrategy, nil)
	viper.Set(configuration.ViperKeyGetCookieSecrets, nil)
	viper.Set(configuration.ViperKeyGetSystemSecret, nil)
	viper.Set(configuration.ViperKeyLogoutRedirectURL, nil)
	viper.Set(configuration.ViperKeyLoginURL, nil)
	viper.Set(configuration.ViperKeyConsentURL, nil)
	viper.Set(configuration.ViperKeyErrorURL, nil)
	viper.Set(configuration.ViperKeyPublicURL, nil)
	viper.Set(configuration.ViperKeyIssuerURL, nil)
	viper.Set(configuration.ViperKeyOAuth2ClientRegistrationURL, nil)
	viper.Set(configuration.ViperKeyAllowTLSTerminationFrom, nil)
	viper.Set(configuration.ViperKeyAccessTokenStrategy, nil)
	viper.Set(configuration.ViperKeySubjectIdentifierAlgorithmSalt, nil)
	viper.Set(configuration.ViperKeyOIDCDiscoverySupportedClaims, nil)
	viper.Set(configuration.ViperKeyOIDCDiscoverySupportedScope, nil)
	viper.Set(configuration.ViperKeyOIDCDiscoveryUserinfoEndpoint, nil)

	viper.Set(configuration.ViperKeyBCryptCost, "4")
	viper.Set(configuration.ViperKeySubjectIdentifierAlgorithmSalt, "00000000")
	viper.Set(configuration.ViperKeyGetSystemSecret, []string{"000000000000000000000000000000000000000000000000"})
	viper.Set(configuration.ViperKeyGetCookieSecrets, []string{"000000000000000000000000000000000000000000000000"})

	viper.Set(configuration.ViperKeyLogLevel, "debug")
}

func NewConfigurationWithDefaults() *configuration.ViperProvider {
	resetConfig()
	return configuration.NewViperProvider(logrusx.New("", ""), true, nil).(*configuration.ViperProvider)
}

func NewConfigurationWithDefaultsAndHTTPS() *configuration.ViperProvider {
	resetConfig()
	return configuration.NewViperProvider(logrusx.New("", ""), false, nil).(*configuration.ViperProvider)
}

func NewRegistryMemory(c *configuration.ViperProvider) *driver.RegistryMemory {
	viper.Set(configuration.ViperKeyLogLevel, "debug")
	r := driver.NewRegistryMemory().WithConfig(c)
	_ = r.Init()
	return r.(*driver.RegistryMemory)
}

func NewRegistrySQLFromDB(c *configuration.ViperProvider, db *sqlx.DB) *driver.RegistrySQL {
	viper.Set(configuration.ViperKeyLogLevel, "debug")
	r := driver.NewRegistrySQL().WithConfig(c).(*driver.RegistrySQL).WithDB(db)
	_ = r.Init()
	return r.(*driver.RegistrySQL)
}

func NewRegistrySQLFromURL(t *testing.T, url string) *driver.RegistrySQL {
	conf := NewConfigurationWithDefaults()
	viper.Set(configuration.ViperKeyDSN, url)
	reg, err := driver.NewRegistry(conf)
	require.NoError(t, err)
	r, ok := reg.(*driver.RegistrySQL)
	require.True(t, ok)
	require.NoError(t, r.Init())
	return r
}

func CleanAndMigrate(reg *driver.RegistrySQL) func(*testing.T) {
	return func(t *testing.T) {
		x.CleanSQL(t, reg.DB())
		require.NoError(t, reg.Persister().MigrateUp(context.Background()))
		t.Log("clean and migrate done")
	}
}

func ConnectToMySQL(t *testing.T) string {
	c := dockertest.ConnectToTestMySQLPop(t)
	url := c.URL()
	if !strings.HasPrefix(url, "mysql://") {
		url = "mysql://" + url
	}
	require.NoError(t, c.Close())
	return url
}

func ConnectToPG(t *testing.T) string {
	c := dockertest.ConnectToTestPostgreSQLPop(t)
	require.NoError(t, c.Close())
	return c.URL()
}

func ConnectToCRDB(t *testing.T) string {
	c := dockertest.ConnectToTestCockroachDBPop(t)
	url := c.URL()
	if !strings.HasPrefix(url, "cockroach") {
		url = "cockroach://" + strings.Split(url, "://")[1]
	}
	require.NoError(t, c.Close())
	return url
}

func ConnectDatabases(t *testing.T) (pg, mysql, crdb *driver.RegistrySQL, clean func(*testing.T)) {
	var pgURL, mysqlURL, crdbURL string
	wg := sync.WaitGroup{}

	wg.Add(3)
	go func() {
		pgURL = ConnectToPG(t)
		t.Log("Pg done")
		wg.Done()
	}()
	go func() {
		mysqlURL = ConnectToMySQL(t)
		t.Log("myssql done")
		wg.Done()
	}()
	go func() {
		crdbURL = ConnectToCRDB(t)
		t.Log("cddb done")
		wg.Done()
	}()
	t.Log("beginning to wait")
	wg.Wait()
	t.Log("done waiting")

	pg = NewRegistrySQLFromURL(t, pgURL)
	mysql = NewRegistrySQLFromURL(t, mysqlURL)
	crdb = NewRegistrySQLFromURL(t, crdbURL)
	dbs := []*driver.RegistrySQL{pg, mysql, crdb}

	clean = func(t *testing.T) {
		wg := sync.WaitGroup{}

		wg.Add(len(dbs))
		for _, db := range dbs {
			go func(db *driver.RegistrySQL) {
				defer wg.Done()
				CleanAndMigrate(db)(t)
			}(db)
		}
		wg.Wait()
	}
	clean(t)
	return
}

func MustEnsureRegistryKeys(r driver.Registry, key string) {
	if err := jwk.EnsureAsymmetricKeypairExists(context.Background(), r, new(veryInsecureRS256Generator), key); err != nil {
		panic(err)
	}
}
