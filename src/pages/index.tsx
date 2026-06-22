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
            AI를 사용하는 올바른 방법
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
              AI에게만 맡기지 않고, 함께 검토하며 만들어가는 도구다.
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
