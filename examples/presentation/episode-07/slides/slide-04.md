# Slide 04: 🏛️ Cocrates Harness의 두 가지 기둥

## Key Message
Cocrates는 Agent(공통 헌법)와 Skills(독립된 전문가 팀)의 분리 구조다. 이 분리를 통해 각 산출물 유형에 최적화 + 개별 개선 + 확장이 가능하다.

## Content
- **Cocrates Agent** (`cocrates.md`): 공통 원칙과 컨트롤 타워 — 최상위 헌법
  - 정체성, 원칙, 구조, 요청 처리, 핵심 활동, 성공 기준
- **Skills** (`.opencode/skills/*/SKILL.md`): 독립된 전문가 팀 — 산출물별 최적화 워크플로우
  - 교육, 문서 작성, ADR, 검증 등 각 활동의 구체적 절차
- **분리 원칙:** 공통 통제 체계는 유지, 개별 워크플로우는 독립적으로 확장
- **장점 3가지:**
  1. 각 산출물 유형에 최적화된 워크플로우 제공
  2. 개별 스킬의 추가/수정/개선이 독립적으로 가능
  3. 새로운 산출물 유형이 필요하면 스킬만 추가

## Visual Elements
- 두 개의 기둥 다이어그램 (좌: Agent / 우: Skills)
- 가운데 "분리 원칙" 레이블
- 장점 3가지는 하단에 번호 리스트
- Agent 기둥 아래: `cocrates.md` / Skills 기둥 아래: `.opencode/skills/*/SKILL.md`

## Layout Notes
- Two-column layout with strong visual separation
- Agent side: smaller, focused (헌법 역할 강조)
- Skills side: larger with expansion暗示 (여러 개의 파일 아이콘)
- 3 benefits at bottom with icons

## Reference
- Script: `scripts/script-04.md`
- Slide plan: `slides.md` → Slide 04
