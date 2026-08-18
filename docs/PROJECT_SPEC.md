# WhyDiag Project Specification

## 1. Назначение

WhyDiag — диагностическая система для Linux-серверов и прикладной инфраструктуры.

Основная задача проекта — не просто собирать технические параметры системы, а определять состояние компонентов, выявлять неисправности, связывать симптомы между подсистемами и формировать понятное инженерное заключение.

WhyDiag должен помогать отвечать на вопрос:

"Почему система или сервис работает неправильно?"

а не только:

"Какое сейчас состояние системы?"

---

## 2. Основные принципы

WhyDiag должен:

1. Получать данные из реальной системы.
2. Выполнять независимые диагностические проверки.
3. Не аварийно завершаться из-за ошибки одной проверки.
4. Различать нормальное состояние, предупреждение и ошибку.
5. Сохранять доказательства диагностического вывода.
6. По возможности указывать вероятную причину проблемы.
7. Коррелировать результаты разных проверок.
8. Формировать человекочитаемый отчёт.
9. Поддерживать машинно-читаемый формат результатов.
10. Оставаться расширяемым по мере добавления новых диагностических модулей.

---

## 3. Уровни состояния

Минимально поддерживаются:

- OK
- WARNING
- ERROR
- UNKNOWN

UNKNOWN используется, если достоверно определить состояние невозможно.

Отсутствие компонента не всегда является ERROR.

Например, отсутствие Docker на сервере без Docker-сервисов является нормальной ситуацией.

---

## 4. Архитектурные слои

### Inventory

Сбор фактической информации о системе.

Примеры:

- hostname
- OS information
- kernel
- uptime
- CPU
- memory
- swap
- filesystem
- timezone
- time synchronization

Inventory не должен самостоятельно принимать сложные диагностические решения.

---

### Checks

Checks анализируют данные Inventory или напрямую безопасно получают данные системы.

Каждая проверка должна возвращать структурированный Result.

Проверка должна быть независимой настолько, насколько это практически возможно.

---

### Core

Core отвечает за:

- регистрацию проверок;
- запуск;
- обработку ошибок;
- агрегирование результатов;
- итоговый статус;
- выполнение диагностического конвейера.

---

### Correlation

В дальнейшем WhyDiag должен поддерживать слой корреляции.

Он связывает несколько результатов.

Пример:

DNS: OK
TCP 443: OK
TLS: OK
Apache proxy: OK
Signaling service: ERROR

Вывод:

Вероятный источник проблемы Nextcloud Talk — signaling service.

---

### Reporting

Отчёты должны поддерживать как минимум:

- CLI text output;
- структурированное внутреннее представление.

В дальнейшем:

- JSON;
- HTML;
- machine-readable reports.

---

## 5. Приоритет развития

Развитие выполняется снизу вверх.

### Stage 1 — Base OS

- hostname
- OS information
- kernel
- uptime
- CPU
- memory
- swap
- filesystem
- timezone
- time synchronization

### Stage 2 — Linux Runtime

- network interfaces
- routes
- DNS
- open/listening ports
- processes
- systemd
- journal
- load average
- resource pressure

### Stage 3 — Storage

- disks
- mounts
- SMART
- filesystem usage
- inode usage
- I/O errors

### Stage 4 — Network Diagnostics

- link state
- addresses
- gateway
- routing
- DNS resolution
- connectivity
- TCP checks
- TLS checks

### Stage 5 — Infrastructure Services

- systemd services
- Docker
- Redis
- MariaDB/MySQL
- Apache
- Nginx
- certificates

### Stage 6 — Nextcloud

- installation detection
- config
- database connectivity
- data directory
- cron/background jobs
- Redis
- trusted domains
- maintenance mode
- app status
- occ health checks

### Stage 7 — Nextcloud Talk

- Talk app
- signaling configuration
- HPB
- websocket connectivity
- TURN
- STUN
- Coturn
- TLS
- DNS
- publisher/subscriber paths

### Stage 8 — Correlation Engine

Корреляция результатов нескольких проверок.

### Stage 9 — Historical / Compare Mode

Сравнение текущего состояния с:

- предыдущим запуском;
- эталонным состоянием;
- reference configuration.

---

## 6. Требования к диагностическим проверкам

Каждая диагностическая проверка должна:

- иметь стабильный ID;
- иметь понятное имя;
- иметь описание;
- возвращать status;
- возвращать summary;
- при необходимости возвращать evidence;
- корректно обрабатывать ошибки;
- иметь тесты, если логика тестируема;
- не изменять конфигурацию сервера;
- не исправлять систему автоматически без отдельного режима.

По умолчанию WhyDiag является READ-ONLY диагностическим инструментом.

---

## 7. Безопасность

Диагностика не должна:

- менять системные настройки;
- перезапускать службы;
- устанавливать пакеты;
- удалять данные;
- менять права файлов;
- изменять firewall;
- изменять Nextcloud config;
- выполнять destructive database operations.

Любая будущая функция remediation должна быть отдельно спроектирована и явно включаться пользователем.

---

## 8. Ошибки

Ошибка одной проверки не должна прерывать весь диагностический запуск.

Проверка должна возвращать ERROR или UNKNOWN с объяснением причины.

Panic внутри проверки должен быть изолирован Core, если это технически целесообразно.

---

## 9. Качество кода

После каждого изменения обязательно выполняются:

go fmt ./...
go vet ./...
go test ./...
go build ./...

Изменения не принимаются, если хотя бы одна обязательная проверка не проходит.

---

## 10. Правило развития проекта

AI Architect не имеет права выдавать задачи только для того, чтобы продолжать работу.

Каждая задача должна иметь доказанную полезность.

Приоритет:

1. исправить сломанное;
2. завершить существующую архитектуру;
3. добавить недостающие базовые проверки;
4. добавить тесты;
5. улучшить диагностику;
6. только затем расширять функциональность.

---

## 11. Конечная цель

WhyDiag должен иметь возможность сформировать инженерное заключение вида:

Component: Nextcloud Talk
Status: ERROR

Symptoms:
- client cannot connect to signaling backend

Evidence:
- DNS resolution: OK
- TLS handshake: OK
- reverse proxy: OK
- signaling service: FAILED
- TCP 8443 local connection: FAILED

Probable cause:
nextcloud-spreed-signaling.service is not running.

Confidence:
HIGH

Suggested investigation:
systemctl status nextcloud-spreed-signaling
journalctl -u nextcloud-spreed-signaling
