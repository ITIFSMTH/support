<template>
  <div>
    <table>
      <thead>
        <tr>
          <th>ID</th>
          <th>Оператор</th>
          <th>Создан</th>
          <th class="right">Действия</th>
        </tr>
      </thead>

      <tbody>
        <tr v-for="operator in operators" :key="operator.ID">
          <td>{{ operator.ID }}</td>
          <td>{{ operator.Login }}</td>
          <td>{{ dateTime(operator.CreatedAt) }}</td>
          <td class="right">
            <span
              @click="deleteModalOpen(operator.ID, operator.Login)"
              class="action red-text"
              >Удалить</span
            >
            |
            <span
              @click="updateLoginModalOpen(operator.ID, operator.Login)"
              class="action"
              >Изменить логин</span
            >
            |
            <span
              @click="updatePasswordModalOpen(operator.ID, operator.Login)"
              class="action"
              >Изменить пароль</span
            >
            |
            <span
              @click="statisticModalOpen(operator.ID, operator.Login)"
              class="action green-text"
            >Статистика</span>
          </td>
        </tr>
      </tbody>
    </table>

    <div ref="deleteModal" id="delete-modal" class="modal">
      <div class="modal-content">
        <h6>Вы уверены, что хотите удалить оператора {{ operatorToDelete.login }}?</h6>
      </div>
      <div class="modal-footer">
        <a class="modal-close white-text green waves-effect btn-flat">Отмена</a>
        <a
          @click="deleteOperator()"
          class="modal-close white-text red waves-effect btn-flat"
        >
          Удалить
        </a>
      </div>
    </div>

    <div ref="updateLoginModal" id="update-login-modal" class="modal">
      <div class="modal-content">
        <h6>
          Какой логин вы хотите назначить для оператора {{ loginOperator.login }}?
        </h6>
        <div class="inputs">
          <div class="input-field">
            <input v-model="login" id="login" type="text" class="input-obj validate" />
            <label
              :class="{ 'red-text': isLoginError, 'green-text': isLoginError === false }"
              class="active"
              for="login"
              >Новый логин</label
            >
          </div>
        </div>
      </div>
      <div class="modal-footer">
        <a
          @click="cancelUpdateLoginOperator()" 
          class="modal-close white-text green waves-effect btn-flat">Отмена</a>
        <a
          @click="updateLoginOperator()"
          class="modal-close white-text red waves-effect btn-flat"
        >
          Обновить
        </a>
      </div>
    </div>

    <div ref="updatePasswordModal" id="update-password-modal" class="modal">
      <div class="modal-content">
        <h6>
          Какой пароль вы хотите назначить для оператора {{ passwordOperator.login }}?
        </h6>
        <div class="inputs">
          <div class="input-field">
            <input v-model="password" id="password" type="password" class="input-obj validate" />
            <label
              :class="{ 'red-text': isPasswordError, 'green-text': isPasswordError === false }"
              class="active"
              for="password"
              >Новый пароль</label
            >
          </div>
        </div>
      </div>
      <div class="modal-footer">
        <a
          @click="cancelUpdatePasswordOperator()" 
          class="modal-close white-text green waves-effect btn-flat">Отмена</a>
        <a
          @click="updatePasswordOperator()"
          class="modal-close white-text red waves-effect btn-flat"
        >
          Обновить
        </a>
      </div>
    </div>

    <div ref="statisticModal" id="statistic-modal" class="modal">
      <div class="modal-content">
        <h5 class="center">
          Статистика оператора {{ statisticOperator.login }}
        </h5>
        <table class="messages-statistic">
          <thead>
            <tr>
              <th class="left">
                Дата
              </th>
              <th
              class="center"
              v-for="(n, i) in statisticOperator.statistic.Messages"
              :key="i">
                {{messagesDate(4 - i)}}
              </th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td class="left">
                Сообщения
              </td>
              <td
              class="center"
              v-for="(statisticDay, i) in statisticOperator.statistic.Messages"
              :key="i">
                {{statisticDay}}
              </td>
            </tr>
          </tbody>
        </table>
        <div class="messages-bottom">
          <div class="messages-today">
            <span class="caption">Сегодня: </span>
            <span class="count">{{getCountToday()}}</span>
          </div>
          <div class="messages-all">
            <span class="caption">5 дней: </span>
            <span class="count">{{getCountAll()}}</span>
          </div>
        </div>
        <table class="logins-statistic">
          <thead>
            <tr>
              <th class="center">
                Дата авторизации
              </th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="login in statisticOperator.statistic.Logins">
              <td class="center">{{dateTime(login*1000)}}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script>
