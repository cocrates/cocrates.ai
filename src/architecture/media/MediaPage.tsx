import type {ReactNode} from 'react';
import Layout from '@theme/Layout';
import Heading from '@theme/Heading';
import Link from '@docusaurus/Link';
import Translate, {translate} from '@docusaurus/Translate';
import PictureBookViewer from '@site/src/components/PictureBookViewer';
import {
  BRUSH_HERO_PAGES,
  CHIKA_PAGES,
  GITHUB_BRUSH_HERO,
  GITHUB_CHIKA_BOOK,
  GITHUB_CHIKA_VIDEO,
  GITHUB_JSONDB_MEDIA,
  YOUTUBE_CHIKA_VIDEO_ID,
} from '@site/src/architecture/media/constants';

import styles from './styles.module.css';

function GithubLink({href}: {href: string}): ReactNode {
  return (
    <Link to={href}>
      <Translate id="media.github">GitHub</Translate>
    </Link>
  );
}

function Section({
  id,
  title,
  children,
}: {
  id: string;
  title: ReactNode;
  children: ReactNode;
}): ReactNode {
  return (
    <section id={id} className={styles.section}>
      <Heading as="h2" className={styles.sectionTitle}>
        {title}
      </Heading>
      {children}
    </section>
  );
}

export default function MediaPage(): ReactNode {
  return (
    <Layout
      wrapperClassName="media-page"
      title={translate({
        id: 'media.page.title',
        message: '미디어',
      })}
      description={translate({
        id: 'media.page.description',
        message:
          'Cocrates로 만든 동영상, 그림책, 다이어그램 미디어 갤러리',
      })}>
      <main className={styles.main}>
        <div className="container">
          <header className={styles.header}>
            <Heading as="h1" className={styles.pageTitle}>
              <Translate id="media.page.title">미디어</Translate>
            </Heading>
            <p className={styles.lead}>
              <Translate id="media.page.lead">
                Cocrates 스킬로 만든 동영상·그림책·구조 설명 미디어입니다.
              </Translate>
            </p>
          </header>

          <Section
            id="chika-video"
            title={
              <Translate id="media.section.chikaVideo.title">
                &apos;치카치카 티라노&apos; 동영상
              </Translate>
            }>
            <p className={styles.paragraph}>
              <Translate id="media.section.chikaVideo.p1">
                AI를 잘 사용하면 동영상 만들기도 어렵지 않다.
              </Translate>
            </p>
            <p className={styles.paragraph}>
              <Translate
                id="media.section.chikaVideo.p2"
                values={{
                  skill: <code>video-authoring</code>,
                }}>
                {
                  '이 동영상은 {skill} 스킬로 만들어졌다: 스토리 비트를 컷 단위로 설계한 뒤, 대사/효과음/BGM을 포함해 여러 클립을 검증·승인 후 조립한다.'
                }
              </Translate>
            </p>
            <p className={styles.paragraph}>
              <Translate
                id="media.section.chikaVideo.p3"
                values={{
                  githubLink: <GithubLink href={GITHUB_CHIKA_VIDEO} />,
                }}>
                {
                  '진짜 어려운 것은 “무엇을, 어떤 순서의 장면으로 보여줄지”를 설계하는 것이다. 설계 문서와 원본은 {githubLink}에서 확인할 수 있다.'
                }
              </Translate>
            </p>
            <div className={styles.videoWrap}>
              <iframe
                className={styles.youtube}
                src={`https://www.youtube.com/embed/${YOUTUBE_CHIKA_VIDEO_ID}`}
                title={translate({
                  id: 'media.video.chika.title',
                  message: '치카치카 티라노',
                })}
                allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share"
                allowFullScreen
                referrerPolicy="strict-origin-when-cross-origin"
              />
            </div>
          </Section>

          <Section
            id="chika-book"
            title={
              <Translate id="media.section.chikaBook.title">
                &apos;치카치카 티라노&apos; 그림책
              </Translate>
            }>
            <p className={styles.paragraph}>
              <Translate id="media.section.chikaBook.p1">
                AI를 잘 사용하면 그림책 구성도 어렵지 않다.
              </Translate>
            </p>
            <p className={styles.paragraph}>
              <Translate
                id="media.section.chikaBook.p2"
                values={{
                  skill: <code>picture-book-authoring</code>,
                }}>
                {
                  '이 그림책은 {skill} 스킬로 작성되었다. 페이지 단위로 스토리·세계관·캐릭터·테마를 설계하고, 검증·승인을 거쳐 전체 흐름의 일관성을 유지한 채 완성본 페이지를 만든다.'
                }
              </Translate>
            </p>
            <p className={styles.paragraph}>
              <Translate
                id="media.section.chikaBook.p3"
                values={{
                  githubLink: <GithubLink href={GITHUB_CHIKA_BOOK} />,
                }}>
                {'설계 문서와 원본은 {githubLink}에서 확인할 수 있다.'}
              </Translate>
            </p>
            <PictureBookViewer
              pages={CHIKA_PAGES}
              altPrefix={translate({
                id: 'media.picturebook.chika.altPrefix',
                message: '치카치카 티라노 페이지',
              })}
            />
          </Section>

          <Section
            id="brush-hero"
            title={
              <Translate id="media.section.brushHero.title">
                &apos;이빨나라의 작은 영웅&apos; 그림책
              </Translate>
            }>
            <p className={styles.paragraph}>
              <Translate
                id="media.section.brushHero.p1"
                values={{
                  skill: <code>picture-book-authoring</code>,
                  githubLink: <GithubLink href={GITHUB_BRUSH_HERO} />,
                }}>
                {
                  '{skill} 스킬로 만든 이닦기 그림책으로, 페이지 단위 설계·검증·승인을 통해 완결성과 흐름을 맞췄다. 설계 문서와 원본은 {githubLink}에서 확인할 수 있다.'
                }
              </Translate>
            </p>
            <PictureBookViewer
              pages={BRUSH_HERO_PAGES}
              altPrefix={translate({
                id: 'media.picturebook.brushHero.altPrefix',
                message: '이빨나라의 작은 영웅 페이지',
              })}
            />
          </Section>

          <Section
            id="jsondb"
            title={
              <Translate id="media.section.jsondb.title">
                jsondb 구조 설명
              </Translate>
            }>
            <p className={styles.paragraph}>
              <Translate
                id="media.section.jsondb.p1"
                values={{
                  diagramSkill: <code>diagram-generating</code>,
                  imageSkill: <code>image-generation</code>,
                  videoSkill: <code>video-generation</code>,
                  githubLink: <GithubLink href={GITHUB_JSONDB_MEDIA} />,
                }}>
                {
                  'jsondb의 동시 요청 병렬 처리 구조를 설명한 미디어다. 구조도는 {diagramSkill} 스킬(excalidraw 생성), 설명 이미지는 {imageSkill} 스킬(google nano banana 모델 사용), 설명 동영상은 {videoSkill} 스킬(google omni 모델 사용)로 제작했다. 원본은 {githubLink}에서 확인할 수 있다.'
                }
              </Translate>
            </p>

            <Heading as="h3" className={styles.subTitle}>
              <Translate id="media.jsondb.diagram">다이어그램</Translate>
            </Heading>
            <figure className={styles.figure}>
              <img
                src="/img/jsondb/jsondb-architecture.png"
                alt={translate({
                  id: 'media.jsondb.diagram.alt',
                  message: 'jsondb 구조',
                })}
                className={styles.mediaImage}
              />
              <figcaption className={styles.caption}>
                <Translate id="media.jsondb.diagram.caption">
                  jsondb 구조
                </Translate>
              </figcaption>
            </figure>

            <Heading as="h3" className={styles.subTitle}>
              <Translate id="media.jsondb.image">이미지</Translate>
            </Heading>
            <figure className={styles.figure}>
              <img
                src="/img/jsondb/jsondb-explanation.png"
                alt={translate({
                  id: 'media.jsondb.image.alt',
                  message: 'jsondb 구조를 설명하는 이미지',
                })}
                className={styles.mediaImage}
              />
              <figcaption className={styles.caption}>
                <Translate id="media.jsondb.image.caption">
                  jsondb 구조를 설명하는 이미지
                </Translate>
              </figcaption>
            </figure>

            <Heading as="h3" className={styles.subTitle}>
              <Translate id="media.jsondb.video">동영상</Translate>
            </Heading>
            <figure className={styles.figure}>
              <div className={styles.videoWrap}>
                <video
                  className={styles.localVideo}
                  controls
                  preload="metadata"
                  src="/img/jsondb/jsondb-explanation.mp4">
                  <Translate id="media.video.unsupported">
                    브라우저가 video 태그를 지원하지 않습니다.
                  </Translate>
                </video>
              </div>
              <figcaption className={styles.caption}>
                <Translate id="media.jsondb.video.caption">
                  jsondb 구조를 설명하는 동영상
                </Translate>
              </figcaption>
            </figure>
          </Section>
        </div>
      </main>
    </Layout>
  );
}
