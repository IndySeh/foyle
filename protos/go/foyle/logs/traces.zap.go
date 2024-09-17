// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: foyle/logs/traces.proto

package logspb

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	_ "github.com/stateful/runme/v3/pkg/api/gen/proto/go/runme/parser/v1"
	_ "github.com/jlewi/foyle/protos/go/foyle/v1alpha1"
	_ "github.com/stateful/runme/v3/pkg/api/gen/proto/go/runme/runner/v1"
	_ "google.golang.org/protobuf/types/known/structpb"
	go_uber_org_zap_zapcore "go.uber.org/zap/zapcore"
	github_com_golang_protobuf_ptypes "github.com/golang/protobuf/ptypes"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (m *Trace) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "id" // field id = 1
	enc.AddString(keyName, m.Id)

	keyName = "end_time" // field end_time = 3
	if t, err := github_com_golang_protobuf_ptypes.Timestamp(m.EndTime); err == nil {
		enc.AddTime(keyName, t)
	}

	keyName = "start_time" // field start_time = 2
	if t, err := github_com_golang_protobuf_ptypes.Timestamp(m.StartTime); err == nil {
		enc.AddTime(keyName, t)
	}

	keyName = "generate" // field generate = 4
	if ov, ok := m.GetData().(*Trace_Generate); ok {
		_ = ov
		if ov.Generate != nil {
			var vv interface{} = ov.Generate
			if marshaler, ok := vv.(go_uber_org_zap_zapcore.ObjectMarshaler); ok {
				enc.AddObject(keyName, marshaler)
			}
		}
	}

	keyName = "eval_mode" // field eval_mode = 6
	enc.AddBool(keyName, m.EvalMode)

	keyName = "spans" // field spans = 8
	enc.AddArray(keyName, go_uber_org_zap_zapcore.ArrayMarshalerFunc(func(aenc go_uber_org_zap_zapcore.ArrayEncoder) error {
		for _, rv := range m.Spans {
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

func (m *Span) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "id" // field id = 1
	enc.AddString(keyName, m.Id)

	keyName = "rag" // field rag = 2
	if ov, ok := m.GetData().(*Span_Rag); ok {
		_ = ov
		if ov.Rag != nil {
			var vv interface{} = ov.Rag
			if marshaler, ok := vv.(go_uber_org_zap_zapcore.ObjectMarshaler); ok {
				enc.AddObject(keyName, marshaler)
			}
		}
	}

	return nil
}

func (m *RAGSpan) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "query" // field query = 1
	enc.AddString(keyName, m.Query)

	keyName = "results" // field results = 2
	enc.AddArray(keyName, go_uber_org_zap_zapcore.ArrayMarshalerFunc(func(aenc go_uber_org_zap_zapcore.ArrayEncoder) error {
		for _, rv := range m.Results {
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

func (m *GenerateTrace) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "request" // field request = 1
	if m.Request != nil {
		var vv interface{} = m.Request
		if marshaler, ok := vv.(go_uber_org_zap_zapcore.ObjectMarshaler); ok {
			enc.AddObject(keyName, marshaler)
		}
	}

	keyName = "response" // field response = 2
	if m.Response != nil {
		var vv interface{} = m.Response
		if marshaler, ok := vv.(go_uber_org_zap_zapcore.ObjectMarshaler); ok {
			enc.AddObject(keyName, marshaler)
		}
	}

	return nil
}

func (m *LogEntries) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "lines" // field lines = 1
	enc.AddArray(keyName, go_uber_org_zap_zapcore.ArrayMarshalerFunc(func(aenc go_uber_org_zap_zapcore.ArrayEncoder) error {
		for _, rv := range m.Lines {
			_ = rv
			aenc.AppendString(rv)
		}
		return nil
	}))

	keyName = "resource_version" // field resource_version = 2
	enc.AddString(keyName, m.ResourceVersion)

	return nil
}

func (m *GetTraceRequest) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "id" // field id = 1
	enc.AddString(keyName, m.Id)

	return nil
}

func (m *GetTraceResponse) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "trace" // field trace = 1
	if m.Trace != nil {
		var vv interface{} = m.Trace
		if marshaler, ok := vv.(go_uber_org_zap_zapcore.ObjectMarshaler); ok {
			enc.AddObject(keyName, marshaler)
		}
	}

	return nil
}

func (m *GetBlockLogRequest) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "id" // field id = 1
	enc.AddString(keyName, m.Id)

	return nil
}

func (m *GetBlockLogResponse) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "block_log" // field block_log = 1
	if m.BlockLog != nil {
		var vv interface{} = m.BlockLog
		if marshaler, ok := vv.(go_uber_org_zap_zapcore.ObjectMarshaler); ok {
			enc.AddObject(keyName, marshaler)
		}
	}

	return nil
}

func (m *GetLLMLogsRequest) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "trace_id" // field trace_id = 1
	enc.AddString(keyName, m.TraceId)

	keyName = "log_file" // field log_file = 2
	enc.AddString(keyName, m.LogFile)

	return nil
}

func (m *GetLLMLogsResponse) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "request_html" // field request_html = 1
	enc.AddString(keyName, m.RequestHtml)

	keyName = "response_html" // field response_html = 2
	enc.AddString(keyName, m.ResponseHtml)

	keyName = "request_json" // field request_json = 3
	enc.AddString(keyName, m.RequestJson)

	keyName = "response_json" // field response_json = 4
	enc.AddString(keyName, m.ResponseJson)

	return nil
}
