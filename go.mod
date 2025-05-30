module github.com/captainpryce/go-authcrunch

go 1.21

require (
	github.com/crewjam/saml v0.4.14
	github.com/emersion/go-sasl v0.0.0-20231106173351-e73c9f7bad43
	github.com/emersion/go-smtp v0.21.0
	github.com/go-ldap/ldap/v3 v3.4.7
	github.com/golang-jwt/jwt/v4 v4.5.0
	github.com/google/go-cmp v0.6.0
	github.com/google/uuid v1.6.0
	github.com/greenpau/versioned v1.0.30
	github.com/iancoleman/strcase v0.3.0
	github.com/skip2/go-qrcode v0.0.0-20200617195104-da1b6568686e
	github.com/urfave/cli/v2 v2.27.1
	go.uber.org/zap v1.27.0
	golang.org/x/crypto v0.22.0
	golang.org/x/net v0.24.0
	gopkg.in/yaml.v3 v3.0.1
  github.com/greenpau/go-authcrunch v1.1.4
)

replace (
  github.com/greenpau/go-authcrunch v1.1.4 => github.com/captainpryce/go-authcrunch v1.1.7
)

require (
	github.com/Azure/go-ntlmssp v0.0.0-20221128193559-754e69321358 // indirect
	github.com/beevik/etree v1.3.0 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.4 // indirect
	github.com/crewjam/httperr v0.2.0 // indirect
	github.com/go-asn1-ber/asn1-ber v1.5.5 // indirect
	github.com/jonboulle/clockwork v0.4.0 // indirect
	github.com/mattermost/xml-roundtrip-validator v0.1.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/rogpeppe/go-internal v1.11.0 // indirect
	github.com/russellhaering/goxmldsig v1.4.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/stretchr/testify v1.8.4 // indirect
	github.com/xrash/smetrics v0.0.0-20240312152122-5f08fbb34913 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/sys v0.19.0 // indirect
)
