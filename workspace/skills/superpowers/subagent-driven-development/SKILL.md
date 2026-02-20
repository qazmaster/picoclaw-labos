---
name: subagent-driven-development
description: "Выполнение планов через субагентов. Используй когда есть plan с независимыми задачами. Триггеры: 'выполни план', 'запусти субагентов', 'subagent', 'делегируй задачи'."
---

# Subagent-Driven Development

Выполняй план, запуская свежий субагент на каждую задачу, с двухэтапным review после каждой: сначала соответствие спеке, потом качество кода.

**Принцип:** Свежий субагент на задачу + два review = высокое качество, быстрая итерация

## Когда использовать

- Есть implementation plan
- Задачи в основном независимы
- Хочешь остаться в этой сессии

## Процесс

```
1. Прочитай план, извлеки все задачи
2. Для каждой задачи:
   a. Запусти implementer субагент
   b. Если вопросы — ответь
   c. Субагент: реализует, тестирует, коммитит, self-review
   d. Запусти spec reviewer — проверка соответствия спеке
   e. Если проблемы — implementer фиксит, re-review
   f. Запусти code quality reviewer
   g. Если проблемы — implementer фиксит, re-review
   h. Отметь задачу выполненной
3. После всех задач — финальный review всей реализации
```

## Промпты для субагентов

### Implementer

```
Выполни Task N из плана:
[полный текст задачи]

Контекст: [описание где эта задача в общей картине]

Действия:
1. Напиши failing test
2. Убедись что падает
3. Напиши минимальный код
4. Убедись что проходит
5. Self-review
6. Commit

Если есть вопросы — спроси ПЕРЕД началом работы.
```

### Spec Reviewer

```
Проверь Task N на соответствие спеке:

Спека: [requirements]
Код: [git diff или файлы]

Проверь:
- Все requirements выполнены?
- Ничего лишнего не добавлено?
- Тесты покрывают requirements?

Ответ: ✅ Spec compliant или ❌ Issues: [список]
```

### Code Quality Reviewer

```
Проверь качество кода для Task N:

Код: [git diff]

Проверь:
- Читаемость
- DRY/YAGNI
- Error handling
- Test quality
- Naming

Ответ: ✅ Approved или ❌ Issues: [severity + description]
```

## Интеграция с OpenClaw

**Запуск субагентов:**

```python
# Implementer
sessions_spawn(
    task="Выполни Task 1: [описание]. Следуй TDD.",
    label="implement-task-1"
)

# Spec Reviewer
sessions_spawn(
    task="Review Task 1 на соответствие спеке: [spec]",
    label="review-spec-1"
)

# Code Quality Reviewer
sessions_spawn(
    task="Review качества кода Task 1",
    label="review-quality-1"
)
```

## Red Flags — НИКОГДА

- ❌ Начинать без плана
- ❌ Пропускать review (spec ИЛИ quality)
- ❌ Запускать несколько implementer параллельно (конфликты)
- ❌ Игнорировать вопросы субагента
- ❌ Принимать "close enough" на spec compliance
- ❌ Code quality review ДО spec compliance
- ❌ Переходить к следующей задаче с открытыми issues

## Связанные скиллы

- `writing-plans` ← создаёт план
- `test-driven-development` → субагенты используют
- `executing-plans` → альтернатива для параллельной сессии
- `dispatching-parallel-agents` → для независимых задач
