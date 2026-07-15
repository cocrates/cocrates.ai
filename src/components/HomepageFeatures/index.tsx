import type {ReactNode} from 'react';
import clsx from 'clsx';
import Link from '@docusaurus/Link';
import useBaseUrl from '@docusaurus/useBaseUrl';
import useDocusaurusContext from '@docusaurus/useDocusaurusContext';
import Heading from '@theme/Heading';
import Translate, {translate} from '@docusaurus/Translate';
import styles from './styles.module.css';

type EpisodeItem = {
  ep: string;
  title: string;
  summary: string;
  to: string;
};

const EpisodeList: EpisodeItem[] = [
  {
    ep: 'EP1',
    title: translate({
      id: 'homepage.ep1.title',
      message: '같은 LLM, 다른 결과',
    }),
    summary: translate({
      id: 'homepage.ep1.summary',
      message:
        '같은 AI를 쓰는데 왜 결과는 다를까? 차이는 모델이 아니라 사용자의 태도에서 시작된다.',
    }),
    to: '/docs/same-llm-different-results',
  },
  {
    ep: 'EP2',
    title: translate({
      id: 'homepage.ep2.title',
      message: '검토하지 않은 산출물의 무가치론',
    }),
    summary: translate({
      id: 'homepage.ep2.summary',
      message:
        'AI가 만든 결과를 이해하고 판단하고 승인하지 못한다면, 그것은 아직 당신의 산출물이 아니다.',
    }),
    to: '/docs/unexamined-code',
  },
  {
    ep: 'EP3',
    title: translate({
      id: 'homepage.ep3.title',
      message: 'Cocrates Harness 설치',
    }),
    summary: translate({
      id: 'homepage.ep3.summary',
      message:
        'OpenCode에 Cocrates Harness를 설치하고 사용해 본다.',
    }),
    to: '/docs/installation',
  },
  {
    ep: 'EP4',
    title: translate({
      id: 'homepage.ep4.title',
      message: '소크라테스식 학습 실습',
    }),
    summary: translate({
      id: 'homepage.ep4.summary',
      message:
        '정답을 받아 적는 대신, 질문에 답하며 능동적으로 학습한다.',
    }),
    to: '/docs/socratic-learning-practice',
  },
  {
    ep: 'EP5',
    title: translate({
      id: 'homepage.ep5.title',
      message: '구조 기반 생성 실습',
    }),
    summary: translate({
      id: 'homepage.ep5.summary',
      message:
        '바로 만들지 않는다. 먼저 구조를 검토하고, 구조를 기반으로 생성/검증한다.',
    }),
    to: '/docs/architecture-driven-generation-practice',
  },
  {
    ep: 'EP6',
    title: translate({
      id: 'homepage.ep6.title',
      message: '구조 기반 스킬 생성 실습',
    }),
    summary: translate({
      id: 'homepage.ep6.summary',
      message:
        '한 번 쓸 결과물이 아니라, 계속 재사용할 생성 스킬을 AI와 함께 설계/생성한다.',
    }),
    to: '/docs/workflow-creation-practice',
  },
  {
    ep: 'EP7',
    title: translate({
      id: 'homepage.ep7.title',
      message: 'Cocrates Harness 구조',
    }),
    summary: translate({
      id: 'homepage.ep7.summary',
      message:
        '하나의 만능 프롬프트가 아니라, Agent와 Skills로 구성된 하네스다.',
    }),
    to: '/docs/cocrates-harness-structure',
  },
  {
    ep: 'EP8',
    title: translate({
      id: 'homepage.ep8.title',
      message: '소크라테스식 학습 원리',
    }),
    summary: translate({
      id: 'homepage.ep8.summary',
      message:
        '왜 Cocrates는 바로 설명하지 않고 질문할까? 성장하게 만드는 학습 원리를 살펴본다.',
    }),
    to: '/docs/socratic-learning-activity',
  },
  {
    ep: 'EP9',
    title: translate({
      id: 'homepage.ep9.title',
      message: '소크라테스식 학습 스킬',
    }),
    summary: translate({
      id: 'homepage.ep9.summary',
      message:
        '배우고, 기록하고, 다시 검증한다. 세 스킬이 지식을 내 것으로 만드는 방식을 본다.',
    }),
    to: '/docs/socratic-learning-skills',
  },
  {
    ep: 'EP10',
    title: translate({
      id: 'homepage.ep10.title',
      message: '구조 기반 생성 원리',
    }),
    summary: translate({
      id: 'homepage.ep10.summary',
      message:
        '좋은 산출물은 구조에서 나온다. 구조 기반으로 산출물을 생성하는 일반적인 워크플로우를 이해한다.',
    }),
    to: '/docs/architecture-driven-generation-activity',
  },
  {
    ep: 'EP11',
    title: translate({
      id: 'homepage.ep11.title',
      message: '구조 기반 생성 스킬',
    }),
    summary: translate({
      id: 'homepage.ep11.summary',
      message:
        '결정, 명세, 생성, 검증을 담당하는 스킬들이 어떻게 품질을 지키는지 파고든다.',
    }),
    to: '/docs/architecture-driven-generation-skills',
  },
  {
    ep: 'EP12',
    title: translate({
      id: 'homepage.ep12.title',
      message: '구조 기반 스킬 생성 원리',
    }),
    summary: translate({
      id: 'homepage.ep12.summary',
      message:
        'AI에게 일을 시키는 법 자체를 만든다. 스킬을 만드는 메타 스킬의 구조를 본다.',
    }),
    to: '/docs/workflow-creation-skill',
  },
  {
    ep: 'EP13',
    title: translate({
      id: 'homepage.ep13.title',
      message: '올바른 Cocrates Harness 활용',
    }),
    summary: translate({
      id: 'homepage.ep13.summary',
      message:
        'Cocrates는 완성품이 아니다. 사용자와 함께 진화하는 하네스다.',
    }),
    to: '/docs/evolution-of-cocrates',
  },
  {
    ep: 'EP14',
    title: translate({
      id: 'homepage.ep14.title',
      message: 'Cocrates 사용자 선언문',
    }),
    summary: translate({
      id: 'homepage.ep14.summary',
      message:
        'AI를 맹신하지 않고, 주도권을 넘기지 않고, 끝까지 성장하겠다는 선언이다.',
    }),
    to: '/docs/user-manifesto',
  },
];

