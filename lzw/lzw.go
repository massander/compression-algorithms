package lzw

import (
	"bytes"
	"fmt"
)

func Encode(data string) []int32 {
	// Задаем изначальный размер словаря
	dictSize := 256

	// Создаем начальный словарь из 256 Ascii символов
	dictionary := make(map[string]int32, dictSize)
	for i := 0; i < dictSize; i++ {
		// Хак чтобы сохранить []byte как строку
		dictionary[string([]byte{byte(i)})] = int32(i)
	}

	// Текущая последовательность
	current := make([]byte, 0)

	// Результат который вернет функция -> string(output)
	output := make([]int32, 0)

	for i := 0; i < len(data); i++ {
		// Считывание нового символа из изходной строки в текущую последовательность.
		current = append(current, data[i])

		// Если текущая последовательность существует в словаре, то
		// считывается следующий символ.
		if _, ok := dictionary[string(current)]; ok {
			continue
		}

		// Добавление к результату текущей последовательности минус последний символ.
		output = append(output, dictionary[string(current[:len(current)-1])])

		// Добавление кода для новой последовательности в словарь.
		dictionary[string(current)] = int32(dictSize)
		dictSize++

		// Подготовка последовательности для переиспользования в следующей итерации.
		// Удаление всех элементов кроме последнего.
		current = current[len(current)-1:]
	}

	// В случае нечетного количества символов в исходной последователности,
	// код последнего символа добавляется к результату.
	if len(current) > 0 {
		output = append(output, dictionary[string(current)])
	}

	return output
}

func Decode(data []int32) (string, error) {
	// Создаем начальный словарь
	dictSize := 256
	dictionary := make(map[int32][]byte, dictSize)
	for i := 0; i < dictSize; i++ {
		dictionary[int32(i)] = []byte{byte(i)}
	}

	// В output будет записываться декодированные символы,
	// которые будут переведен в строку
	output := bytes.NewBufferString("")

	current := make([]byte, 0)

	for _, k := range data {
		var entry []byte
		if x, ok := dictionary[k]; ok {
			entry = x[:len(x):len(x)]
		} else if k == int32(dictSize) && len(current) > 0 {
			entry = append(current, current[0])
		} else {
			return output.String(), fmt.Errorf("lzw: bad compressed symbol %v", int(k))
		}
		output.Write(entry)

		if len(current) > 0 {
			current = append(current, entry[0])
			dictionary[int32(dictSize)] = current
			dictSize++
		}
		current = entry
	}
	return output.String(), nil
}
