---
name: ontology
description: Typed vocabulary + constraint system for representing knowledge as a verifiable graph. Entities, relations, type constraints, schema validation. Use when remembering, linking, or querying structured information.
---

# Ontology

A knowledge graph with typed entities, relations, and schema constraints.

## Core Concept

Everything is an entity with a type, properties, and relations. Every mutation validates against type constraints before committing.

```
Entity: { id, type, properties, relations, created, updated }
Relation: { from_id, relation_type, to_id, properties }
```

## Triggers

| User Says | Action |
|---|---|
| "Remember that..." | Create/update entity |
| "What do I know about X?" | Query graph |
| "Link X to Y" | Create relation |
| "Show all tasks for project Z" | Graph traversal |
| "What depends on X?" | Dependency query |

## Core Types

```yaml
Person: { name, email?, phone?, notes? }
Organization: { name, type?, members[] }
Project: { name, status, goals[], owner? }
Task: { title, status, due?, priority?, assignee?, blockers[] }
Event: { title, start, end?, location?, attendees[] }
Document: { title, path?, url?, summary? }
Note: { content, tags[], refs[] }
```

## Storage

Default: `memory/ontology/graph.jsonl` (append-only)

```jsonl
{"op":"create","entity":{"id":"p_001","type":"Person","properties":{"name":"Alice"}}}
{"op":"relate","from":"proj_001","rel":"has_owner","to":"p_001"}
```

## Commands

```bash
# Create
python3 scripts/ontology.py create --type Person --props '{"name":"Alice"}'

# Query
python3 scripts/ontology.py query --type Task --where '{"status":"open"}'
python3 scripts/ontology.py related --id proj_001 --rel has_task

# Link
python3 scripts/ontology.py relate --from proj_001 --rel has_task --to task_001

# Validate
python3 scripts/ontology.py validate
```

## Schema Constraints

Define in `memory/ontology/schema.yaml`:

```yaml
types:
  Task:
    required: [title, status]
    status_enum: [open, in_progress, blocked, done]
relations:
  has_owner:
    from_types: [Project, Task]
    to_types: [Person]
    cardinality: many_to_one
  blocks:
    from_types: [Task]
    to_types: [Task]
    acyclic: true
```

## Cross-Skill Integration

```python
# Email skill creates commitment → Task skill picks it up
commitment = ontology.create("Commitment", {"description": "Send report by Friday"})
tasks = ontology.query("Commitment", {"status": "pending"})
```
