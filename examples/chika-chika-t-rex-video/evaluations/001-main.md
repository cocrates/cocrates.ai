# Segment 001 evaluation — 메인 스토리 (치카치카 티라노)

## 1. Criteria Check

| Criterion | Result | Evidence |
|-----------|--------|----------|
| 타겟 연령(3~6세)이 영상을 따라갈 수 있다 | ✅ | 각 클립 6~12초, 단순한 스토리 라인, 그림책 원작 기반. 시각적 스토리텔링 + 짧은 대사로 연령 적합. |
| 아이가 양치질을 긍정적으로 느끼게 한다 | ✅ | Clip 12(현실 복귀)에서 아이가 스스로 "내일도 치카치카 티라노 할 거야!"라고 말하며 양치질을 즐거운 모험으로 프레이밍. 교훈을 직접 말하지 않음. |
| 양치질의 올바른 방법이 자연스럽게 녹아 있다 | ✅ | Clip 06(위아래), Clip 07(앞뒤), Clip 08(안쪽), Clip 10(헹굼) — 4가지 동작이 액션 클립에 각각 할당되어 자연스럽게 전달됨. |
| 그림책 원작 스토리와 캐릭터를 충실히 따름 | ✅ | 원작 `episodes/001-치카치카-대모험.md`의 10페이지 스토리 비트, 대사, 캐릭터 성격을 13개 클립으로 충실히 재구성. 주요 대사 원문 유지. |
| 3개 목소리(엄마/주인공/몬스터)가 명확히 구분된다 | ✅ | 엄마(내레이션/액션 묘사), 주인공(투정/당당한 대사/다짐), 충균 몬스터(건방짐→당황→울먹) — 역할과 톤이 뚜렷이 구분됨. TTS에서 다른 목소리 모델 사용 예정. |
| BGM/SFX가 이야기 흐름을 돕고 방해하지 않음 | ✅ | BGM은 장면 분위기에 맞게 변화(잔잔→긴장→액션→승리→포근). SFX는 변신/액션/마무리에만 사용되어 대사를 묻히지 않음. |
| 금기사항 위반 없음 | ✅ | 치과 공포 유발 요소 없음. 충균 몬스터는 코믹 악당(무섭지 않음). 이빨 빠짐/치료 장면 없음. |

## 2. Craft Checks

| Check | Result | Evidence |
|-------|--------|----------|
| All required clip fields present | ✅ | 모든 클립에 message, direction guide, required tracks, hook 있음. |
| Clip message alone makes the beat imaginable | ✅ | 각 클립 메시지만 읽어도 해당 비트의 스토리/감정이 명확히 전달됨. |
| Track choice ↔ message / duration budget | ✅ | 액션 클립(06~08, 10)에만 선택적 motion video. 나머지는 still image + TTS로 TTS-led clip 예상 길이(~30s) 이내. |
| Hooks / arc role | ✅ | 모든 클립에 hook 존재. 도입→상승→클라이맥스→해소→여운의 명확한 아크. |
| Short: sequence.md includes segment-level content | ✅ | sequence.md에 13개 클립 전체 설계 포함. |

## 3. Reference Integrity

| Check | Result | Evidence |
|-------|--------|----------|
| Matches references.md plan | ✅ | characters(티라노/충균-몬스터), locations(현실-화장실/치카치카-왕국)만 사용. references.md 선언과 일치. |
| No undeclared entities | ✅ | 참조 이미지에 없는 엔터티 사용하지 않음. 모든 클립이 선언된 캐릭터/장소만 참조. |
| state vs clip direction separation | ✅ | state(base/t-rex, big/small)는 물리적 정체성 유지. 표정/포즈/구도는 clip direction에서 별도 지시. |

## 4. Audience / Craft Notes

- **Viewer (3~6세):** 짧은 클립 길이, 단순한 문장, 반복되는 액션 패턴(위아래→앞뒤→안쪽)이 인지 부담을 줄임. 의성어/효과음이 집중력을 유지시킴.
- **Stakeholder (부모/교육자):** 양치질을 긍정적 경험으로 전환. 교훈을 직접 주입하지 않고 모험으로 자연스럽게 연결.
- **Rhythm / audio:** 엄마(내레이션) → 주인공(대사) → 몬스터(대사) → 엄마 → ... 패턴으로 목소리가 번갈아 나와 단조롭지 않음. 액션 클립에서 BGM/SFX가 박진감을 더함.

## 5. Revisions (Design-First)

| # | Finding | Severity | Design file | Fix | Status |
|---|---------|----------|-------------|-----|--------|
| — | 발견된 이슈 없음 | — | — | — | — |

## 6. Design Lock Readiness (G3)

- [x] Criteria / Craft / Reference checks acceptable
- [x] Clip messages and direction frozen for Stage ④–⑤
- [x] 4가지 양치 동작 모두 커버됨
- [x] 금기사항 위반 없음
- [x] 참조 이미지 로컬 복사 완료
