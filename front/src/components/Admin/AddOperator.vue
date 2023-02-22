<template>
  <form 
  @submit.prevent="addOperator" 
  class="controller-card card small light-blue lighten-5 col s12 m12 l4 xl4">
    <div class="card-content">
      <span class="card-title">Добавить оператора</span>

      <div class="inputs">
        <div class="input-field">
          <input v-model="login" class="input-obj validate" id="login" type="text" />
          <label
            :class="{ 'red-text': isLoginError, 'green-text': isLoginError === false }"
            class="active"
            for="login"
            >Логин</label
          >
        </div>

        <div class="input-field">
          <input
            v-model="password"
            class="input-obj validate"
            id="password"
            type="password"
          />
          <label
            :class="{
              'red-text': isPasswordError,
              'green-text': isPasswordError === false,
            }"
            class="active"
            for="password"
            >Пароль</label
          >
        </div>
      </div>

      <div class="buttons center">
        <button class="button-login btn waves-effect green waves-light" type="submit">
          <span class="button-text">Создать</span>
          <i class="button-icon material-icons">person_add</i>
        </button>
      </div>
    </div>
  </form>
</template>

<script>
export default {
  name: 'addOperator',
  computed: {
    isLoginError() {return this.$store.state.admin.addOperator.loginError},
    isPasswordError() {return this.$store.state.admin.addOperator.passwordError},
    login: {
      get() {
        return this.$store.state.admin.addOperator.login;
      },
      set(login) {
        this.$store.commit("updateAddOperatorLogin", login);
      },
    },
    password: {
      get() {
        return this.$store.state.admin.addOperator.password;
      },
      set(password) {
        this.$store.commit("updateAddOperatorPassword", password);
      },
    },
  },
  methods: {
    async addOperator() {
      await this.$store.dispatch("addOperator")
    }
  }
}
</script>