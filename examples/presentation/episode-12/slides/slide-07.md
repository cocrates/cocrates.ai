# Slide Content Specification — Slide 7

## Slide 7: 🏗️ Meta Snowflake 7단계 개관

**Section:** 본론 2 — Meta Snowflake 7단계
**Type:** 개념/개관
**Time:** ~1분 30초

### Visual Design
- **Background:** 시리즈 공통 템플릿
- **Layout:** 상단 7단계 파이프라인 + 하단 매핑 표
  - 상단: 제목 "Meta Snowflake 7단계" (32pt)
  - 중단 상부: 7단계 수평 파이프라인 다이어그램
    - Kernel → Components → Frame → **🚨 Outline** → Spec → Skill → Verification
  - 중단 하부: Snowflake 매핑 표 (Meta 단계 | Snowflake 대응 | 핵심 활동)
  - 하단: 게이트 강조 + 💡 포인트

### Content Elements
| 요소 | 내용 | 비고 |
|------|------|------|
| 제목 | 🏗️ Meta Snowflake 7단계 개관 | 상단 |
| 파이프라인 | Kernel → Components → Frame → **🚨 Outline** → Spec → Skill → Verification | 7단계 수평 배열 |
| 매핑: Kernel | define | 생성 대상을 **한 문장**으로 정의 | 1행 |
| 매핑: Components | plan | 5대 차원 식별 | 2행 |
| 매핑: Frame | architecture design | 워크플로우 + 파일 구조 + 승인 지점 설계 | 3행 |
| 매핑: **🚨 Outline** | **detail design** | 단계별 파일 산출물, 입력, 완료 기준, 승인 조건 정의 | 4행 (강조) |
| 매핑: Spec→Skill→Verification | generation | 명세 확정 → 파일 생성 → 검증 | 5-7행 |
| 🚨 게이트 | "Outline(detail design)이 완전히 확정되기 전까지 Spec/Skill(generation)로 넘어가지 않음" | 강조 |
| 소설 비유 연결 | "Outline = 장면 리스트. 장면 리스트 없이 집필하지 않음" | 게이트 옆 |
| 💡 포인트 | "이 7단계는 이 발표자료 슬라이드의 구조를 설계할 때와 동일한 논리로 작동한다" | 하단 박스 |

### Layout Notes
- 7단계 파이프라인은 상단에 크게 표시. Outline 단계는 빨간 테두리/색상으로 강조.
- 매핑 표는 3열. Meta 단계 열은 굵게, Snowflake 열은 중간, 활동 열은 일반.
- 게이트 강조는 🚨 아이콘과 함께 별도 박스.

### Reference
- Script: `scripts/script-07.md`
- Slide plan: `slides.md` → Slide 7
