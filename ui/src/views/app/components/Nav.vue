<template>
  <div style="position: relative">
    <v-app-bar app color="primary" dark>
      <a @click="goHome"><h1 style="color: aliceblue">{{appName}}</h1></a>
      <LanguageSelector  class="ml-12 language-selector pt-4" />
      <GithubLink class="mx-4"/>
      <SwaggerLink class="mx-4"/>
      <v-spacer/>
      <Avatar v-if="login" :login="login"/>
      <v-btn data-test="logout" v-if="login" class="ms-5" @click="logout" >{{$t('logout')}}</v-btn>
    </v-app-bar>
  </div>
</template>

<script lang="ts">
import {
  Component, Vue,
} from 'vue-property-decorator';
import LanguageSelector from '@/views/app/components/LanguageSelector.vue';
import GithubLink from '@/views/app/components/GithubLink.vue';
import SwaggerLink from '@/views/app/components/SwaggerLink.vue';
import Avatar from '@/views/app/components/Avatar.vue';

@Component({
  components: {
    LanguageSelector,
    GithubLink,
    Avatar,
    SwaggerLink,
  },
})
export default class Nav extends Vue {
  drawer = false

  group = null

  appName = 'BZDACS'

  goHome() {
    this.$router.push({ name: 'Home', query: { lang: this.$route.query.lang ? this.$route.query.lang : 'en' } });
  }

  get isAuthorized(): boolean {
    return this.$store.direct.getters.isAuthorized;
  }

  get login(): string {
    return this.$store.direct.getters.login;
  }

  logout(): Promise<void> {
    return this.$store.direct.dispatch.logout();
  }
}
</script>

<style scoped lang="scss">
.language-selector {
  width: 100px;
}
</style>

<i18n>
{
  "en": {
    "logout": "logout"
  },
  "ru": {
    "logout": "Выйти"
  }
}
</i18n>
