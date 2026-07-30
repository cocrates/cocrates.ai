# Reference model plan

## Intent summary
치카치카 티라노 그림책 스토리를 3~5분 애니메이션 스토리텔링 영상으로 재해석. 캐릭터(주인공/몬스터)와 장소(화장실/왕국)의 시각적 일관성이 필수적.

## Reference needed?
- Needed: **yes**
- Why: 캐릭터 외형(눈/색상 연결고리)과 장소(화장실/왕국)가 클립 전체에서 일관되어야 함. 그림책 원작의 비주얼 아이덴티티 유지.

## Catalog kinds

| Kind | Consistency goal | Entity count |
|------|------------------|-------------|
| `characters` | 캐릭터 외형/색상/비율 일관성 유지 | 2 (티라노, 충균-몬스터) |
| `locations` | 배경 장소의 스타일/색감/구성 일관성 유지 | 2 (현실-화장실, 치카치카-왕국) |

### Characters

#### 티라노 (Tyranno)
| Field | Value |
|-------|-------|
| States | `base` (인간 꼬마), `t-rex` (티라노사우루스) |
| Visual anchors | ①커다란 검은 눈+긴 속눈썹(공통) ②검정 앞머리↔이마 깃 ③노란 잠옷↔배/꼬리 노란 무늬 |
| Reference image | `images/references/characters/티라노.png` (base) |
| Reference image | `images/references/characters/티라노-t-rex.png` (t-rex) |

#### 충균-몬스터 (Cavity Monster)
| Field | Value |
|-------|-------|
| States | `big` (초기 등장, 큰 상태), `small` (작아진 상태) |
| Visual anchors | 둥글고 통통한 보라색 몸체, 큰 외눈박이, 짧은 통통한 손발 |
| Reference image | `images/references/characters/충균-몬스터-big.png` (big) |
| Reference image | `images/references/characters/충균-몬스터-small.png` (small) |

### Locations

#### 현실-화장실
| Field | Value |
|-------|-------|
| Views | `세면대-앞/거울-뷰` |
| State | `base` |
| Reference image | `images/references/locations/현실-화장실-세면대-앞-거울-뷰.png` |

#### 치카치카-왕국
| Field | Value |
|-------|-------|
| Positions | `칫솔숲-입구`, `이빨성-앞마당` |
| Views | `성-전망`, `성문-뷰` |
| States | `base`, `오염`, `회복` |
| Reference images | `images/references/locations/치카치카-왕국-칫솔숲-입구-성-전망-오염.png` |
| | `images/references/locations/치카치카-왕국-이빨성-앞마당-성문-뷰-오염.png` |
| | `images/references/locations/치카치카-왕국-이빨성-앞마당-성문-뷰-회복.png` |

## Stage ④ impact
- Phase 0 reference images: **done** — `images/references/` 에 로컬 복사 완료
- Phase 1 clip key images: **required** — 각 클립의 키 이미지 생성 필요 (참조 이미지 기반)
- Phase 2 motion video: **선택적** — 액션 클립에 한해 모션 적용 고려
- Phase 3 speech: **required** — 모든 클립에 TTS 필요 (엄마/주인공/몬스터)
- Phase 4 bgm/sfx: **required** — BGM + 효과음
