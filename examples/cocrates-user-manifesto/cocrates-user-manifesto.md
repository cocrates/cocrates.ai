---
marp: true
theme: uncover
paginate: true
size: 16:9
---

<style>
/* === from marps/template-01.md === */
section.tpl-01 {
  background: #0a0a0a;
  color: #ffffff;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  text-align: center;
  padding: 10% 8%;
  font-family: 'Noto Sans KR', 'Pretendard', sans-serif;
}
section.tpl-01 h1 {
  font-size: 54pt;
  font-weight: 900;
  color: #f0f0f0;
  letter-spacing: -0.02em;
  margin: 0 0 0.3em 0;
  line-height: 1.2;
}
section.tpl-01 h2 {
  font-size: 22pt;
  font-weight: 400;
  color: #888888;
  margin: 0 0 1.5em 0;
  line-height: 1.4;
}
section.tpl-01 .tpl-01__quote {
  font-size: 14pt;
  font-weight: 300;
  color: #666666;
  margin-top: auto;
  padding-top: 2em;
}
/* === from marps/template-02.md === */
section.tpl-02 {
  background: #0a0a0a;
  color: #ffffff;
  display: flex;
  flex-direction: column;
  align-items: center;
  text-align: center;
  padding: 6% 10%;
  font-family: 'Noto Sans KR', 'Pretendard', sans-serif;
  justify-content: flex-start;
}
section.tpl-02 h1 {
  font-size: 48pt;
  font-weight: 800;
  color: #f0f0f0;
  margin: 0 0 0.4em 0;
  text-align: center;
  line-height: 1.3;
}
section.tpl-02 .tpl-02__title-divider {
  width: 40%;
  height: 1px;
  background: #333333;
  margin: 0 auto 1.2em auto;
}
section.tpl-02 .tpl-02__body {
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  width: 100%;
  max-width: 90%;
  gap: 0.3em;
}
section.tpl-02 .tpl-02__footer {
  font-size: 11pt;
  font-weight: 300;
  color: #666666;
  text-align: center;
  padding-top: 1em;
  margin-top: auto;
}
section.tpl-02 .cmd {
  font-size: 42pt;
  font-weight: 900;
  color: #ffffff;
  text-align: center;
  line-height: 1.2;
  margin: 0.2em 0;
}
section.tpl-02 .cmd-ko {
  font-size: 24pt;
  font-weight: 500;
  color: #999999;
  text-align: center;
  margin: 0.1em 0;
}
section.tpl-02 .desc {
  font-size: 16pt;
  font-weight: 400;
  color: #cccccc;
  text-align: center;
  line-height: 1.5;
  margin: 0.3em 0;
  max-width: 85%;
}
section.tpl-02 .point {
  font-size: 24pt;
  font-weight: 400;
  color: #ffffff;
  text-align: center;
  margin: 0.5em 0;
}
section.tpl-02 .transition-lg {
  font-size: 54pt;
  font-weight: 900;
  color: #a00000;
  text-align: center;
  margin: 0.3em 0;
}
section.tpl-02 .accent-red {
  color: #ff3b30;
  font-weight: 800;
}
section.tpl-02 .declaration {
  font-size: 24pt;
  font-weight: 700;
  color: #ffffff;
  text-align: center;
  line-height: 1.4;
  margin: 0.5em 0;
  max-width: 85%;
}
section.tpl-02 .signal {
  font-size: 14pt;
  font-weight: 300;
  color: #666666;
  text-align: center;
  letter-spacing: 0.1em;
}
section.tpl-02 .highlight {
  font-weight: 900;
  color: #ffffff;
}
section.tpl-02 .eq-left {
  font-size: 22pt;
  font-weight: 700;
  color: #888888;
}
section.tpl-02 .eq-op {
  font-size: 22pt;
  font-weight: 400;
  color: #666666;
}
section.tpl-02 .eq-right {
  font-size: 22pt;
  font-weight: 900;
  color: #ffffff;
}
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
section.tpl-02 .preamble {
  font-size: 14pt;
  font-weight: 400;
  color: #999999;
  text-align: center;
  font-style: italic;
  margin: 0 0 0.4em 0;
  line-height: 1.3;
}
section.tpl-02 .manifesto-item {
  font-size: 15pt;
  font-weight: 700;
  color: #ffffff;
  text-align: left;
  line-height: 1.6;
}
section.tpl-02.compact .manifesto-item {
  font-size: 13pt;
  line-height: 1.4;
}
section.tpl-02.compact .tpl-02__body {
  gap: 0.1em;
}
section.tpl-02 .tpl-02__quote {
  font-size: 14pt;
  font-weight: 300;
  color: #666666;
  margin-top: auto;
  padding-top: 2em;
}

