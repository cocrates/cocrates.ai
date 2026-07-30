# Stage ③ — Evaluate (Criteria + Craft + Visual Reference Integrity + Story Lock)

**Prerequisites:** Approved Stage ② design artifacts (`world-bible.md`, `characters*`, `locations*`, `series.md`, `episodes*`).

**Gate artifacts:**
- `evaluations/{nnn}-{episode-slug}.md` (story lock decision per episode)

**Next stage:** `04-generate.md` (after user approves story lock gates)

---

## Procedure

### 3.1 Evaluate per episode

For each `episodes/{nnn}-{episode-slug}.md`:

1. Load `overview.md` and extract **Validation Criteria**.

2. Check **episode/page craft completeness** (design correctness):
   - Page 0/표지 규칙 준수
   - 모든 `Page {N}`에 `페이지 스토리`, `일러스트레이션 가이드`, `렌더링 텍스트`, `텍스트–이미지 분업`, `페이지 넘김 훅`(마지막 페이지 제외)이 존재
   - Page-turn hooks exist on non-final pages
   - Text–image split is not redundant captioning
   - Rendering text fits target age (read-aloud rhythm + word density)
   - No didactic closing monologue (theme emerges via final scene/image)

3. Check **Reference Model Integrity** (그림책 장면 일관성의 핵심):
   - 캐릭터 참조 모델(= state) 무결성
     - 각 page에서 사용하는 캐릭터는 `characters/{character-slug}.md`에 정의된 **state**(또는 base) 중 하나를 명시하는가
     - 표정/자세/동작/감정 표현은 “참조 모델 state”가 아니라 페이지 가이드(페이지 이미지 생성 결정)로 남아 있는가
     - 캐릭터 state는 이야기에서 “물리적 외형/복장 변화”가 있을 때만 갱신되는가
   - 장소 참조 모델(= state) 무결성
     - 각 page에서 사용하는 장소는 `locations/{location-slug}.md`에 정의된 **scene 축(position + view)** 및 해당 scene의 **state(base 또는 state-slug)**를 명시하는가
     - 조명/시간대/날씨/분위기/촬영 방향/구도/카메라 앵글/동적 요소는 참조 모델 state가 아닌 페이지별 연출로 취급되는가
     - “창밖 풍경”처럼 보이는 내용이 **동적 요소 수준**이면 state 변경 없이 페이지 프롬프트로 처리되는가
     - “새 건물 완공/벽 파괴”처럼 물리 구조가 고정적으로 바뀌면 새 state로 갱신되는가

4. Check **Scene Continuity & Visual Consistency**:
   - 연속된 페이지에서 “같은 장면”으로 묶이는 구간은 동일한 캐릭터 state / 동일한 장소 scene(position+view) / 동일한 state를 유지하는가
   - 변경이 필요한 순간에는 어떤 변경인지(캐릭터: 복장/장비 state, 장소: 물리 구조 state)를 명확히 표시했는가
   - `position`/`view` 의미가 “무엇이 보이는가(보이는 장면)” 관점으로 정의되어 있는가

5. Check **Text–Image Collaboration**:
   - 페이지 스토리의 핵심 정보가 `렌더링 텍스트 + 이미지가 담아야 할 정보` 합으로 충분히 전달되는가
   - “겹침 없음”은 강제하지 않되, 중복으로 인해 중요한 정보가 사라지지 않는가
   - 페이지 넘김 훅이 실제로 다음 페이지 내용을 자연스럽게 기대하게 만드는가

6. Persona checks (전문가 시각):
   - 아이 관점(독자 몰입/이해/넘김 욕구)
   - 부모 관점(안전/가치/설교 과잉 없음)
   - 평론가 관점(품질/협응/리듬)
   - 일러스트 전문가 관점(시각 일관성/참조 모델 연결/장면 명확성)

7. Record results into `evaluations/{nnn}-{episode-slug}.md`.

---

### 3.2 Evaluation Record Template (put on disk)

`evaluations/{nnn}-{episode-slug}.md`:

