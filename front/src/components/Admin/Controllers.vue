<template>
  <div class="admin-controllers">
    <AddOperator />
    <Mailing />
  </div>
</template>

<script>
import AddOperator from "./AddOperator";
import Mailing from "./Mailing";

export default {
  name: "controllers",
  components: {
    AddOperator,
    Mailing
  },
};
</script>

<style>
@media (min-width: 992px) {
  .controllers > div {
    display: flex;
    justify-content: space-between;
    width: 100%;
  }

  .controllers form {
    margin-left: 0 !important;
    width: 100% !important;
  }

  .controllers form:first-child {
    margin-right: 20px;
  }

  .card-title {
    font-size: 22px !important;
  }

  .card-content {
    height: 100%;
  }

  .admin-controllers .card.small {
    min-height: 300px !important;
  }
}

.card.small {
  height: auto !important;
}

.card-title {
  font-weight: 500 !important;
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
