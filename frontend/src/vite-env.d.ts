/// <reference types="vite/client" />

interface ImportMetaEnv {
  readonly VITE_MEDIA_DIRECTORY: string
}

interface ImportMeta {
  readonly env: ImportMetaEnv
}