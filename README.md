# Google Token Source Chain

gcp-token-chain package contains an implementation to have multiple `oauth2.TokenSource`s

## Installation

~~~~
go get github.com/dakimura/gcp-token-chain
~~~~

Or you can manually git clone the repository to
`$(go env GOPATH)/src/github.com/dakimura/gcp-token-chain`.

## Usage

inject some token sources for testing, caching or for some other reasons

```
package example

import (
	"github.com/dakimura/gcp-token-chain/oauth2"
	googleoauth2 "golang.org/x/oauth2"
)
func main(){
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
```