function Section({
  id,
  title,
  children,
}: {
  id: string;
  title: string;
  children: ReactNode;
}) {
  return (
    <section id={id} className={styles.section}>
      <div className="container">
        <Heading as="h2" className={styles.sectionTitle}>
          {title}
        </Heading>
        {children}
      </div>
    </section>
  );
}

function EpisodeCard({ep, title, summary, to}: EpisodeItem) {
  return (
    <Link to={to} className={clsx('card', styles.episodeCard)}>
      <div className="card__header">
        <span className={styles.episodeBadge}>{ep}</span>
        <Heading as="h3" className={styles.episodeTitle}>
          {title}
        </Heading>
      </div>
      <div className="card__body">
        <p className={styles.episodeSummary}>{summary}</p>
      </div>
    </Link>
  );
}

function InflearnBanner() {
  const {
    i18n: {currentLocale},
  } = useDocusaurusContext();
  const isKorean = currentLocale === 'ko';
  const bannerSrc = useBaseUrl(
    isKorean
      ? '/img/cocrates-inflearn-banner.png'
      : '/img/cocrates-inflearn-banner-en.png',
  );
  const href = isKorean ? 'https://inf.run/nm2dW' : 'https://inf.run/she9D';
  const alt = translate({
    id: 'homepage.inflearn.banner.alt',
    message: '인프런 강의: AI를 사용하는 올바른 방법 — Cocrates',
  });

  return (
    <a
      className={styles.inflearnBanner}
      href={href}
      target="_blank"
      rel="noopener noreferrer">
      <img src={bannerSrc} alt={alt} width={1279} height={321} />
    </a>
  );
}

