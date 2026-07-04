# Validation — Cocrates User Manifesto

## 1. Purpose Fit

| Criterion | Result | Note |
|-----------|--------|------|
| Purpose aligned with slides? | ✅ | "7계명을 처음 접하는 개발자에게 소개" — Slide 01~13 모두 부합 |
| Audience (초~중급 개발자) 적합? | ✅ | on-slide copy와 스크립트 모두 초급 개발자가 이해할 수 있는 수준 |
| 3가지 구체 목표(overview.md L7-9) 반영? | ✅ | Cocrates 철학(Slide 02), 7계명(Slides 05~11), 도구+의지 결합(Slides 12~13) |

## 2. Narrative Arc

| Section | Slides | Duration | Flow | 평가 |
|---------|--------|----------|------|------|
| 도입 | 01~04 | ~5.5분 | 타이틀 → Cocrates 정의 → 문제 제기 → 선언 | ✅ 자연스러운 도입 |
| 7계명 | 05~11 | ~14분 | 4❌ + 3✅ 계명 순차 선언 | ✅ 단일 구조로 일관성 |
| 종합 | 12~13 | ~2.5분 | AI-native Engineer 선언 → 원문 제시 | ✅ 강력한 마무리 |
| **Total** | **13장** | **~22분** | | ⚠️ 20분 초과 가능성 |

**평가:** 전체 흐름은 완결성 있음. 도입→선언→종합의 3막 구조가 명확. 단, 시간이 20분을 2분 정도 초과할 수 있어 발표 시 속도 조절 또는 일부 슬라이드 압축 필요.

## 3. Clarity

| Check | Result | Evidence |
|-------|--------|----------|
| 1 slide = 1 message? | ✅ | 각 슬라이드가 하나의 포커스 메시지를 가짐 |
| on-slide copy 충분? | ✅ | 영문 계명 + 한글 + 설명 + 포인트 — 슬라이드만 봐도 내용 전달됨 |
| Script와 중복 없는가? | ✅ | 스크립트는 on-slide copy를 확장/해설할 뿐, 동일 내용 반복 아님 |
| 7계명 구조 일관성? | ✅ | Slide 05~11 모두 동일한 focal 레이아웃 + 순차 등장 |

## 4. Visual Design — Spec & Template Consistency

### Template assignment consistency

| Slide | Template | slides.md | slide-*.md | Match? |
|-------|----------|-----------|------------|--------|
| 01 | template-01 | ✅ | ✅ | ✅ |
| 02 | template-02 | ✅ | ✅ | ✅ |
| 03 | template-02 | ✅ | ✅ | ✅ |
| 04 | template-02 | ✅ | ✅ | ✅ |
| 05 | template-02 | ✅ | ✅ | ✅ |
| 06 | template-02 | ✅ | ✅ | ✅ |
| 07 | template-02 | ✅ | ✅ | ✅ |
| 08 | template-02 | ✅ | ✅ | ✅ |
| 09 | template-02 | ✅ | ✅ | ✅ |
| 10 | template-02 | ✅ | ✅ | ✅ |
| 11 | template-02 | ✅ | ✅ | ✅ |
| 12 | template-02 | ✅ | ✅ | ✅ |
| 13 | template-02 | ✅ | ✅ | ✅ |

### Template spec compliance

