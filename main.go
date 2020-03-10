package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"sort"
	"strings"
)

// input - общая переменная для ввода текста
var input string

// арреи с маленькими буквами русского алфавита
var russianLowerCase = []string{"а", "б", "в", "г", "д", "е", "ё", "ж", "з", "и", "й", "к", "л", "м", "н", "о", "п", "р", "с", "т", "у", "ф", "х", "ц", "ч", "ш", "щ", "ъ", "ы", "ь", "э", "ю", "я"}

// арреи с большими буквами русского алфавита
var russianUpperCase = []string{"А", "Б", "В", "Г", "Д", "Е", "Ё", "Ж", "З", "И", "Й", "К", "Л", "М", "Н", "О", "П", "Р", "С", "Т", "У", "Ф", "Х", "Ц", "Ч", "Ш", "Щ", "Ъ", "Ы", "Ь", "Э", "Ю", "Я"}

// копилка буковок-структур
var bukvyVnalichii = []Bukovka{}

// Bukovka - свойства каждой буквы: ее символ, порядковый номер в алфавите-аррее, кол-во в тексте
type Bukovka struct {
	Simvol      string
	Poziciya    int
	Kolichestvo int
}

// хз какая проверка для файла
func check(e error) {
	if e != nil {
		panic(e)
	}
}

//
// получает текст в строке, считает сколько всего символов
func skolkoBukv(text string) {
	// пихает все руны символов в массив
	textArray := []rune(text)
	// считает сколько получилось рун в массиве
	dlinaText := len(textArray)
	// закоменченное говно
	// fmt.Println("текст")
	// fmt.Printf("\"%v\"\n\n", text)
	fmt.Printf("имеет\n%v символов\n\n", dlinaText)
}

// считает количество повторов буквы в стринге, bukvaNow это подсчитывающаяся буква, а poiskSchet - число повторов буквы
func bukvaCount(bukvaNow string) int {
	poiskSchet := strings.Count(input, bukvaNow)
	return poiskSchet
}

// сортирует элементы массива bukvyVnalichii, начиная от самого часто повторяющегося
func sorter() {
	sort.Slice(bukvyVnalichii, func(i, j int) bool { return bukvyVnalichii[i].Kolichestvo > bukvyVnalichii[j].Kolichestvo })
}

// выводит текст, используя данные буковок-структур массива bukvyVnalichii, k - означает общее количество повторений всех букв в массиве
func printer(k int) {
	fmt.Printf("было найдено\n")
	fmt.Printf("%v букв алфавита\n\n", len(bukvyVnalichii))
	// пока не кончатся буквы-структуры массива bukvyVnalichii
	for _, bukAr := range bukvyVnalichii {
		// пора считать процентное соотношение каждой буквы относительно остальных
		var percent float64
		percent = 100 / float64(k) * float64(bukAr.Kolichestvo)
		fmt.Printf("\"%v\" = %v (%v %%)\n", bukAr.Simvol, bukAr.Kolichestvo, math.Round(percent))
	}
}

func alfavit() {
	// 1. создай массив алфавита, где каждая буква имеет 3 свойства: имя, позицию и частоту встречания

	// пока не кончатся элементы массива russianLowerCase (буквы в стрингах), i это индекс текущего элемента в массиве
	for i, bukva := range russianLowerCase {
		// b1 считает кол-во повторений буквы в аррее для маленьких букв
		b1 := bukvaCount(bukva)
		// b2 использует текущую позицию маленькой буквы, чтобы проверить кол-во повторений для ее большой версии
		b2 := bukvaCount(russianUpperCase[i])
		// временная переменная буквы, суммарное повторение малых и больших версий буквы в тексте
		kolichestvoBukvy := b1 + b2
		// если буковка хоть раз появилась в тексте, ее данные следует оформить в структуру и добавить в аррей bukvyVnalichii
		if kolichestvoBukvy > 0 {
			k := Bukovka{Simvol: bukva, Poziciya: i, Kolichestvo: kolichestvoBukvy}
			bukvyVnalichii = append(bukvyVnalichii, k)
			// fmt.Println(k)
		}
	}
	// fmt.Println(bukvyVnalichii)

	// 2. заставь эту сучку считать процентики
	// переменная для накопления суммы всех повторений, нужна для расчета удельной единицы
	var kolichestvoVsehBukv int
	// пока не кончатся элементы массива bukvyVnalichii (буковки-стуктуры), i это индекс текущего элемента в массиве
	for i, bukAr := range bukvyVnalichii {
		// плюсуй кол-во текущей буквы к количеству всех имеющихся сейчас буквы
		kolichestvoVsehBukv += bukAr.Kolichestvo
		// а как только дошел до последней, врубай сортировщик массива
		if i == len(bukvyVnalichii)-1 {
			// вот, родименький
			sorter()
			// ну теперь когда все посчитали, можно и выводить выводим тут, т.к надо построчно
			printer(kolichestvoVsehBukv)
			// fmt.Printf("а вот эта буковка %v - последняя в аррее\n", bukvyVnalichii[i].Simvol)
			// fmt.Printf("общее кол-во найденных букв - %v\n", kolichestvoVsehBukv)

		}
	}

}

func main() {
	fileName := os.Args[1]
	data, err := ioutil.ReadFile(fileName)
	check(err)
	input = string(data)
	// fmt.Println(input)
	// считаем количество символов в нем
	skolkoBukv(input)

	// считаем конкретные буквы, создаем алфавит
	alfavit()
}
