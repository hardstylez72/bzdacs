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
  Component, Vue, Watch,
} from 'vue-property-decorator';

@Component
export default class LoginPage extends Vue {
  login = ''

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
    this.$store.direct.dispatch.guestLogin(this.login)
      .finally(() => {
        console.log('guest login');
        console.log('this.isAuthorized', this.isAuthorized);
        if (this.isAuthorized) {
          this.$store.direct.dispatch.userSession();
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
