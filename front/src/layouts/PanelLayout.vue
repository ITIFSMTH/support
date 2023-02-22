<template>
  <div class="white panel-layout">
    <Navbar
      @openChatBar="openChatBar"
    />
    
    <div class="panel-layout-wrapper">
      <router-view v-slot="{ Component }">
        <component ref="chats" :is="Component" />
      </router-view>
    </div>
  </div>
</template>

<script>
import Navbar from '../components/Navbar'

export default {
  name: 'panel-layout',
  methods: {
    openChatBar() {
      this.$refs.chats.openChatBar()
    }
  },
  components: {
    Navbar
  }
}
</script>

<style scoped>
  .panel-layout {
    width: 100%;
    min-height: 100vh;
    color: #263238;
  }

  .panel-layout-wrapper {
    padding-top: 64px;
  }
</style>