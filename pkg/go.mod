module github.com/pulumi/pulumi/pkg/v3

go 1.17

replace (
	github.com/Sirupsen/logrus => github.com/sirupsen/logrus v1.5.0
	github.com/pulumi/pulumi/sdk/v3 => ../sdk
)

require (
	cloud.google.com/go/logging v1.0.0
	cloud.google.com/go/storage v1.15.0
	github.com/StackExchange/wmi v1.2.1 // indirect
	github.com/aws/aws-sdk-go v1.38.35
	github.com/blang/semver v3.5.1+incompatible
	github.com/davecgh/go-spew v1.1.1
	github.com/djherbis/times v1.2.0
	github.com/docker/docker v1.4.2-0.20200319182547-c7ad2b866182
	github.com/dustin/go-humanize v1.0.0
	github.com/gedex/inflector v0.0.0-20170307190818-16278e9db813
	github.com/gofrs/uuid v3.3.0+incompatible
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.5.2
	github.com/google/go-querystring v1.0.0
	github.com/google/pprof v0.0.0-20210226084205-cbba55b83ad5
	github.com/gorilla/mux v1.7.4
	github.com/grpc-ecosystem/grpc-opentracing v0.0.0-20180507213350-8e809c8a8645
	github.com/hashicorp/go-multierror v1.1.1
	github.com/hashicorp/hcl/v2 v2.3.0
	github.com/ijc/Gotty v0.0.0-20170406111628-a8b993ba6abd
	github.com/json-iterator/go v1.1.9
	github.com/mitchellh/copystructure v1.0.0
	github.com/mxschmitt/golang-combinations v1.0.0
	github.com/nbutton23/zxcvbn-go v0.0.0-20180912185939-ae427f1e4c1d
	github.com/opentracing/opentracing-go v1.1.0
	github.com/pgavlin/goldmark v1.1.33-0.20200616210433-b5eb04559386
	github.com/pulumi/pulumi/sdk/v3 v3.19.0
	github.com/rjeczalik/notify v0.9.2
	github.com/santhosh-tekuri/jsonschema/v5 v5.0.0
	github.com/sergi/go-diff v1.1.0
	github.com/shirou/gopsutil v3.21.7+incompatible
	github.com/shurcooL/httpfs v0.0.0-20190707220628-8d4bc4ba7749 // indirect
	github.com/shurcooL/vfsgen v0.0.0-20200824052919-0d455de96546 // indirect
	github.com/skratchdot/open-golang v0.0.0-20200116055534-eef842397966
	github.com/spf13/cobra v1.0.0
	github.com/spf13/pflag v1.0.3
	github.com/stretchr/testify v1.6.1
	github.com/tklauser/go-sysconf v0.3.8 // indirect
	github.com/tweekmonster/luser v0.0.0-20161003172636-3fa38070dbd7
	github.com/xeipuuv/gojsonschema v1.2.0
	github.com/zclconf/go-cty v1.3.1
	gocloud.dev v0.23.0
	gocloud.dev/secrets/hashivault v0.23.0
	golang.org/x/crypto v0.0.0-20210506145944-38f3c27a63bf
	golang.org/x/net v0.0.0-20210505214959-0714010a04ed
	golang.org/x/oauth2 v0.0.0-20210427180440-81ed05c6b58c
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	google.golang.org/api v0.46.0
	google.golang.org/genproto v0.0.0-20210506142907-4a47615972c2
	google.golang.org/grpc v1.37.0
	gopkg.in/AlecAivazis/survey.v1 v1.8.9-0.20200217094205-6773bdf39b7f
	gopkg.in/src-d/go-git.v4 v4.13.1
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c
	pgregory.net/rapid v0.4.7
	sourcegraph.com/sourcegraph/appdash v0.0.0-20190731080439-ebfcffb1b5c0
	sourcegraph.com/sourcegraph/appdash-data v0.0.0-20151005221446-73f23eafcf67 // indirect
)

