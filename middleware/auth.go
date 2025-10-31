package middleware

import (
	"fmt"
	"strings"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	/*
		c.Flags().StringVar(&a.clientID, "client-id", "example-app", "OAuth2 client ID of this application.")
		c.Flags().StringVar(&a.clientSecret, "client-secret", "example-app-secret", "OAuth2 client secret of this application.")
		c.Flags().BoolVar(&a.pkce, "pkce", true, "Use PKCE flow for the code exchange.")
		c.Flags().StringVar(&a.redirectURI, "redirect-uri", "http://127.0.0.1:5555/callback", "Callback URL for OAuth2 responses.")
		c.Flags().StringVar(&issuerURL, "issuer", "http://127.0.0.1:5556/dex", "URL of the OpenID Connect issuer.")
		c.Flags().StringVar(&listen, "listen", "http://127.0.0.1:5555", "HTTP(S) address to listen at.")
		c.Flags().StringVar(&tlsCert, "tls-cert", "", "X509 cert file to present when serving HTTPS.")
		c.Flags().StringVar(&tlsKey, "tls-key", "", "Private key for the HTTPS cert.")
		c.Flags().StringVar(&rootCAs, "issuer-root-ca", "", "Root certificate authorities for the issuer. Defaults to host certs.")
		c.Flags().BoolVar(&debug, "debug", false, "Print all request and responses from the OpenID Connect issuer.")
	*/

	return func(c *gin.Context) {
		provider, err := oidc.NewProvider(c, "http://127.0.0.1:5556/dex")
		if err != nil {
			fmt.Println("Could not get provider")
		}

		var verifier = provider.Verifier(&oidc.Config{ClientID: "example-app"})

		authHeader := c.GetHeader("Authorization")

		accToken := strings.TrimPrefix(authHeader, "Bearer")
		//fmt.Println(accToken)
		//fmt.Println(verifier)

		idToken, err := verifier.Verify(c, accToken)

		if err != nil {
			fmt.Println(err.Error())
			c.Next()
			return
		}

		var claims struct {
			Subject string `json:"sub"`
			Name    string `json:"name"`
		}

		if err := idToken.Claims(&claims); err != nil {
			// handle error
			fmt.Println(err.Error())
		} else {
			fmt.Print("Hello ")
			fmt.Println(claims.Name)
		}

		c.Next()
	}

}
