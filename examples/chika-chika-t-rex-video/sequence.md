# Sequence: 치카치카 티라노

## Overall story / message
양치질하기 싫어하는 꼬마가 칫솔을 들자 티라노사우루스로 변신해 치카치카 왕국으로 모험을 떠난다. 충균 몬스터에게 더럽혀진 이빨 성을 구하기 위해 칫솔 검으로 위아래, 앞뒤, 안쪽까지 양치 액션을 펼치고, 마지막 물 헹굼으로 몬스터를 물리친다. 현실로 돌아와 거울 속 반짝이는 이빨을 확인하며 양치질의 즐거움을 깨닫는다.

영상은 **엄마(내레이터)** 가 이야기를 이끌고, **주인공(꼬마/티라노)** 과 **충균 몬스터**의 대사가 번갈아 나오며, BGM과 효과음이 분위기를 살리는 구조. 각 클립은 키 이미지(참조 이미지 기반 생성)를 배경으로 하고, 액션 클립은 모션/효과로 움직임을 표현한다.

## Craft Notes
- Genre rhythm: 그림책 스토리텔링 → 도입(평화) → 갈등 등장 → 액션 상승 → 클라이맥스 → 해결 → 여운
- Core message: 양치질은 재미있는 모험이다! (교훈을 직접 말하지 않고 액션으로 보여줌)
- Climax clip: Clip 09 (몬스터가 가장 작아지는 순간 + 승리)
- 목소리 구성:
  - **엄마 (내레이터)**: 이야기 서술, 액션 묘사 — 차분하고 따뜻한 톤
  - **주인공**: 꼬마의 투정, 티라노의 당당함 — 밝고 활기찬 톤
  - **충균 몬스터**: 건방지다가 점점 당황 — 코믹 악당 톤

---

## Clips

---

### Clip 00 — Title

#### Clip message
영상의 시작. 치카치카 왕국을 배경으로 타이틀을 보여준다. 모험의 시작을 예고하는 기대감과 호기심 유발.

#### Direction guide
- Visual: **still image** — 치카치카 왕국 배경(칫솔숲/이빨성 전망) 위에 타이틀 텍스트 중첩. 왕국의 밝고 반짝이는 분위기. 티라노의 실루엣이나 당당한 포즈가 보이면 좋음.
- References: `locations/치카치카-왕국` (position: 칫솔숲-입구, view: 성-전망, state: base)
- Clip-specific direction: 타이틀 텍스트 "치카치카 티라노"가 중앙에 페이드인. 부드러운 반짝임 효과.
- On-screen text: **치카치카 티라노** (중앙, 큰 글씨)

#### Required tracks
- image: **yes** — 타이틀 키 이미지
- video: no (still)
- voice (speech): **yes** — 엄마: "치카치카 티라노"
- bgm/sfx: **yes** — BGM 인트로 (경쾌하고 신나는 분위기)

#### Hook to next clip
"이 티라노는 누구일까? 어떤 모험이 펼쳐질까?"

---

### Clip 01 — 양치하기 싫어!

#### Clip message
현실 화장실. 꼬마가 칫솔을 들고 심드렁한 표정. 양치질을 싫어하는 투정. 하지만 갑자기 주변이 반짝이기 시작하며 변신을 암시하는 긴장감.

#### Direction guide
- Visual: **still image** — 현실-화장실, 세면대 앞. 꼬마(인간=base)가 칫솔을 손에 들고 심드렁/울상. 뒤로 거울에 비친 모습. 주변에 은은한 반짝임 효과가 퍼지기 시작.
- References: `locations/현실-화장실` (세면대-앞/거울-뷰, state: base), `characters/티라노` (state: base)
- Clip-specific direction: 꼬마가 칫솔을 들고 있는 모습. 처음에는 평범한 분위기에서 점점 반짝임이 더해짐.
- On-screen text: 없음

#### Required tracks
- image: **yes** — 클립 키 이미지
- video: no (still + 효과 애니메이션)
- voice (speech): **yes**
  - 주인공: "앙! 양치하기 싫어!"
  - 엄마: "그런데 갑자기 —"
