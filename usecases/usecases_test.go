package usecases_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/ibrohimislam/ivy/domains"
	"github.com/ibrohimislam/ivy/usecases"
)

type FakeDataRepo struct {
	data domains.Entity
}

func (fdr *FakeDataRepo) Store(e domains.Entity) error {
	fdr.data = e
	return nil
}

func (fdr *FakeDataRepo) FindById(_ string) *domains.Entity {
	return &fdr.data
}

type FakeUserRepo struct{}

func (fur *FakeUserRepo) Store(_ domains.User) error {
	return nil
}

func (fur *FakeUserRepo) FindById(_ string) *domains.User {
	return &domains.User{"", "", false}
}

type FakeLogger struct{}

func (fl *FakeLogger) Log(_ string) error {
	return nil
}

func TestUsecasesSpec(t *testing.T) {
	fdr := &FakeDataRepo{domains.Entity{"", make([]domains.Data, 0), make([]domains.MetaData, 0), ""}}
	interactor := usecases.NewDataInteractor(&FakeUserRepo{}, fdr, &FakeLogger{})

	Convey("testing PutEntity", t, func() {
		data := make([]domains.Data, 0)
		metaData := make([]domains.MetaData, 0)
		ownerId := "dummy_userid"

		interactor.PutEntity(ownerId, data, metaData)

		Convey("should be able to set id on pushed item", func() {
			So(fdr.data.Id, ShouldNotBeBlank)
		})

		Convey("user should owns their data", func() {
			So(fdr.data.OwnerId, ShouldEqual, ownerId)
		})
	})
}
