module iunite.club

require (
	github.com/DataDog/datadog-go v0.0.0-20180822151419-281ae9f2d895 // indirect
	github.com/Microsoft/go-winio v0.4.11 // indirect
	github.com/NYTimes/gziphandler v1.0.1 // indirect
	github.com/StackExchange/wmi v0.0.0-20180725035823-b12b22c5341f // indirect
	github.com/armon/go-metrics v0.0.0-20180917152333-f0300d1749da // indirect
	github.com/armon/go-radix v1.0.0 // indirect
	github.com/beorn7/perks v0.0.0-20180321164747-3a771d992973 // indirect
	github.com/boltdb/bolt v1.3.1 // indirect
	github.com/circonus-labs/circonus-gometrics v2.2.4+incompatible // indirect
	github.com/circonus-labs/circonusllhist v0.0.0-20180430145027-5eb751da55c6 // indirect
	github.com/coredns/coredns v1.2.3 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/elazarl/go-bindata-assetfs v1.0.0 // indirect
	github.com/emicklei/go-restful v2.8.0+incompatible
	github.com/emicklei/go-restful-openapi v1.0.0
	github.com/go-log/log v0.1.0
	github.com/go-ole/go-ole v1.2.1 // indirect
	github.com/go-openapi/spec v0.17.0
	github.com/golang/protobuf v1.2.0
	github.com/hashicorp/go-discover v0.0.0-20180831154906-f9c9239562a8 // indirect
	github.com/hashicorp/go-immutable-radix v1.0.0 // indirect
	github.com/hashicorp/go-memdb v0.0.0-20180223233045-1289e7fffe71 // indirect
	github.com/hashicorp/go-msgpack v0.0.0-20150518234257-fa3f63826f7c // indirect
	github.com/hashicorp/go-retryablehttp v0.0.0-20180718195005-e651d75abec6 // indirect
	github.com/hashicorp/go-sockaddr v0.0.0-20180320115054-6d291a969b86 // indirect
	github.com/hashicorp/go-syslog v0.0.0-20170829120034-326bf4a7f709 // indirect
	github.com/hashicorp/go-version v1.0.0 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/hashicorp/hil v0.0.0-20170627220502-fa9f258a9250 // indirect
	github.com/hashicorp/memberlist v0.1.0 // indirect
	github.com/hashicorp/net-rpc-msgpackrpc v0.0.0-20151116020338-a14192a58a69 // indirect
	github.com/hashicorp/raft v1.0.0 // indirect
	github.com/hashicorp/raft-boltdb v0.0.0-20171010151810-6e5ba93211ea // indirect
	github.com/iron-kit/go-ironic v0.1.1
	github.com/iron-kit/monger v0.2.1
	github.com/juju/ratelimit v1.0.1
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/micro/go-api v0.3.1
	github.com/micro/go-log v0.0.0-20170512141327-cbfa9447f9b6
	github.com/micro/go-micro v0.11.0
	github.com/micro/go-plugins v0.14.1
	github.com/micro/go-web v0.4.0
	github.com/micro/micro v0.14.0 // indirect
	github.com/mitchellh/cli v1.0.0 // indirect
	github.com/mitchellh/copystructure v1.0.0 // indirect
	github.com/mitchellh/go-testing-interface v1.0.0 // indirect
	github.com/pascaldekloe/goe v0.0.0-20180627143212-57f6aae5913c // indirect
	github.com/prometheus/client_golang v0.9.0 // indirect
	github.com/prometheus/client_model v0.0.0-20180712105110-5c3871d89910 // indirect
	github.com/prometheus/common v0.0.0-20181015124227-bcb74de08d37 // indirect
	github.com/prometheus/procfs v0.0.0-20181005140218-185b4288413d // indirect
	github.com/qiniu/api.v7 v7.2.4+incompatible
	github.com/qiniu/x v7.0.8+incompatible // indirect
	github.com/satori/go.uuid v1.2.0
	github.com/shirou/gopsutil v2.17.12+incompatible // indirect
	github.com/shirou/w32 v0.0.0-20160930032740-bb4de0191aa4 // indirect
	github.com/stretchr/objx v0.1.1 // indirect
	github.com/tv42/httpunix v0.0.0-20150427012821-b75d8614f926 // indirect
	github.com/uber-go/atomic v1.3.2 // indirect
	go.uber.org/ratelimit v0.0.0-20180316092928-c15da0234277
	golang.org/x/crypto v0.0.0-20180927165925-5295e8364332
	gopkg.in/mgo.v2 v2.0.0-20180705113604-9856a29383ce
	gopkg.in/vmihailenco/msgpack.v2 v2.9.1 // indirect
	qiniupkg.com/x v7.0.8+incompatible // indirect
)

replace github.com/micro/go-api => github.com/iron-kit/go-api v0.3.2

replace go.uber.org/ratelimit v0.0.0-20180316092928-c15da0234277 => github.com/uber-go/ratelimit v0.0.0-20180316092928-c15da0234277

replace github.com/iron-kit/go-ironic v0.1.1 => ../github.com/iron-kit/go-ironic

replace github.com/iron-kit/monger v0.2.1 => ../github.com/iron-kit/monger
