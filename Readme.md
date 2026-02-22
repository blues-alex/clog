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

### Экспортированные переменные-флаги

- `EnableError`
- `EnableWarning`
- `EnableInfo`
- `EnableSuccess`
- `EnableDebug`

### Цветовые функции

- `Red`, `HiRed`, `Yellow`, `HiYellow`, `Green`, `HiGreen`, `Blue`, `Cyan`, `White`