- bgm/sfx: **yes** — BGM (잔잔→긴장 상승) + SFX (반짝! 효과음 — 클립 끝에서)

#### Hook to next clip
"반짝임이 점점 커진다! 무슨 일이 일어날까?"

---

### Clip 02 — 치카치카 티라노!

#### Clip message
변신! 꼬마가 초록색 티라노사우루스로 변신한다. 놀라면서도 신난 표정. 치카치카 왕국이 배경으로 펼쳐지며 모험의 시작을 알린다.

#### Direction guide
- Visual: **still image** — 반짝이는 폭죽 같은 효과 가운데, 티라노(t-rex)가 당당히 서 있는 모습. 배경은 치카치카 왕국(칫솔숲/이빨성). 역동적인 포즈, 칫솔 검은 아직 없고 두 팔을 벌린 포효하는 모습.
- References: `locations/치카치카-왕국` (칫솔숲-입구/성-전망, state: base), `characters/티라노` (state: t-rex)
- Clip-specific direction: 변신 순간의 반짝임 효과. 티라노의 역동적인 포효 포즈.
- On-screen text: 없음

#### Required tracks
- image: **yes** — 클립 키 이미지
- video: no (still + 효과)
- voice (speech): **yes** — 주인공(티라노): "크아아아! 치카치카 티라노다!"
- bgm/sfx: **yes** — BGM (신나고 경쾌하게 전환) + SFX (변신/포효 효과)

#### Hook to next clip
"티라노가 된 주인공. 도착한 곳은 어디일까?"

---

### Clip 03 — 더러워진 왕국

#### Clip message
치카치카 왕국에 도착한 티라노. 그런데 평소 반짝이는 하얀 성이 보라색과 검은색으로 더럽혀져 있다. 티라노가 당황하고 놀란다. 위기가 감지된다.

#### Direction guide
- Visual: **still image** — 치카치카-왕국, 칫솔숲 입구에서 바라본 성 전망. 성이 보라색/검은색으로 더럽혀짐(오염). 티라노(t-rex)가 왼쪽 아래에서 성을 올려다보며 깜짝 놀란 표정.
- References: `locations/치카치카-왕국` (칫솔숲-입구/성-전망, state: 오염), `characters/티라노` (state: t-rex, 놀란 표정)
- Clip-specific direction: 밝았던 하늘이 부분적으로 흐려짐. 오염된 성의 모습을 강조.
- On-screen text: 없음

#### Required tracks
- image: **yes** — 클립 키 이미지
- video: no (still)
- voice (speech): **yes** — 주인공: "어?! 누가 우리 치카치카 왕국을 이렇게 만들었어?"
- bgm/sfx: **yes** — BGM (긴장/위기감) + SFX (불길한 효과)

#### Hook to next clip
"대체 누가 왕국을 이렇게 만든 걸까?"

---

### Clip 04 — 충균 몬스터 등장

#### Clip message
충균 몬스터가 성 위에서 나타난다. 자신만만하게 깔깔대며 왕국을 더럽힌 것이 자신이라고 자랑한다. 악당의 정체가 드러나는 순간.

#### Direction guide
- Visual: **still image** — 이빨성 위에 올라선 충균 몬스터(big). 활짝 웃으며 양팔 벌린 건방진 포즈. 보라색 몬스터가 성 위에서 당당하게 서 있음. 티라노는 아래에서 올려다봄.
- References: `locations/치카치카-왕국` (이빨성-앞마당/성문-뷰, state: 오염), `characters/충균-몬스터` (state: big)
- Clip-specific direction: 몬스터의 우쭐한 표정과 큰 체구 강조.
- On-screen text: 없음

#### Required tracks
- image: **yes** — 클립 키 이미지
- video: no (still)
- voice (speech): **yes** — 충균 몬스터: "히히히! 내가 다 더럽혀 주겠어!"
- bgm/sfx: **yes** — BGM (악당 테마/코믹 위협) + SFX (몬스터 웃음 효과)

#### Hook to next clip
"티라노가 맞서기로 결심한다! 어떻게 몬스터를 막을까?"

---

