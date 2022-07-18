// Copyright (c) All respective contributors to the Peridot Project. All rights reserved.
// Copyright (c) 2021-2022 Rocky Enterprise Software Foundation, Inc. All rights reserved.
// Copyright (c) 2021-2022 Ctrl IQ, Inc. All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice,
// this list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice,
// this list of conditions and the following disclaimer in the documentation
// and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its contributors
// may be used to endorse or promote products derived from this software without
// specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
// ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE
// LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
// CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
// SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
// INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
// CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
// ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
// POSSIBILITY OF SUCH DAMAGE.

package temporalutils

import (
	"context"
	"crypto/tls"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.temporal.io/api/workflowservice/v1"
	"go.temporal.io/sdk/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"os"
	"strings"
	"time"
)

func AddFlags(pflags *pflag.FlagSet) {
	pflags.String("temporal.hostport", "", "Temporal connection string")
}

func NewClient(opts client.Options) (client.Client, error) {
	temporalHostPort := "temporal:7233"

	customTemporalHostPort := viper.GetString("temporal.hostport")
	if customTemporalHostPort != "" {
		temporalHostPort = customTemporalHostPort
		viper.Set("temporal.hostport", temporalHostPort)
	}

	if strings.Contains(temporalHostPort, ":443") {
		opts.ConnectionOptions = client.ConnectionOptions{
			DialOptions: []grpc.DialOption{
				grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})),
			},
		}
	}

	logrus.Infof("Using Temporal instance at %s", temporalHostPort)

	opts.HostPort = temporalHostPort

	bycNs := os.Getenv("BYC_NS")
	temporalNamespace := os.Getenv("TEMPORAL_NAMESPACE")
	if temporalNamespace != "" {
		bycNs = temporalNamespace
	}
	if opts.Namespace != "" {
		bycNs = opts.Namespace
	}
	if bycNs == "" {
		bycNs = "default"
	}

	nscl, err := client.NewNamespaceClient(opts)
	if err != nil {
		return nil, err
	}
	dur := 5 * 24 * time.Hour
	err = nscl.Register(context.TODO(), &workflowservice.RegisterNamespaceRequest{
		Namespace:                        bycNs,
		WorkflowExecutionRetentionPeriod: &dur,
	})
	if err != nil && !strings.Contains(err.Error(), "Namespace already exists") {
		return nil, err
	}
	opts.Namespace = bycNs
	viper.Set("temporal.namespace", bycNs)

	return client.NewClient(opts)
}
