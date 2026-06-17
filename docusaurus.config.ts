import {themes as prismThemes} from 'prism-react-renderer';
import type {Config} from '@docusaurus/types';
import type * as Preset from '@docusaurus/preset-classic';

const config: Config = {
  title: 'Cocrates',
  tagline: '생각하고, 기록하고, 나누는 공간',
  favicon: 'img/logo.svg',

  future: {
    v4: true,
  },

  url: 'https://cocrates.github.io',
  baseUrl: '/cocrates.ai/',

  organizationName: 'cocrates',
  projectName: 'cocrates',

  onBrokenLinks: 'throw',

  i18n: {
    defaultLocale: 'ko',
    locales: ['ko'],
  },

  presets: [
    [
      'classic',
      {
        docs: {
          sidebarPath: './sidebars.ts',
          editUrl: 'https://github.com/cocrates/cocrates.ai/tree/main/',
        },
        blog: false,
        theme: {
          customCss: './src/css/custom.css',
        },
      } satisfies Preset.Options,
    ],
  ],

  themeConfig: {
    image: 'img/logo.svg',
    colorMode: {
      respectPrefersColorScheme: true,
    },
    navbar: {
      title: 'Cocrates',
      logo: {
        alt: 'Cocrates Logo',
        src: 'img/logo.svg',
      },
      items: [
        {
          type: 'docSidebar',
          sidebarId: 'docsSidebar',
          position: 'left',
          label: '문서',
        },
        {
          href: 'https://github.com/cocrates/cocrates.ai',
          label: 'GitHub',
          position: 'right',
        },
      ],
    },
    footer: {
      style: 'dark',
      links: [
        {
          title: '문서',
          items: [
            {
              label: '시작하기',
              to: '/docs/intro',
            },
          ],
        },
        {
          title: '링크',
          items: [
            {
              label: 'GitHub',
              href: 'https://github.com/cocrates/cocrates.ai',
            },
          ],
        },
      ],
      copyright: `Copyright © ${new Date().getFullYear()} Cocrates. Built with Docusaurus.`,
    },
    prism: {
      theme: prismThemes.github,
      darkTheme: prismThemes.dracula,
    },
  } satisfies Preset.ThemeConfig,
};

export default config;
