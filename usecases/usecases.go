package usecases

import (
	"fmt"
	"strings"

	"github.com/nu7hatch/gouuid"

	"github.com/ibrohimislam/ivy/domains"
)

type Logger interface {
	Log(message string) error
}

type DataInteractor struct {
	userRepository domains.UserRepository
	dataRepository domains.DataRepository
	logger         Logger
}

func NewDataInteractor(userRepository domains.UserRepository, dataRepository domains.DataRepository, logger Logger) *DataInteractor {
	return &DataInteractor{userRepository: userRepository, dataRepository: dataRepository, logger: logger}
}

func (interactor *DataInteractor) Entity(userId, entityId string) (*domains.Entity, error) {
	user := interactor.userRepository.FindById(userId)
	entity := interactor.dataRepository.FindById(entityId)

	if user.DepartementId != entity.DepartementId {
		message := "User #%i is not allowed to see items entity #%i"
		err := fmt.Errorf(message, user.Id, entity.Id)
		interactor.logger.Log(err.Error())
		return &domains.Entity{}, err
	}

	return &domains.Entity{entity.Id, entity.DataSet, entity.MetaData, entity.DepartementId}, nil
}

func (interactor *DataInteractor) PutEntity(userId string, dataSet []domains.Data, metaData []domains.MetaData) {
	user := interactor.userRepository.FindById(userId)
	err := interactor.dataRepository.Store(domains.Entity{interactor.getSlug(), dataSet, metaData, user.DepartementId})

	if err != nil {
		interactor.logger.Log(err.Error())
	}
}

func (interactor *DataInteractor) getSlug() string {
	uuid, _ := uuid.NewV4()
	return strings.Replace(uuid.String(), "-", "", -1)
}
