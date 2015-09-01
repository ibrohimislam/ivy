package ivy_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/ibrohimislam/ivy"
)

type FakeDataRepo struct{}

func (ftr *FakeDataRepo) Save() error {
	return nil
}

func (ftr *FakeDataRepo) Load(_ string) error {
	return nil
}

func TestDataManagerSpec(t *testing.T) {
	manager := ivy.NewDataManager(&FakeDataRepo{})

	Convey("testing data manager", t, func() {
		Convey("error should be passed from repo", func() {
			err := manager.Load("test")
			So(err, ShouldBeNil)
			err = manager.Save()
			So(err, ShouldBeNil)
		})
	})
}