/* Override uncover theme: 선명한 페이지 번호 (text-shadow 블러 제거) */
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

<!-- Slide 01 — Cover -->
<!-- _class: tpl-01 -->
<!-- _paginate: false -->

# Cocrates User Manifesto

## AI 시대, 지적 주권을 선언한다

<div class="tpl-01__quote">The unexamined code is not worth generating.</div>

<!--
(천천히, 무게를 실어서)

"**The unexamined code is not worth generating.**"

검토되지 않은 산출물은 생성할 가치가 없다.

이 문장은 Cocrates의 철학을 한 줄로 압축한 것입니다. 오늘 이 자리에서 저는 여러분과 함께 선언하고자 합니다 — **Cocrates User Manifesto.**

뭐든 AI가 뚝딱 만들어주는 시대입니다. 그렇기 때문에, 우리는 지금 질문해야 합니다.

"AI 리더인가, 내가 리더인가?"
-->

---

<!-- Slide 02 — Cocrates란 무엇인가? -->
<!-- _class: tpl-02 -->
<!-- _paginate: true -->

# Cocrates란 무엇인가?

<div class="tpl-02__title-divider"></div>

<div class="tpl-02__body">

<div class="declaration highlight"><strong>Co + Socrates</strong> = 함께 배우는 스승</div>

<div class="desc">AI 시대, 사용자가 주도적으로<br> 학습하고 설계하고 검토하고 생성하도록 하는 에이전트 하네스입니다.</div>

<div class="declaration highlight">지적 주도권을 놓치지 않게 하는 도구</div>

</div>

<!--
Cocrates는 **Co + Socrates**의 합성어입니다. 소크라테스와 함께 배운다는 뜻이죠. 소크라테스는 제자들에게 답을 주지 않았습니다. 끊임없이 질문을 던져서, 그들이 스스로 답을 찾게 만들었어요.

Cocrates도 똑같이 작동합니다.

이 도구는 여러분이 직접 **학습하고, 설계하고, 검토하고, 생성하도록 강제합니다.** AI가 여러분의 두뇌를 대체하지 않습니다. 대신 — 여러분의 생각을 구조화하고, 검증하고, 완성하도록 도와줍니다. 여러분이 AI와 함께 올바르게 일하도록 강제합니다.

핵심은 이것입니다:

Cocrates는 여러분이 **지적 주권을 빼앗기지 않도록 지켜주는 도구**라는 것입니다.
-->

---

<!-- Slide 03 — 문제 제기 -->
<!-- _class: tpl-02 -->
<!-- _paginate: true -->

# 도구만으로는 부족하다

<div class="tpl-02__title-divider"></div>

<div class="tpl-02__body">

<div class="declaration highlight">Cocrates는 훌륭한 도구입니다.</div>

<div class="transition-lg">하지만...</div>

<div class="desc">도구가 아무리 좋아도, 제대로 사용하지 않으면 아무 소용없습니다.</div>

<div class="declaration highlight">당신은 여전히 AI의 노예입니다.</div>

</div>

<!--
인정합니다. Cocrates는 훌륭한 도구입니다. 정말 잘 만들어졌어요.

**하지만...**

(잠시 멈춤. 시선으로 청중을 훑는다)

도구가 아무리 좋아도, 쓰는 사람이 제대로 사용하지 않으면 아무 소용없습니다.

여러분이 Cocrates가 던져주는 제안에 "오케이, 잘했어." 하고 대충 승인만 누르고, AI가 생성한 결과물을 "알아서 잘했겠지" 하고 검토 없이 넘어간다면 — 여러분은 Cocrates를 쓰나, Copilot을 쓰나, ChatGPT를 쓰나 뭐가 다릅니까?

**당신은 여전히 AI의 노예입니다.**

