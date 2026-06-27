# Slide 08: 🔄 Merge 전략 — 지식 중복 방지

## Key Message
Knowledge Capture는 새로운 파일을 무분별하게 양산하지 않는다. 기존 지식을 검색하고, 발견 시 덮어쓰지 않고 Merge한다. 지식은 교체가 아니라 누적되어야 한다.

## Content
### Headlines / text
- **문제:** 매번 새 파일 → KB 수천 개 → 회상 불가능
- **Merge 전략 3단계:**
  1. 저장 요청 시 동일 주제 기존 KB 검색 (파일명, 제목, Tags 비교)
  2. 발견 시 **덮어쓰지 않음** — 새 인사이트만 `## Update History`와 함께 추가
  3. 미발견 시 새 파일 생성
- **병합 예시:**
  ```markdown
  ## Update History
  - 2026-06-26: Push/Pull 전략의 판단 규칙 추가
  - 2026-06-20: 최초 작성 (3블록 구조)
  ```
- **💡 포인트:** 지식은 **교체**가 아니라 **누적**되어야 한다. Merge 전략은 지식의 연속성을 보장한다.

### Visual elements
- 상단: 🔄 타이틀과 문제 정의
- 중앙: 3단계 프로세스 다이어그램 (1→2→3 화살표)
- 업데이트 히스토리 예시 코드 블록
- 하단: 💡 강조 박스

## Layout Notes
- 3단계는 수평 프로세스 바 형태
- "덮어쓰지 않음"을 강조 (빨간색 X 표시)
- 업데이트 히스토리는 실제 예시로 신뢰감

## Reference
- Script: `scripts/script-08.md`
- Slide plan: `slides.md` → Slide 08
