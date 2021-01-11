<template>
  <v-form class="login-form">
      <v-text-field
        v-model="login"
        required
        :label="$t('input-label')"
      />
    <v-btn @click="loginAction">{{ $t('login-btn-text') }}</v-btn>
    </v-form>
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

<i18n>
{
  "en": {
    "input-label": "Enter a login",
    "login-btn-text": "login"
  },
  "ru": {
    "input-label": "Введите логин",
    "login-btn-text": "войти"
  }
}
</i18n>
