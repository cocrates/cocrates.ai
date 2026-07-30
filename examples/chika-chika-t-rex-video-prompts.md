# Prompt Session Log: 치카치카 티라노 (영상)

> **Project:** 치카치카 티라노 — 그림책 읽어주는 애니메이션 영상
> **Date:** 2026-07-30
> **Skill:** video-authoring, todo, image-generation, video-generation
> **Total turns:** 20

---

## Session Summary

치카치카 티라노 그림책(10페이지, 이미지 완료)을 애니메이션 스토리텔링 영상으로 제작하는 세션. Cocrates는 video-authoring 스킬을 로드하고 5단계 워크플로(Define → Design → Evaluate → Generate Components → Assemble)를 따라 진행. 기존 그림책의 참조 이미지를 로컬로 복사하고, 13개 클립으로 스토리를 재구성한 후 클립 키 이미지 13장을 생성. 이후 전략을 전환하여 video-generation으로 각 클립을 대사/효과음/BGM이 포함된 단일 동영상으로 제작하기로 결정. 음성 일관성을 위해 voice-spec.md에 3개 목소리의 상세 사양을 정의하고 모든 video YAML에 동일한 영문 prompt 템플릿을 적용. transition-plan.md로 클립 간 전환을 설계하고 각 YAML에 반영. 13개 비디오 클립을 gemini-omni-flash-preview로 생성한 후, `cocrates-video-edit` MCP로 `assembly/001-main.yaml` edit-spec을 검증 및 렌더링하여 최종 `output/치카치카-티라노.mp4` (~112초, 68MB)를 완성.

---

## Prompt Log

### 1. 최초 요청 — 그림책 읽어주는 동영상

**User:** chika-chika-t-rex의 그림책을 읽어주는 동영상을 만들어줘. chika-chika-t-rex-video 폴더에

**Cocrates:**
- 사용자 요청을 video-authoring(동영상 제작)으로 식별하고 todo 스킬과 함께 로드
- workspace 구조 확인 — `chika-chika-t-rex/` 디렉토리 존재 확인 (그림책 프로젝트)
- 그림책의 에피소드(`001-치카치카-대모험.md` — 10페이지 분량), 참조 이미지(캐릭터 8장, 장소 10장), 페이지 이미지(10장)가 모두 완료된 상태임을 확인
- video-authoring 워크플로 파일 6개(01-define ~ 05-assemble + consistency)를 모두 로드하여 전체 프로세스 파악

---

### 2. 영상 스타일 구체화

**User:** 그림책을 단순히 읽어주는 것이 아니라, 나레이션 -> 페이지 애니메이션 -> 나레이션 -> ... 등으로 구성된 동영상이면 좋겠어.
---
적당한 BGM이나 효과음도 있으면 좋지. TTS는 한국어로. 엄마, 주인공, 하얀이, 설탕 괴물... 다른 목소리로 하는 것이 좋을 듯.

**Cocrates:**
- 영상의 형식(나레이션↔애니메이션 번갈아), BGM/효과음, 3개 목소리 TTS가 포함된 풍성한 영상으로 구체화
- 사용자가 언급한 "하얀이", "설탕 괴물"이 기존 그림책의 캐릭터(충균 몬스터)와 일치하는지 질문 — 명확하지 않은 용어는 추정하지 않고 확인

---

### 3. 캐릭터 목소리 정정

**User:** 하얀이/설탕 괴물이 아니라 충균-몬스터.

**Cocrates:**
- 캐릭터명 정정 수용
- 최종 목소리 구성 확정: **엄마**(내레이터) / **주인공**(꼬마/티라노 대사) / **충균 몬스터**(악당 대사)
- Stage ① Define의 첫 번째 질문: 예상 길이(1분 30초~2분 30초 제시) 및 나레이션↔페이지 애니메이션의 구체적 흐름(Page 1 예시: 꼬마 대사 → 엄마 대사 → 페이지 이미지 + 효과음)을 예시로 들어 사용자의 의도 확인

---

### 4. 비전 확인 — 페이지 이미지 미사용

**User:** 1. 그림책에 만들어진 페이지 이미지는 사용하지 않아.
2. character/location 참조 이미지는 그대로 사용할 것.
3. 페이지1의 경우, 주인공이 양치하기 싫어... 라고 말하고 울상을 짓는데, 갑자기 번쩍 효과가 나면서 티라노로 변하고 치카치카 왕국에서 나타남. 그림책의 스토리로 동영상을 다시 제작한다고 보면 됨.

