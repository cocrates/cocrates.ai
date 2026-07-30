# Voice Specification — 치카치카 티라노

이 문서는 모든 비디오 클립에서 목소리 일관성을 유지하기 위한 음성 사양을 정의한다.
각 video YAML의 `params.prompt`에 동일한 음성 설명을 포함하여 생성 일관성을 확보한다.

---

## 1. 엄마 (Mom / Narrator)

| 항목 | 사양 |
|------|------|
| 언어 | 한국어 (Korean) |
| 화자 | 성인 여성, 엄마 목소리 |
| 톤 | 따뜻하고 부드러운 (warm and gentle) |
| 전달 방식 | 이야기를 읽어주듯 차분하면서도 약간의 생동감 있게 |
| 감정 범위 | 잔잔한 내레이션 → 신나는 액션 묘사 → 따뜻한 마무리 |
| 음높이 | 중간, 부드러움 |

**영문 prompt 템플릿:**
```
A warm Korean female narrator voice, gentle motherly tone,
speaks clearly in Korean.
```

---

## 2. 주인공 (Main Character — Kid / T-Rex)

| 항목 | 사양 |
|------|------|
| 언어 | 한국어 (Korean) |
| 화자 | 5~6세 남자아이 (young boy) |
| 기본 톤 | 밝고 활기찬 (bright and playful) |
| 변환 | 인간 상태와 티라노 상태 모두 같은 아이 목소리 유지 |
| 감정 상태 | • 투정/심술 (whiny, reluctant): 양치하기 싫을 때<br>• 신남/자신감 (excited, confident): 티라노로 변신 후<br>• 놀람/당황 (surprised): 더러워진 왕국을 발견했을 때<br>• 기쁨/다짐 (happy, determined): 마지막 장면 |

**영문 prompt 템플릿:**
```
A young Korean boy's voice (around 5-6 years old), bright and playful
natural child tone, speaks in Korean.
```

**상태별 감정 가이드:**

| Clip | 감정 | 설명 |
|------|------|------|
| 01 | 투정 (whiny) | "앙! 양치하기 싫어!" — 심드렁하고 짜증 섞인 톤 |
| 02 | 신남 (excited) | "크아아아! 치카치카 티라노다!" — 당차고 자신감 넘치는 톤 |
| 03 | 놀람 (surprised) | "어?! 누가..." — 당황하고 놀란 톤 |
| 05 | 자신감 (confident) | "칫솔 검, 준비 완료!" — 결의에 찬 톤 |
| 10 | 명령 (commanding) | "자, 마무리! 물로 입안을 헹궈라~!" — 당당하고 경쾌한 톤 |
| 12 | 기쁨/다짐 (happy) | "내일도 치카치카 티라노 할 거야!" — 환하고 즐거운 톤 |

---

## 3. 충균 몬스터 (Cavity Monster)

| 항목 | 사양 |
|------|------|
| 언어 | 한국어 (Korean) |
| 화자 | 코믹 악당 목소리 (comical cartoon villain) |
| 기본 톤 | 약간 저음, 건방지고 과장된 (slightly deep, smug and exaggerated) |
| 전달 방식 | 오버액팅, 연극적인 톤 (over-the-top, theatrical) |
| 감정 상태 | • 건방짐 (smug): 첫 등장 시 "히히히!"<br>• 아픔/당황 (pain/panic): 공격 맞을 때 "아야야!" "히이이! 멈춰!"<br>• 간지러움/웃음 (ticklish/laughing): "하하하! 간지러워!!!"<br>• 울먹/패배 (whimpering/defeated): 클라이맥스 "으으..." |

**영문 prompt 템플릿:**
```
A comical cartoon villain voice in Korean, slightly deep and raspy,
over-the-top theatrical tone, speaks in Korean.
```

---

## 4. 공통 오디오 규칙

- **BGM:** 상황에 맞는 배경음악 (잔잔→긴장→액션→승리→포근)
- **SFX:** 효과음은 대사와 겹치지 않도록 간격 조절
- **대사 우선:** BGM/SFX는 대사를 묻지 않는 볼륨으로

---

## 5. 영문 Prompt 통합 템플릿

각 비디오 YAML의 prompt에서 다음 형식으로 음성을 지정한다:

```
Sound: [화자 설명] says in Korean "[대사]" with [감정/톤 설명].
[추가 효과음 설명]. Background music: [BGM 설명].
```
