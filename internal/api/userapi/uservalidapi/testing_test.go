package uservalidapi

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	pb "gitlab.ozon.dev/Hostile359/homework-1/pkg/api"
	mock_usergrpc "gitlab.ozon.dev/Hostile359/homework-1/internal/api/userapi/uservalidapi/mocks"
)

type userValidApiFixture struct {
	Ctx context.Context
	client *mock_usergrpc.MockUserClient
	userValidApi pb.UserServer
}

func setUp(t *testing.T) userValidApiFixture {
	var fixture userValidApiFixture
	fixture.client = mock_usergrpc.NewMockUserClient(gomock.NewController(t))
	fixture.userValidApi = New(fixture.client)

	return fixture
}
