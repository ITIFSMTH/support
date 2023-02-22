<template>
  <nav class="nav light-blue lighten-5">
    <div class="nav-wrapper">
      <ul class="nav-ul">
        <li
          class="menu-link"
          v-if="isChats"
          @click="openChatBar"
        >
          <a class="top-nav sidenav-trigger full"><i class="material-icons">menu</i></a>
        </li>
      
        <router-link
        v-if="isUserAdmin"
        to="/bots"
        custom
        v-slot="{ href, navigate, isActive }">
          <li
          :class="{'active': isActive}"
          class="nav-li">
            <a
            :href="href"
            @click="navigate"
            class="nav-a">
              <i class="nav-icon material-icons">precision_manufacturing</i>
              <span class="nav-text">Боты</span>
            </a>
          </li>
        </router-link>
        <router-link
        v-if="isUserAdmin"
        to="/admin"
        custom
        v-slot="{ href, navigate, isActive }">
          <li
          :class="{'active': isActive}"
          class="nav-li">
            <a
            :href="href"
            @click="navigate"
            class="nav-a">
              <i class="nav-icon material-icons">shield</i>
              <span class="nav-text">Админ</span>
            </a>
          </li>
        </router-link>
        <router-link
        to="/chats"
        custom
        v-slot="{ href, navigate, isActive }">
          <li
          :class="{'active': isActive}" 
          class="nav-li">
            <a 
            :href="href"
            @click="navigate"
            class="nav-a">
              <i class="nav-icon material-icons">forum</i>
              <span class="nav-text">Чаты</span>
            </a>
          </li>
        </router-link>

        <li
          class="logout-link"
          @click="logout"
        >
          <a class="top-nav full"><i class="nav-logout material-icons">logout</i></a>
        </li>
      </ul>
    </div>
  </nav>
</template>

<script>
  export default {
    name: 'navbar',
    methods: {
      openChatBar() {
        this.$emit('openChatBar')
      },
      logout() {
        this.$store.commit('logout')
        this.$router.push('/login')
      }
    },
    computed: {
      isChats() {
        return this.$route.name === 'chats';
      },
      isUserAdmin() {
        return this.$store.getters.isUserAdmin
      }
    }
  }
</script>

<style scoped>
  .menu-link {
    position: absolute;
  }

  .logout-link {
    position: absolute;
    right: 15px;
    top: 0;
  }

  .nav {
    position: fixed;
    z-index: 50;
  }

  .nav-ul {
    text-align: center;
    height: 100%;
  }

  .nav-li {
    display: inline-block;
    float: none;
    height: 100%;
  }

  .nav-a {
    height: 100%;
  }

  .nav-icon {
    height: 10px;
    line-height: 42px;
  }

  .nav-li.active {
    background-color: rgba(0,0,0,0.25) ;
  }

  .nav-icon,
  .nav-text {
    color: #263238;
  }

  .nav-logout {
    color: #263238;
  }
</style>
