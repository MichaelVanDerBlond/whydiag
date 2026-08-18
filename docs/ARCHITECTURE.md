# WhyDiag Architecture

## Dependency direction

cmd
 |
 v
app
 |
 v
core
 |
 +----------------+
 |                |
 v                v
checks         inventory

Основная бизнес-логика не должна зависеть от CLI.

Inventory собирает данные.

Checks анализируют данные.

Core управляет выполнением.

App собирает приложение.

cmd содержит entry point.

---

## Package responsibilities

### cmd/whydiag

Только запуск приложения и обработка CLI entry point.

### internal/app

Композиция приложения.

Регистрация collectors/checks.

### internal/core

Базовые интерфейсы и структуры:

- Check
- Descriptor
- Result
- Runner
- Context
- execution summary

Core не должен содержать специфическую логику Nextcloud, Linux, Apache и т.д.

### internal/inventory

Сбор фактического состояния ОС.

Inventory по возможности должен возвращать данные, а не готовый пользовательский диагноз.

### internal/checks

Диагностическая логика.

Предпочтительная структура:

internal/checks/linux
internal/checks/network
internal/checks/systemd
internal/checks/storage
internal/checks/apache
internal/checks/redis
internal/checks/database
internal/checks/nextcloud
internal/checks/talk

Новые каталоги создаются только когда существует реальная необходимость.

---

## Result model

Диагностический Result должен со временем иметь возможность представить:

- check ID
- status
- summary
- details
- evidence
- cause
- recommendations

Расширение Result должно выполняться обратно совместимо, если это разумно.

---

## Diagnostics are read-only

Collector и Check не должны изменять состояние диагностируемой системы.

Команды используются только для чтения.

---

## External command execution

В будущем рекомендуется единый abstraction layer для запуска внешних команд.

Это необходимо для:

- тестирования;
- timeout;
- capture stdout/stderr;
- exit code;
- безопасной обработки отсутствующих команд.

Не следует создавать собственный os/exec вызов в каждом диагностическом модуле, если уже существует общий механизм.

---

## Testability

Предпочтение отдаётся архитектуре, где парсинг отделён от выполнения команды.

Пример:

command output
    |
    v
parser
    |
    v
structured model
    |
    v
diagnostic check

Это позволяет тестировать большинство логики без реального systemd, Nextcloud или Docker.

---

## Correlation

Checks не должны напрямую зависеть друг от друга.

Корреляция должна выполняться отдельным механизмом над готовыми Result.

Это позволит избежать цепочки жёстких зависимостей между диагностическими модулями.