### Clip 05 — 칫솔 검, 준비 완료!

#### Clip message
티라노가 칫솔 검을 꺼내 든다. 빛나는 칫솔 검이 등장하며 결의에 찬 표정. 승리를 확신하는 자신감. 액션의 시작을 알린다.

#### Direction guide
- Visual: **still image** — 티라노(t-rex)가 칫솔 검을 오른손(짧은 팔)에 높이 치켜들고 있는 당당한 포즈. 칫솔 검이 반짝임. 배경은 이빨성 앞마당(오염).
- References: `locations/치카치카-왕국` (이빨성-앞마당/성문-뷰, state: 오염), `characters/티라노` (state: t-rex, 칫솔 검 든 포즈)
- Clip-specific direction: 칫솔 검에 반짝이는 하이라이트 효과. 티라노의 자신만만한 표정.
- On-screen text: 없음

#### Required tracks
- image: **yes** — 클립 키 이미지
- video: no (still)
- voice (speech): **yes** — 주인공: "칫솔 검, 준비 완료!"
- bgm/sfx: **yes** — BGM (액션/전투 모드로 전환) + SFX (칼집에서 빼는 듯한 효과)

#### Hook to next clip
"티라노의 첫 번째 공격! 과연 몬스터를 이길 수 있을까?"

---

### Clip 06 — 위아래로 치카치카!

#### Clip message
티라노의 첫 공격! 칫솔 검을 위아래로 휘두르며 몬스터를 향해 돌진한다. 칫솔 검에서 하얀 치약 거품 광선이 나와 몬스터를 때린다. 몬스터가 아파하며 움찔한다. 액션의 시작!

#### Direction guide
- Visual: **still image or motion** — 티라노(t-rex)가 칫솔 검을 위에서 아래로 휘두르는 역동적인 포즈. 칫솔 검에서 치약 거품 광선이 몬스터(big)를 향해 발사됨. 몬스터가 맞고 움찔하는 표정. 속도감 있는 구도.
- References: `locations/치카치카-왕국` (이빨성-앞마당/성문-뷰, state: 오염), `characters/티라노` (state: t-rex, 액션 포즈), `characters/충균-몬스터` (state: big)
- Clip-specific direction: 액션 폭발! 칫솔 검의 움직임 궤적을 선으로 표현. 속도감과 역동감.
- On-screen text: "위아래로 — 치카치카!" (텍스트 애니메이션)

#### Required tracks
- image: **yes** — 클립 키 이미지
- video: **optional** — 액션 모션 클립 고려 가능 (<10s)
- voice (speech): **yes**
  - 엄마: "위아래로 — 치카치카!"
  - 충균 몬스터: "아야야! 뭐야 그거!"
- bgm/sfx: **yes** — BGM (액션 박진감) + SFX (칫솔 검 휘두름/광선/몬스터 피격)

#### Hook to next clip
"위아래로 닦으니까 몬스터가 아파한다! 다음 공격은?"

---

### Clip 07 — 앞뒤로도 치카치카!

#### Clip message
두 번째 공격! 칫솔 검을 앞뒤로 휘두르며 몬스터를 감싸 문지른다. 몬스터가 점점 작아지기 시작한다. 액션이 먹히고 있다!

#### Direction guide
- Visual: **still image or motion** — 티라노(t-rex)가 칫솔 검을 앞뒤로 움직이는 동작. 칫솔 검의 좌우 궤적이 흔들림선으로 표현. 충균 몬스터가 이전보다 약간 작아짐(small로 변하는 중간). 움츠린 표정.
- References: `locations/치카치카-왕국` (이빨성-앞마당/성문-뷰, state: 오염→회복 중), `characters/티라노` (state: t-rex), `characters/충균-몬스터` (state: small)
- Clip-specific direction: 몬스터의 크기 변화(줄어듦)를 시각적으로 강조. 칫솔 검의 좌우 궤적 표현.
- On-screen text: "앞뒤로도 — 치카치카!" (텍스트 애니메이션)

#### Required tracks
- image: **yes** — 클립 키 이미지
- video: **optional** — 액션 모션 클립 고려
- voice (speech): **yes**
  - 엄마: "앞뒤로도 — 치카치카!"
  - 충균 몬스터: "히이이! 멈춰!"
