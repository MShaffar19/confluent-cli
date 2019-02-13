package kafka

import (
	"fmt"
	"strings"

	"github.com/Shopify/sarama"

	"github.com/confluentinc/cli/shared"
)

// NewSaramaProducer returns a sarama.ClusterConsumerconfigured for the CLI config
func NewSaramaConsumer(group string, config *shared.Config) (sarama.ConsumerGroup, error) {
	kafka, err := config.KafkaClusterConfig()
	if err != nil {
		return nil, err
	}

	client, err := sarama.NewClient(strings.Split(kafka.Bootstrap, ","), saramaConf(kafka))
	if err != nil {
		return nil, err
	}
	return sarama.NewConsumerGroupFromClient(group, client)
}

// NewSaramaProducer returns a sarama.ClusterProducer configured for the CLI config
func NewSaramaProducer(config *shared.Config) (sarama.SyncProducer, error) {
	kafka, err := config.KafkaClusterConfig()
	if err != nil {
		return nil, err
	}
	return sarama.NewSyncProducer(strings.Split(kafka.Bootstrap, ","), saramaConf(kafka))
}

// ConsumerGroupHandler instances are used to handle individual topic-partition claims.
type GroupHandler struct{}

// Setup is run at the beginning of a new session, before ConsumeClaim.
func (*GroupHandler) Setup(_ sarama.ConsumerGroupSession) error { return nil }

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
// but before the offsets are committed for the very last time.
func (*GroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
// Once the Messages() channel is closed, the Handler must finish its processing
// loop and exit.
func (h GroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Println(string(msg.Value))
		sess.MarkMessage(msg, "")
	}
	return nil
}

// saramaConf converts KafkaClusterConfig to sarama.Config
func saramaConf(kafka shared.KafkaClusterConfig) *sarama.Config {
	saramaConf := sarama.NewConfig()
	saramaConf.Version = sarama.V1_1_0_0
	saramaConf.Net.TLS.Enable = true
	saramaConf.Net.SASL.Enable = true
	saramaConf.Net.SASL.User = kafka.APIKey
	saramaConf.Net.SASL.Password = kafka.APISecret

	saramaConf.Producer.Return.Successes = true
	saramaConf.Producer.Return.Errors = true

	saramaConf.Consumer.Offsets.Initial = sarama.OffsetOldest

	return saramaConf
}
