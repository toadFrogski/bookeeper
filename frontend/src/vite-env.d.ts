/// <reference types="vite/client" />
/// <reference types="vite-plugin-svgr/client" />

interface ImportMetaEnv {
  readonly VITE_MEDIA_DIRECTORY: string
}

interface ImportMeta {
  readonly env: ImportMetaEnv
}