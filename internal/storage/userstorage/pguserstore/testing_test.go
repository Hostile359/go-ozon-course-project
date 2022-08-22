package pguserstore

import (
	"testing"

	"github.com/driftprogramming/pgxpoolmock"
	"github.com/golang/mock/gomock"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/app/userapp"
)

type userStoreFixture struct {
	userStorage userapp.Storage
	ctrl *gomock.Controller
	mockPool *pgxpoolmock.MockPgxPool
}

func setUp(t *testing.T) userStoreFixture {
	var fixture userStoreFixture
	fixture.ctrl = gomock.NewController(t)
	fixture.mockPool = pgxpoolmock.NewMockPgxPool(fixture.ctrl)
	fixture.userStorage = New(fixture.mockPool)

	return fixture
}

func (f *userStoreFixture) tearDown() {
	f.ctrl.Finish()
}
