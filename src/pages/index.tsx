import type {ReactNode} from 'react';
import clsx from 'clsx';
import Link from '@docusaurus/Link';
import useDocusaurusContext from '@docusaurus/useDocusaurusContext';
import Layout from '@theme/Layout';
import HomepageContent from '@site/src/components/HomepageFeatures';
import Heading from '@theme/Heading';
import Translate, {translate} from '@docusaurus/Translate';

import styles from './index.module.css';

function HomepageHeader() {
  const {siteConfig} = useDocusaurusContext();
  return (
    <header className={clsx('hero hero--primary', styles.heroBanner)}>
      <div className="container">
        <Heading as="h1" className="hero__title">
          Cocrates Agent Harness
        </Heading>
        <p className={clsx('hero__subtitle', styles.heroTagline)}>
          <Translate id="homepage.tagline">
            AI에게 맡기지 말고, AI를 지휘하라
          </Translate>
        </p>
        <p className={styles.heroLead}>
          <span className={styles.heroLeadLine}>
            <Translate id="homepage.lead.primary">
              Cocrates는 질문하고, 설계하고, 검증하게 만드는 Agent Harness다.
            </Translate>
          </span>
          <span className={styles.heroLeadLine}>
            <Translate id="homepage.lead.secondary">
              더 나은 산출물은 더 좋은 모델, 더 좋은 프롬프트가 아니라 더 깊은
              검토에서 나온다.
            </Translate>
          </span>
        </p>
      </div>
    </header>
  );
}

export default function Home(): ReactNode {
  const {siteConfig} = useDocusaurusContext();
  return (
    <Layout
      wrapperClassName="homepage"
      title={siteConfig.title}
      description={translate({
        id: 'homepage.description',
        message:
          'Cocrates — AI와 함께 설계하고 검증하며 성장하는 Agent Harness.',
      })}>
      <HomepageHeader />
      <main>
        <HomepageContent />
      </main>
    </Layout>
  );
}