(말을 천천히, 선명하게)

도구는 도구일 뿐입니다. 중요한 것은 **당신의 태도**입니다.
-->

---

<!-- Slide 04 — 선언 -->
<!-- _class: tpl-02 -->
<!-- _paginate: true -->

# Cocrates User Manifesto

<div class="tpl-02__title-divider"></div>

<div class="tpl-02__body">

<div class="declaration">To safeguard my intellectual sovereignty in the age of AI, I hereby declare<br> my intellectual independence:</div>

</div>

<div class="tpl-02__footer">
<div class="signal">——— 7계명 ———</div>
</div>

<!--
(스크린을 바라보며, 선언문을 읽듯이 천천히)

**"To safeguard my intellectual sovereignty in the age of AI,**
**I hereby declare my intellectual independence:"**

AI 시대, 나의 지적 주권을 지키기 위해 — 나는 이에 나의 지적 독립을 선언한다.

여러분, 이 문장은 단순한 치장이 아닙니다. 이것은 **결의**입니다. 앞으로 우리가 AI와 마주할 모든 순간, 우리의 뇌를 AI에게 아웃소싱하지 않겠다는 결의입니다.

이제부터 여러분께 Cocrates 사용자 선언, 7계명을 소개하겠습니다. 우리가 AI와 함께 일할 때 반드시 지켜야 할 원칙입니다. 하나씩 선언해봅시다.
-->

---

<!-- Slide 05 — 제1계명 -->
<!-- _class: tpl-02 -->
<!-- _paginate: true -->

# 제1계명

<div class="tpl-02__title-divider"></div>

<div class="tpl-02__body">

<div class="cmd">I do not blindly request.</div>

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

---

<!-- Slide 06 — 제2계명 -->
<!-- _class: tpl-02 -->
<!-- _paginate: true -->

# 제2계명

<div class="tpl-02__title-divider"></div>

<div class="tpl-02__body">

<div class="cmd">I do not blindly trust.</div>

<div class="cmd-ko">검증 없는 신뢰를 거부한다.</div>

<div class="desc">"알아서 잘했겠지" 하고 넘어가는 순간,<br>당신은 AI의 그럴싸한 요리를 그대로 삼킨 것이다.</div>

<div class="point">진정한 신뢰는 엄격한 검증 끝에 쟁취하는 전리품이다.</div>

</div>

<!--
**"I do not blindly trust."**

검증 없는 신뢰를 거부한다.

여러분, AI가 생성한 코드를 읽지 않고 프로덕션에 배포한 적 있나요? AI가 쓴 문서를 검토하지 않고 고객에게 보낸 적 있나요?

없다고 말하면 좋겠습니다. 하지만 솔직히, 다들 한 번쯤은 있지 않을까요?

AI의 출력물은 언제나 그럴싸합니다. 완벽한 문장, 자신만만한 어조, 오류 없는 포맷. 당신의 뇌는 속습니다. "아, 이건 검증할 필요 없겠다. 완벽한데?"

그 순간, 당신은 AI가 만들어낸 **그럴싸한 요리를 그대로 삼킨 것입니다.**

우리는 이것을 경계해야 합니다.

진정한 신뢰는 **엄격한 검증 끝에 쟁취하는 전리품입니다.** AI의 출력물을 처음부터 의심하세요. 깨뜨려보세요. 잘못된 부분을 찾아내세요.

그래도 당신이 믿기로 결정했다면 — 그 신뢰는 가치가 있습니다. 공짜로 주어진 맹목적 신뢰는 위험합니다.
-->

---

<!-- Slide 07 — 제3계명 -->
<!-- _class: tpl-02 -->
<!-- _paginate: true -->

# 제3계명

<div class="tpl-02__title-divider"></div>

<div class="tpl-02__body">

<div class="cmd">I do not outsource<br>my <span class="highlight">SOVEREIGNTY</span>.</div>

<div class="cmd-ko">맥락의 주권을 넘겨주지 않는다.</div>

<div class="desc">AI가 제안하는 쉬운 지름길과 편향된 기본값은 달콤한 독이다.<br>AI는 팀원, 내가 팀장이다.</div>

<div class="point">편리함과 주권을 맞바꾸지 않겠다.</div>

</div>

