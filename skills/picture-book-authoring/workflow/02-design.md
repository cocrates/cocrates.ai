# Stage ② — Design (World/Characters/Locations + Series/Episode Design)

**Prerequisites:** Approved `overview.md`

**Gate artifacts (design layer):**
- `{project-root}/world-bible.md`
- `{project-root}/characters.md` + `characters/{character-slug}.md`
- `{project-root}/locations.md` + `locations/{location-slug}.md`
- `{project-root}/series.md`
- `episodes/{nnn}-{episode-slug}.md` (episode files include page design)

**Next stage:** `03-evaluate.md` (after user approves design gates)

---

## Procedure

### 2.0 Design constraints (global)

- **아키텍처-퍼스트**: 이미지 생성은 Stage ④에서만.
- **일관성 소스**:
  - 캐릭터 일관성: `characters/{character-slug}.md` (기본 외형/얼굴/실루엣 + 참조 모델 목록)
  - 장소 일관성: `locations/{location-slug}.md` (참조 모델은 `state` 변화 중심, `position`·`view`는 프레이밍 앵커)
  - 이야기/연출 일관성: `episodes/{nnn}-{episode-slug}.md` (페이지별 렌더링 텍스트 + 일러스트 가이드)
- **Design-first fixes**: Evaluate가 문제를 지적하면 “그림 프롬프트만”으로 고치지 말고, 먼저 Design 파일을 수정한 뒤 재평가.

#### 참조 모델(Reference Model)이란?

**참조 모델 = 이야기 전체에서 지속·유지되는 물리적 외형/구조의 정의.**
Stage ④에서 image-generation 스킬을 통해 "참조 이미지"로 생성되며, 이후 모든 페이지 이미지가 이 참조 이미지를 기반으로 일관성을 유지한다.

| 구분 | 참조 모델에 **포함** (= 참조 이미지 생성 대상) | 참조 모델에 **불포함** (= 페이지 일러스트레이션 가이드에서 지정) |
|------|------|------|
| **캐릭터** | 얼굴/체형/실루엣, 복장/장비 state별 차이 | 표정, 포즈, 동작, 감정 표현 |
| **장소** | 공간의 물리적 구조·배치·고정 요소 (가구, 건물, 지형 등) | 조명, 시간대, 날씨, 분위기, 촬영 방향/구도/카메라 앵글, 동적 요소 (지나가는 사람 등) |

**참조 모델 변경이 필요한 경우 (= 새 참조 이미지 필요):**
- 캐릭터: 복장/장비가 바뀜 → 새 state 슬러그 추가
- 장소: 물리적 구조가 영구히 변함 (예: 창밖에 새 건물 완공, 벽이 파괴됨) → 새 state 슬러그 추가
  - 예: 창밖에 보이던 “외부 건물(물리 구조)”이 새로 들어와 화면에 새 요소가 고정적으로 보임 → state 갱신
  - 반대로: 창밖에서 멀리 인물이 지나가거나, 커튼/사람/물체가 잠깐 가렸다가 다시 사라지는 정도(일시·가역) → state 아님(페이지 프롬프트로 처리)

**참조 모델 변경이 불필요한 경우 (= 페이지 프롬프트로 처리):**
- 커튼 개폐, 문 열림/닫힘 등 일시적·가역적 변화
- 조명/시간대/날씨/분위기 변화
- 동적 요소 (지나가는 캐릭터, 날아가는 새 등)

---

### 2.1 World Design

1. Write `world-bible.md` using the following structure:

```markdown
# 세계관

## 세계 이름
{name}

## 세계 설명
{description of the world — rules, atmosphere, key features}

## 핵심 규칙
- {rule 1}
- {rule 2}

## 역사/배경
{background story of the world}
```

---

### 2.2 Character Catalog Design

1. Create `characters.md` as an index (table only).
2. For each character, create `characters/{character-slug}.md` using this structure:

```markdown
# {Character Name}

## 기본 정보
- 역할(주요/조연/악역 등): {role}
- 핵심 욕구/Core Drive: {what they want most}
- 중앙 갈등/Central Conflict: {what prevents them from getting it}

## 외형 및 거동 특징
- 얼굴/체형/실루엣의 “변하면 안 되는 고정 포인트”: {describe in detail}
- 시그니처 제스처/버릇/행동 패턴(거동): {tics, posture habits, typical movement}
- 표정/말투에서 반복되는 느낌(캐릭터성): {expression + speaking style notes}

## 주요 인물과의 관계
- {other-character-slug}: {relationship type, power dynamic, conflict seed}
- {other-character-slug}: ...

## 참조 모델 (외형 변화 + state 목록)
{각 state가 하나의 참조 이미지에 대응됨 — Stage ④ Phase 0에서 image-generation으로 생성}

- base: {기본 복장/무장 상태의 정의}
- {state-slug}: {state가 바뀌는 요소(의상/장비/무장/휴대물) + 얼굴/체형 규칙 유지}
- {state-slug}: ...

참조 모델 규칙:
- state = 물리적 외형 변화만 (복장/장비/무장/휴대물)
- 얼굴/체형/고정 포인트는 모든 state에서 동일하게 유지
- 표정/포즈/동작/감정은 state가 아님 → 페이지 일러스트레이션 가이드에서 지정
- 이야기 진행 중 복장/장비가 바뀌면 새 state 슬러그를 추가
```

