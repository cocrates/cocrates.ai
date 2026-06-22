# Cocrates Icon

소크라테스의 철학적 유산과 Cocrates Harness의 핵심 가치(**질문, 구조적 통제, 인간 지성의 주도권**)를 형상화한 SVG 아이콘입니다.

소크라테스를 상징하는 그리스식 기둥(Architecture/Structure)의 상단이 생각을 일깨우는 말풍선(Dialogue/Question)과 **지혜의 불꽃(Ignorance to Light)** 형태로 피어오르며, 이를 단단한 외곽 프레임(Harness)이 감싸고 제어하는 모습을 미니멀하고 현대적인 라인 아트 스타일로 설계했습니다.

```xml
<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 256 256" width="100%" height="100%">
    <!-- 배경 정의 (다크 모드 기준의 깊이감 있는 그라데이션) -->
    <defs>
        <linearGradient id="harnessGradient" x1="0%" y1="0%" x2="100%" y2="100%">
            <stop offset="0%" stop-color="#4F46E5" /> <!-- Royal Blue/Indigo -->
            <stop offset="100%" stop-color="#06B6D4" /> <!-- Cyan -->
        </linearGradient>
        <linearGradient id="innerGradient" x1="0%" y1="100%" x2="0%" y2="0%">
            <stop offset="0%" stop-color="#312E81" stop-opacity="0.2" />
            <stop offset="100%" stop-color="#06B6D4" stop-opacity="0.05" />
        </linearGradient>
    </defs>

    <!-- 1. External Harness (외곽 제어 프레임: 육각형의 단단한 통제 구조) -->
    <polygon points="128,16 224,72 224,184 128,240 32,184 32,72" 
             fill="url(#innerGradient)" 
             stroke="url(#harnessGradient)" 
             stroke-width="8" 
             stroke-linejoin="round" />

    <!-- 2. Socratic Pillar & Dialogue (소크라테스 기둥과 구조적 대화의 융합) -->
    <!-- 기둥 받침대 (기반, 원칙) -->
    <path d="M 80,190 L 176,190" stroke="#FFFFFF" stroke-width="8" stroke-linecap="round" />
    <path d="M 72,204 L 184,204" stroke="url(#harnessGradient)" stroke-width="8" stroke-linecap="round" />
    
    <!-- 기둥 줄기 (세 줄의 엄격한 기둥선 / 3대 철학을 상징) -->
    <path d="M 104,120 L 104,190" stroke="#FFFFFF" stroke-width="6" stroke-linecap="round" />
    <path d="M 128,132 L 128,190" stroke="#FFFFFF" stroke-width="6" stroke-linecap="round" />
    <path d="M 152,120 L 152,190" stroke="#FFFFFF" stroke-width="6" stroke-linecap="round" />

    <!-- 말풍선 및 불꽃 형태로 확장되는 기둥 상단 머리 (질문과 메타인지) -->
    <!-- 좌측 기둥 머리 곡선 -->
    <path d="M 96,120 C 96,105 76,105 76,84 C 76,56 100,44 128,44" 
          fill="none" 
          stroke="url(#harnessGradient)" 
          stroke-width="8" 
          stroke-linecap="round" />
          
    <!-- 우측 기둥 머리 곡선 (말풍선의 꼬리 꼬임 형상 결합) -->
    <path d="M 160,120 C 160,105 180,105 180,84 C 180,56 156,44 128,44" 
          fill="none" 
          stroke="url(#harnessGradient)" 
          stroke-width="8" 
          stroke-linecap="round" />

    <!-- 3. Ignorance to Light (무지의 자각을 상징하는 중앙의 빛/스파크 코어) -->
    <circle cx="128" cy="84" r="10" fill="#FFFFFF" />
    <path d="M 128,62 L 128,70" stroke="#06B6D4" stroke-width="4" stroke-linecap="round" />
    <path d="M 128,98 L 128,106" stroke="#06B6D4" stroke-width="4" stroke-linecap="round" />
    <path d="M 106,84 L 114,84" stroke="#06B6D4" stroke-width="4" stroke-linecap="round" />
    <path d="M 142,84 L 150,84" stroke="#06B6D4" stroke-width="4" stroke-linecap="round" />
</svg>

```

## 💡 디자인 디테일 설명

* **육각형 프레임 (Harness):** 무분별한 AI 생성을 가두고 사용자의 승인 하에 점진 구체화하는 Cocrates Harness의 **엄격한 게이트(Gate)와 제어 시스템**을 의미합니다.
* **그리스 신전 기둥 (Structure):** 아무리 복잡한 컨텍스트라도 8대 보편 범주(ASR)와 아키텍처(Spec)를 통해 단단하게 지탱하는 구조적 접근법을 형상화했습니다. 기둥의 세로선은 Cocrates의 3대 철학을 상징합니다.
* **상단의 곡선 (Dialogue):** 굳어있는 기둥 상단이 유연한 곡선의 말풍선 형태로 이어지며, 정답을 주지 않고 끊임없이 질문을 던지는 **소크라테스식 대화법**을 나타냅니다.
* **중앙의 불꽃/코어 (Light):** 자신의 무지를 깨닫는 순간(**"I know I know nothing"**) 번쩍이는 지성의 불꽃이자 메타인지의 각성을 시각화했습니다.