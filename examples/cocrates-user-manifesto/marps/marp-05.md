---
marp: true
theme: uncover
---

<!-- _class: tpl-02 -->
<!-- _paginate: true -->

<style>
/* === from marps/template-02.md === */
section.tpl-02 {
  background: #0a0a0a; color: #ffffff; display: flex; flex-direction: column;
  align-items: center; text-align: center; padding: 6% 10%;
  font-family: 'Noto Sans KR', 'Pretendard', sans-serif; justify-content: flex-start;
}
section.tpl-02 h1 { font-size: 32pt; font-weight: 800; color: #ffffff; margin: 0 0 0.4em 0; text-align: center; line-height: 1.3; }
section.tpl-02 .tpl-02__title-divider { width: 40%; height: 1px; background: #333333; margin: 0 auto 1.2em auto; }
section.tpl-02 .tpl-02__body { flex: 1; display: flex; flex-direction: column; align-items: center; justify-content: center; width: 100%; max-width: 90%; gap: 0.3em; }
section.tpl-02 .cmd { font-size: 42pt; font-weight: 900; color: #ffffff; text-align: center; line-height: 1.2; margin: 0.2em 0; }
section.tpl-02 .cmd-ko { font-size: 24pt; font-weight: 500; color: #999999; text-align: center; margin: 0.1em 0; }
section.tpl-02 .desc { font-size: 16pt; font-weight: 400; color: #cccccc; text-align: center; line-height: 1.5; margin: 0.3em 0; max-width: 85%; }
section.tpl-02 .point { font-size: 15pt; font-weight: 400; color: #888888; text-align: center; margin: 0.5em 0; }
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

# 제1계명

<div class="tpl-02__title-divider"></div>

<div class="tpl-02__body">

<div class="cmd">1. "I do not blindly request"</div>

<div class="cmd-ko">설계 없는 요구를 거부한다.</div>

<div class="desc">설계 없는 요청은 인지적 나태함의 산물일 뿐이다.</div>

<div class="point">좋은 설계와 좋은 질문만이 가치 있는 답변을 만든다.</div>

</div>

<!--
**"I do not blindly request."**

설계 없는 요구를 거부한다.

여러분, 솔직히 말해보세요. 혹시 AI에게 이런 식으로 말한 적 있나요?

"이커머스 앱 만들어줘."
"이거 분석해줘."
"블로그 글 써줘."

이것은 **맹목적 요청**입니다. AI에게 당신의 뇌를 빌려주는 행위에요.

진짜 개발자는 이렇게 말합니다.

"나는 이런 도메인에서 이런 비즈니스 로직을 처리해야 해. 내가 생각한 아키텍처는 이렇고, 고려 중인 트레이드오프는 이것들이야. 이 구조에 맞춰 생성해줘."

차이가 보이시나요? 전자는 AI가 모든 걸 결정합니다. 후자는 **당신이 결정하고 AI가 따릅니다.**

생각 없이 요청하지 마세요. **좋은 설계와 좋은 질문만이 가치 있는 답변을 만듭니다.**
-->
