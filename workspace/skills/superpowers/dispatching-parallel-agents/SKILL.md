---
name: dispatching-parallel-agents
description: "Параллельный запуск субагентов для независимых задач. Используй когда 2+ задачи без shared state. Триггеры: 'параллельно', 'одновременно', 'parallel agents', 'несколько задач сразу'."
---

# Dispatching Parallel Agents

## Обзор

Когда есть несколько независимых проблем (разные тестовые файлы, разные подсистемы), последовательное расследование — трата времени. Каждое расследование независимо и может идти параллельно.

**Принцип:** Один агент на независимую проблему. Работают параллельно.

## Когда использовать

**Используй когда:**

- 3+ файлов с тестами падают по разным причинам
- Несколько подсистем сломаны независимо
- Каждая проблема понятна без контекста других
- Нет shared state между расследованиями

**НЕ используй когда:**

- Проблемы связаны (фикс одной может пофиксить другие)
- Нужно понимание всей системы
- Агенты будут мешать друг другу

## Паттерн

### 1. Идентифицируй независимые домены

Группируй по тому что сломано:

- File A: Tool approval flow
- File B: Batch completion
- File C: Abort functionality

### 2. Создай focused задачи

Каждый агент получает:

- **Scope:** Один файл или подсистема
- **Goal:** Сделай эти тесты зелёными
- **Constraints:** Не меняй другой код
- **Output:** Саммари что нашёл и пофиксил

### 3. Запусти параллельно

```python
# OpenClaw
sessions_spawn(task="Fix agent-tool-abort.test.ts failures", label="fix-abort")
sessions_spawn(task="Fix batch-completion.test.ts failures", label="fix-batch")
sessions_spawn(task="Fix tool-approval.test.ts failures", label="fix-approval")
```

### 4. Review и интеграция

Когда агенты вернутся:

- Прочитай каждый summary
- Verify что фиксы не конфликтуют
- Запусти полный test suite
- Интегрируй все изменения

## Структура промпта для агента

```markdown
Пофикси 3 падающих теста в src/agents/agent-tool-abort.test.ts:

1. "should abort tool with partial output" - expects 'interrupted at'
2. "should handle mixed completed and aborted" - fast tool aborted
3. "should track pendingToolCount" - expects 3 results, gets 0

Это timing/race condition issues. Задача:

1. Прочитай тест файл
2. Найди root cause
3. Пофикси (event-based waiting, не увеличивай timeouts)

НЕ просто увеличивай timeouts.

Return: Summary что нашёл и что пофиксил.
```

## Типичные ошибки

**❌ Слишком широко:** "Пофикси все тесты"
**✅ Специфично:** "Пофикси agent-tool-abort.test.ts"

**❌ Нет контекста:** "Пофикси race condition"
**✅ Контекст:** Вставь error messages и имена тестов

**❌ Нет constraints:** Агент может рефакторить всё
**✅ Constraints:** "НЕ меняй production code" или "Fix tests only"

## Интеграция с OpenClaw

```python
# Запуск трёх агентов параллельно
sessions_spawn(
    task="Fix abort tests in file X. Return summary.",
    label="fix-abort",
    runTimeoutSeconds=600
)
sessions_spawn(
    task="Fix batch tests in file Y. Return summary.",
    label="fix-batch",
    runTimeoutSeconds=600
)
sessions_spawn(
    task="Fix approval tests in file Z. Return summary.",
    label="fix-approval",
    runTimeoutSeconds=600
)

# Проверить статус
sessions_list(kinds=["isolated"])

# Получить результаты
sessions_history(sessionKey="fix-abort")
```

## Связанные скиллы

- `systematic-debugging` → каждый агент использует
- `subagent-driven-development` → для sequential задач
- `test-driven-development` → агенты пишут тесты
