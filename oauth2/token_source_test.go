package oauth2

import (
	"errors"
	"golang.org/x/oauth2"
	"reflect"
	"testing"
)

var SampleToken = &oauth2.Token{AccessToken: "foo"}

// errorTokenSource always returns an error
type errorTokenSource struct {
	errMsg string
}

func (ets *errorTokenSource) Token() (*oauth2.Token, error) {
	return nil, errors.New(ets.errMsg)
}

func TestChainedTokenSource_Token(t *testing.T) {
	tests := []struct {
		name         string
		tokenSources []oauth2.TokenSource
		want         *oauth2.Token
		wantErr      bool
	}{
		{
			name:         "single tokenSource",
			tokenSources: []oauth2.TokenSource{oauth2.StaticTokenSource(SampleToken)},
			want:         SampleToken,
			wantErr:      false,
		},
		{
			name: "multiple tokenSources, use the 2nd tokenSource",
			tokenSources: []oauth2.TokenSource{
				&errorTokenSource{errMsg: "test"},
				oauth2.StaticTokenSource(SampleToken),
			},
			want:    SampleToken,
			wantErr: false,
		},
		{
			name: "multiple tokenSources, all error",
			tokenSources: []oauth2.TokenSource{
				&errorTokenSource{errMsg: "test"},
				&errorTokenSource{errMsg: "test2"},
				&errorTokenSource{errMsg: "test3"},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			cts := &ChainedTokenSource{
				tokenSources: tt.tokenSources,
			}
			got, err := cts.Token()
			if (err != nil) != tt.wantErr {
				t.Errorf("Token() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Token() got = %v, want %v", got, tt.want)
			}
		})
	}
}