```markdown
# Episode {nnn} 평가

## 1. Criteria Check (from overview.md Validation Criteria)
| Criterion | Result | Evidence |
|-----------|--------|----------|
| {criterion from overview} | ✅ / ⚠️ / ❌ | {page / quote / note} |

## 2. Craft Checks (Design correctness)
| Check | Result | Evidence |
|-------|--------|----------|
| 모든 Page 필수 필드 존재 | ✅ / ⚠️ / ❌ | {page} |
| Page 0은 cover | ✅ / ⚠️ / ❌ | {page} |
| Page-turn hook (비-마지막 페이지) | ✅ / ⚠️ / ❌ | {page} |
| Text–image split (중복 캡션/군더더기 없음) | ✅ / ⚠️ / ❌ | {page} |
| Read-aloud rhythm / age density | ✅ / ⚠️ / ❌ | {page} |
| No didactic closing monologue | ✅ / ⚠️ / ❌ | {last page} |
| 캐릭터/장소는 catalogs 등록 엔터티만 사용 | ✅ / ⚠️ / ❌ | {entities} |

## 3. Reference Model Integrity Checks (Visual lock)

### 3.1 캐릭터 참조 모델(= state) 검사
| Check | Result | Evidence |
|-------|--------|----------|
| 모든 page에 캐릭터 state(또는 base)가 명시됨 | ✅ / ⚠️ / ❌ | {page + character-slug} |
| state는 “복장/장비/무장/휴대물” 같은 물리 변화만 반영 | ✅ / ⚠️ / ❌ | {page} |
| 표정/자세/동작은 page 가이드에 있고 state로 고정되지 않음 | ✅ / ⚠️ / ❌ | {page} |
| 물리 변화가 있을 때만 state가 갱신됨 | ✅ / ⚠️ / ❌ | {page} |

### 3.2 장소 참조 모델(= state) + scene 정의 검사
| Check | Result | Evidence |
|-------|--------|----------|
| 모든 page에 location + position + view가 명시됨 | ✅ / ⚠️ / ❌ | {page} |
| 조명/시간/날씨/분위기/촬영 방향/구도/카메라 앵글은 참조 모델 state가 아님 (view는 “무엇이 보이는가” 기준) | ✅ / ⚠️ / ❌ | {page} |
| “창밖 풍경”의 변화 기준이 올바름
|  - 동적(일시/가역) 변화는 state 아님 | ✅ / ⚠️ / ❌ | {page + 예시} |
|  - 물리 구조(새 건물/벽 파괴 등) 변화는 새 state로 갱신 | ✅ / ⚠️ / ❌ | {page + 예시} |
| 연속 장면 구간에서 위치(position)·보이는 뷰(view)·state 유지가 일관됨 | ✅ / ⚠️ / ❌ | {page range} |


### 3.3 Staging reference model (= continuing-situation blocking)
| Check | Result | Evidence |
|-------|--------|----------|
| 연속 상황에 staging이 있는가 | ✅ / ⚠️ / ❌ | {span + staging-slug} |
| staging 참조 뷰 2–3 계획 | ✅ / ⚠️ / ❌ | {staging file} |
| 페이지가 staging을 인용하는가 | ✅ / ⚠️ / ❌ | {page} |
| 무단 L/R·좌석·스테이션 변경 없음 | ✅ / ⚠️ / ❌ | {page range} |

### 3.4 Illustration Guide Completeness
| Check | Result | Evidence |
|-------|--------|----------|
| illustration guide가 “캐릭터 참조 모델 + 장소 scene”을 모두 포함 | ✅ / ⚠️ / ❌ | {page} |
| page story가 illustration guide의 표시 요소와 모순되지 않음 | ✅ / ⚠️ / ❌ | {page} |

### 3.5 일러스트 텍스트(시각 효과) 가이드 적합성
| Check | Result | Evidence |
|-------|--------|----------|
| `렌더링 텍스트`의 읽기 순서(좌→우, 위→아래)가 자연스러운가 | ✅ / ⚠️ / ❌ | {page} |
| 대사/감탄사 강조(가장 크게+굵게, 가능하면 글로우/아웃라인)가 명확한가 | ✅ / ⚠️ / ❌ | {page} |
| 내레이션 스타일(더 작게+둥근 세리프+따뜻한 그림자)이 일관적인가 | ✅ / ⚠️ / ❌ | {page} |
| 중요 단어/표현이 크기 또는 색상으로 구분되어 강조되는가 | ✅ / ⚠️ / ❌ | {page} |
| 텍스트가 얼굴/핵심 비주얼을 가리지 않도록 `anchor area`가 제안되는가 | ✅ / ⚠️ / ❌ | {page} |
| 8세 이상 & 텍스트가 많은 페이지는 본문을 “그림 속 이야기 섹션(패널/박스)”으로 분리하는가(권장) | ✅ / ⚠️ / ❌ | {page} |

## 4. Scene Continuity & Visual Consistency Checks
| Check | Result | Evidence |
|-------|--------|----------|
| 연속 페이지에서 동일성(캐릭터/장소/구조) 유지 | ✅ / ⚠️ / ❌ | {page range} |
| 변경 시점에 변경 종류(캐릭터 state vs 장소 state vs 동적 요소)가 구분됨 | ✅ / ⚠️ / ❌ | {page} |
| position/view가 “무엇이 보이는가” 기준으로 정의되어 장면이 명확함 | ✅ / ⚠️ / ❌ | {page} |

## 5. Text–Image Collaboration Checks
| Check | Result | Evidence |
|-------|--------|----------|
| 페이지 스토리의 핵심 정보가 text+image로 충분히 전달 | ✅ / ⚠️ / ❌ | {page} |
| 중복/누락으로 인해 다음 페이지 기대가 무너지는지 | ✅ / ⚠️ / ❌ | {page} |
| 페이지 넘김 훅이 자연스러운 긴장/호기심으로 연결 | ✅ / ⚠️ / ❌ | {page} |

## 6. Persona Checks
### 아이 관점
- 재미도: {rating}
- 이해도: {rating}
- 페이지 넘기고 싶은가: {rating}
- 피드백: {feedback}

### 부모 관점
- 가치/공유 의향: {rating}
- 안전성: {rating}
- 설교 과잉 없음: {rating}
- 피드백: {feedback}

### 평론가 관점
- 스토리 완결성: {rating}
- 캐릭터 매력: {rating}
- 텍스트–이미지 협응: {rating}
- 피드백: {feedback}

### 일러스트 전문가 관점
- 참조 모델 일관성: {rating}
- 장면 명확성(무엇이 보이는가): {rating}
- 페이지 프레이밍 고정 준수: {rating}
- 텍스트 시각 효과 가이드 제공(폰트/강조/읽기 순서/anchor area/패널 분리): {rating}
- 텍스트 오버레이 가이드 (Stage④ YAML 반영용): {feedback}
  - Reading order anchor plan: {좌→우/위→아래 배치 규칙 + 텍스트 영역 맵}
  - TextBox list: {line-index → anchor area / size / font style / color / emphasis}
  - Dialogue/exclamation style: {largest+bold + (글로우/아웃라인 등) 적용 방식}
  - Narration style: {smaller font + (serif/sans) + 그림자 적용 방식}
  - Key word emphasis: {어떤 단어/표현을 어떻게 강조하는지}
  - 8세+ 텍스트 많은 경우: {본문 패널/박스 분리 yes/no + 위치/크기}

## 7. 수정 사항 (Design-First)
| # | Finding | Severity | 적용 위치(Design 파일 / Stage④ YAML overlay 반영) | 개선안(무엇을 바꿀지) | Action Status |
|---|---------|-----------|---------------------------|---------------------------|----------------|
| 1 | {finding} | High/Med/Low | characters/... / locations/... / episodes/... | {proposed edit} | {todo/done} |
| 2 | {finding} | ... | ... | ... | ... |

## 8. Story Lock Readiness (G3 체크)
- [ ] Reference Model Integrity: ✅/⚠️/❌ 모든 항목
- [ ] Scene Continuity & Visual Consistency: ✅/⚠️/❌ 모든 항목
- [ ] Illustration Guide Completeness: ✅/⚠️/❌ 모든 항목
- [ ] Illustration Text Styling/Effects Guide: ✅/⚠️/❌ 모든 항목
- [ ] Craft Checks: ✅/⚠️/❌ 모든 항목
```

