---
name: executing-plans
description: "Пошаговое выполнение планов с review checkpoints. Используй когда есть план для выполнения в отдельной сессии. Триггеры: 'выполни план', 'execute plan', 'реализуй по плану'."
---

# Executing Plans

## Обзор

Загрузи план, критически review, выполняй батчами, отчитывайся между батчами.

**Принцип:** Batch execution с checkpoints для архитектурного review.

## Процесс

### Step 1: Загрузи и review план

1. Прочитай файл плана
2. Критически оцени — есть вопросы или concerns?
3. Если concerns: Подними их ПЕРЕД началом
4. Если нет: Создай TodoWrite и продолжай

### Step 2: Выполни batch

**Default: первые 3 задачи**

Для каждой задачи:

1. Отметь in_progress
2. Следуй каждому шагу точно
3. Выполни verifications
4. Отметь completed

### Step 3: Отчёт

Когда batch завершён:

- Покажи что реализовано
- Покажи verification output
- Скажи: "Готов к feedback."

### Step 4: Продолжение

На основе feedback:

- Применить изменения если нужно
- Выполнить следующий batch
- Повторять до завершения

### Step 5: Завершение

После всех задач:

- Verify все тесты
- Предложи options (merge/PR/keep/discard)
- Выполни выбор

## Когда остановиться

**СТОП немедленно когда:**

- Blocker mid-batch (missing dependency, test fails, unclear instruction)
- Критические gaps в плане
- Не понимаешь instruction
- Verification падает повторно

**Спроси разъяснение вместо угадывания.**

## Интеграция с OpenClaw

**Spawn для выполнения плана:**

```python
sessions_spawn(
    task="""
    Выполни план из docs/plans/2026-02-06-feature.md

    Процесс:
    1. Прочитай план
    2. Выполняй по 3 задачи
    3. После каждого batch — report
    4. После всех — verify и report
    """,
    label="execute-feature-plan",
    runTimeoutSeconds=3600
)
```

## Связанные скиллы

- `writing-plans` ← создаёт план
- `subagent-driven-development` → альтернатива в той же сессии
- `test-driven-development` → используется в каждой задаче
