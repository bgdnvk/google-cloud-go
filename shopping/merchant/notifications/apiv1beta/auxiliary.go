// Copyright 2024 Google LLC
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

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package notifications

import (
	notificationspb "cloud.google.com/go/shopping/merchant/notifications/apiv1beta/notificationspb"
	"google.golang.org/api/iterator"
)

// NotificationSubscriptionIterator manages a stream of *notificationspb.NotificationSubscription.
type NotificationSubscriptionIterator struct {
	items    []*notificationspb.NotificationSubscription
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*notificationspb.NotificationSubscription, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *NotificationSubscriptionIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *NotificationSubscriptionIterator) Next() (*notificationspb.NotificationSubscription, error) {
	var item *notificationspb.NotificationSubscription
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *NotificationSubscriptionIterator) bufLen() int {
	return len(it.items)
}

func (it *NotificationSubscriptionIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}