package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Field = zapcore.Field
type Level = zapcore.Level

const (
	// DebugLevel logs are typically voluminous, and are usually disabled in
	// production.
	DebugLevel = zapcore.DebugLevel
	// InfoLevel is the default logging priority.
	InfoLevel = zapcore.InfoLevel
	// WarnLevel logs are more important than Info, but don't need individual
	// human review.
	WarnLevel = zapcore.WarnLevel
	// ErrorLevel logs are high-priority. If an application is running smoothly,
	// it shouldn't generate any error-level logs.
	ErrorLevel = zapcore.ErrorLevel
	// DPanicLevel logs are particularly important errors. In development the
	// logger panics after writing the message.
	DPanicLevel = zapcore.DPanicLevel
	// PanicLevel logs a message, then panics.
	PanicLevel = zapcore.PanicLevel
	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel = zapcore.FatalLevel
)

var (

	// Binary constructs a field that carries an opaque binary blob.
	//
	// Binary data is serialized in an encoding-appropriate format. For example,
	// zap's JSON encoder base64-encodes binary blobs. To log UTF-8 encoded text,
	// use ByteString.
	Binary = zap.Binary

	// ByteString constructs a field that carries UTF-8 encoded text as a []byte.
	// To log opaque binary blobs (which aren't necessarily valid UTF-8), use
	// Binary.
	ByteString = zap.ByteString

	// String adds a string-valued key:value pair to a Span.LogFields() record
	String = zap.String

	// Bool adds a bool-valued key:value pair to a Span.LogFields() record
	Bool = zap.Bool

	// Int adds an int-valued key:value pair to a Span.LogFields() record
	Int = zap.Int

	// Int32 adds an int32-valued key:value pair to a Span.LogFields() record
	Int32 = zap.Int32

	// Int64 adds an int64-valued key:value pair to a Span.LogFields() record
	Int64 = zap.Int64

	// Uint32 adds a uint32-valued key:value pair to a Span.LogFields() record
	Uint32 = zap.Uint32

	// Uint64 adds a uint64-valued key:value pair to a Span.LogFields() record
	Uint64 = zap.Uint64

	// Float32 adds a float32-valued key:value pair to a Span.LogFields() record
	Float32 = zap.Float32

	// Float64 adds a float64-valued key:value pair to a Span.LogFields() record
	Float64 = zap.Float64

	// Error adds an error with the key "error" to a Span.LogFields() record
	Error = zap.Error

	// Object adds an object-valued key:value pair to a Span.LogFields() record
	Object = zap.Object

	// Noop creates a no-op log field that should be ignored by the tracer.
	Noop = zap.Skip

	// Any adds an any-valued key:value pair to a Span.LogFields() record
	Any = zap.Any
)
