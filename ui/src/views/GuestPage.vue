<template>
  <div class="login-form">
    <v-form>
        <v-text-field
          v-model="login"
          required
          label="Логин"
        />
      <v-btn @click="loginAction">Войти</v-btn>
    </v-form>
  </div>
</template>

<script lang="ts">
import {
  Component, Vue,
} from 'vue-property-decorator';

@Component

export default class LoginPage extends Vue {
  login = ''

  loginAction() {
    this.$store.direct.dispatch.guestLogin(this.login)
      .then(() => {
        if (this.$store.direct.getters.isAuthorized) {
          this.$router.push({ name: 'Home' });
        }
      });
  }

  created() {
    this.loginAction();
  }
}
</script>

<style scoped lang="scss">
.login-form {
  display: flex;
  justify-content: center;
  flex-direction: column;
  align-items: center;
  width: 100%;
  text-align: center;
  margin: 20% auto;
}
</style>
