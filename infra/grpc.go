// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package infra holds useful helpers for csi driver server plugin
package infra

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"k8s.io/klog/v2"
)

// LogInterceptor returns a new unary server interceptors that performs request
// and response logging.
func LogInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		start := time.Now()
		if klog.V(5).Enabled() {
			deadline, _ := ctx.Deadline()
			klog.V(5).InfoS("request", "method", info.FullMethod, "deadline", time.Until(deadline).String())
		}
		resp, err := handler(ctx, req)
		if klog.V(5).Enabled() {
			s, _ := status.FromError(err)
			klog.V(5).InfoS("response", "method", info.FullMethod, "duration", time.Since(start).String(), "code", s.Code(), "message", s.Message())
		}
		return resp, err
	}
}
