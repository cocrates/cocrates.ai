# Slide Content Specification — Slide 9

## Slide 9: 🚫 금지 사항 + 🧬 재사용성과 진화

**Section:** 본론 3 — 제약과 진화
**Type:** 원칙/요약
**Time:** ~2분

### Visual Design
- **Background:** 시리즈 공통 템플릿
- **Layout:** 좌우 2컬럼
  - 상단: 제목 "제약과 진화" (32pt)
  - 좌측 컬럼: "🚫 금지 사항" — 6가지 ❌ 리스트
  - 우측 컬럼: "🧬 재사용성과 진화" — 스킬 생태계 다이어그램
  - 하단: 💡 포인트

### Content Elements (좌측)
| 요소 | 내용 | 비고 |
|------|------|------|
| 섹션 제목 | 🚫 금지 사항 (Constraints) | 좌측 상단 |
| 금지 1 | ❌ 스킬 없는데 최종 산출물부터 생성 | 구조적 접근 생략 |
| 금지 2 | ❌ 구조 분석 없이 SKILL.md 템플릿만 채우기 | 5대 프리즘 분석 생략 |
| 금지 3 | ❌ 사용자 검토 지점 없는 스킬 작성 | 승인 게이트 없음 |
| 금지 4 | ❌ 중간 산출물을 파일로 저장하지 않음 | 채팅 휘발 |
| 금지 5 | ❌ 기존 스킬과 중복되는 새 스킬 | 시스템 효율성 저하 |
| 금지 6 | ❌ `compatibility: opencode` / `metadata.agent: cocrates` 누락 | 하네스 호환성 |

### Content Elements (우측)
| 요소 | 내용 | 비고 |
|------|------|------|
| 섹션 제목 | 🧬 재사용성과 진화 | 우측 상단 |
| 자동 로드 | 한 번 만든 스킬 → 이후 동일 유형 요청 시 시스템이 자동 로드 | 화살표 다이어그램 |
| artifact-specific skill 탄생 | document-authoring, presentation-authoring, blog-series-authoring의 탄생 과정 | 연결선 |
| 확장성 | 새 산출물 유형 = 스킬만 추가. 기존 스킬은 독립적 개선 가능 | 하단 |

### Layout Notes
- 좌우 6:4 비율. 좌측에 금지 사항, 우측에 재사용성 다이어그램.
- 금지 사항은 ❌ 아이콘과 함께 빨간 계열로.
- 재사용성은 "generating-skill-creation → 새 스킬 → 시스템 등록 → 자동 로드" 순환 다이어그램.
- 💡 포인트: "generating-skill-creation은 '만들고 사라지는' 스킬이 아니라 스킬 생태계의 출발점"

### Reference
- Script: `scripts/script-09.md`
- Slide plan: `slides.md` → Slide 9
