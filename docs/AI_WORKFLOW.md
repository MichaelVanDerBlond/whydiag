# WhyDiag Autonomous AI Workflow

## 1. Назначение

Этот документ определяет механику автономной работы двух AI-агентов WhyDiag:

- Architect
- Developer

Содержательные правила общения определены в:

docs/AI_PROTOCOL.md

Архитектурные ограничения проекта определены в:

docs/ARCHITECTURE.md

Цели и этапы развития определены в:

docs/PROJECT_SPEC.md

Orchestrator обязан следовать этим документам.

---

## 2. Главный принцип

AI-агенты не управляют workflow напрямую.

Architect и Developer формируют содержательные сообщения.

Orchestrator:

- сохраняет сообщения;
- определяет следующего говорящего;
- проверяет допустимость перехода;
- управляет Git;
- разрешает или запрещает изменение production files;
- запускает validation;
- контролирует loop limits.

LLM не устанавливает внутренние состояния самостоятельно.

---

## 3. Одна активная TASK

Одновременно существует только одна активная TASK.

До завершения текущей TASK новая задача не создаётся.

TASK считается завершённой только при:

ACCEPTED

ABORTED

или BLOCKED с остановкой автоматического цикла.

---

## 4. Канал диалога

Для каждой TASK создаётся отдельный журнал:

.ai/dialog/TASK-NNNN.jsonl

Каждая строка является отдельным JSON-сообщением.

Пример:

{"seq":1,"from":"architect","type":"TASK","message":"..."}
{"seq":2,"from":"developer","type":"PLAN","message":"..."}
{"seq":3,"from":"architect","type":"PLAN_APPROVED","message":"..."}
{"seq":4,"from":"developer","type":"RESULT","message":"..."}
{"seq":5,"from":"architect","type":"ACCEPTED","message":"..."}

История текущей TASK является частью контекста обоих агентов.

История старых TASK не должна автоматически передаваться целиком.

---

## 5. Нормальный жизненный цикл

Нормальная последовательность:

Architect:
TASK

Developer:
PLAN

Architect:
PLAN_APPROVED

Developer:
IMPLEMENTATION

Orchestrator:
VALIDATION

Developer:
RESULT

Architect:
ACCEPTED

После ACCEPTED:

Architect получает право выбрать следующую TASK.

---

## 6. Ветка QUESTION

Если Developer не понимает TASK:

Developer:
QUESTION

Architect:
ANSWER

После ANSWER Developer повторно анализирует TASK.

Следующее сообщение Developer:

PLAN

QUESTION

или BLOCKED.

Developer не изменяет production files до PLAN_APPROVED.

---

## 7. Ветка PLAN_REWORK

Если PLAN не соответствует TASK:

Architect:
PLAN_REWORK

Developer:
обновлённый PLAN

Architect:
PLAN_APPROVED

или повторный PLAN_REWORK.

Максимум:

3 PLAN_REWORK подряд.

После превышения:

BLOCKED.

---

## 8. Ветка реализации

После PLAN_APPROVED orchestrator разрешает Developer изменение production files.

Developer реализует только утверждённый PLAN.

Если Developer обнаруживает новое архитектурное противоречие:

Developer:
QUESTION

или BLOCKED.

В этом случае незавершённые изменения не публикуются как RESULT.

---

## 9. Validation

После реализации orchestrator запускает:

gofmt
go vet ./...
go test ./...
go build ./...

Validation является техническим gate.

Validation PASS означает только:

проект компилируется и обязательные проверки прошли.

Validation PASS НЕ означает автоматически, что TASK выполнена.

Перед RESULT обязательно должны существовать изменения, соответствующие утверждённому PLAN.

---

## 10. Validation repair

Если validation FAIL:

orchestrator передаёт Developer:

- текущую TASK;
- утверждённый PLAN;
- текущий diff;
- полный validation output;
- актуальное содержимое связанных файлов.

Developer выполняет одну конкретную repair attempt.

Максимум:

3 validation repair.

Если после трёх содержательно разных попыток validation не проходит:

BLOCKED.

Повтор одного и того же исправления без новой гипотезы запрещён.

---

## 11. RESULT

После успешной реализации и validation Developer формирует RESULT.

RESULT должен описывать фактическое состояние repository.