**Cocrates:**
- 사용자의 비전을 정확히 파악: 그림책 페이지 이미지는 사용하지 않고, 참조 이미지를 기반으로 장면을 새로 구성하여 애니메이션 스타일로 재해석
- 사용할 에셋(캐릭터/장소 참조 이미지, 스토리)과 새로 만들 것(클립 키 이미지, TTS, BGM, SFX)을 구분하여 정리
- Page 1을 예시로 3개의 클립 분할案(현실 화장실→변신→왕국 등장) 제시

---

### 5. Define 확정 질문

**User:** 그렇지.

**Cocrates:**
- Stage ① Define의 마무리 질문 4가지:
  - 예상 길이 (~3~5분)
  - 분위기/톤 (그림책과 동일 — 유쾌하고 신나는 모험 액션, 코믹)
  - 타겟 시청자 (3~6세)
  - 금기사항 (치과 공포 금지, 무서운 악당 금지)
- `overview.md` 작성: Basic info, Intent/message summary, Reference need sketch, Style direction, Constraints, Validation Criteria 7개 항목

---

### 6. G1 승인

**User:** 오케이

**Cocrates:**
- G1 승인 → TODO.md 초기화 (T-001 in_progress 상태로 시작)
- `overview.md` 기반으로 `TODO.md` 작성 (T-001~T-005)
- 프로젝트 디렉토리 구조 생성: `evaluations/`, `images/references/`, `images/clips/`, `speech/`, `music/`, `assembly/`, `output/`

---

### 7. Stage ② Design — 참조 이미지 복사 요청

**User:** 참조 이미지의 yaml/png는 모두 chika-chika-t-rex/의 것을 복사해서 사용해줘. 원본 그림책과 같은 느낌이 되도록.

**Cocrates:**
- `chika-chika-t-rex/images/characters/`와 `chika-chika-t-rex/images/locations/`의 모든 YAML+PNG 파일을 `chika-chika-t-rex-video/images/references/`로 복사
- 캐릭터 4개(티라노 base/t-rex 4개, 충균 몬스터 big/small 4개) + 장소 5개(화장실/칫솔숲/이빨성/회복 오염 각 2개씩) = 총 18개 파일 복사 완료
- `references.md`의 경로를 `../chika-chika-t-rex/images/`에서 로컬 `images/references/`로 전면 수정
- `sequence.md` 작성: 13개 클립(Clip 00~12) 설계 — 각 Clip에 message, direction guide, required tracks, hook 포함

---

### 8. G2 승인

**User:** 오케이

**Cocrates:**
- G2 승인 → TODO.md 업데이트 (T-002 done, T-003 in_progress)
- Stage ③ Evaluate: `evaluations/001-main.md` 작성
- Validation Criteria 7개 항목 전부 ✅ 통과 확인
  - 4가지 양치 동작(위아래/앞뒤/안쪽/헹굼)이 Clip 06/07/08/10에 각각 할당됨
  - 3개 목소리(엄마/주인공/몬스터) 역할 명확히 구분
  - 금기사항 위반 없음
- **Design Lock (G3)** 승인 요청

---

### 9. G3 승인 — Generate 진행

**User:** 다음 진행해줘.

**Cocrates:**
- G3 승인 → TODO.md 업데이트 (T-003 done, T-004 in_progress)
- image-generation 스킬 로드 및 기존 페이지 YAML(`00.yaml`) 분석 — model: `gemini-3.1-flash-image`, size: `1K`, aspectRatio: `3:4` (세로) 확인
- 13개 클립 키 이미지 YAML 작성 (16:9 와이드 영상에 맞게 aspectRatio 변경)
- Batch 1(Clip 00~03) YAML 제시: Title / 양치하기 싫어! / 변신! / 더러워진 왕국

---

### 10. Clip 이미지 Batch 1 생성

**User:** 생성해줘

**Cocrates:**
- 4개 YAML을 병렬로 MCP generate 호출 (동기 모드)
- 전부 성공적으로 생성 완료:
  - `00-title.png` (955KB)
  - `01-dont-wanna-brush.png` (715KB)
  - `02-transformation.png` (871KB)
  - `03-dirty-kingdom.png` (911KB)
