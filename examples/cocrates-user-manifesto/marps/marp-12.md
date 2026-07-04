---
marp: true
theme: uncover
---

<!-- _class: tpl-02 -->
<!-- _paginate: true -->

<style>
section.tpl-02 {
  background: #0a0a0a; color: #ffffff; display: flex; flex-direction: column;
  align-items: center; text-align: center; padding: 6% 10%;
  font-family: 'Noto Sans KR', 'Pretendard', sans-serif; justify-content: flex-start;
}
section.tpl-02 h1 { font-size: 32pt; font-weight: 800; color: #ffffff; margin: 0 0 0.4em 0; text-align: center; line-height: 1.3; }
section.tpl-02 .tpl-02__title-divider { width: 40%; height: 1px; background: #333333; margin: 0 auto 1.2em auto; }
section.tpl-02 .tpl-02__body { flex: 1; display: flex; flex-direction: column; align-items: center; justify-content: center; width: 100%; max-width: 90%; gap: 0.3em; }
section.tpl-02 .tpl-02__footer { font-size: 11pt; font-weight: 300; color: #666666; text-align: center; padding-top: 1em; margin-top: auto; }
section.tpl-02 .eq-left { font-size: 22pt; font-weight: 700; color: #888888; }
section.tpl-02 .eq-op { font-size: 22pt; font-weight: 400; color: #666666; margin: 0 0.3em; }
section.tpl-02 .eq-right { font-size: 22pt; font-weight: 900; color: #ffffff; }
section.tpl-02 .desc { font-size: 16pt; font-weight: 400; color: #cccccc; text-align: center; line-height: 1.5; margin: 0.3em 0; max-width: 85%; }
section.tpl-02 .highlight { font-weight: 900; color: #ffffff; }
/* Slide 12 equation grid — 3행 2열 */
section.tpl-02 .s-12__eq {
  display: grid;
  grid-template-columns: auto auto;
  grid-template-rows: auto auto auto;
  align-items: center;
  justify-items: center;
  width: 100%;
  max-width: 85%;
  gap: 0.3em 1em;
}
section.tpl-02 .s-12__eq-harness { grid-column: 1; grid-row: 1; justify-self: start; }
section.tpl-02 .s-12__eq-plus { grid-column: 1; grid-row: 2; justify-self: center; font-size: 26pt; }
section.tpl-02 .s-12__eq-manifesto { grid-column: 1; grid-row: 3; justify-self: start; margin-left: 1.5em; }
section.tpl-02 .s-12__eq-result { grid-column: 2; grid-row: 2; justify-self: start; white-space: nowrap; }
section.tpl-02 .s-12__footer {
  margin-top: 0.8em;
  text-align: center;
}
section.tpl-02 .s-12__footer .highlight {
  display: block;
  font-size: 16pt;
  font-weight: 900;
  color: #ffffff;
  margin-bottom: 0.2em;
}
section.tpl-02 .s-12__footer .desc {
  font-size: 14pt;
  color: #888888;
  margin: 0;
  max-width: 100%;
}
/* 선명한 페이지 번호 (uncover text-shadow 블러 제거) */
section::after {
  text-shadow: none !important;
  background: none !important;
  color: rgba(255,255,255,0.45);
  font-size: 0.45em;
  width: auto;
  height: auto;
  padding: 12px 25px;
}
</style>

# 도구와 의지의 결합

<div class="tpl-02__title-divider"></div>

<div class="tpl-02__body">

<div class="s-12__eq">
  <span class="eq-left s-12__eq-harness">[ Cocrates Harness ]</span>
  <span class="eq-op s-12__eq-plus">+</span>
  <span class="eq-left s-12__eq-manifesto">[ User Manifesto ]</span>
  <span class="s-12__eq-result"><span class="eq-op">=</span> <span class="eq-right">[ AI-native Engineer ]</span></span>
</div>

<div class="s-12__footer">
  <span class="highlight">당신의 능력은 무한 확장된다.</span>
  <span class="desc">100명의 AI 팀원을 이끄는 팀장.</span>
</div>

</div>

<!--
자, 이제 모든 것을 종합해봅시다.

(잠시 멈춤)

**Cocrates Harness**는 여러분이 AI 팀원을 어떻게 이끌어야 하는지 설계와 검토를 강제합니다.
**사용자 선언 7계명**는 지적 주도권을 유지하고 최고의 아키텍트가 되기 위한 우리의 의지입니다.

이 둘이 하나가 될 때, 우리는 AI를 10명, 100명의 팀원으로 이끄는 진정한 AI-native Engineer가 되는 것입니다.

이제 마지막 슬라이드에서 전체 선언문을 함께 읽어보겠습니다."
-->