<!--
**"I do not outsource my SOVEREIGNTY."**

맥락의 주권을 외주 주지 않는다.

여러분, 주권(sovereignty)이라는 말이 거창하게 들리나요?

아닙니다. 이것은 아주 실용적인 문제입니다. 우리가 AI에게 '어떻게' 할지를 결정하게 하는 순간, 우리는 주권을 넘겨주는 겁니다.

LLM은 친절합니다. "이걸 이렇게 하면 되지 않을까?" 하고 지름길을 제안하죠. 그리고 그 기본값은 언제나 **편향되어 있습니다.** 가장 보편적인 답, 가장 안전한 답, 가장 무난한 답을 던져줍니다.

그리고 그 답을 당신이 "그래 그게 맞겠지" 하고 받아들인다면 — 당신은 더 이상 결정권자가 아닙니다.

이것을 기억하세요:

**AI는 팀원입니다. 당신이 팀장입니다.**

팀장이 "이 스프린트에 이 기능을 빼죠"라고 말하는데, 팀원이 "아니요, 이건 제가 판단할 문제입니다"라고 말한다면? 말도 안 되죠. 결정권자는 팀장입니다. AI와의 관계도 마찬가지입니다.

AI가 아무리 훌륭한 분석을 내놓아도, 최종 결정은 **당신의 손에** 있어야 합니다.

절대 편리함과 주권을 맞바꾸지 마십시오.
-->

---

<!-- Slide 08 — 제4계명 -->
<!-- _class: tpl-02 -->
<!-- _paginate: true -->

# 제4계명

<div class="tpl-02__title-divider"></div>

<div class="tpl-02__body">

<div class="cmd">I do not hide my ignorance.</div>

<div class="cmd-ko">무지를 아는 척 기만하지 않는다.</div>

<div class="desc">모르면서 아는 척, AI의 날림 생성에 묻어가는 것이 가장 부끄러운 일이다.</div>

<div class="point">모르는 것을 모른다고 말하는 용기가 가장 강력한 지적 무기다.</div>

</div>

<!--
**"I do not hide my ignorance."**

무지를 아는 척 기만하지 않는다.

여러분, 이 계명이 가장 실천하기 어려울 겁니다.

AI가 모든 걸 알고 있는 시대에, "저는 이것을 모릅니다"라고 말하는 것은 부끄러운 일처럼 느껴지거든요. AI가 1초 만에 대답하는 것을, 왜 내가 모르고 있어야 하는가? 자존심이 상합니다.

그러나 저는 말합니다.

**모르면서 아는 척 AI의 날림 생성에 묻어가는 것이 가장 부끄러운 일이다.**

AI가 던져준 답을 '그런 게 있었지' 하고 넘어가는 순간, 당신은 배울 기회를 영원히 잃습니다. 모르는 것은 부끄러운 게 아닙니다. **모르는 것을 아는 척하는 것이 부끄러운 겁니다.**

소크라테스는 말했습니다. "I know I know nothing." 내가 모른다는 것을 안다.

이것이 지혜의 시작입니다.

AI에게 물어보세요. 그리고 그 답을 이해할 때까지 다시 물어보세요. "이게 무슨 뜻이야?", "왜 이렇게 되는 거야?", "더 깊이 설명해줘."

모르는 것을 모른다고 말하는 용기가 가장 강력한 지적 무기입니다. 이 무지의 무기를 단단히 쥐십시오.
-->

---

<!-- Slide 09 — 제5계명 -->
<!-- _class: tpl-02 -->
<!-- _paginate: true -->

# 제5계명

<div class="tpl-02__title-divider"></div>

<div class="tpl-02__body">

<div class="cmd">I relentlessly <span class="highlight">fight</span><br>against cognitive laziness.</div>

<div class="cmd-ko">인지적 나태함에 맞서 끝까지 투쟁한다.</div>

<div class="desc">AI가 제공하는 빠른 답의 중력은 무시무시하다.<br>끌려가지 않도록 끝까지 싸워야 한다.</div>

<div class="point">멈추고 자문하라 — 내가 이 구조를 진짜 이해하고 있는가?</div>

</div>

<!--
**"I relentlessly fight against cognitive laziness."**

인지적 나태함에 맞서 끝까지 투쟁한다.

