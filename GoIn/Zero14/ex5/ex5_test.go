package ex5

import (
    "errors"
    "testing"
)

func TestGetUTFLength(t *testing.T) {
    tests := []struct {
        input    []byte
        expected int
        err      error
    }{
        {[]byte("hello"), 5, nil},                          // обычная ASCII строка
        {[]byte("Привет"), 6, nil},                         // строка на русском языке
        {[]byte{0xff, 0xfe, 0xfd}, 0, ErrInvalidUTF8},      // некорректная строка
        {[]byte("こんにちは"), 5, nil},                      // строка на японском языке
        {[]byte("👍"), 1, nil},                             // emoji
        {[]byte(""), 0, nil},                               // пустая строка

        // Неполные многобайтовые последовательности
        {[]byte{0xc2}, 0, ErrInvalidUTF8},                  // Неполная 2-битная последовательность
        {[]byte{0xe2, 0x82}, 0, ErrInvalidUTF8},            // Неполная 3-битная последовательность
        {[]byte{0xf0, 0x9f, 0x92}, 0, ErrInvalidUTF8},      // Неполная 4-битная последовательность

        // Неправильные стартовые байты
        {[]byte{0x80}, 0, ErrInvalidUTF8},                  // Неправильный одиночный байт

        // Слишком большие значения для символов
        {[]byte{0xf8, 0x88, 0x80, 0x80}, 0, ErrInvalidUTF8},// Недопустимый символ за пределами диапазона UTF-8

        // Слишком длинные последовательности (переполнение)
        {[]byte{0xc1, 0x81}, 0, ErrInvalidUTF8},            // Переполнение для ASCII-символа

        // Одиночные байты, зарезервированные для продолжений
        {[]byte{0x80}, 0, ErrInvalidUTF8},                  // Недопустимый одиночный байт

        // Переполнение кодовой точки (overlong encoding)
        {[]byte{0xc0, 0x80}, 0, ErrInvalidUTF8},            // Переполнение для U+0000

        // Недопустимые диапазоны (суррогатные пары)
        {[]byte{0xed, 0xa0, 0x80}, 0, ErrInvalidUTF8},      // U+D800, начало верхнего диапазона
        {[]byte{0xed, 0xa0, 0x81}, 0, ErrInvalidUTF8},      // U+D801
        {[]byte{0xed, 0xaf, 0xbf}, 0, ErrInvalidUTF8},      // U+DBFF, конец верхнего диапазона
        {[]byte{0xed, 0xb0, 0x80}, 0, ErrInvalidUTF8},      // U+DC00, начало нижнего диапазона
        {[]byte{0xed, 0xb0, 0x81}, 0, ErrInvalidUTF8},      // U+DC01
        {[]byte{0xed, 0xbf, 0xbf}, 0, ErrInvalidUTF8},      // U+DFFF, конец нижнего диапазона
    }

    for _, tt := range tests {
        result, err := GetUTFLength(tt.input)
        if result != tt.expected || !errors.Is(err, tt.err) {
            t.Errorf("For input %v, expected length %d and error %v, got length %d and error %v",
                tt.input, tt.expected, tt.err, result, err)
        }
    }
}
