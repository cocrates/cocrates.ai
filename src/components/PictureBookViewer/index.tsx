import {useCallback, useState, type ReactNode} from 'react';
import clsx from 'clsx';
import Translate, {translate} from '@docusaurus/Translate';

import styles from './styles.module.css';

type PictureBookViewerProps = {
  pages: string[];
  altPrefix: string;
  className?: string;
};

export default function PictureBookViewer({
  pages,
  altPrefix,
  className,
}: PictureBookViewerProps): ReactNode {
  const [index, setIndex] = useState(0);
  const total = pages.length;
  const current = pages[index];

  const goPrev = useCallback(() => {
    setIndex((i) => Math.max(0, i - 1));
  }, []);

  const goNext = useCallback(() => {
    setIndex((i) => Math.min(total - 1, i + 1));
  }, [total]);

  if (total === 0) {
    return null;
  }

  return (
    <div className={clsx(styles.viewer, className)}>
      <div className={styles.frame}>
        <img
          src={current}
          alt={`${altPrefix} ${index + 1} / ${total}`}
          className={styles.page}
          draggable={false}
        />
      </div>
      <div className={styles.controls}>
        <button
          type="button"
          className={clsx('button button--secondary button--sm', styles.navButton)}
          onClick={goPrev}
          disabled={index === 0}
          aria-label={translate({
            id: 'picturebook.prev.aria',
            message: '이전 페이지',
          })}>
          <Translate id="picturebook.prev">← 이전</Translate>
        </button>
        <span className={styles.pageIndicator} aria-live="polite">
          {index + 1} / {total}
        </span>
        <button
          type="button"
          className={clsx('button button--secondary button--sm', styles.navButton)}
          onClick={goNext}
          disabled={index === total - 1}
          aria-label={translate({
            id: 'picturebook.next.aria',
            message: '다음 페이지',
          })}>
          <Translate id="picturebook.next">다음 →</Translate>
        </button>
      </div>
    </div>
  );
}
