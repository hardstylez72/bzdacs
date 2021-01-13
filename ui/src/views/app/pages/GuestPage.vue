<template>
  <v-form class="guest-login-form" ref="form">
      <v-text-field
        v-model="login"
        required
        data-test="login-input"
        :label="$t('input-label')"
        :rules="codeRules"
      />
    <v-btn data-test="login-btn" @click="loginAction" >{{ $t('login-btn-text') }}</v-btn>
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

  codeRules = [
    (v: string) => !!v || this.$t('required'),
  ]

  @Watch('isAuthorized')
  onUserAuthChange(isAuthorized: boolean) {
    console.log('isAuthorized', isAuthorized);
    if (isAuthorized) {
      this.$router.push({ name: 'Home', query: { tab: 'routes', lang: this.$route.query.lang ? this.$route.query.lang : 'en' } });
    }
  }

  loginAction() {
    this.validate();
    if (!this.login) {
      return;
    }

    this.$store.direct.dispatch.guestLogin(this.login)
      .finally(() => {
        if (this.isAuthorized) {
          this.$router.push({ name: 'Home', query: { tab: 'routes', lang: this.$route.query.lang ? this.$route.query.lang : 'en' } });
        }
      });
  }

  validate() {
    // eslint-disable-next-line @typescript-eslint/ban-ts-ignore
    // @ts-ignore
    this.$refs.form.validate();
  }
}
</script>

<style scoped lang="scss">
.guest-login-form {
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
    "input-label": "Enter login",
    "login-btn-text": "login",
    "required": "required"
  },
  "ru": {
    "input-label": "Введите логин",
    "login-btn-text": "войти",
    "required": "Обязательное поле"
  }
}
</i18n>