| Rule | Status | Note |
|------|--------|------|
| Template-02 배경(#0a0a0a) 일관? | ✅ | 모든 slides/slide-*.md에서 `#0a0a0a` 명시 |
| 타이포그래피 계층 일관? | ✅ | 영문 계명(큼), 한글(중간), 설명(보통), 포인트(작게) |
| 색상 팔레트 일관? | ✅ | #ffffff / #cccccc / #888888 / #666666 계층 유지 |
| 선 굵은 미니멀 스타일? | ✅ | 장식 최소화, 텍스트 중심, 넉넉한 여백 |
| ❌/✅ 아이콘 분류 (1~4 ❌, 5~7 ✅)? | ✅ | slide-*.md Type과 content elements 모두 일치 |
| Template에 slide 번호 미기재? | ✅ | `templates/template-01.md`, `template-02.md` 모두 slide 번호 불포함 |
| Template-02 의도 범위 일치? | ✅ | Slides 02~13 사용, template-02 Intended Use 업데이트 완료 |

## 5. Script / Timing

| Slide | Script exists? | Estimated | Script tone 적합? |
|-------|----------------|-----------|-------------------|
| 01 | ✅ | 0.5분 | ✅ 무게감 있는 오프닝 |
| 02 | ✅ | 2분 | ✅ 개념 설명 또렷 |
| 03 | ✅ | 2분 | ✅ 드라마틱한 반전 |
| 04 | ✅ | 1분 | ✅ 선언적 엄숙함 |
| 05 | ✅ | 2분 | ✅ 투쟁적, 공감 유도 |
| 06 | ✅ | 2분 | ✅ 경계심 고취 |
| 07 | ✅ | 2분 | ✅ 주권 강조 |
| 08 | ✅ | 2분 | ✅ 용기와 겸손 |
| 09 | ✅ | 2분 | ✅ 투지 폭발 |
| 10 | ✅ | 2분 | ✅ 최적 추구 집중 |
| 11 | ✅ | 2분 | ✅ 자만심 경계 |
| 12 | ✅ | 1.5분 | ✅ 종합 + 확장 비유 |
| 13 | ✅ | 1분 | ✅ 원문 낭독, 여운 |

**총 예상 시간: ~22분** (20분 목표 대비 2분 초과)

**권장사항:**
- Slide 02~04에서 각 20초씩 단축 → 약 1분 확보
- 또는 Slide 05~11 각각 10초씩 단축 → 약 1분 10초 확보
- 실제 발표 리허설 후 조정 권장

## 6. Completeness

| Stage | Artifact | Status |
|-------|----------|--------|
| ① Define | `overview.md` | ✅ |
| ② Plan | `outline.md` | ✅ |
| ③ Architecture | `slides.md` | ✅ |
| ④ Scripts | `scripts/script-01~13.md` | ✅ |
| ⑤ Template design | `templates/template-01.md`, `template-02.md` | ✅ |
| ⑥ Slide design | `slides/slide-01~13.md` | ✅ |
| ⑦ Validation | `q-and-a.md` + `validation.md` | ✅ |

**Phase A 완료 조건 모두 충족:** ✅

## 7. 사소한 발견

| # | 항목 | 상태 | 조치 |
|---|------|------|------|
| 1 | template-02 Intended Use에 "Slides 02–12"로 기재 → Slide 13 누락 | ✅ 수정완료 | `templates/template-02.md` |
| 2 | template-02 accent color 설명에 "거부 계명 1~5" → 현재 1~4와 불일치 | ✅ 수정완료 | 숫자 제거, 단순화 |
| 3 | overview.md L7 "Cocrates의 3대 철학" — 슬라이드에서 '3대 철학'이 명시적으로 드러나지 않음 | ⚠️ 정보 | Slide 02에서 Cocrates 정의 설명 시 포함되므로 큰 문제 아님 |
| 4 | overview.md L18 "13장 (~1분 30초/슬라이드)" — 실제 추정 합계 22분 | ⚠️ 정보 | 리허설 필요 |

## 8. 최종 결론

| 항목 | 판정 |
|------|------|
| 목적 적합성 | ✅ **PASS** |
| 내러티브 완결성 | ✅ **PASS** |
| 명확성 | ✅ **PASS** |
| 시각적 일관성 | ✅ **PASS** (사소한 수정 완료) |
| 스크립트 완비 | ✅ **PASS** |
| 전체 완성도 | ✅ **PASS** (시간 관리만 유의) |

**Phase A 완료.** 사용자 승인 시 Phase B(렌더링)로 진행 가능.
