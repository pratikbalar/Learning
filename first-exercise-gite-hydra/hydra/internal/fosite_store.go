package internal

import (
	"context"

	"github.com/ory/hydra/client"
	"github.com/ory/hydra/driver"
)

func AddFositeExamples(r driver.Registry) {
	for _, c := range []client.Client{
		{
			ID:            "my-client",
			Secret:        "foobar",
			RedirectURIs:  []string{"http://localhost:3846/callback"},
			ResponseTypes: []string{"id_token", "code", "token"},
			GrantTypes:    []string{"implicit", "refresh_token", "authorization_code", "password", "client_credentials"},
			Scope:         "fosite,openid,photos,offline",
		},
		{
			ID:            "encoded:client",
			Secret:        "encoded&password",
			RedirectURIs:  []string{"http://localhost:3846/callback"},
			ResponseTypes: []string{"id_token", "code", "token"},
			GrantTypes:    []string{"implicit", "refresh_token", "authorization_code", "password", "client_credentials"},
			Scope:         "fosite,openid,photos,offline",
		},
	} {
		if err := r.ClientManager().CreateClient(context.Background(), &c); err != nil {
			panic(err)
		}
	}
}
