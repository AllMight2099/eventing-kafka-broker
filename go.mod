module knative.dev/eventing-kafka-broker

go 1.16

require (
	github.com/Shopify/sarama v1.31.1
	github.com/cloudevents/sdk-go/protocol/kafka_sarama/v2 v2.4.1
	github.com/cloudevents/sdk-go/v2 v2.8.0
	github.com/google/go-cmp v0.5.6
	github.com/google/uuid v1.3.0
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/pierdipi/sacura v0.0.0-20210302185533-982357fc042b
	github.com/rickb777/date v1.14.1
	github.com/stretchr/testify v1.7.0
	github.com/xdg-go/scram v1.1.0
	go.uber.org/atomic v1.9.0
	go.uber.org/multierr v1.6.0
	go.uber.org/zap v1.19.1
	google.golang.org/protobuf v1.27.1
	k8s.io/api v0.23.5
	k8s.io/apiextensions-apiserver v0.23.4
	k8s.io/apimachinery v0.23.5
	k8s.io/apiserver v0.23.4
	k8s.io/client-go v0.23.5
	k8s.io/utils v0.0.0-20220210201930-3a6ce19ff2f9
	knative.dev/eventing v0.31.1-0.20220516052256-d7a8a95792bd
	knative.dev/eventing-kafka v0.31.1-0.20220516052256-0dc7832e36b4
	knative.dev/hack v0.0.0-20220512014059-f4972b4daff9
	knative.dev/pkg v0.0.0-20220512013937-2d8305b2e59a
	knative.dev/reconciler-test v0.0.0-20220513144654-fdb392439fb5
	sigs.k8s.io/yaml v1.3.0
)