- bgm/sfx: **yes** — BGM (액션 지속) + SFX (칫솔 검 휘두름/몬스터 축소)

#### Hook to next clip
"몬스터가 작아졌지만 아직 끝나지 않았다! 마지막 기술!"

---

### Clip 08 — 안쪽까지 샥샥!

#### Clip message
세 번째 공격! 칫솔 검을 섬세하게 움직이며 몬스터 구석구석을 닦아낸다. 몬스터가 간지러워하며 웃고 발버둥친다. 몬스터가 더 작아졌다!

#### Direction guide
- Visual: **still image or motion** — 티라노(t-rex)가 집중하는 표정으로 칫솔 검을 섬세하게 조종. 칫솔 검에서 나온 작은 광선들이 몬스터 주변을 감쌈. 충균 몬스터(small)가 간지러워하며 웃고 발버둥치는 모습. 몬스터가 티라노 무릎 높이 정도.
- References: `locations/치카치카-왕국` (이빨성-앞마당/성문-뷰, state: 회복 중), `characters/티라노` (state: t-rex, 집중), `characters/충균-몬스터` (state: small)
- Clip-specific direction: 섬세한 움직임 강조. 몬스터의 코믹한 표정(간지러움)이 포인트.
- On-screen text: "안쪽까지 — 샥샥!" (텍스트 애니메이션)

#### Required tracks
- image: **yes** — 클립 키 이미지
- video: **optional** — 액션 모션 클립 고려
- voice (speech): **yes**
  - 엄마: "안쪽까지 — 샥샥!"
  - 충균 몬스터: "하하하! 간지러워!!!"
- bgm/sfx: **yes** — BGM (코믹 액션) + SFX (간지러움/섬세한 효과)

#### Hook to next clip
"몬스터가 많이 작아졌다! 드디어 끝이 보인다!"

---

### Clip 09 — 치카치카 티라노의 승리! (클라이맥스)

#### Clip message
클라이맥스. 꼼꼼히 닦은 덕분에 충균 몬스터가 티라노의 발목 높이만큼 조그마하게 줄어들었다. 티라노가 당당히 칫솔 검을 어깨에 메고 서 있다. 성의 더러움도 거의 다 사라졌다. 승리의 순간!

#### Direction guide
- Visual: **still image** — 티라노(t-rex)가 당당하고 자신만만한 포즈, 칫솔 검을 어깨에 메고 있음. 충균 몬스터(small)가 발목 높이에서 울먹이는 표정. 배경의 이빨 성이 거의 깨끗해짐(회복). 크기 대비가 극명하게 드러남.
- References: `locations/치카치카-왕국` (이빨성-앞마당/성문-뷰, state: 회복), `characters/티라노` (state: t-rex), `characters/충균-몬스터` (state: small)
- Clip-specific direction: 티라노는 크게, 몬스터는 작게 대비 강조. 성이 반짝이기 시작하는 효과.
- On-screen text: 없음

#### Required tracks
- image: **yes** — 클립 키 이미지
- video: no (still)
- voice (speech): **yes**
  - 충균 몬스터: "으으... 너무 닦여버렸잖아..."
  - 엄마: "치카치카 티라노의 승리!"
- bgm/sfx: **yes** — BGM (승리의 팡파르/웅장) + SFX (반짝임 효과)

#### Hook to next clip
"드디어 승리! 하지만 아직 마지막 한 가지가 남았다!"

---

### Clip 10 — 물로 입안을 헹궈라!

#### Clip message
마무리 동작. 티라노가 칫솔 검을 크게 휘둘러 치약 거품을 쓸어낸다. 물과 거품의 소용돌이가 몬스터를 휩쓸어 사라지게 한다. 성이 완전히 깨끗해지고 하늘이 다시 맑아진다. 충균 몬스터의 최후.

