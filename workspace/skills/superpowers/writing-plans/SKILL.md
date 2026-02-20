---
name: writing-plans
description: "Создание детальных планов реализации. Используй когда есть спека или требования для multi-step задачи, ПЕРЕД кодом. Триггеры: 'напиши план', 'составь план', 'implementation plan', 'разбей на задачи'."
---

# Writing Plans: Детальные планы реализации

## Обзор

Пиши планы так, будто исполнитель — джуниор без контекста проекта. Документируй всё: какие файлы трогать, код, как тестировать. Разбивай на bite-sized задачи. DRY. YAGNI. TDD. Частые коммиты.

**Сохраняй планы в:** `docs/plans/YYYY-MM-DD-<feature-name>.md`

## Гранулярность задач

**Каждый шаг — одно действие (2-5 минут):**

- "Напиши failing test" — шаг
- "Запусти, убедись что падает" — шаг
- "Напиши минимальный код для прохождения" — шаг
- "Запусти тесты" — шаг
- "Закоммить" — шаг

## Шаблон заголовка плана

```markdown
# [Feature Name] Implementation Plan

> **Для агента:** Используй skill `executing-plans` для пошагового выполнения.

**Цель:** [Одно предложение]

**Архитектура:** [2-3 предложения о подходе]

**Tech Stack:** [Ключевые технологии]

---
```

## Структура задачи

```markdown
### Task N: [Component Name]

**Файлы:**

- Create: `exact/path/to/file.py`
- Modify: `exact/path/to/existing.py:123-145`
- Test: `tests/exact/path/to/test.py`

**Step 1: Напиши failing test**
\`\`\`python
def test_specific_behavior():
result = function(input)
assert result == expected
\`\`\`

**Step 2: Запусти тест (должен упасть)**
Run: `pytest tests/path/test.py::test_name -v`
Expected: FAIL

**Step 3: Напиши минимальную реализацию**
\`\`\`python
def function(input):
return expected
\`\`\`

**Step 4: Запусти тест (должен пройти)**
Run: `pytest tests/path/test.py::test_name -v`
Expected: PASS

**Step 5: Commit**
\`\`\`bash
git add . && git commit -m "feat: add specific feature"
\`\`\`
```

## Handoff к выполнению

После сохранения плана предложи выбор:

**"План готов. Два варианта выполнения:**

1. **Subagent-Driven** (эта сессия) — субагент на каждую задачу, review между
2. **Parallel Session** — отдельная сессия с `executing-plans`

**Какой подход?"**

## Интеграция с OpenClaw

**Spawn субагента для выполнения:**

```
sessions_spawn(task="Выполни план из docs/plans/<file>.md", label="execute-<feature>")
```

**Связанные скиллы:**

- `brainstorming` ← идея перед планом
- `subagent-driven-development` → выполнение
- `executing-plans` → пошаговое выполнение
- `test-driven-development` → TDD в каждой задаче
