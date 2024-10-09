// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: foyle/logs/sessions.proto

package logspb

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	_ "github.com/jlewi/foyle/protos/go/foyle/v1alpha1"
	_ "github.com/stateful/runme/v3/pkg/api/gen/proto/go/runme/runner/v1"
	go_uber_org_zap_zapcore "go.uber.org/zap/zapcore"
	github_com_golang_protobuf_ptypes "github.com/golang/protobuf/ptypes"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (m *Session) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "context_id" // field context_id = 1
	enc.AddString(keyName, m.ContextId)

	keyName = "start_time" // field start_time = 2
	if t, err := github_com_golang_protobuf_ptypes.Timestamp(m.StartTime); err == nil {
		enc.AddTime(keyName, t)
	}

	keyName = "end_time" // field end_time = 3
	if t, err := github_com_golang_protobuf_ptypes.Timestamp(m.EndTime); err == nil {
		enc.AddTime(keyName, t)
	}

	keyName = "log_events" // field log_events = 4
	enc.AddArray(keyName, go_uber_org_zap_zapcore.ArrayMarshalerFunc(func(aenc go_uber_org_zap_zapcore.ArrayEncoder) error {
		for _, rv := range m.LogEvents {
			_ = rv
			if rv != nil {
				var vv interface{} = rv
				if marshaler, ok := vv.(go_uber_org_zap_zapcore.ObjectMarshaler); ok {
					aenc.AppendObject(marshaler)
				}
			}
		}
		return nil
	}))

	keyName = "full_context" // field full_context = 5
	if m.FullContext != nil {
		var vv interface{} = m.FullContext
		if marshaler, ok := vv.(go_uber_org_zap_zapcore.ObjectMarshaler); ok {
			enc.AddObject(keyName, marshaler)
		}
	}

	keyName = "total_input_tokens" // field total_input_tokens = 6
	enc.AddInt32(keyName, m.TotalInputTokens)

	keyName = "total_output_tokens" // field total_output_tokens = 7
	enc.AddInt32(keyName, m.TotalOutputTokens)

	keyName = "generate_trace_ids" // field generate_trace_ids = 8
	enc.AddArray(keyName, go_uber_org_zap_zapcore.ArrayMarshalerFunc(func(aenc go_uber_org_zap_zapcore.ArrayEncoder) error {
		for _, rv := range m.GenerateTraceIds {
			_ = rv
			aenc.AppendString(rv)
		}
		return nil
	}))

	return nil
}

func (m *GetSessionRequest) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "context_id" // field context_id = 1
	enc.AddString(keyName, m.ContextId)

	return nil
}

func (m *GetSessionResponse) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "session" // field session = 1
	if m.Session != nil {
		var vv interface{} = m.Session
		if marshaler, ok := vv.(go_uber_org_zap_zapcore.ObjectMarshaler); ok {
			enc.AddObject(keyName, marshaler)
		}
	}

	return nil
}

func (m *ListSessionsRequest) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	return nil
}

func (m *ListSessionsResponse) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "sessions" // field sessions = 1
	enc.AddArray(keyName, go_uber_org_zap_zapcore.ArrayMarshalerFunc(func(aenc go_uber_org_zap_zapcore.ArrayEncoder) error {
		for _, rv := range m.Sessions {
			_ = rv
			if rv != nil {
				var vv interface{} = rv
				if marshaler, ok := vv.(go_uber_org_zap_zapcore.ObjectMarshaler); ok {
					aenc.AppendObject(marshaler)
				}
			}
		}
		return nil
	}))

	return nil
}

func (m *DumpExamplesRequest) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "output" // field output = 1
	enc.AddString(keyName, m.Output)

	return nil
}

func (m *DumpExamplesResponse) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "num_examples" // field num_examples = 1
	enc.AddInt32(keyName, m.NumExamples)

	keyName = "num_sessions" // field num_sessions = 2
	enc.AddInt32(keyName, m.NumSessions)

	return nil
}
