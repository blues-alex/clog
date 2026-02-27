# clog

Цветной логгер для Go с поддержкой нескольких уровней логирования, записи в файл, ротацией и прогресс-барами.

## Возможности

- 5 уровней логирования: Error, Warning, Info, Success, Debug
- Цветной вывод в консоль
- Запись в файл (дублирование)
- Ротация логов по размеру с построчным обрезанием
- Progress wrappers для прогресс-баров
- Потокобезопасность (mutex)
- 9 цветов: Red, HiRed, Yellow, HiYellow, Green, HiGreen, Blue, Cyan, White

## Установка

```bash
go get github.com/blues-alex/clog
```

## Быстрый старт

```go
import "github.com/blues-alex/clog"

func main() {
    clog.SetEnableAll()
    clog.Error("Ошибка")
    clog.Warning("Предупреждение")
    clog.Info("Информация")
    clog.Success("Успех")
    clog.Debug("Отладка")
}
```

## Уровни логирования

| Функция | Описание | Цвет |
|---------|----------|------|
| `Error` | Ошибки | Красный |
| `Warning` | Предупреждения | Желтый |
| `Info` | Информация | Белый |
| `Success` | Успешные операции | Зеленый |
| `Debug` | Отладочная информация | Синий |

По умолчанию включены: `Error`, `Warning`

## Управление уровнями

```go
// Отдельные сеттеры
clog.SetEnableError(true)
clog.SetEnableWarning(false)
clog.SetEnableInfo(true)
clog.SetEnableSuccess(true)
clog.SetEnableDebug(false)

// Все сразу
clog.SetEnableAll()   // включить все
clog.SetDisableAll()  // выключить все
```

## Запись в файл

```go
// Запись в файл (дублируется в консоль)
clog.SetLogFile("app.log")
clog.Error("запишется в файл и консоль")

// Закрыть файл (только консоль)
clog.SetLogFile("")
```

## Ротация логов

При достижении максимального размера файла, начало файла обрезается построчно:

```go
// Максимальный размер файла (0 - выключено)
clog.SetMaxLogFileSize(10 * 1024 * 1024)  // 10 MB

// Процент от начала файла для удаления (по умолчанию 20%)
clog.SetLogFileTrimPercent(30)

// После этого установить файл
clog.SetLogFile("app.log")
```

## Цветной вывод

### Цветовые функции

```go
msg := clog.Red("текст")
msg := clog.Green("текст")
// и др.: HiRed, Yellow, HiYellow, HiGreen, Blue, Cyan, White
```

### Print-обёртки

Для каждого цвета доступны 6 функций:

```go
// Print - вывод без переноса строки
clog.PrintRed("текст")

// Printf - форматированный вывод
clog.PrintfRed("значение: %d", 42)

// Println - вывод с переносом строки
clog.PrintlnRed("текст")

// Sprint - вернуть строку
msg := clog.SprintRed("текст")

// Sprintf - вернуть отформатированную строку
msg := clog.SprintfRed("значение: %d", 42)

// Sprintln - вернуть строку с переносом
msg := clog.SprintlnRed("текст")
```

Цвета: `Red`, `HiRed`, `Yellow`, `HiYellow`, `Green`, `HiGreen`, `Blue`, `Cyan`, `White`

## Progress (прогресс-бар)

Wrapper'ы с методами для вывода на одну строку (с `\r`):

```go
// Доступные wrapper'ы
clog.ErrorWrapper
clog.WarningWrapper
clog.InfoWrapper
clog.SuccessWrapper
clog.DebugWrapper

// Методы
clog.DebugWrapper.Progress("Step ", i)
clog.WarningWrapper.Progressf("Progress: %d%%", percent)
clog.InfoWrapper.Progressln("done")

// Обычные методы (без \r)
clog.DebugWrapper.Print("text")
clog.DebugWrapper.Printf("format: %d", 42)
clog.DebugWrapper.Println("text")

// Пример прогресс-бара
for i := 0; i <= 100; i += 10 {
    clog.InfoWrapper.Progressf("Загрузка: %d%%", i)
    time.Sleep(100 * time.Millisecond)
}
fmt.Println() // перенос строки после завершения
```

## Потокобезопасность

Пакет безопасен для использования в горутинах — использует mutex для записи.

## Тесты

```bash
go test ./...
```

## Примеры

```bash
go run ./examples/
```

## API

### Функции логирования

- `Err(err error)` — логировать ошибку (если не nil)
- `Error(s ...any)`, `Warning(s ...any)`, `Info(s ...any)`, `Success(s ...any)`, `Debug(s ...any)`

### Сеттеры

- `SetEnableError(bool)`, `SetEnableWarning(bool)`, `SetEnableInfo(bool)`, `SetEnableSuccess(bool)`, `SetEnableDebug(bool)`
- `SetEnableAll()`, `SetDisableAll()`
- `SetLogFile(filename string) error`
- `SetMaxLogFileSize(size int64)`, `SetLogFileTrimPercent(percent int)`

### Флаги (экспортированные переменные)

- `EnableError`, `EnableWarning`, `EnableInfo`, `EnableSuccess`, `EnableDebug`

### Цветовые функции

- `Red`, `HiRed`, `Yellow`, `HiYellow`, `Green`, `HiGreen`, `Blue`, `Cyan`, `White`

### Progress wrappers

- `ErrorWrapper`, `WarningWrapper`, `InfoWrapper`, `SuccessWrapper`, `DebugWrapper`

Методы: `Print()`, `Printf()`, `Println()`, `Progress()`, `Progressf()`, `Progressln()`