import moment from "moment";

export default {
  name: "operators",
  data() {
    return {
      deleteModal: {},
      loginModal: {},
      passwordModal: {},
      statisticModal: {},
      operatorToDelete: {
        id: 0,
        login: "",
      },
      loginOperator: {
        id: 0,
        login: "",
        newLogin: "",
        isError: false,
      },
      passwordOperator: {
        id: 0,
        login: "",
        newPassword: "",
      },
      statisticOperator: {
        id: 0,
        login: "",
        statistic: {},
      }
    };
  },
  methods: {
    dateTime(value) {
      return moment(value).format("YYYY-MM-DD hh:mm:ss");
    },
    messagesDate(minus) {
      return moment(this.statisticOperator.statistic.LastDate).locale('ru').subtract(minus, 'days').format("DD MMMM");
    },
    getCountAll() {
      if (!this.statisticOperator.id) return
      let sum = 0
      for (const i of this.statisticOperator.statistic.Messages) sum += i
      return sum
    },
    getCountToday() {
      if (!this.statisticOperator.id) return
      return this.statisticOperator.statistic.Messages.at(-1)
    },
    deleteModalOpen(id, login) {
      this.operatorToDelete = {
        id: id,
        login: login,
      };
      this.deleteModal.open();
    },
    async deleteOperator() {
      await this.$store.dispatch("deleteOperator", this.operatorToDelete.id);
    },
    updateLoginModalOpen(id, login) {
      this.loginOperator.id = id
      this.loginOperator.login = login
      this.updateLoginModal.open();
    },
    cancelUpdateLoginOperator() {
      this.$store.commit("cancelUpdateOperatorLogin");
    },
    async updateLoginOperator() {
      await this.$store.dispatch(
        "updateLoginOperator",
        this.loginOperator.id
      );
    },
    updatePasswordModalOpen(id, login) {
      this.passwordOperator = {
        id: id,
        login: login,
      };
      this.updatePasswordModal.open();
    },
    cancelUpdatePasswordOperator() {
      this.$store.commit("cancelUpdateOperatorPassword");
    },
    async updatePasswordOperator() {
      await this.$store.dispatch(
        "updatePasswordOperator",
        this.passwordOperator.id,
      );
    },
    async statisticModalOpen(id, login) {
      await this.$store.dispatch(
        "fetchOperatorStatistic",
        id
      );
      
      this.statisticOperator = {
        id: id,
        login: login,
        statistic: this.$store.getters.operatorStatistic
      };
      this.statisticModal.open();
    }
  },
  computed: {
    operators() {
      return this.$store.getters.operators;
    },
    isLoginError() {
      return this.$store.state.admin.loginError;
    },
    isPasswordError() {
      return this.$store.state.admin.passwordError;
    },
    login: {
      get() {
        return this.$store.state.admin.login;
      },
      set(login) {
        this.$store.commit("updateOperatorLoginState", login);
      },
    },
    password: {
      get() {
        return this.$store.state.admin.password;
      },
      set(password) {
        this.$store.commit("updateOperatorPasswordState", password);
      },
    },
  },
  mounted() {
    this.$store.dispatch("fetchOperators");

    this.deleteModal = M.Modal.init(this.$refs.deleteModal, {
      dismissible: false
    });
    this.updateLoginModal = M.Modal.init(this.$refs.updateLoginModal, {
      dismissible: false
    });
    this.updatePasswordModal = M.Modal.init(this.$refs.updatePasswordModal, {
      dismissible: false
    });
    this.statisticModal = M.Modal.init(this.$refs.statisticModal);
  },
};
</script>

<style scoped>
.messages-bottom {
  padding-top: 5px;
  display: flex;
  justify-content: space-between;
  font-size: 19px;
}

th:first-child,
td:first-child,
th:last-child,
td:last-child {
  padding: 15px;
}

tbody tr:nth-child(2n + 1) {
  background-color: #e1f5fe;
}

.action {
  cursor: pointer;
}

.modal,
.modal-footer {
  background-color: white;
}

.modal-footer {
  padding-left: 24px;
  padding-right: 24px;
}

.modal-footer a {
  margin: 0 0 0 20px !important;
}

.input-obj {
  color: white;
  background: none;
}

.inputs {
  padding-top: 10px;
}
</style>
