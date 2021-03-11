/*
Copyright 2021 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package versioned

import (
	"fmt"

	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
	eventingv1 "knative.dev/eventing/pkg/client/clientset/versioned/typed/eventing/v1"
	eventingv1beta1 "knative.dev/eventing/pkg/client/clientset/versioned/typed/eventing/v1beta1"
	flowsv1 "knative.dev/eventing/pkg/client/clientset/versioned/typed/flows/v1"
	flowsv1beta1 "knative.dev/eventing/pkg/client/clientset/versioned/typed/flows/v1beta1"
	messagingv1 "knative.dev/eventing/pkg/client/clientset/versioned/typed/messaging/v1"
	messagingv1beta1 "knative.dev/eventing/pkg/client/clientset/versioned/typed/messaging/v1beta1"
	sourcesv1 "knative.dev/eventing/pkg/client/clientset/versioned/typed/sources/v1"
	sourcesv1alpha1 "knative.dev/eventing/pkg/client/clientset/versioned/typed/sources/v1alpha1"
	sourcesv1alpha2 "knative.dev/eventing/pkg/client/clientset/versioned/typed/sources/v1alpha2"
	sourcesv1beta1 "knative.dev/eventing/pkg/client/clientset/versioned/typed/sources/v1beta1"
	sourcesv1beta2 "knative.dev/eventing/pkg/client/clientset/versioned/typed/sources/v1beta2"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	EventingV1beta1() eventingv1beta1.EventingV1beta1Interface
	EventingV1() eventingv1.EventingV1Interface
	FlowsV1beta1() flowsv1beta1.FlowsV1beta1Interface
	FlowsV1() flowsv1.FlowsV1Interface
	MessagingV1beta1() messagingv1beta1.MessagingV1beta1Interface
	MessagingV1() messagingv1.MessagingV1Interface
	SourcesV1alpha1() sourcesv1alpha1.SourcesV1alpha1Interface
	SourcesV1alpha2() sourcesv1alpha2.SourcesV1alpha2Interface
	SourcesV1beta1() sourcesv1beta1.SourcesV1beta1Interface
	SourcesV1beta2() sourcesv1beta2.SourcesV1beta2Interface
	SourcesV1() sourcesv1.SourcesV1Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	eventingV1beta1  *eventingv1beta1.EventingV1beta1Client
	eventingV1       *eventingv1.EventingV1Client
	flowsV1beta1     *flowsv1beta1.FlowsV1beta1Client
	flowsV1          *flowsv1.FlowsV1Client
	messagingV1beta1 *messagingv1beta1.MessagingV1beta1Client
	messagingV1      *messagingv1.MessagingV1Client
	sourcesV1alpha1  *sourcesv1alpha1.SourcesV1alpha1Client
	sourcesV1alpha2  *sourcesv1alpha2.SourcesV1alpha2Client
	sourcesV1beta1   *sourcesv1beta1.SourcesV1beta1Client
	sourcesV1beta2   *sourcesv1beta2.SourcesV1beta2Client
	sourcesV1        *sourcesv1.SourcesV1Client
}

// EventingV1beta1 retrieves the EventingV1beta1Client
func (c *Clientset) EventingV1beta1() eventingv1beta1.EventingV1beta1Interface {
	return c.eventingV1beta1
}

// EventingV1 retrieves the EventingV1Client
func (c *Clientset) EventingV1() eventingv1.EventingV1Interface {
	return c.eventingV1
}

// FlowsV1beta1 retrieves the FlowsV1beta1Client
func (c *Clientset) FlowsV1beta1() flowsv1beta1.FlowsV1beta1Interface {
	return c.flowsV1beta1
}

// FlowsV1 retrieves the FlowsV1Client
func (c *Clientset) FlowsV1() flowsv1.FlowsV1Interface {
	return c.flowsV1
}

// MessagingV1beta1 retrieves the MessagingV1beta1Client
func (c *Clientset) MessagingV1beta1() messagingv1beta1.MessagingV1beta1Interface {
	return c.messagingV1beta1
}

// MessagingV1 retrieves the MessagingV1Client
func (c *Clientset) MessagingV1() messagingv1.MessagingV1Interface {
	return c.messagingV1
}

// SourcesV1alpha1 retrieves the SourcesV1alpha1Client
func (c *Clientset) SourcesV1alpha1() sourcesv1alpha1.SourcesV1alpha1Interface {
	return c.sourcesV1alpha1
}

// SourcesV1alpha2 retrieves the SourcesV1alpha2Client
func (c *Clientset) SourcesV1alpha2() sourcesv1alpha2.SourcesV1alpha2Interface {
	return c.sourcesV1alpha2
}

// SourcesV1beta1 retrieves the SourcesV1beta1Client
func (c *Clientset) SourcesV1beta1() sourcesv1beta1.SourcesV1beta1Interface {
	return c.sourcesV1beta1
}

// SourcesV1beta2 retrieves the SourcesV1beta2Client
func (c *Clientset) SourcesV1beta2() sourcesv1beta2.SourcesV1beta2Interface {
	return c.sourcesV1beta2
}

// SourcesV1 retrieves the SourcesV1Client
func (c *Clientset) SourcesV1() sourcesv1.SourcesV1Interface {
	return c.sourcesV1
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfig will generate a rate-limiter in configShallowCopy.
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		if configShallowCopy.Burst <= 0 {
			return nil, fmt.Errorf("burst is required to be greater than 0 when RateLimiter is not set and QPS is set to greater than 0")
		}
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.eventingV1beta1, err = eventingv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.eventingV1, err = eventingv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.flowsV1beta1, err = flowsv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.flowsV1, err = flowsv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.messagingV1beta1, err = messagingv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.messagingV1, err = messagingv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.sourcesV1alpha1, err = sourcesv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.sourcesV1alpha2, err = sourcesv1alpha2.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.sourcesV1beta1, err = sourcesv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.sourcesV1beta2, err = sourcesv1beta2.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.sourcesV1, err = sourcesv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	var cs Clientset
	cs.eventingV1beta1 = eventingv1beta1.NewForConfigOrDie(c)
	cs.eventingV1 = eventingv1.NewForConfigOrDie(c)
	cs.flowsV1beta1 = flowsv1beta1.NewForConfigOrDie(c)
	cs.flowsV1 = flowsv1.NewForConfigOrDie(c)
	cs.messagingV1beta1 = messagingv1beta1.NewForConfigOrDie(c)
	cs.messagingV1 = messagingv1.NewForConfigOrDie(c)
	cs.sourcesV1alpha1 = sourcesv1alpha1.NewForConfigOrDie(c)
	cs.sourcesV1alpha2 = sourcesv1alpha2.NewForConfigOrDie(c)
	cs.sourcesV1beta1 = sourcesv1beta1.NewForConfigOrDie(c)
	cs.sourcesV1beta2 = sourcesv1beta2.NewForConfigOrDie(c)
	cs.sourcesV1 = sourcesv1.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.eventingV1beta1 = eventingv1beta1.New(c)
	cs.eventingV1 = eventingv1.New(c)
	cs.flowsV1beta1 = flowsv1beta1.New(c)
	cs.flowsV1 = flowsv1.New(c)
	cs.messagingV1beta1 = messagingv1beta1.New(c)
	cs.messagingV1 = messagingv1.New(c)
	cs.sourcesV1alpha1 = sourcesv1alpha1.New(c)
	cs.sourcesV1alpha2 = sourcesv1alpha2.New(c)
	cs.sourcesV1beta1 = sourcesv1beta1.New(c)
	cs.sourcesV1beta2 = sourcesv1beta2.New(c)
	cs.sourcesV1 = sourcesv1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