```markdown
<!-- `characters.md` index template -->
# 캐릭터 목록

| 이름 | 역할 | 핵심 특징 |
|------|------|----------|
| {name} | {role} | {key trait} |

## 관계도
{character relationships}
```

---

### 2.3 Location Catalog Design

1. Create `locations.md` as an index (table only).
2. For each location, create `locations/{location-slug}.md` using the template below.

```markdown
# {Location Name}

## 기본 정보
- 유형: {world / region / city / building / room / outdoor / ...}
- 서사적 역할: {why this place matters to the story}

## 공간 구성
- 구조/레이아웃: {spatial layout — rooms, corridors, open areas, landmarks}
- 규모감: {size impression — cramped / vast / intimate / etc.}

## 감각 환경
- 조명: {natural light / artificial / dim / harsh / flickering / ...}
- 온도/습도: {cold / warm / humid / dry / seasonal shift}
- 냄새: {dominant scent — dust, food, chemicals, nature, decay, ...}
- 소리/소음: {ambient sound — silence, traffic, machinery, birdsong, echoes, ...}
- 질감/촉감: {what surfaces feel like if touched — rough stone, polished wood, ...}

## 분위기 노트
{이 장소가 주는 심리적 인상 — 안전 / 위협 / 향수 / 고립 / ...}

## 장면 목록 (location → position → view → state)
{이 목록은 Stage ④ "참조 이미지" 생성에 그대로 사용되는 장면 슬러그 목록입니다.}
{각 행이 하나의 참조 이미지에 대응됩니다.}

| position | view | state | 설명 |
|------|------|-------|------|
| {position-slug} | {view-slug} | base | {기본 상태 설명} |
| {position-slug} | {view-slug} | {state-slug} | {변화된 상태 — 화재/파괴/신축/계절 등} |
| ... | ... | ... | ... |

참조 모델 규칙:
- position: `location` 내부의 특정 위치 (예: 거실-창쪽, 거실-문쪽, 침대-머리맡 쪽)
- view: 해당 position에서 **보이는 장면** (예: 창이 보이는 뷰, 침대가 보이는 뷰, 책장이 보이는 뷰). 3D 공간을 직접 관리할 수 없으므로, "무엇이 보이는가"를 기준으로 뷰를 정의한다. 촬영 방향/구도/카메라 앵글은 view가 아니며, 페이지 이미지 생성에서 결정됨.
- state: 공간의 물리적 구조가 영구히 바뀐 상태 (base = 기본). 참조 모델은 state 변화로만 갱신됨.
  - 포함 예: 새 건물 완공으로 창밖 풍경(물리 구조)이 바뀜, 벽 파괴, 가구 재배치
  - 불포함 예: 커튼 개폐, 문 열림/닫힘, 조명 변화, 날씨/시간대, 일시적 등장인물/동적 요소 → 페이지 프롬프트로 처리
- 각 행(position × view × state)이 하나의 참조 이미지에 대응. 참조 모델(물리적 동일성) 갱신은 `state` 변화로만 판단한다.
```

```markdown
<!-- `locations.md` index template -->
# 위치 목록

| 이름 | 유형 | 핵심 특징 |
|------|------|----------|
| {name} | {world/region/location/position} | {key trait} |

## 계층 관계
{location hierarchy}
```

---

### 2.4 Series (Episode List) Design

1. Create `series.md`:
   - `Episode List` (번호/제목/요약/페이지 수(예상))
   - 전체 구조 서술 (감정 피크 + 페이지 넘김 긴장 흐름)

```markdown
# Episode List

| 번호 | 제목 | 요약 | 페이지 수(예상) |
|------|------|------|------------------|
| 1 | {title} | {summary} | {pages} |

## 전체 구조
{story arc overview — emotional peaks, where page-turn tension builds}
```

---

### 2.5 Episode Page Design

For each `episodes/{nnn}-{episode-slug}.md`, design **all pages**:

#### 필수 페이지 필드
각 `Page {idx}`마다:
- `Page 0`은 표지(cover)로 고정
- `페이지 스토리`: 장면이 전달하는 이야기(배경/상황/감정/전개)
- `일러스트레이션 가이드`: 이미지에 반드시 들어갈 “구성요소”를 간결히 기록
  - 등장 캐릭터/장소의 **슬러그**
  - 캐릭터의 **state(복장/장비/무장 상태)** 또는 “기본(base)”
  - 장소의 **position + view** (참조 모델 기준) 및 **state** (물리적 구조 변화가 있는 경우)
  - 촬영 방향/구도/카메라 앵글 (참조 모델이 아닌 페이지별 연출)
- `렌더링 텍스트`: 실제 페이지에 표시될 텍스트(언어 포함) (동화체/읽기 리듬)
- `텍스트–이미지 분업`: 텍스트가 맡는 정보 vs 그림이 맡는 정보
- `페이지 넘김 훅`: 마지막 페이지를 제외하고 다음 장을 당기는 궁금증/긴장

