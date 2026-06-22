import React, {type ReactNode} from 'react';
import useDocusaurusContext from '@docusaurus/useDocusaurusContext';
import {useLocation} from '@docusaurus/router';
import {translate} from '@docusaurus/Translate';
import {mergeSearchStrings, useHistorySelector} from '@docusaurus/theme-common';
import {applyTrailingSlash} from '@docusaurus/utils-common';
import DropdownNavbarItem from '@theme/NavbarItem/DropdownNavbarItem';
import IconLanguage from '@theme/Icon/Language';
import type {Props} from '@theme/NavbarItem/LocaleDropdownNavbarItem';
import styles from './styles.module.css';

function stripLocalePrefixFromSuffix(
  pathnameSuffix: string,
  defaultLocale: string,
  locales: string[],
): string {
  for (const locale of locales) {
    if (locale === defaultLocale) {
      continue;
    }
    const prefix = `${locale}/`;
    if (pathnameSuffix === locale) {
      return '';
    }
    if (pathnameSuffix.startsWith(prefix)) {
      return pathnameSuffix.slice(prefix.length);
    }
  }
  return pathnameSuffix;
}

function useLocaleDropdownUtils() {
  const {
    siteConfig,
    i18n: {defaultLocale, locales, localeConfigs},
  } = useDocusaurusContext();
  const {pathname} = useLocation();
  const search = useHistorySelector((history) => history.location.search);
  const hash = useHistorySelector((history) => history.location.hash);

  const getLocaleConfig = (locale: string) => {
    const localeConfig = localeConfigs[locale];
    if (!localeConfig) {
      throw new Error(
        `Docusaurus bug, no locale config found for locale=${locale}`,
      );
    }
    return localeConfig;
  };

  const getPathnameSuffix = () => {
    const canonicalPathname = applyTrailingSlash(pathname, {
      trailingSlash: siteConfig.trailingSlash,
      baseUrl: siteConfig.baseUrl,
    });
    const rawSuffix = canonicalPathname.replace(siteConfig.baseUrl, '');
    return stripLocalePrefixFromSuffix(rawSuffix, defaultLocale, locales);
  };

  const getBaseURLForLocale = (locale: string) => {
    const localeConfig = getLocaleConfig(locale);
    const pathnameSuffix = getPathnameSuffix();
    const alternatePath = `${localeConfig.baseUrl}${pathnameSuffix}`;
    const isSameDomain = localeConfig.url === siteConfig.url;
    if (isSameDomain) {
      return `pathname://${alternatePath}`;
    }
    return `${localeConfig.url}${alternatePath}`;
  };

  return {
    getURL: (locale: string, options: {queryString?: string}) => {
      const finalSearch = mergeSearchStrings(
        [search, options.queryString],
        'append',
      );
      return `${getBaseURLForLocale(locale)}${finalSearch}${hash}`;
    },
    getLabel: (locale: string) => getLocaleConfig(locale).label,
    getLang: (locale: string) => getLocaleConfig(locale).htmlLang,
  };
}

export default function LocaleDropdownNavbarItem({
  mobile,
  dropdownItemsBefore,
  dropdownItemsAfter,
  queryString,
  ...props
}: Props): ReactNode {
  const utils = useLocaleDropdownUtils();
  const {
    i18n: {currentLocale, locales},
  } = useDocusaurusContext();

  const localeItems = locales.map((locale) => ({
    label: utils.getLabel(locale),
    lang: utils.getLang(locale),
    to: utils.getURL(locale, {queryString}),
    target: '_self',
    autoAddBaseUrl: false,
    className:
      locale === currentLocale
        ? mobile
          ? 'menu__link--active'
          : 'dropdown__link--active'
        : '',
  }));

  const items = [...dropdownItemsBefore, ...localeItems, ...dropdownItemsAfter];
  const dropdownLabel = mobile
    ? translate({
        message: 'Languages',
        id: 'theme.navbar.mobileLanguageDropdown.label',
        description: 'The label for the mobile language switcher dropdown',
      })
    : utils.getLabel(currentLocale);

  return (
    <DropdownNavbarItem
      {...props}
      mobile={mobile}
      label={
        <>
          <IconLanguage className={styles.iconLanguage} />
          {dropdownLabel}
        </>
      }
      items={items}
    />
  );
}
