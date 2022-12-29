module github.com/EverlongProject/i18n4go

go 1.16

require (
	github.com/EverlongProject/go-i18n v1.8.1
	github.com/cloudfoundry/jibber_jabber v0.0.0-20151120183258-bcc4c8345a21 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/golang/protobuf v0.0.0-20160817174113-f592bd283e9e // indirect
	github.com/hpcloud/tail v1.0.0 // indirect
	github.com/nicksnyder/go-i18n v1.10.1 // indirect
	github.com/onsi/ginkgo v1.11.0
	github.com/onsi/gomega v1.4.1
	github.com/pivotal-cf-experimental/jibber_jabber v0.0.0-20151120183258-bcc4c8345a21
	github.com/urfave/cli v1.22.7
	golang.org/x/net v0.4.0 // indirect
	golang.org/x/sync v0.0.0-20220722155255-886fb9371eb4
	gopkg.in/fsnotify.v1 v1.4.7 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

// Old versions of this library (anything before 0.3.8) have a known security vulnerability, see https://github.com/golang/go/issues/56152.
// replace golang.org/x/net => golang.org/x/net v0.4.0
replace golang.org/x/text/encoding/unicode => golang.org/x/text/encoding/unicode v0.3.3

replace golang.org/x/net/http2 => golang.org/x/net/http2 v0.0.0-20211209124913-491a49abca63

replace golang.org/x/crypto => golang.org/x/crypto v0.0.0-20201216223049-8b5274cf687f

replace golang.org/x/crypto/ssh => golang.org/x/crypto/ssh v0.0.0-20200220183623-bac4c82f6975

replace golang.org/x/text => golang.org/x/text v0.3.3

replace golang.org/x/net => golang.org/x/net v0.4.0