여러분, 여기까지 오는 동안 다들 느끼셨죠? 이 계명들이 말이 쉽지, 실천은 정말 어렵다는 것을.

AI가 당신에게 말합니다.

"이 코드를 최적화해드릴까요?" — "아니, 됐어. 돌아가니까."
"이 구조를 리팩토링하는 게 좋겠습니다" — "다음에 하지 뭐."
"이 개념을 설명해드릴까요?" — "뭐... 대충 알겠어."

이것이 인지적 나태함입니다. AI가 제공하는 **빠른 답의 중력은 엄청납니다.** 한 번 편하게 가는 길을 맛보면, 다시 돌아오기 어렵습니다.

투쟁하세요. 끝까지 싸우세요. AI가 던져준 답을 받아들이기 전에 잠시 멈추고, 스스로에게 물어보세요.

"내가 이 구조를 진짜 이해하고 있는가?"
"이것을 다른 방식으로도 할 수 있는가?"
"AI가 틀릴 가능성은 없는가?"

이 싸움에서 지는 유일한 방법은 **싸우지 않는 것**입니다.
-->

---

<!-- Slide 10 — 제6계명 -->
<!-- _class: tpl-02 -->
<!-- _paginate: true -->

# 제6계명

<div class="tpl-02__title-divider"></div>

<div class="tpl-02__body">

<div class="cmd">I always strive for the <span class="highlight">optimal</span>.</div>

<div class="cmd-ko">언제나 최적의 구조를 갈망한다.</div>

<div class="desc">AI가 던져주는 첫 번째 결과물에 안주하지 마라. 아키텍처에 '대충'은 없다.</div>

<div class="point">최적을 추구하지 않으면, 최고가 될 수 없다.</div>

</div>

<!--
**"I always strive for the optimal."**

언제나 최적의 구조를 갈망한다.

여러분, ChatGPT에게 질문을 던져보면 10초 안에 답이 나옵니다. 깔끔하고, 그럴싸하고, 바로 쓸 수 있을 것처럼 보이죠.

**그게 '최적'일까요?**

아닙니다. 그건 가장 평균적이고 무난한 답일 뿐입니다. LLM은 수십억 개의 데이터에서 '가장 확률이 높은' 답을 골라줄 뿐, '가장 좋은' 답을 주지 않습니다.

진짜 실력은 그 다음에서 드러납니다.

(잠시 멈춤)

AI의 첫 번째 결과물을 받아든 순간, 당신은 선택해야 합니다. "됐다, 이걸로 충분하다"고 말할 것인가, 아니면 "더 나은 방법이 없는가"라고 질문할 것인가.

이것이 이 계명의 핵심입니다.

"더 나은 방법이 없는가?"
"이 구조의 트레이드오프는 무엇인가?"
"다른 접근 방식과 비교했을 때 이것이 정말 최적인가?"

아키텍처에 **'대충'은 없습니다.** '적당히'는 존재하지 않습니다. 당신이 '됐다'고 말하는 순간, 당신의 성장은 멈춥니다.

**최적을 추구하지 않으면 최고가 될 수 없습니다.**

이것은 완벽주의가 아닙니다. 이것은 끊임없는 개선의 의지입니다. AI가 제공하는 '충분히 좋은' 답을 거부하고, 당신 스스로 '더 나은' 답을 찾아가는 태도 말이죠.
-->

---

<!-- Slide 11 — 제7계명 -->
<!-- _class: tpl-02 -->
<!-- _paginate: true -->

# 제7계명

<div class="tpl-02__title-divider"></div>

<div class="tpl-02__body">

<div class="cmd">I <span class="highlight">never</span> build a <span class="highlight">closed</span> fortress.</div>

<div class="cmd-ko">나만의 닫힌 성벽을 쌓지 않는다.</div>

<div class="desc">내가 최고라는 착각에 갇히는 순간, 고인 물이 된다.</div>

<div class="point">혼자 고군분투하지 말고, 함께 성장하라.</div>

</div>

<!--
**"I never build a closed fortress."**

나만의 닫힌 성벽을 쌓지 않는다.

일곱 번째이자 마지막 계명입니다. 이 계명은 가장 위험한 적을 경계합니다 — 바로 **자만심**입니다.