---

### 3.3 Evaluation perspectives (use as check rubric)
| Perspective | Focus | Typical questions |
|---|---|---|
| 아이(독자) | 이해/즐거움 | 재밌는가? 따라갈 수 있는가? 넘기고 싶어지는가? |
| 부모(구매자) | 안전/가치 | 안전한가? 공유하고 싶은가? 설교 과잉은 없는가? |
| 평론가(품질) | 완성/협응 | 스토리 완결성, 텍스트-이미지 협응, 리듬 적합성 |
| 일러스트 전문가 | 시각 일관성 | 참조 모델 연결이 끊기지 않는가? 장면이 명확히 “보이는가”? |

---

## Design-First Revision Loop

If Evaluation finds issues:

1. Return to **Stage ② (Design)** and update the relevant Design files:
   - `characters/{character-slug}.md` (캐릭터 state)
   - `locations/{location-slug}.md` (location scene: position/view/state)
   - `episodes/{nnn}-{episode-slug}.md` (페이지별 illustration guide/렌더링 텍스트)
2. Re-run affected checks and update `evaluations/{nnn}-{episode-slug}.md`.
3. Only after user approves the story lock (G3), proceed to Stage ④.

---

## Approval Gate G3 — Story Lock (per episode)

User confirms:

1. Criteria Check — all items addressed?
2. Craft Checks — satisfactory?
3. Reference Model Integrity — 캐릭터 state / 장소 position-view-state가 catalogs와 모순 없이 정의되어 있는가?
4. Scene Continuity & Visual Consistency — 연속 장면의 동일성이 유지되고, 변경은 필요한 곳에서만 발생하는가?
5. Persona feedback items — 잔여 리스크가 허용 가능한 수준인가?
6. **Story lock:** episode page stories, rendering text, text–image split, page-turn hooks, and illustration guides are now frozen for image generation.

**Do not proceed to Stage ④ until G3 is approved.**

After G3:
- Treat the episode’s design as locked.
- Any later change to story/text/craft/characters/locations/world rules requires rollback to Stage ② and a new G3 approval before regenerating images.
