package uservalidapi

import (
	"context"
	"io/ioutil"
	"testing"

	"github.com/Shopify/sarama"
	mock_sarama "github.com/Shopify/sarama/mocks"
	log "github.com/sirupsen/logrus"
	"github.com/golang/mock/gomock"
	mock_usergrpc "gitlab.ozon.dev/Hostile359/homework-1/internal/api/userapi/uservalidapi/mocks"
	pb "gitlab.ozon.dev/Hostile359/homework-1/pkg/api"
)

type userValidApiFixture struct {
	Ctx context.Context
	client *mock_usergrpc.MockUserClient
	producer *mock_sarama.SyncProducer
	userValidApi pb.UserServer
}

func setUp(t *testing.T) userValidApiFixture {
	log.SetOutput(ioutil.Discard)
	var fixture userValidApiFixture
	fixture.client = mock_usergrpc.NewMockUserClient(gomock.NewController(t))
	saramaCfg := sarama.NewConfig()
	saramaCfg.Producer.Return.Successes = true
	fixture.producer = mock_sarama.NewSyncProducer(t, saramaCfg)
	fixture.userValidApi = New(fixture.client, fixture.producer)

	return fixture
}
