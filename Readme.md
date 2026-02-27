# clog

Цветной логгер для Go с поддержкой нескольких уровней логирования и записи в файл.

## Установка

```bash
go get github.com/blues-alex/clog
```

## Быстрый старт

```go
package main

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

| Функция | Описание | Цвет по умолчанию |
|---------|----------|-------------------|
| `Error` | Ошибки | Красный |
| `Warning` | Предупреждения | Желтый |
| `Info` | Информация | Белый |
| `Success` | Успешные операции | Зеленый |
| `Debug` | Отладочная информация | Синий |

По умолчанию включены: `Error`, `Warning`

## Управление уровнями

```go
// Включить/выключить отдельный уровень
clog.SetEnableError(true)
clog.SetEnableWarning(false)
clog.SetEnableInfo(true)
clog.SetEnableSuccess(true)
clog.SetEnableDebug(false)

// Включить все уровни
clog.SetEnableAll()

// Выключить все уровни
clog.SetDisableAll()
```

## Запись в файл

```go
// Запись логов в файл (дублируется в консоль)
clog.SetLogFile("app.log")

// Вернуться к выводу только в консоль
clog.SetLogFile("")
```

## Потокобезопасность

Пакет безопасен для использования в горутинах — использует мьютекс для записи.

## Примеры

См. `examples/main.go` для полного демо:

```bash
go run ./examples/
```

## API

### Функции логирования

- `Err(err error)` — логировать ошибку (если не nil)
- `Error(s ...any)` — логировать ошибку
- `Warning(s ...any)` — логировать предупреждение
- `Info(s ...any)` — логировать информацию
- `Success(s ...any)` — логировать успех
- `Debug(s ...any)` — логировать отладку

### Сеттеры

- `SetEnableError(bool)`
- `SetEnableWarning(bool)`
- `SetEnableInfo(bool)`
- `SetEnableSuccess(bool)`
- `SetEnableDebug(bool)`
- `SetEnableAll()`
- `SetDisableAll()`
- `SetLogFile(filename string) error`
- `SetMaxLogFileSize(size int64)`
- `SetLogFileTrimPercent(percent int)`

### Экспортированные переменные-флаги

- `EnableError`
- `EnableWarning`
- `EnableInfo`
- `EnableSuccess`
- `EnableDebug`

### Цветовые функции

- `Red`, `HiRed`, `Yellow`, `HiYellow`, `Green`, `HiGreen`, `Blue`, `Cyan`, `White`

### Цветной вывод (Print wrappers)

Для каждого цвета доступны функции:
- `PrintX(a ...any)`, `PrintlnX(a ...any)`, `PrintfX(format, a ...any)`
- `SprintX(a ...any)`, `SprintlnX(a ...any)`, `SprintfX(format, a ...any)`

Где X: `Red`, `HiRed`, `Yellow`, `HiYellow`, `Green`, `HiGreen`, `Blue`, `Cyan`, `White`

Пример:
```go
clog.PrintRed("Ошибка: ")
clog.PrintfGreen("Значение: %d\n", 42)
clog.PrintlnBlue("Информация")
msg := clog.SprintfHiYellow("Форматированное: %s", "значение")
```

## Ротация логов

При достижении максимального размера файла, начало файла обрезается:

```go
clog.SetMaxLogFileSize(10 * 1024 * 1024) // 10 MB (по умолчанию 0 - без ротации)
clog.SetLogFileTrimPercent(20)            // Удалять 20% от начала (по умолчанию 20%)
clog.SetLogFile("app.log")
```

## Progress (прогресс-бар)

Обёртки с методом `.Progress()` для вывода на одну строку (с `\r`):

```go
clog.DebugWrapper.Progress("Progress: ", 50, "%")
clog.WarningWrapper.Progressf("Step %d of %d", current, total)
```

Доступные wrapper'ы:
- `clog.ErrorWrapper`
- `clog.WarningWrapper`
- `clog.InfoWrapper`
- `clog.SuccessWrapper`
- `clog.DebugWrapper`

Методы: `Print()`, `Printf()`, `Println()`, `Progress()`, `Progressf()`
