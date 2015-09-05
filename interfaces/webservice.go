package interfaces

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ibrohimislam/ivy/domains"
	"github.com/ibrohimislam/ivy/usecases"
)

type DataInteractor interface {
	Entity(userId, entityId string) (*domains.Entity, error)
	PutEntity(userId string, dataSet []domains.Data, metaData []domains.MetaData)
}

type WebserviceHandler struct {
	dataInteractor usecases.DataInteractor
}

func NewWebserviceHandler(dataInteractor usecases.DataInteractor) *WebserviceHandler {
	return &WebserviceHandler{dataInteractor: dataInteractor}
}

func (handler WebserviceHandler) ReadData(res http.ResponseWriter, req *http.Request) {
	userId := req.FormValue("userId")
	entityId := req.FormValue("entityId")

	entity, _ := handler.dataInteractor.Entity(userId, entityId)
	entityMarshal, _ := json.Marshal(entity)

	fmt.Fprintln(res, entityMarshal)
}