// Generated by go mod tidy -go=1.17
require (
	cloud.google.com/go v0.81.0 // indirect
	github.com/Azure/azure-pipeline-go v0.2.3 // indirect
	github.com/Azure/azure-sdk-for-go v54.0.0+incompatible // indirect
	github.com/Azure/azure-storage-blob-go v0.13.0 // indirect
	github.com/Azure/go-ansiterm v0.0.0-20170929234023-d6e3b3328b78 // indirect
	github.com/Azure/go-autorest v14.2.0+incompatible // indirect
	github.com/Azure/go-autorest/autorest v0.11.18 // indirect
	github.com/Azure/go-autorest/autorest/adal v0.9.13 // indirect
	github.com/Azure/go-autorest/autorest/azure/auth v0.5.7 // indirect
	github.com/Azure/go-autorest/autorest/azure/cli v0.4.2 // indirect
	github.com/Azure/go-autorest/autorest/date v0.3.0 // indirect
	github.com/Azure/go-autorest/autorest/to v0.4.0 // indirect
	github.com/Azure/go-autorest/autorest/validation v0.3.1 // indirect
	github.com/Azure/go-autorest/logger v0.2.1 // indirect
	github.com/Azure/go-autorest/tracing v0.6.0 // indirect
	github.com/agext/levenshtein v1.2.1 // indirect
	github.com/apparentlymart/go-textseg v1.0.0 // indirect
	github.com/cheggaaa/pb v1.0.18 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.0 // indirect
	github.com/dimchansky/utfbom v1.1.1 // indirect
	github.com/emirpasic/gods v1.12.0 // indirect
	github.com/form3tech-oss/jwt-go v3.2.2+incompatible // indirect
	github.com/go-ole/go-ole v1.2.5 // indirect
	github.com/gofrs/flock v0.7.1 // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/snappy v0.0.3 // indirect
	github.com/google/go-cmp v0.5.5 // indirect
	github.com/google/uuid v1.2.0 // indirect
	github.com/google/wire v0.5.0 // indirect
	github.com/googleapis/gax-go/v2 v2.0.5 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.0 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
	github.com/hashicorp/go-sockaddr v1.0.2 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/hashicorp/vault/api v1.1.0 // indirect
	github.com/hashicorp/vault/sdk v0.2.0 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/jbenet/go-context v0.0.0-20150711004518-d14ea06fba99 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/jstemmer/go-junit-report v0.9.1 // indirect
	github.com/kballard/go-shellquote v0.0.0-20180428030007-95032a82bc51 // indirect
	github.com/kevinburke/ssh_config v0.0.0-20190725054713-01f96b0aa0cd // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.1 // indirect
	github.com/kr/pretty v0.3.0 // indirect
	github.com/mattn/go-colorable v0.1.6 // indirect
	github.com/mattn/go-ieproxy v0.0.1 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/mattn/go-runewidth v0.0.8 // indirect
	github.com/mgutz/ansi v0.0.0-20170206155736-9520e82c474b // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/go-ps v1.0.0 // indirect
	github.com/mitchellh/go-wordwrap v1.0.0 // indirect
	github.com/mitchellh/mapstructure v1.4.1 // indirect
	github.com/mitchellh/reflectwalk v1.0.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/opentracing/basictracer-go v1.0.0 // indirect
	github.com/pierrec/lz4 v2.6.0+incompatible // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/russross/blackfriday/v2 v2.0.1 // indirect
	github.com/ryanuber/go-glob v1.0.0 // indirect
	github.com/sabhiram/go-gitignore v0.0.0-20180611051255-d3107576ba94 // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/sirupsen/logrus v1.4.2 // indirect
	github.com/src-d/gcfg v1.4.0 // indirect
	github.com/texttheater/golang-levenshtein v0.0.0-20191208221605-eb6844b05fc6 // indirect
	github.com/tklauser/numcpus v0.2.3 // indirect
	github.com/uber/jaeger-client-go v2.22.1+incompatible // indirect
	github.com/uber/jaeger-lib v2.2.0+incompatible // indirect
	github.com/xanzy/ssh-agent v0.2.1 // indirect
	github.com/xeipuuv/gojsonpointer v0.0.0-20180127040702-4e3ac2762d5f // indirect
	github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415 // indirect
	go.opencensus.io v0.23.0 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	golang.org/x/lint v0.0.0-20201208152925-83fdc39ff7b5 // indirect
	golang.org/x/mod v0.4.2 // indirect
	golang.org/x/sys v0.0.0-20210817190340-bfb29a6856f2 // indirect
	golang.org/x/term v0.0.0-20201126162022-7de9c90e9dd1 // indirect
	golang.org/x/text v0.3.6 // indirect
	golang.org/x/time v0.0.0-20210220033141-f8bda1e9f3ba // indirect
	golang.org/x/tools v0.1.0 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.26.0 // indirect
	gopkg.in/square/go-jose.v2 v2.5.1 // indirect
	gopkg.in/src-d/go-billy.v4 v4.3.2 // indirect
	gopkg.in/warnings.v0 v0.1.2 // indirect
	gopkg.in/yaml.v2 v2.2.8 // indirect
)