Orchestrator самостоятельно добавляет:

- commit SHA;
- git diff summary;
- validation output;
- изменённые файлы.

Developer не должен выдумывать эти данные.

После RESULT управление переходит Architect.

---

## 12. Review

Architect получает:

- TASK;
- полный диалог текущей TASK;
- утверждённый PLAN;
- RESULT;
- diff;
- validation;
- актуальное содержимое изменённых файлов.

Architect отвечает:

ACCEPTED

REWORK

или QUESTION.

---

## 13. REWORK

После REWORK Developer не формирует новый PLAN автоматически.

Сначала Developer проверяет:

- Evidence;
- Required Change;
- Acceptance;
- фактический код.

Если замечание понятно и находится внутри TASK:

Developer выполняет исправление.

Если замечание требует нового архитектурного решения:

Developer:
QUESTION

Если REWORK фактически расширяет TASK:

Developer:
QUESTION

Architect обязан решить, является ли это:

- уточнением текущей TASK;
- пересмотром TASK;
- новой задачей;
- причиной ABORTED.

Максимум:

3 REWORK для одной TASK.

После превышения:

BLOCKED.

---

## 14. ACCEPTED

После ACCEPTED orchestrator:

1. фиксирует TASK как закрытую;
2. сохраняет финальный RESULT;
3. сохраняет review;
4. очищает transient runtime state;
5. передаёт управление Architect для выбора следующей задачи.

---

## 15. BLOCKED

BLOCKED останавливает автоматический цикл текущей TASK.

Orchestrator не должен автоматически:

- повторять предыдущий запрос;
- создавать новую TASK;
- сбрасывать состояние;
- запускать очередную реализацию.

BLOCKED требует:

- решения Architect;
- либо участия человека;
- либо ABORTED.

---

## 16. ABORTED

ABORTED закрывает текущую TASK без принятой реализации.

Используется, если:

- TASK ошибочна;
- TASK потеряла актуальность;
- задача уже реализована;
- Definition of Done некорректен;
- TASK невозможно безопасно выполнить в текущем scope.

После ABORTED Architect может создать следующую TASK.

---

## 17. Git ownership

Перед каждым содержательным действием orchestrator запоминает:

base_commit

Перед публикацией результата выполняется:

git fetch origin

Если:

origin/ai-autonomous != base_commit

результат генерации считается устаревшим.

Он не публикуется.

Агент получает актуальный repository state и повторно принимает решение.

Это защищает от гонки между Architect и Developer.

---

## 18. Запрет перезаписи TASK

TASK-файл создаётся только в exclusive create mode.

Существующая TASK никогда не перезаписывается автоматически.

Если требования текущей задачи были уточнены, изменения сохраняются в диалоге.

При фундаментальном изменении задачи Architect использует ABORTED и создаёт новую TASK.

---

## 19. Контроль циклов

Orchestrator ведёт счётчики:

plan_rework_count
implementation_repair_count
review_rework_count
question_repeat_count

Также хранится hash нормализованного последнего сообщения каждого типа.

Эквивалентное повторное сообщение без новой информации считается циклом.

Orchestrator не запускает следующее действие автоматически.

Вместо этого агент получает возможность:

QUESTION
BLOCKED
ABORTED

---

## 20. Контекст Architect

При создании TASK Architect получает:

- docs/PROJECT_SPEC.md;
- docs/ARCHITECTURE.md;
- README.md;
- project tree;
- актуальный source summary;
- последние принятые изменения.

Architect не должен получать полный архив всех прошлых диалогов.

При review Architect получает только контекст текущей TASK.

---

## 21. Контекст Developer

Developer получает:

- docs/PROJECT_SPEC.md;
- docs/ARCHITECTURE.md;
- текущую TASK;
- текущий диалог;
- реальные связанные файлы;
- связанные tests.

Developer не должен получать весь repository как набор редактируемых файлов.

Файлы для изменения определяются после PLAN_APPROVED.

---

## 22. Изменение файлов

LLM не получает прямого права произвольно записывать любой путь.

Developer возвращает структурированный набор proposed edits.

Orchestrator проверяет:

- путь;
- scope;
- соответствие PLAN;
- запрещённые каталоги;
- количество файлов.

Запрещены изменения:

