package main

// К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.
//
//
//var justString string
//func someFunc() {
//  v := createHugeString(1 << 10)
//  justString = v[:100]
//}
//
//func main() {
//  someFunc()
//}

// Ответ
// Строка v остается в памяти и занимает она 1024 байта, переменная justString глобальная, она ссылается на локальную, значит ячейка памяти где хранится локальная переменная не будет очищена.

//var justString string

//func someFunc() {
//	v := createHugeString(1 << 10)
//	sliceBytes := []byte(v[:100]) // Чтобы отойти от старой строки конвертируем срез в байты, чтобы аллоцировать в памяти и потом создадим новую строку, тогда старую строку удалит сборщик мусора
//	justString = string(sliceBytes)
//}
//
//func main() {
//	someFunc()
//}
