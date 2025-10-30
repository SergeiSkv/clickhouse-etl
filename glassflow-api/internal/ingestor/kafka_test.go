package ingestor

import (
	"testing"

	"github.com/IBM/sarama"
	"github.com/nats-io/nats.go"
	"github.com/stretchr/testify/assert"
)

func TestConvertKafkaToNATSHeaders(t *testing.T) {
	processor := &KafkaMsgProcessor{}

	headers := []sarama.RecordHeader{
		{Key: []byte("trace-id"), Value: []byte("123")},
		{Key: []byte("empty"), Value: []byte{}},
		{Key: []byte("nil"), Value: nil},
	}

	got := processor.convertKafkaToNATSHeaders(headers)

	assert.Equal(t, nats.Header{
		"trace-id": []string{"123"},
		"empty":    []string{""},
	}, got)
}
