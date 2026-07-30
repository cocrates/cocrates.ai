# Stage ④ — Generate (Reference Images → Page Images)

**Prerequisites:** Stage ③ story lock approved (per episode).

**Gate artifacts:**
- reference images and page images generated into `images/characters`, `images/locations`, `images/{nnn}-{episode-slug}/`
- final assembly confirmation in `output/{book-slug}-final/` (user approval)

**Next stage:** none (this is the terminal stage for the skill)

---

## Hard Boundary

To preserve story lock:
- Allowed in ④:
  - image YAML prompts that implement locked illustration guides
  - re-rendering reference/page PNGs for quality/consistency (max 2 retries per image)
  - rendering the locked `렌더링 텍스트` inside each page image
- Not allowed in ④:
  - rewriting episode/page story or rendering text
  - changing illustration guide meaning (only prompt phrasing/tightening)

If a problem is a **Design gap** (wrong cast/place/state/unclear guide):
- Stop generation
- Roll back to Stage ②
- Re-run Stage ③ for affected episode(s)

---

## Phase 0: Reference Image Generation (Characters + Locations)

Goal: 참조 모델(캐릭터·장소·staging — `workflow/reference-models.md`)을 참조 이미지로 생성하여, 이후 모든 페이지 이미지의 일관성 기반을 확보한다.

**포함:** 캐릭터 얼굴/체형/복장·장비(state별), 장소 세트 구조(state별), staging 앙상블 배치(2–3뷰).
**불포함:** 표정/포즈, 조명/시간/날씨/분위기, 일회성 카메라, 동적 요소 — 페이지 연출.

### 0.1 Generation order (per catalogs)

1. Character base references: `images/characters/{character-slug}.yaml/.png`
2. Character state variants: `images/characters/{character-slug}-{state-slug}.yaml/.png`
3. Location base references: `images/locations/{location-slug}.yaml/.png`
4. Location position/view references: `images/locations/{location-slug}-{position-slug}-{view-slug}.yaml/.png`
5. Location state variants: `images/locations/{location-slug}-{position-slug}-{view-slug}-{state-slug}.yaml/.png`
6. Staging ensemble refs (2–3 views): `images/stagings/{staging-slug}-{establishing|reverse|detail}.yaml/.png` — see `workflow/reference-models.md`

### 0.2.1 샘플: reference image YAML 구조

**Character base YAML structure** (`images/characters/{character-slug}.yaml`):

```yaml
type: image
model: gemini-3.1-flash-image

params:
  prompt: |
    {picture-ready character description in English}
    A children's book illustration style character reference sheet...
  size: 1K
  aspectRatio: "1:1"
  seed: null

output: "./images/characters/{character-slug}.png"
```

**Location base YAML structure** (`images/locations/{location-slug}.yaml`):

```yaml
type: image
model: gemini-3.1-flash-image

params:
  prompt: |
    {picture-ready location description in English}
    A children's book illustration background reference (location base / establishing view)...
  size: 1K
  aspectRatio: "4:3"
  seed: null

output: "./images/locations/{location-slug}.png"
```

**Location position/view YAML structure** (`images/locations/{location-slug}-{position-slug}-{view-slug}.yaml`):

```yaml
type: image
model: gemini-3.1-flash-image

params:
  prompt: |
    {picture-ready location description in English}
    A children's book illustration background reference...
  size: 1K
  aspectRatio: "4:3"
  seed: null

output: "./images/locations/{location-slug}-{position-slug}-{view-slug}.png"
```

### 0.2 Per-image YAML approval (MANDATORY)

For each reference image:
1. Create image-generation YAML that reflects the approved Design notes.
   - YAML `title/message/design` role boundary:
     - `message`: 참조 이미지의 “의미(정체성/구조 일관성)”만 간단히(스토리/배치/폰트 디테일 중복 금지)
     - `design`: 참조 이미지의 “시각 구현”만(캐릭터 얼굴/체형/복장 state, 장소 물리 구조 state, 프레이밍 앵커; on-image text 없음)
2. Show the full YAML to the user and request explicit approval.
3. Only after explicit approval, call the MCP image generation.
4. Verify visual quality (do not treat it as story redesign).

---

## Phase 1: Page Image Generation (Include text overlay)

For each episode:
- For each page index (e.g. `{00}`, `{01}`, ...), generate:
  - `images/{nnn}-{episode-slug}/{page-idx}.yaml`
  - `images/{nnn}-{episode-slug}/{page-idx}.png`

### 1.1 Per-page YAML approval (MANDATORY)