#### Direction guide
- Visual: **still image or motion** — 티라노(t-rex)가 칫솔 검을 크게 휘두르는 마무리 동작. 칫솔 검에서 물결/거품 소용돌이가 휘몰아침. 충균 몬스터(small)가 거품에 휩쓸려 허둥대는 모습. 배경의 이빨 성이 완전히 깨끗해지고 하늘도 맑아짐.
- References: `locations/치카치카-왕국` (이빨성-앞마당/성문-뷰, state: 회복), `characters/티라노` (state: t-rex), `characters/충균-몬스터` (state: small)
- Clip-specific direction: 물/거품이 휘몰아치는 역동적인 마무리 장면. 몬스터가 소용돌이에 빨려들어감.
- On-screen text: 없음

#### Required tracks
- image: **yes** — 클립 키 이미지
- video: **optional** — 모션 클립 고려
- voice (speech): **yes**
  - 주인공: "자, 마무리! 물로 입안을 헹궈라~!"
  - 충균 몬스터: "아아아악! 휩쓸려간다!"
- bgm/sfx: **yes** — BGM (클라이맥스 해소) + SFX (물/거품 소용돌이, 몬스터 소멸)

#### Hook to next clip
"몬스터가 사라졌다! 이제 치카치카 왕국은 어떻게 되었을까?"

---

### Clip 11 — 반짝이는 왕국

#### Clip message
치카치카 왕국이 다시 반짝반짝 빛난다. 이빨 성이 새하얗게 정화되었고, 하늘에는 무지개가 떴다. 티라노가 성 앞에서 미소 짓는다. 모두가 기뻐하는 축제 분위기. 모험의 성공적인 완료.

#### Direction guide
- Visual: **still image** — 반짝이는 이빨 성(완전 회복), 하늘에 무지개, 곳곳에 작은 반짝임(칫솔 요정들). 티라노(t-rex)가 행복한 미소를 지으며 칫솔 검을 하늘로 들어 올림. 전체적으로 밝고 따뜻한 축제 분위기.
- References: `locations/치카치카-왕국` (이빨성-앞마당/성문-뷰, state: 회복), `characters/티라노` (state: t-rex, 행복)
- Clip-specific direction: 전체적으로 밝고 따뜻한 분위기. 반짝임과 무지개 효과.
- On-screen text: 없음

#### Required tracks
- image: **yes** — 클립 키 이미지
- video: no (still)
- voice (speech): **yes**
  - 엄마: "반짝반짝 — 치카치카 왕국이 되살아났다!"
  - 환호 (여럿): "만세! 치카치카 티라노 만세!"
- bgm/sfx: **yes** — BGM (따뜻한 승리/축제) + SFX (반짝임, 환호)

#### Hook to next clip
"모험이 끝났다. 이제 현실로 돌아갈 시간. 그런데...?"

---

### Clip 12 — 내일도 치카치카 티라노!

#### Clip message
현실 화장실로 돌아온 꼬마. 거울 속 자신의 이빨이 반짝반짝 빛난다. 아이가 환하게 웃으며 다음 양치질을 기약한다. 양치질이 더 이상 하기 싫은 일이 아니라 즐거운 모험이 되었음을 암시. 따뜻하고 포근한 마무리.

#### Direction guide
- Visual: **still image** — 현실-화장실, 세면대 앞. 꼬마(인간=base)가 거울을 보며 환하게 웃고 있음. 거울 속 아이의 이빨이 살짝 반짝이는 효과. 포근한 실내등, 따뜻한 분위기.
- References: `locations/현실-화장실` (세면대-앞/거울-뷰, state: base), `characters/티라노` (state: base, 웃는 표정)
- Clip-specific direction: 거울 속 비치는 아이의 얼굴 강조. 이빨 반짝임 효과. 포근한 마무리.
- On-screen text: 없음

#### Required tracks
- image: **yes** — 클립 키 이미지
- video: no (still)
- voice (speech): **yes** — 주인공: "내일도 치카치카 티라노 할 거야!"
- bgm/sfx: **yes** — BGM (포근하고 따뜻한 엔딩) + SFX (미세한 반짝임)

#### Hook to next clip
**(마지막 클립)** 해소/잔상 — 아이의 웃는 얼굴에 반짝이는 이빨이 오래도록 기억에 남는다.
