package main

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/envoyproxy/envoy/contrib/golang/filters/http/source/go/pkg/api"
)

var UpdateUpstreamBody = "upstream response body updated by the simple plugin"

type filter struct {
	callbacks api.FilterCallbackHandler
	path      string
	config    *config
}

func (f *filter) sendLocalReplyInternal() api.StatusType {
	headers := make(map[string]string)
	body := fmt.Sprintf("%s, path: %s\r\n", f.config.echoBody, f.path)
	f.callbacks.SendLocalReply(200, body, headers, -1, "test-from-go")
	return api.LocalReply
}

func (f *filter) DecodeHeaders(header api.RequestHeaderMap, endStream bool) api.StatusType {
	f.path, _ = header.Get(":path")
	if f.path == "/localreply_by_config" {
		return f.sendLocalReplyInternal()
	}
	intn := rand.Intn(100)
	if intn%2 == 1 {
		header.Set("x-harness-cluster", "prod1")
		_ = fmt.Sprint("prod1 header set")
	} else {
		header.Set("x-harness-cluster", "prod2")
		_ = fmt.Sprint("prod2 header set")
	}
	return api.Continue
}

func (f *filter) DecodeData(buffer api.BufferInstance, endStream bool) api.StatusType {
	return api.Continue
}

func (f *filter) DecodeTrailers(trailers api.RequestTrailerMap) api.StatusType {
	return api.Continue
}

func (f *filter) EncodeHeaders(header api.ResponseHeaderMap, endStream bool) api.StatusType {
	if f.path == "/update_upstream_response" {
		header.Set("Content-Length", strconv.Itoa(len(UpdateUpstreamBody)))
	}
	header.Set("Rsp-Header-From-Go", "bar-test")
	return api.Continue
}

func (f *filter) EncodeData(buffer api.BufferInstance, endStream bool) api.StatusType {
	if f.path == "/update_upstream_response" {
		if endStream {
			buffer.SetString(UpdateUpstreamBody)
		} else {
			// TODO implement buffer->Drain, buffer.SetString means buffer->Drain(buffer.Len())
			buffer.SetString("")
		}
	}
	return api.Continue
}

func (f *filter) EncodeTrailers(trailers api.ResponseTrailerMap) api.StatusType {
	return api.Continue
}

func (f *filter) OnDestroy(reason api.DestroyReason) {
}