#### Craft rules (적용)
- Page-turn hook: 마지막 페이지 제외
- Text–image split: “겹침 없음”은 필수 아님
  - 대신 **페이지 스토리의 필수 정보가 텍스트+이미지 합으로 충분히 전달되는지**를 검증
- Read-aloud rhythm: 타겟 연령에 맞는 문장/어휘 밀도
- No didactic close: 마지막 페이지는 “설교/교훈 독백”이 아니라 장면으로 마무리
- Age density: `overview.md`의 연령 범위를 기준으로 단어 선택 조정

```markdown
<!-- `episodes/{nnn}-{episode-slug}.md` template -->
# Episode {nnn}: {Title}

## 요약
{episode summary — what happens, emotional arc}

## Craft Notes
- 연령 밀도: {words/vocabulary fit for target age}
- 주제 전달: {how theme is shown in scenes — not stated as sermon}
- 클라이맥스 페이지: {page number}

## 페이지 구성

### Page {N}

#### 페이지 스토리
{이 페이지가 전달하는 이야기 — 배경, 상황, 감정, 전개}
{중요: 페이지 스토리만 읽어도 이 페이지(및 해당 에피소드의 전개)가 명확히 상상되어야 함. 이것이 Stage ④ image-generation의 “message”가 됨.}
{등장 캐릭터·장소는 catalogs에 등록된 이름만}

#### 일러스트레이션 가이드
- 캐릭터: {character-slug} (state: {state-slug 또는 base})
- 장소: {location-slug} / {position-slug} / {view-slug} (state: {state-slug 또는 base})
- 페이지 연출: {조명/시간대/날씨/분위기, 촬영 방향/구도/카메라 앵글, 캐릭터 표정/포즈 등 — 참조 모델이 아닌 이 페이지 고유 연출}
{상세한 이미지 프롬프트는 Stage ④ image-generation YAML 작성 시}

#### 렌더링 텍스트
{실제 페이지에 표시될 텍스트 — 동화체, 소리 내어 읽기 리듬}

#### 텍스트–이미지 분업
- 텍스트가 맡는 것: {…}
- 그림이 맡는 것: {…}
{검증: 페이지 스토리의 필수 정보가 텍스트+이미지 합으로 충분히 전달되는가? (yes / fix)}

#### 페이지 넘김 훅
{다음 장을 넘기게 하는 궁금증·긴장 — 마지막 페이지는 "해소/잔상"으로 표기}
```

---

### 2.x Episode Design Internal Feedback Loop

Episode/page 설계 중 필요한 캐릭터/장소/세계 규칙이 누락되면:
1. 해당 Design 파일을 먼저 수정:
   - 캐릭터 필요 → `characters/{character-slug}.md` 수정
   - 장소 필요 → `locations/{location-slug}.md` 수정
   - 세계 규칙 필요 → `world-bible.md` 수정
2. 수정 후, 영향을 받는 episode 파일을 다시 검토

추가 기준:
- page가 참조하는 **캐릭터 state**(복장/무장 상태)가 catalogs에 정의돼 있는지
- page가 참조하는 **장소 position/view/state**가 catalogs에 정의돼 있는지

---

## Completeness Check (Stage ②)

- [ ] `world-bible.md`, `characters.md`, `locations.md`, `series.md` 존재
- [ ] 모든 `characters/{character-slug}.md`, `locations/{location-slug}.md` 작성됨
- [ ] 모든 `episodes/{nnn}-{episode-slug}.md`에 페이지별 렌더링 텍스트 + 일러스트 가이드가 존재
- [ ] 텍스트–이미지 분업(필수 정보 충분성)과 페이지 넘김 훅 규칙이 적용됨
- [ ] 각 page에서 사용한 캐릭터 state / 장소 position-view-state가 catalogs에 모두 정의돼 있어 Stage ④ Phase 0 참조 이미지 생성이 가능함

---

## Gate G2 (User Approval)

다음에 대해 사용자 승인 전까지 Stage ③으로 넘어가지 않습니다.
- 세계관/룰/배경: 일관적이고 충분히 구체적인가?
- 캐릭터 카탈로그: state별 의상/장비가 맞고, 얼굴/체형 일관성이 유지되는가?
- 장소 카탈로그: location → position → view → state 계층이 명확한가?
- 에피소드/페이지 설계:
  - 페이지 스토리가 감정을 잘 전달하는가?
  - 페이지 스토리만 읽어도 페이지 전개가 명확히 상상되는가?
  - 렌더링 텍스트가 연령·리듬에 맞는가?
  - 텍스트–이미지 분업(필수 정보 충분성)과 페이지 넘김 훅이 있는가?
  - 마지막 페이지가 설교가 아닌 장면으로 끝나는가?
  - 일러스트레이션 가이드가 필요한 장면 요소를 담는가?
  - 참조 이미지 준비성(캐릭터 state/장소 position-view-state가 catalogs에 정의돼 있는가?)

**Do not proceed until G2 is approved.**

