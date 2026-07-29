# Stage ① — Define (`overview.md`)

**Prerequisites:** None — first stage of a new picture book.

**Gate artifact:** `{project-root}/overview.md`

**Next stage:** `02-design.md` (after user approves `overview.md`)

---

## Publication Model (unit of generation)

- **Episode** is the generation unit: each `episodes/{nnn}-{episode-slug}.md` is evaluated and then generates page images into `images/{nnn}-{episode-slug}/`.
- **Page** design lives inside the episode file (story + rendering text + illustration guide).

---

## Procedure

1. Use Socratic dialogue to lock the book foundation (one question at a time).
2. After the user answers, write `overview.md` (do not leave it as chat-only).
3. Include **Validation Criteria** as concrete, testable checks used in Stage ③ Evaluate.
4. Do not design episodes or images in this stage.

### What to ask (one question at a time)

| Question | Purpose |
|---|---|
| “한 줄 로글라인(스토리의 한 문장 핵심)은?” | Core concept |
| “타겟 연령(예: 3–5세)과 읽기/그림 난이도 감각?” | Age range + constraints |
| “에피소드는 몇 개로 나눌까?” | Episode List scale |
| “에피소드당 페이지 수는 대략 어느 정도로 할까? (예: 10/20/30/40 내외)” | Episode size (approx only) — actual episode design may change |
| “테마/장르와 톤(따뜻함/몽환/코믹 등)은?” | Craft direction |
| “일러스트 스타일(수채/디지털페인팅/콜라주 등)과 참고작이 있나?” | Visual style constraints |
| “금기사항/범위 밖은?” | Safety |
| “이 그림책이 성공했다고 느끼는 기준(Validation Criteria)은?” | Evaluate criteria |

---

## `overview.md` Template

```markdown
# {Title}

## 기본 정보
- 제목: {title}
- 타겟 연령: {age range}
- 에피소드 수: {episode count}
- 에피소드당 페이지 수(대략): {pages per episode} (예: 10/20/30/40 내외) — 실제 episode 설계에서 조정 가능
- 테마/장르: {theme}
- 톤/분위기: {tone}

## 스토리 요약
{story summary — beginning, middle, end}

## 캐릭터 개요
- 주인공: {name} — {one-line description}
- 조연: {name} — {one-line description}

## 배경 개요
- 세계: {world name} — {one-line description}
- 주요 지역: {region} — {one-line description}
- 주요 장소: {location} — {one-line description}

## 일러스트 스타일
- 미디어: {watercolor / digital painting / collage / ...}
- 팔레트: {color palette description}
- 라인: {outline style}
- 분위기: {mood}
- 참고작: {reference artists or styles if any}

## 제약 조건
- 금기사항: {taboos}
- 범위 밖: {deliberately excluded}

## Validation Criteria
{Checks used in Stage ③ Evaluate — concrete, testable success conditions.}

| Criterion | How to verify |
|-----------|---------------|
| {e.g. Target age can follow the plot from pictures + short text} | {e.g. Each page readable aloud in ≤20 seconds} |
| {e.g. Page-turn curiosity through climax} | {e.g. Pages 1–N-1 have a turn hook} |
| {e.g. Theme earned by scene, not closing sermon} | {e.g. Final page shows outcome; no moral monologue} |
| {e.g. Safe / taboo-compliant} | {e.g. No content from 금기사항} |
```

---

## Completeness Check

- [ ] Title/age/scale fields are filled in episode/page terms
- [ ] Validation Criteria are concrete enough to be checked in Stage ③
- [ ] Story summary matches the user’s intent
- [ ] Illustration style and taboos are specific

---

## Gate

User approves `overview.md`.

**Do not proceed to Stage ② without approval.**

