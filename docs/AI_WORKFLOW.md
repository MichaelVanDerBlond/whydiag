# WhyDiag Autonomous AI Workflow

## Roles

Система использует две независимые AI-роли.

### Architect

Architect:

- читает документацию;
- анализирует проект;
- выбирает следующую задачу;
- пишет TASK;
- проверяет результат Developer;
- принимает или отклоняет изменение.

Architect НЕ редактирует production source code.

---

### Developer

Developer:

- получает одну TASK;
- изучает существующий код;
- реализует только поставленную задачу;
- пишет/обновляет тесты;
- запускает обязательные проверки;
- создаёт RESULT;
- создаёт git commit.

Developer не определяет roadmap.

---

## State machine

Допустимые состояния:

NEED_TASK
DEV_WORK
NEED_REVIEW
DEV_REWORK
PAUSED
FAILED

### NEED_TASK

Architect создаёт ровно одну новую TASK.

После этого:

DEV_WORK

### DEV_WORK

Developer реализует TASK.

После успешной реализации:

NEED_REVIEW

### NEED_REVIEW

Architect выполняет review.

Если принято:

NEED_TASK

Если отклонено:

DEV_REWORK

### DEV_REWORK

Developer исправляет только замечания текущего review.

После этого:

NEED_REVIEW

---

## Task format

Каждая TASK должна содержать:

- ID
- Title
- Motivation
- Current behavior
- Expected behavior
- Scope
- Allowed changes
- Forbidden changes
- Definition of Done
- Tests
- Documentation impact

---

## Task limits

Одна задача должна представлять одно логическое изменение.

По умолчанию Architect должен избегать задач:

- более 5-7 исходных файлов;
- более 500 строк изменения;
- затрагивающих несколько независимых подсистем.

Исключения допустимы только с техническим обоснованием.

---

## Review rules

Architect обязан проверить:

1. соответствие TASK;
2. git diff;
3. отсутствие несвязанного рефакторинга;
4. качество реализации;
5. tests;
6. documentation;
7. архитектурную совместимость;
8. отсутствие выдуманных API/команд/данных.

Результат:

ACCEPTED

или:

REJECTED

---

## Autonomous safety

AI запрещено:

- merge в main;
- force push;
- rewrite git history;
- delete repository;
- изменять SSH credentials;
- менять system configuration;
- устанавливать системные пакеты;
- выполнять destructive commands.

Ветка автономной разработки:

ai-autonomous

main контролируется человеком.

---

## Loop limits

Одна TASK допускает максимум 3 последовательных REJECTED review.

После третьего REJECTED:

FAILED

Автоматический цикл останавливается для этой задачи.

Это предотвращает бесконечные циклы между Architect и Developer.
