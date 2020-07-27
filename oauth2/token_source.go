package oauth2

import (
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
)

var baseErr = errors.New("token error")

// ChainedTokenSource is an implementation of Token Source to try to get a token from multiple TokenSources and
// returns the first found token
type ChainedTokenSource struct {
	tokenSources []oauth2.TokenSource
}

// NewChainedTokenSource returns a newly initialized ChainedTokenSource struct
func NewChainedTokenSource(tokenSources ...oauth2.TokenSource) oauth2.TokenSource {
	return &ChainedTokenSource{tokenSources: tokenSources}
}

// Token tries to get a token from each TokenSource and returns the first found token.
func (cts *ChainedTokenSource) Token() (*oauth2.Token, error) {
	err := baseErr

	for _, ts := range cts.tokenSources {
		token, err2 := ts.Token()
		if err2 == nil {
			return token, nil
		}
		err = errors.Wrap(err, err2.Error())
	}

	return nil, err
}
