package data_service

type AnalyticalDataService interface {
	SendXmlData()
}

type XmlDocument struct {
}

func (doc *XmlDocument) SendXmlData() {
	println("Отправка данных")
}
