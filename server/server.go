// Copyright © 2017 Mike Hudgins <mchudgins@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package server

import (
	"context"
	"time"

	"github.com/coreos/etcd/clientv3"
	"go.uber.org/zap"
)

type webmonitor struct {
	etcd       *clientv3.Client
	cert       string
	key        string
	clientCert string
	clientKey  string
}

func New(logger *zap.Logger, listenPort string, cert string, key string, fInsecure bool) (*webmonitor, error) {

	etcd, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"localhost:2379"},
		DialTimeout: 2 * time.Second,
		AutoSyncInterval: 60 * time.Second,
	})
	if err != nil {
		return nil, err
	}

	mon := &webmonitor{etcd: etcd}

	return mon, nil
}

func (w *webmonitor) Run(ctx context.Context) {

}