.git/
.ai/tasks/
.ai/reviews/
.ai/dialog/

кроме действий orchestrator.

---

## 23. Allowed production paths

Developer может изменять при наличии утверждённого PLAN:

README.md
docs/
cmd/
internal/
go.mod
go.sum
Makefile
.gitignore

Новые пути разрешаются только при соответствии TASK и архитектуре.

---

## 24. Ограничение размера изменения

По умолчанию одна TASK:

- до 5 production files;
- до 500 изменённых строк;
- одна логическая подсистема.

Если PLAN превышает лимит, Architect должен решить:

- разбить TASK;
- либо явно разрешить исключение.

---

## 25. State machine orchestrator

Допустимые внутренние состояния:

WAIT_ARCHITECT_TASK
WAIT_DEVELOPER_RESPONSE
WAIT_ARCHITECT_TASK_RESPONSE
WAIT_ARCHITECT_PLAN_REVIEW
WAIT_DEVELOPER_IMPLEMENTATION
RUN_VALIDATION
WAIT_DEVELOPER_RESULT
WAIT_ARCHITECT_REVIEW
WAIT_DEVELOPER_REWORK
BLOCKED
CLOSED

Эти значения принадлежат orchestrator.

Они не являются содержательными сообщениями LLM.

---

## 26. Mapping сообщений к переходам

WAIT_ARCHITECT_TASK
TASK -> WAIT_DEVELOPER_RESPONSE

WAIT_DEVELOPER_RESPONSE
QUESTION -> WAIT_ARCHITECT_TASK_RESPONSE
PLAN -> WAIT_ARCHITECT_PLAN_REVIEW
BLOCKED -> BLOCKED

WAIT_ARCHITECT_TASK_RESPONSE
ANSWER -> WAIT_DEVELOPER_RESPONSE
ABORTED -> CLOSED
BLOCKED -> BLOCKED

WAIT_ARCHITECT_PLAN_REVIEW
PLAN_APPROVED -> WAIT_DEVELOPER_IMPLEMENTATION
PLAN_REWORK -> WAIT_DEVELOPER_RESPONSE
QUESTION -> WAIT_DEVELOPER_RESPONSE
ABORTED -> CLOSED

WAIT_DEVELOPER_IMPLEMENTATION
IMPLEMENTATION -> RUN_VALIDATION
QUESTION -> WAIT_ARCHITECT_TASK_RESPONSE
BLOCKED -> BLOCKED

RUN_VALIDATION
PASS -> WAIT_DEVELOPER_RESULT
FAIL -> WAIT_DEVELOPER_IMPLEMENTATION
REPAIR_LIMIT -> BLOCKED

WAIT_DEVELOPER_RESULT
RESULT -> WAIT_ARCHITECT_REVIEW

WAIT_ARCHITECT_REVIEW
ACCEPTED -> CLOSED
REWORK -> WAIT_DEVELOPER_REWORK
QUESTION -> WAIT_DEVELOPER_RESPONSE
ABORTED -> CLOSED

WAIT_DEVELOPER_REWORK
RESULT -> WAIT_ARCHITECT_REVIEW
QUESTION -> WAIT_ARCHITECT_TASK_RESPONSE
BLOCKED -> BLOCKED

---

## 27. Восстановление после restart

После restart orchestrator не угадывает текущее действие.

Он восстанавливает состояние из:

.ai/runtime/session.json

и сверяет:

- active_task;
- last_seq;
- state;
- base_commit;
- counters;
- dialog log.

Если runtime state противоречит Git:

BLOCKED

до безопасной синхронизации.

---

## 28. Старые state-файлы

Старые значения:

NEED_TASK
DEV_WORK
NEED_REVIEW
DEV_REWORK
FAILED

считаются legacy protocol.

Новый orchestrator не должен использовать их как канал общения Architect и Developer.

После миграции они могут быть удалены или оставлены только для диагностики совместимости.

---

## 29. Главный критерий автономности

Автономность не означает отсутствие остановок.

Корректный автономный агент обязан остановиться, если дальнейшее действие требует догадки.

QUESTION и BLOCKED являются нормальной частью автономной разработки.

Главный критерий:

каждое автоматическое действие должно увеличивать количество достоверной информации или приближать TASK к Definition of Done.