export default function HomepageFeatures(): ReactNode {
  return (
    <>
      <Section
        id="meaning"
        title={translate({
          id: 'homepage.section.meaning',
          message: 'AI 시대의 주도권을 되찾는 방법',
        })}>
        <div className={styles.meaningGrid}>
          <div className={styles.meaningBlock}>
            <Heading as="h3" className={styles.meaningHeading}>
              Cocrates = Co-Socrates
            </Heading>
            <ul className={styles.list}>
              <li>
                <Translate id="homepage.meaning.user">
                  사용자는 Socrates가 되어 AI가 놓치는 것, 모르는 것, 쉽게
                  넘어가려는 것을 끝까지 질문한다.
                </Translate>
              </li>
              <li>
                <Translate id="homepage.meaning.ai">
                  AI도 Socrates가 되어 사용자의 막연한 요청을 구조화하고, 함께
                  설계한 기준으로 산출물을 만들고 검증한다.
                </Translate>
              </li>
            </ul>
          </div>
          <blockquote className={styles.quote}>
            <Translate id="homepage.meaning.quote">
              Cocrates는 AI가 대신 생각하게 만드는 도구가 아니다. 사용자가 더
              깊이 생각하고, 더 정확히 결정하고, 더 빠르게 성장하게 만드는
              하네스다.
            </Translate>
          </blockquote>
        </div>
      </Section>

      <Section
        id="install"
        title={translate({
          id: 'homepage.section.install',
          message: '지금 바로 시작하세요',
        })}>
        <p className={styles.paragraph}>
          <Translate
            id="homepage.install.intro"
            values={{
              opencodeLink: (
                <a
                  href="https://opencode.ai/download"
                  target="_blank"
                  rel="noopener noreferrer">
                  OpenCode
                </a>
              ),
            }}>
            {
              '{opencodeLink}만 설치한 뒤, AI에게 아래처럼 Cocrates Harness 설치를 요청하면 된다.'
            }
          </Translate>
        </p>
        <div className={styles.codeBlock}>
          <code>
            <Translate id="homepage.install.prompt">
              https://cocrates.ai/install.md 에 따라 Cocrates Harness를
              설치해줘
            </Translate>
          </code>
        </div>
        <p className={styles.paragraph}>
          <Translate id="homepage.install.detail">
            상세 설치는
          </Translate>{' '}
          <Link to="/docs/installation">
            <Translate id="homepage.install.guide">설치 가이드</Translate>
          </Link>
          <Translate id="homepage.install.detailSuffix">를 참고할 것.</Translate>
        </p>
      </Section>

      <Section
        id="usage"
        title={translate({
          id: 'homepage.section.usage',
          message: '이렇게 써보면 차이가 보인다',
        })}>
        <ul className={clsx(styles.list, styles.usageList)}>
          <li>
            <Translate id="homepage.usage.learn">
              궁금한 것을 물어보라. Cocrates는 답을 쏟아내기 전에, 당신이
              스스로 이해하도록 질문한다. 질문에 답하며 당신은 깊이 이해하게 된다.
            </Translate>
          </li>
          <li>
            <Translate id="homepage.usage.build">
              개발을 요청해 보라. Cocrates는 바로 코드를 쓰기 전에, 무엇을
              결정해야 하는지부터 드러낸다. 결정을 검토하며 당신이 원하는 것을 만들게 된다.
            </Translate>
          </li>
          <li>
            <Translate id="homepage.usage.growth">
              그 과정에서 당신은 AI를 리딩할 수 있는 사람으로 성장한다.
            </Translate>
          </li>
        </ul>
      </Section>

      <Section
        id="docs"
        title={translate({
          id: 'homepage.section.docs',
          message: '14편으로 따라가는 Cocrates 여정',
        })}>
        <div className={styles.episodeGrid}>
          {EpisodeList.map(item => (
            <EpisodeCard key={item.ep} {...item} />
          ))}
        </div>
      </Section>

      <Section
        id="media"
        title={translate({
          id: 'homepage.section.media',
          message: '강의와 영상',
        })}>
        <div className={styles.mediaStack}>
          <InflearnBanner />
        </div>
      </Section>
    </>
  );
}
