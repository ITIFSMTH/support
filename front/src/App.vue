<template>
  <div id="app">
    <component :is="layout">
      <router-view />
    </component>
  </div>
</template>

<script>
  import AuthLayout from '@/layouts/AuthLayout'
  import PanelLayout from '@/layouts/PanelLayout'

  export default {
    computed: {
      layout() {
        return (this.$route.meta.layout || 'auth') + '-layout'
      }
    },
    components: {
      AuthLayout,
      PanelLayout
    },
    metaInfo: {
      title: 'Orb11ta - Поддержка',
      titleTemplate: '%s | Orb11ta'
    }
  }
</script>

<style>
@import '~materialize-css/dist/css/materialize.min.css';
</style>
