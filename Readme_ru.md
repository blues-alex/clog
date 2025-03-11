## Обзор

### Основные функции

- `SetEnableError`, `SetEnableWarning`, `SetEnableInfo`, `SetEnableSuccess`, `SetEnableDebug`: Управляйте уровнем логирования, включая или отключая определенные типы сообщений.
- `SetEnableAll`, `SetDisableAll`: Включите все или отключите все уровни логирования одновременно.
- `SetLogFile`: Укажите файл для хранения логов.
- `Err`, `Error`, `Warning`, `Info`, `Success`, `Debug`: Основные функции для записи сообщений в лог, каждая предназначена для разных типов и цветов.

### Структура пакета

Проект организован вокруг одного пакета: `clog`.

### Используемые библиотеки и фреймворки

- `fmt`: Для форматирования строк.
- `os`: Для управления файловыми операциями.
- `time`: Для захвата текущего временного метки.
- `github.com/fatih/color`: Предоставляет цветной вывод для консольных логов.

## Документация кода

### clog.go

#### Переменные

```go
var (
    // Функции для окрашивания строк в разных цветах
    Blue, Cyan, Green, HiGreen, Yellow, HiYellow, Red, HiRed, White func(string) string

    // Логические флаги для контроля уровней логирования
    EnableError   bool = true
    EnableWarning bool = true
    EnableInfo    bool = false
    EnableSuccess bool = false
    EnableDebug   bool = false

    logFile *os.File // Указатель на файл лога
)
```

- `Blue`, `Cyan`, и т.д.: Функции для окрашивания выводов в консоль.
- `EnableError`, `EnableWarning`, и т.д.: Флаги, указывающие, активны ли каждый тип логирования.
- `logFile`: Ссылка на файл, куда записываются логи.

#### Установщики для уровней логирования

```go
func SetEnableError(enable bool) { EnableError = enable }
func SetEnableWarning(enable bool) { EnableWarning = enable }
func SetEnableInfo(enable bool) { EnableInfo = enable }
func SetEnableSuccess(enable bool) { EnableSuccess = enable }
func SetEnableDebug(enable bool) { EnableDebug = enable }
```

- Эти функции позволяют устанавливать активный статус для каждого уровня логирования.

#### Функция для установки файла лога

```go
func SetLogFile(filename string) error {
    if filename == "" {
        // Закрыть существующий файл лога и оставить только stdout, если не указан новый файл
        if logFile != nil {
            logFile.Close()
            logFile = nil
        }
        return nil
    }
    var err error
    logFile, err = os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        // Ошибка, если открытие файла не удалось
        return err
    }
    return nil // Успешно настроили файл лога
}
```

- Эта функция устанавливает новый файл лога или сбрасывает логирование в стандартный вывод, если имя файла не указано.

#### Вспомогательная функция для форматирования сообщений

```go
func formatMessage(s ...any) string {
    // Удаление скобок из начала и конца сообщения
    return fmt.Sprintf("%v", s...)[1 : len(fmt.Sprint(s...))-1]
}
```

- Эта вспомогательная функция очищает сообщения, удаляя любые окружающие квадратные скобки, делая их более читаемыми.

#### Вспомогательная функция для логирования

```go
func loggerFunc(consoleLine, fileLine string) {
    // Вывод сообщения в консоль и, при необходимости, в файл лога, если он установлен
    fmt.Println(consoleLine)
    if logFile != nil {
        logFile.WriteString(fileLine + "\n")
    }
}
```

- Эта функция регистрирует сообщения как в консоль, так и, если применимо, в указанный файл лога.

#### Основные функции для логирования

Каждая из этих функций обрабатывает логирование для разных уровней с соответствующими цветами:

```go
func Err(e error) {
    if e != nil && EnableError {
        // Если произошла ошибка и уровень Error включен, вызовите функцию Error
        // Для упрощения кода (если не требуется обработка ошибки)
        Error(e)
    }
}

func Error(s ...any) {
    if !EnableError { return } // Пропустить, если логирование ошибок отключено
    mess := formatMessage(s...)
    ts := time.Now().Format("02.01.06 15:04:05.000")
    // Формирует строки для консоли с цветным текстом и для файла
    consoleLine := fmt.Sprintf("[%s] %s %s\n", Green(ts), Red("ERROR:"), HiRed(mess))
    fileLine := fmt.Sprintf("[%s] ERROR: %s\n", ts, mess)
    loggerFunc(consoleLine, fileLine)
}

// Аналогично для других уровней логирования (Warning, Info, Success, Debug)
```

- Эти функции обеспечивают правильную форматировку и окрашивание сообщений в соответствии с их уровнем.

## Примеры использования

### Пример использования основных функций логирования

Этот пример демонстрирует использование пакета clog путем включения всех уровней логирования и направления их в файл с именем "app.log". Затем показаны различные типы сообщений, записываемых в лог.

```go
package main

import (
    "github.com/blues-alex/clog"
)

func main() {
    // Включите все уровни логирования
    clog.SetEnableAll()
    // Установите файл лога на "app.log"
    err := clog.SetLogFile("app.log")
    if err != nil {
        panic(err) // Обработайте любые ошибки, возникающие при настройке логирования
    }

    // Запишите различные типы сообщений в лог
    clog.Debug("Это отладочное сообщение.")
    clog.Info("Это информационное сообщение.")
    clog.Warning("Это предупреждение.")
    clog.Error("Это ошибка.")
    clog.Success("Операция выполнена успешно.")
}
```

- Этот пример показывает, как настроить пакет clog для полного логирования, включая установку файла лога и использование различных функций регистратора.

### Пример использования установщиков

Вот еще один способ использования пакета clog путем выборочного включения определенных уровней логирования и направления их в файл "error.log".

```go
package main

import (
    "clog"
)

func main() {
    // Включите только уровень логирования ошибок
    clog.SetEnableError(true)
    clog.SetEnableWarning(false) // Отключите уровень логирования предупреждений
    err := clog.SetLogFile("error.log")
    if err != nil {
        panic(err)
    }

    // Запишите только сообщения об ошибках в указанный файл
    clog.Error("Это сообщение будет записано в 'error.log'.")
    clog.Warning("Это предупреждение не появится в 'error.log'.") // Не будет записано из-за отключенного настройки
}
```

- Этот пример иллюстрирует, как настроить уровни логирования и файл назначения, фокусируясь только на сообщениях об ошибках.