- 확인 요청 및 Batch 2(Clip 04~07) 제시

---

### 11. Batch 2 생성

**User:** 좋아

**Cocrates:**
- Batch 2(Clip 04~07) 제시: 몬스터 등장 / 칫솔 검 / 위아래 공격 / 앞뒤 공격 — 각 YAML 설명 및 승인 요청

---

### 12. Batch 2 생성

**User:** 생성해줘.

**Cocrates:**
- 4개 YAML 병렬 MCP generate 호출:
  - `04-monster-appears.png` (853KB)
  - `05-toothbrush-sword.png` (843KB)
  - `06-up-down-attack.png` (935KB)
  - `07-front-back-attack.png` (895KB)
- Batch 3(Clip 08~12) 마지막 5개 제시: 안쪽 샥샥 / 승리 / 헹굼 마무리 / 왕국 회복 / 현실 복귀

---

### 13. Batch 3 생성

**User:** 생성해줘.

**Cocrates:**
- 5개 YAML 병렬 MCP generate 호출:
  - `08-inside-scrub.png` (894KB)
  - `09-victory.png` (856KB)
  - `10-rinse-finale.png` (681KB)
  - `11-kingdom-restored.png` (744KB)
  - `12-back-to-reality.png` (638KB)
- **13개 클립 키 이미지 전부 생성 완료** ✅
- Phase 3 (TTS 음성 생성) 안내 — 3개 목소리 × 13개 클립의 대사 스크립트 제시

---

### 14. 전략 변경 — Video Generation으로 통합

**User:** TTS를 생성하지 말고, 참조 이미지를 사용해서 video를 생성해줘. video 생성할 때, 누가 무슨 말을 하는지, 효과음 등도 명확히 하면, 오디오를 따로 만들지 않아도 됨.

**Cocrates:**
- video-generation 스킬 로드 및 분석
- 전략 전환 결정: 별도 TTS + 이미지 조립 → 각 클립을 video-generation으로 **대사+효과음+BGM이 내장된 단일 mp4**로 생성
- 각 클립의 키 이미지를 `params.references`의 `type: image` (FIRST_FRAME)로 사용
- Clip 00 (Title)용 첫 video YAML 작성 — 음성(엄마 "치카치카 티라노"), BGM, 반짝임 효과 포함, 8초, 16:9
- Batch 1(Clip 00~03) video YAML 제시

---

### 15. Clip 01 대사 간소화

**User:** 01에서... 엄마 목소리 "그런데 갑자기 —" 가 필요할까? 삭제하면 10초 이내로 될 것 같은데

**Cocrates:**
- 사용자 의견 수용 — Clip 01에서 엄마 대사 "그런데 갑자기 —" 삭제
- Clip 01 video YAML 업데이트:
  - design: 오디오 항목에서 엄마 대사 제거, "클립 끝에서 반짝이는 효과음으로 다음 장면을 암시"로 변경
  - params.prompt: Sound 섹션에서 엄마 대사 제거, durationSeconds 12→10으로 단축
  - BGM 설명: "잔잔→긴장감 상승" 유지

---

### 16. 목소리 일관성 방안 논의

**User:** 계속되는 비디오 생성에서 목소리를 고정하려면 어떻게 해야 하지? 목소리에 대해 좀 더 명확히 정의해서 일관성이 유지되도록 해줘.

**Cocrates:**
- 3가지 접근법 제시:
  1. **상세 음성 사양을 모든 YAML에 동일하게 명시** — 각 캐릭터 목소리를 구체적으로 정의하여 모든 프롬프트에 포함
  2. 음성 참조 오디오 생성 → 비디오 생성 시 audio 참조로 사용
  3. 접근법 1 + 2 결합
- 접근법 2(음성 참조 오디오)를 권장하며 선택 요청

---

### 17. 접근법 1 선택

**User:** omni 모델이 audio는 지원하지 않더라고. 접근법1로 해줘.

**Cocrates:**
- `voice-spec.md` 작성 — 3개 목소리의 상세 사양 정의:
  - **엄마:** `A warm Korean female narrator voice, gentle motherly tone, speaks clearly in Korean`
  - **주인공:** `A young Korean boy's voice (around 5-6 years old), bright and playful natural child tone, speaks in Korean`
  - **충균 몬스터:** `A comical cartoon villain voice in Korean, slightly deep and raspy, over-the-top theatrical tone, speaks in Korean`
