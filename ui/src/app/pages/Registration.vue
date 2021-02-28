<template>
  <div class="login-form">

    <v-form ref="form">
        <v-text-field
          v-model="login"
          :rules="rules"
          :label="$t('input-login-label')"
          @keyup.enter.native="loginAction"
        />
        <v-text-field
          v-model="password"
          :rules="rules"
          type="password"
          :label="$t('input-password-label')"
          @keyup.enter.native="loginAction"
        />
      <v-btn @click="loginAction">{{ $t('register-btn-text')}}</v-btn>
      <span>or <a @click="$router.push({name: 'Login'})">login</a></span>
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

  rules = [
    (v: string) => !!v || this.$t('required'),
  ]

  get isAuthorized(): boolean {
    return this.$store.direct.getters.sysUser.isAuthorized;
  }

  goHome() {
    this.$router.push({ name: 'Home' });
  }

  @Watch('isAuthorized')
  onUserAuthChange(isAuthorized: boolean) {
    if (isAuthorized) {
      this.goHome();
    }
  }

  validate(): boolean {
    // @ts-ignore
    this.$refs.form.validate();
    return !((!this.login) || (!this.password));
  }

  loginAction() {
    if (!this.validate()) {
      return;
    }
    this.$store.direct.dispatch.sysUser.register({ login: this.login, password: this.password });
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
    "input-login-label": "Enter login",
    "input-password-label": "Enter password",
    "register-btn-text": "register",
    "required": "required"
  },
  "ru": {
    "input-login-label": "Введите логин",
    "input-password-label": "Введите пароль",
    "register-btn-text": "Зарегистрироваться",
    "required": "Обязательно для заполнения"
  }
}
</i18n>