여섯 개의 계명을 완벽하게 지켜온 당신. AI에게 요청할 때마다 설계하고, 매 출력을 검증하고, 결정권을 놓지 않고, 모르는 것을 인정하고, 인지적 나태와 싸우고, 최적을 추구했습니다.

그러면 어느 순간 이런 생각이 듭니다.

"나는 이제 AI를 다룰 줄 아는 개발자다."
"내가 최고다."

**내가 최고라는 착각에 갇히는 순간, 진화는 멈추고 — 당신은 고인 물이 됩니다.**

(잠시 멈춤)

이 계명은 말합니다. 혼자서 완벽해지려 하지 말고, 함께 진화하라고. 당신의 지식과 경험을 나누고, 동료의 피드백을 경청하고, 더 나은 구조가 있다면 자존심을 버리고 유연하게 받아들이라고.

AI 시대의 진정한 강자는 '모든 것을 아는 사람'이 아닙니다. **끊임없이 배우고, 나누고, 함께 성장하는 사람입니다.**

여러분, 닫힌 성벽 안에 갇히지 마십시오. 밖으로 나와서 소통하세요.
-->

---

<!-- Slide 12 — AI-native Engineer -->
<!-- _class: tpl-02 -->
<!-- _paginate: true -->

# 도구와 의지의 결합

<div class="tpl-02__title-divider"></div>

<div class="tpl-02__body">

<div class="s-12__eq">
  <span class="eq-left s-12__eq-harness">[ Cocrates Harness ]</span>
  <span class="eq-op s-12__eq-plus">+</span>
  <span class="eq-left s-12__eq-manifesto">[ User Manifesto ]</span>
  <span class="s-12__eq-result"><span class="eq-op">=</span> <span class="eq-right">[ AI-native Engineer ]</span></span>
</div>
<br>
  <span class="point">당신의 능력은 무한 확장된다.</span>
  <span class="desc">100명의 AI 팀원을 이끄는 팀장.</span>

</div>

<!--
자, 이제 모든 것을 종합해봅시다.

(잠시 멈춤)

**Cocrates Harness**는 여러분이 AI 팀원을 어떻게 이끌어야 하는지 설계와 검토를 강제합니다.
**사용자 선언 7계명**는 지적 주도권을 유지하고 최고의 아키텍트가 되기 위한 우리의 의지입니다.

이 둘이 하나가 될 때, 우리는 AI를 10명, 100명의 팀원으로 이끄는 진정한 AI-native Engineer가 되는 것입니다.

이제 마지막 슬라이드에서 전체 선언문을 함께 읽어보겠습니다."
-->

---

<!-- Slide 13 — The Manifesto -->
<!-- _class: tpl-02 -->
<!-- _paginate: true -->

# Cocrates User Manifesto

<div class="tpl-02__title-divider"></div>

<div class="tpl-02__body">

<div class="preamble">To safeguard my intellectual sovereignty in the age of AI,<br> I hereby declare my intellectual independence:</div>

<div class="manifesto-item">
1. I do not blindly request.<br>
2. I do not blindly trust.<br>
3. I do not outsource my sovereignty.<br>
4. I do not hide my ignorance.<br>
5. I relentlessly fight against cognitive laziness.<br>
6. I always strive for the optimal.<br>
7. I never build a closed fortress.
</div>

<div class="tpl-02__quote">The unexamined code is not worth generating.</div>

</div>

<!--
(천천히, 한 줄 한 줄 의미를 음미하며 읽듯이)

**"To safeguard my intellectual sovereignty in the age of AI, I hereby declare my intellectual independence:"**

① **I do not blindly request.**
② **I do not blindly trust.**
③ **I do not outsource my sovereignty.**
④ **I do not hide my ignorance.**
⑤ **I relentlessly fight against cognitive laziness.**
⑥ **I always strive for the optimal.**
⑦ **I never build a closed fortress.**

(잠시 멈춤, 청중을 바라보며)

**"The unexamined code is not worth generating."**

이 선언을 기억하십시오. 오늘 이 자리에서 시작된 이 다짐을, 내일 코드를 작성할 때, AI와 대화할 때, 결정을 내려야 할 때마다 떠올리십시오.

감사합니다.
-->
