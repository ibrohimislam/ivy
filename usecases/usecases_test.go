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

func (fur *FakeUserRepo) FindById(id string) *domains.User {
	if id == "1" {
		return &domains.User{"1", "satu", false, "dummyDepartementId1"}
	}

	return &domains.User{"2", "dua", false, "dummyDepartementId2"}
}

type FakeLogger struct{}

func (fl *FakeLogger) Log(_ string) error {
	return nil
}

func TestUsecasesSpec(t *testing.T) {
	fur := &FakeUserRepo{}
	fdr := &FakeDataRepo{domains.Entity{"", make([]domains.Data, 0), make([]domains.MetaData, 0), ""}}
	interactor := usecases.NewDataInteractor(fur, fdr, &FakeLogger{})

	Convey("testing PutEntity(...)", t, func() {
		data := make([]domains.Data, 0)
		metaData := make([]domains.MetaData, 0)
		userId := "1"

		interactor.PutEntity(userId, data, metaData)

		Convey("should be able to set id on pushed item", func() {
			So(fdr.data.Id, ShouldNotBeBlank)
		})

		Convey("user should owns their data", func() {
			So(fdr.data.DepartementId, ShouldEqual, fur.FindById(userId).DepartementId)
		})
	})

	Convey("testing Entity(...)", t, func() {
		departementId := "dummyDepartementId1"
		fdr.data = domains.Entity{"3", make([]domains.Data, 0), make([]domains.MetaData, 0), departementId}

		_, err := interactor.Entity("1", "")

		Convey("user should be able to get their data", func() {
			So(err, ShouldBeNil)
		})

		_, err = interactor.Entity("2", "")

		Convey("user should not be able to get another's data ", func() {
			So(err, ShouldNotBeNil)
		})
	})
}
