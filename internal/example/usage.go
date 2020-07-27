package example

import (
	"github.com/dakimura/gcp-token-chain/oauth2"
	googleoauth2 "golang.org/x/oauth2"
)

type CachedSomeTokenSource struct{}

func (csts *CachedSomeTokenSource) Token() (*googleoauth2.Token, error) {
	return nil, nil
}

type SomeTokenSource struct{}

func (csts *SomeTokenSource) Token() (*googleoauth2.Token, error) {
	return nil, nil
}

func main() {
	chainedTokenSource := oauth2.NewChainedTokenSource(
		&CachedSomeTokenSource{},
		&SomeTokenSource{},
		googleoauth2.StaticTokenSource(&googleoauth2.Token{AccessToken: "e.g. staticAccesstokenForTest"}),
	)

	// try to get token from each token source in order.
	token, err := chainedTokenSource.Token()
	if err != nil {
		print(err)
	}
	print(token)
}
