package data

type JsonDocument struct {
}

func (doc *JsonDocument) ConvertToXml() string {
	return "<xml></xml>"
}

type JsonDocumentAdapter struct {
	jsonDocument *JsonDocument
}

func (adapter *JsonDocumentAdapter) SendXmlData() {
	adapter.jsonDocument.ConvertToXml()
	println("Отправка данных")
}

// Данный пример показывает, что есть внешняя библиотека работающая с XMl,
// но у нас все данные JSON, следовательно надо создать адаптер, внутри
// которого мы бы кастили JSON в XML для работы с внешней библиотекой.
// Создав адаптер, теперь у нас есть структура, подходящая под интерфейс внешнего пакета.