- 각 목소리의 감정 상태별 가이드라인 포함 (주인공: 투정/신남/놀람/자신감/기쁨, 몬스터: 건방짐/아픔/당황/간지러움/패배)
- 공통 오디오 규칙: BGM/SFX는 대사를 묻지 않는 볼륨으로
- 13개 video YAML 전면 재작성 — 모든 params.prompt에 동일한 음성 사양 문구를 캐릭터별로 일관되게 적용
- 영문 prompt 통합 템플릿 정의: `Sound: [화자 설명] says in Korean "[대사]" with [감정/톤 설명]. [효과음]. Background music: [BGM 설명].`

---

### 18. Transition Plan 수립

**User:** 각각 비디오 생성에서 다음 비디오 클립으로의 연결을 고려해줘.

**Cocrates:**
- 각 클립이 단독으로 존재하는 것이 아니라 연결되어 하나의 영상으로 흘러가야 함을 인식
- **`transition-plan.md`** 작성:
  - **BGM Arc**: 13개 클립을 관통하는 BGM 연속성 설계 (메인 테마 인트로 → 잔잔 → 모험 → 서스펜스 → 악당 → 액션(3클립 연속) → 승리 → 축제 → 엔딩 피아노 변주)
  - **Clip-to-Clip 전환 상세**: 00→01부터 11→12까지 각 전환의 끝/전환/시작/후크 정의
  - 후크 디자인: 각 클립 끝에서 시청자의 궁금증을 유발하여 다음 클립으로 자연스럽게 유도
- `transition-plan.md`의 전환 설계를 각 비디오 YAML의 `design` 필드와 `params.prompt`에 반영
- 모든 YAML에 다음 패턴 적용:
  - **`design`**: 클립 시작(이전 클립과의 연결점) + 클립 끝(다음 클립으로의 후크)을 명시
  - **`params.prompt`**: 시작 시 이전 클립의 마지막 상태를 언급, 끝에서 다음 클립을 암시하는 시각/청각 요소 포함
- 예시 — Clip 02 (변신):
  - 시작: "In a single continuous shot... The scene begins with a burst of sparkles transitioning from the previous clip's sparkling buildup."
  - 끝: "At the end: the sparkles settle, and the T-rex's expression shifts slightly to curiosity... leading to the next scene where he discovers the dirty castle."
- **BGM 연속성** 적용: Clip 06~08(액션 3연속)은 BGM이 끊기지 않고 지속됨을 명시
- **디자인 → 프롬프트 직접 매핑**: `design`의 전환 설명을 `params.prompt`에 자연어로 그대로 반영하여 omni 모델이 전환을 이해하고 구현할 수 있도록 함

---

### 19. 비디오 생성 (13개 클립)

**User:** 모두 생성해줘.

**Cocrates:**
- Batch 1 (Clip 00~03): Title / 양치하기 싫어! / 변신 / 더러워진 왕국 — 4개 생성 ✅
- Batch 2 (Clip 04~08): 몬스터 등장 / 칫솔 검 / 위아래 / 앞뒤 / 안쪽 — 5개 생성 ✅
- Batch 3 (Clip 09~12): 승리 / 헹굼 / 왕국 회복 / 현실 복귀 — 4개 생성 ✅
- 모든 클립에 대사(3개 목소리) + SFX + BGM이 내장됨 (gemini-omni-flash-preview)
- `videos/001-main/` 아래 13개 mp4 총 ~27MB

---

### 20. Stage ⑤ Assemble — 최종 mp4 생성

**User:** 최종 비디오 생성해줘.

**Cocrates:**
- video-authoring skill 재확인: Stage ⑤는 `cocrates-video-edit` MCP 사용 (`validate_spec` → `render_video`)
- `edit-spec.md` 스키마에 따라 `assembly/001-main.yaml` 작성:
  - 13개 클립을 순서대로 연결 (hard cut, transition 제거 — 모든 전환은 비디오 클립 내에 내장됨)
  - 각 클립은 전체 길이로 재생 (startFrom/endAt 생략)
- MCP `validate_spec` (inline spec) ✅ → `render_video` ✅
- 최종 출력: `output/치카치카-티라노.mp4`
  - 길이: ~112초 (약 1분 52초)
  - 해상도: 1280×720 (소스 클립 네이티브)
  - 용량: 68MB
- G5 gate 대기 중

---
