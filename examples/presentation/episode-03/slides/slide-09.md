# Slide 09 — 설치 확인 — Cocrates = Plugin + Skill

**Section:** 본론 — 설치 + 검증
**Script reference:** script-09.md

---

## Layout

```
┌─────────────────────────────────────────────────┐
│  설치 확인 — Cocrates = Plugin + Skill            │
│                                                   │
│  ┌────────────────────┐ ┌────────────────────┐    │
│  │ 🔌 Plugin          │ │ 📁 Skill           │    │
│  │                    │ │                    │    │
│  │ opencode.jsonc     │ │ ~/.config/opencode │    │
│  │                    │ │   /skills/          │    │
│  │ 확인:              │ │                    │    │
│  │ "cocrates-harness" │ │ adr-writing.md     │    │
│  │ 플러그인 포함?      │ │ spec-writing.md    │    │
│  │                    │ │ education.md 등     │    │
│  └────────────────────┘ └────────────────────┘    │
│                                                   │
│  ✅ 둘 다 정상 → opencode 재시작 → Cocrates 선택      │
│  과정: U(이해) → A(분석) → E(평가) → A(승인)        │
└─────────────────────────────────────────────────┘
```

## Visual Elements

| Element | Description | Source |
|---------|-------------|--------|
| 배경 | 다크 템플릿 | |
| 제목 | "설치 확인 — Cocrates = Plugin + Skill" — 28pt | |
| 2열 레이아웃 | 좌: Plugin (설정 아이콘 🔌), 우: Skill (폴더 아이콘 📁) | |
| 체크리스트 | 각 열 안에 확인할 항목 리스트 | |
| 하단 강조 | ✅ 확인 완료 시 opencode 재시작 → Cocrates 선택 | |
| U→A→E→A | 하단에 작게 표시 (원칙 리마인더) | |

## Audio / Narration

- **Tone:** 분석적, 명확하게
- **Key phrases:**
  - "Cocrates Harness는 Plugin과 Skill, 두 가지로 구성됩니다"
  - "Plugin 확인: opencode.jsonc 안에 cocrates-harness가 있는가?"
  - "Skill 확인: skills/ 디렉토리에 스킬 파일들이 있는가?"
  - "둘 다 정상이면, opencode를 재시작하고 Cocrates Agent를 선택하세요"

## Transitions

- **Enter:** 좌측(Plugin) → 우측(Skill) 순차적으로 나타남
- **Exit:** → Slide 10

## Notes

- 검토 항목이 단 2개뿐이라는 점 강조 (Plugin + Skill만 보면 됨).
- "확인하는 법을 몰라서 불안한 상황"을 방지.
- U→A→E→A를 슬라이드 하단에 작게 리마인더로 표시.
