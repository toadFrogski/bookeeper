import i18n from "i18next";
import Backend from "i18next-http-backend";
import LanguageDetector from "i18next-browser-languagedetector";
import { initReactI18next } from "react-i18next";

export enum I18Namespace {
  Bookeeper = "bookeeper",
}

export enum Language {
  English = "en",
  Romanian = "ro",
  Russian = "ru",
}

i18n
  .use(Backend)
  .use(LanguageDetector)
  .use(initReactI18next)
  .init({
    backend: {
      loadPath: "/locales/{{lng}}/{{ns}}.json",
    },
    detection: {
      order: ["querystring", "localStorage", "navigator"],
      lookupQuerystring: "lng",
      lookupCookie: "lng",
      lookupLocalStorage: "lng",
      caches: ["localStorage"],
      convertDetectedLanguage: (lng: string) => lng.split("-")[0],
    },
    debug: import.meta.env.DEV,
    react: {
      // wait: true,
      transSupportBasicHtmlNodes: true,
      transKeepBasicHtmlNodesFor: ["br"],
    },
    defaultNS: I18Namespace.Bookeeper,
    ns: I18Namespace.Bookeeper,
    load: "currentOnly",
  });
export default i18n;
