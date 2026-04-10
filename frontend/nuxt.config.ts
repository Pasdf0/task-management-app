export default defineNuxtConfig({
  css: [
    '~/assets/scss/main.scss',
  ],
  app: {
    head: {
      htmlAttrs: {
        'data-theme': 'light'
      }
    }
  },
  
  runtimeConfig: {
    apiInternalBaseUrl: '',
    public: {
      apiBaseUrl: '',
      apiTimeout: 10000
    }
  }
})