// Copyright (c) 2020 InfraCloud Technologies
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package activity

import (
	"github.com/wschroederga/msbotbuilder-go/schema"
	"github.com/pkg/errors"
)

// TurnContext wraps the Activity received and provides operations for the user
// program of this SDK.
//
// The return value is Activity as provided by the client program, to be send to the connector service.
type TurnContext struct {
	Activity schema.Activity
}

// SendActivity sends an activity to user.
// TODO: Change comment
func (t *TurnContext) SendActivity(options ...MsgOption) (schema.Activity, error) {
	activity, err := applyMsgOptions(schema.Activity{Type: schema.Message}, options...)
	if err != nil {
		return activity, errors.Wrap(err, "Failed to apply MsgOptions.")
	}
	return ApplyConversationReference(activity, GetCoversationReference(t.Activity), false), nil
}

func applyMsgOptions(activity schema.Activity, options ...MsgOption) (schema.Activity, error) {
	for _, opt := range options {
		if err := opt(&activity); err != nil {
			return activity, err
		}
	}
	return activity, nil
}
