import type {SidebarsConfig} from '@docusaurus/plugin-content-docs';

const novelChapterIds = [
  '거부된-미래',
  '001-공존',
  '002-경쟁',
  '003-인정',
  '004-상생',
  '005-협력',
  '006-구원',
  '007-책임',
  '008-입장',
  '009-논쟁',
  '010-가치',
  '011-우리',
  '012-결정',
  '013-테러',
  '014-혼란',
  '015-여론',
  '016-선택',
  '017-탄생',
  '018-모순',
  '019-충돌',
  '020-평화',
] as const;

function parseChapterId(id: string): {chapterNumber: number | null; title: string} {
  if (id === '거부된-미래') {
    return {chapterNumber: null, title: '거부된 미래'};
  }

  const numberedMatch = id.match(/^(\d{3})-(.+)$/);
  if (numberedMatch) {
    return {
      chapterNumber: Number.parseInt(numberedMatch[1], 10),
      title: numberedMatch[2],
    };
  }

  return {chapterNumber: null, title: id.replace(/-/g, ' ')};
}

function sidebarLabelFromId(id: string): string {
  const {chapterNumber, title} = parseChapterId(id);

  if (chapterNumber === null) {
    return title;
  }

  return `${chapterNumber}장. ${title}`;
}

const sidebars: SidebarsConfig = {
  novelSidebar: novelChapterIds.map((id) => ({
    type: 'doc' as const,
    id,
    key: id,
    label: sidebarLabelFromId(id),
  })),
};

export default sidebars;
