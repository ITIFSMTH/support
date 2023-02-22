<template>
  <div class="row valign-wrapper">
    <div class="col s12 m8 offset-m2 l8 offset-l2 xl4 offset-xl4">
      <form @submit.prevent="onLogin" class="card small light-blue lighten-5">
        <div class="card-content">
          <span class="card-title">Вход</span>
          
          <div class="inputs">
            <div class="input-field">
              <input v-model="login" class="input-obj validate" id="login" type="text" />
              <label
              :class="{'red-text': isLoginError, 'green-text': isLoginError === false}"
              class="active" 
              for="login">Логин</label>
            </div>

            <div class="input-field">
              <input
                v-model="password"
                class="input-obj validate"
                id="password"
                type="password"
              />
              <label 
              :class="{'red-text': isPasswordError, 'green-text': isPasswordError === false}"
              class="active" 
              for="password">Пароль</label>
            </div>
          </div>

          <div class="buttons center">
            <button class="button-login btn waves-effect green waves-light" type="submit">
              <span class="button-text">Войти</span>
              <i class="button-icon material-icons">login</i>
            </button>
          </div>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
export default {
  name: "login",
  computed: {
    isLoginError() {return this.$store.state.login.loginError},
    isPasswordError() {return this.$store.state.login.passwordError},
    login: {
      get() {
        return this.$store.state.login.login;
      },
      set(login) {
        this.$store.commit("updateLogin", login);
      },
    },
    password: {
      get() {
        return this.$store.state.login.password;
      },
      set(password) {
        this.$store.commit("updatePassword", password);
      },
    },
  },
  methods: {
    async onLogin() {
      const success = await this.$store.dispatch("login")
      if (success) this.$router.push("/chats")
    },
  },
  created() {
    document.title = "Orb11ta | Авторизация"
  }
};
</script>

<style scoped>
.row {
  height: calc(100% - 20px);
}

.card-title {
  font-weight: 500;
}

.input-obj {
  color: #263238;
  background: none;
}

.inputs {
  padding-top: 10px;
}

.button-login {
  width: 100%;
}

.button-text {
  font-family: "Roboto", sans-serif;
  font-weight: 500;
  text-transform: none;
  vertical-align: top;
}

.button-icon {
  margin-left: 5px;
}
</style>
