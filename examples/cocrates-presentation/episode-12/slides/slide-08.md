# Slide Content Specification — Slide 8

## Slide 8: 📋 각 단계별 상세 — Kernel부터 Verification까지

**Section:** 본론 2 — Meta Snowflake 7단계
**Type:** 개념/상세
**Time:** ~1분 30초

### Visual Design
- **Background:** 시리즈 공통 템플릿
- **Layout:** 수직 리스트 (7단계 순서대로)
  - 상단: 제목 "각 단계별 상세" (32pt)
  - 중단: 7단계를 수직으로 나열. 각 단계: 단계명 + 간단한 설명 + 아이콘
  - Outline 단계는 별도 박스로 확대 (5가지 항목 코드 블록 포함)
  - 하단: 💡 포인트

### Content Elements
| 요소 | 내용 | 비고 |
|------|------|------|
| 제목 | 📋 각 단계별 상세 — Kernel부터 Verification까지 | 상단 |
| Kernel | "이 스킬은 [대상 상황]에서 [산출물 유형]을 [검토 가능한 단계]를 거쳐 생성하도록 돕는다" — Kernel이 흔들리면 모든 후속 단계 흔들림 | define |
| Components | 5대 차원을 산출물 유형에 맞게 명명하고 범위 설정 | plan |
| Frame | 5단계 Snowflake 워크플로우 + 디렉토리 구조 + 승인 지점 설계 | arch design |
| **🚨 Outline** | **각 단계별 파일 산출물 정의 (5가지 항목):** | detail design (강조) |
| Outline Format | • Input: 의존하는 산출물<br>• Creation activity: 생성 작업<br>• Completion criteria: 완료 조건<br>• Review questions: 검토 질문<br>• Approval criteria: 승인 조건 | 코드 블록 형태 |
| Spec | frontmatter, 원칙, 사용 시점, 절차, 대화 규칙, 금지 사항, 완료 조건 확정 | generation |
| Skill | `.opencode/skills/{slug}/SKILL.md` 파일 생성 | generation |
| Verification | 스킬이 Cocrates 생성 하네스로 작동하는지 검증 | generation |
| 💡 포인트 | "Outline Format이 중요한 이유 — 중간 산출물이 채팅에 휘발되지 않고 파일로 관리되도록 보장" | 하단 박스 |

### Layout Notes
- 7단계 수직 리스트. 각 단계는 작은 수평 구분선으로 분리.
- Outline 단계는 더 넓은 박스로 확대. 5가지 항목은 코드 블록 스타일로 표시.
- 다른 단계는 간결하게 — Outline만 상세 표시.

### Reference
- Script: `scripts/script-08.md`
- Slide plan: `slides.md` → Slide 8