For each page image:
1. Create `{page}.yaml` that:
   - uses character/location/staging reference PNGs; keep seating/L-R from staging when cited
   - implements only the locked illustration guide action
   - renders the locked `렌더링 텍스트` inside the image as part of illustration
   - uses the locked `렌더링 텍스트` verbatim (언어 포함 그대로 / 번역·재서술·요약·어순 변경 금지)
   - (참고) if `evaluations/{nnn}-{episode-slug}.md`에 일러스트 전문가의 텍스트 시각 효과 가이드가 있다면, 해당 가이드를 그대로 텍스트 박스 anchor area / 폰트/강조 규칙에 반영
   - YAML `title/message/design`은 아래 역할 경계를 지켜 중복을 피함:
     - `title`: `{episode-slug} / Page {idx}` 같은 짧은 이미지 이름(서사/샷리스트 금지)
     - `message`: `페이지 스토리`에서 “무엇이 전달돼야 하는가(의미/정서/관계/전개 비트)”만 1~2문장으로 요약(shot list/배치/폰트 중복 금지)
     - `design`: `일러스트레이션 가이드` + (TextBox/anchor area/폰트/색/강조/패널 분리)처럼 “그림으로 어떻게 구현할지”만 작성(episode 설계를 근거로만 사용)
2. Show YAML to user for explicit review and approval
3. On explicit approval only → call MCP generate
4. Verify:
   - visual fidelity to the locked guide
   - text readability and correct placement

### 1.2 Prompt pattern (guideline)

Using the provided reference images:
- Keep character appearance, setting, lighting, and style exactly the same
- Change only the character’s action described in the locked illustration guide
- Add extra elements only if the locked guide lists them

TEXT OVERLAY (Episode rendering text):
- Render every line from locked `렌더링 텍스트` directly into the image
- 배치/읽기 순서(중요):
  - 텍스트 박스는 원칙적으로 `왼쪽 → 오른쪽 → 위 → 아래`의 읽기 흐름이 되도록 배치
  - 여러 줄일 경우, 같은 영역에서는 줄이 자연스럽게 다음 줄로 이어지게 배치
- 대사/감탄사(exclamation)는 가장 크게 + 굵게(가능하면 얇은 글로우/아웃라인) 강조
- 내레이션(narration)은 더 작게 + 소프트 둥근 세리프(또는 유사) + 따뜻한 그림자
- 중요 단어/표현은 (a) 크기 확대 또는 (b) 색상 변경(또는 볼드) 중 하나로 강조
- Do not obscure faces/key visuals with text

**(참고) SKILL 템플릿 프롬프트 패턴**

```
Using the provided reference images:
- Keep the character's appearance, setting, lighting, and style exactly the same
- Change only the character's action: {action from locked illustration guide}
- [Add any new elements only if listed in the locked illustration guide]

TEXT OVERLAY — render these episode `렌더링 텍스트` lines directly into the image as part of the illustration:
{For each line in 렌더링 텍스트, specify:}
1. {TextBox (ReadingOrder)}: {text} — {Anchor area (예: top-left / top-right / bottom-left / bottom-right or 좌상단/우상단 등), size, font style, color, emphasis level}
2. ...
{Place text in areas that do not obscure characters' faces}
{Key dialogue or exclamations should be largest and boldest}
{Narration lines should be smaller, in soft rounded serif/sans appropriate to the rendering language}
```

**Text overlay design rules (must follow):**
- All rendering text from the locked episode design must appear in the image
- No text mutation: `렌더링 텍스트`의 문구/철자/구두점은 그대로 사용
- Text must feel integrated into the illustration, not a plain overlay
- Text placement must respect reading order and avoid blocking key visuals
- Key dialogue / exclamations → largest, boldest, with glow
- Narration lines → smaller, soft rounded serif/sans appropriate to the rendering language, white with warm shadow
- Place text in areas that do not obscure characters' faces or key visual elements
- For 8세 이상 & 텍스트가 많은 페이지:
  - 본문(그림 속 이야기)을 별도 텍스트 영역(예: 연한 박스/패널)으로 분리해 가독성을 높이는 레이아웃을 권장
  - 대사/핵심 표현은 패널 안에서 강조 규칙을 우선 적용
- The final image should look like a finished picture book page where illustration and text work together

---

## Phase 2: Visual Consistency Review

1. Compare generated page images within each episode.
2. If **visual inconsistencies** (palette/style/identity) appear:
   - regenerate only the affected image(s)
   - do not change locked story/text
   - max 2 retries per image
3. If inconsistencies indicate a **Design gap** (wrong state/place/unclear guide):
   - rollback to Stage ②
   - re-evaluate and re-generate affected episode page images

---

## Approval Gate G4 — Final Result

Confirm with user:
1. All reference images — quality acceptable?
2. All page images — consistent, appealing, and faithful to locked guides?
3. Final result in `output/{book-slug}-final/` — ready to deliver?

**Do not deliver until G4 is approved.**

