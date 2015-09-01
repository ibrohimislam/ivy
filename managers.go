package ivy

type DataAdapter interface {
	Save() error
	Load(key string) error
}

type DataManager struct {
	dataRepo DataAdapter
}

func NewDataManager(dataRepo DataAdapter) *DataManager {
	return &DataManager{dataRepo: dataRepo}
}

func (dm *DataManager) Load(key string) error {
	return dm.dataRepo.Load(key)
}

func (dm *DataManager) Save() error {
	return dm.dataRepo.Save();
}