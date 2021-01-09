<template>
  <div class="login-form">
    <v-form>
        <v-text-field
          v-model="login"
          required
          label="Логин"
        />
        <v-text-field
          v-model="password"
          required
          type="password"
          label="Пароль"
        />
      <v-btn @click="loginAction">Войти</v-btn>
    </v-form>
  </div>
</template>

<script lang="ts">
import {
  Component, Vue, Watch,
} from 'vue-property-decorator';

@Component

export default class LoginPage extends Vue {
  login = ''

  password = ''

  get isAuthorized(): boolean {
    return this.$store.direct.getters.isAuthorized;
  }

  @Watch('isAuthorized')
  onUserAuthChange(isAuthorized: boolean) {
    if (isAuthorized) {
      this.$router.push({ name: 'Home' });
    }
  }

  loginAction() {
    this.$store.direct.dispatch.adminLogin({ login: this.login, password: this.password });
  }

  created() {
    this.$store.direct.dispatch.adminLogin()
      .finally(() => {
        if (this.isAuthorized) {
          this.$router.push({ name: 'Home' });
        }
      });
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
