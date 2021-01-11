<template>
  <div class="locale-changer">
    <v-select v-model="lang" :items="langs" label="language" dark filled>
    </v-select>
  </div>
</template>

<script lang="ts">
import {
  Component, Vue, Watch,
} from 'vue-property-decorator';

@Component
export default class LanguageSelector extends Vue {
  langs: string[] = ['ru', 'en']

  lang = 'en'

  defaultLang = 'en'

  @Watch('lang')
  onChange(lang: string) {
    if (this.lang === this.getLangFromQuery()) {
      return;
    }
    this.$router.push({ query: { ...this.$route.query, lang } });
    // eslint-disable-next-line no-undef
    window.location.reload();
  }

  getLangFromQuery(): string {
    const { lang } = this.$route.query;
    if (!lang) {
      this.$router.push({ query: { ...this.$route.query, lang: this.defaultLang } });
      return this.defaultLang;
    }

    if (this.langs.includes(lang)) {
      return lang;
    }

    this.$router.push({ query: { ...this.$route.query, lang: this.defaultLang } });
    return this.defaultLang;
  }

  created() {
   this.$i18n.locale = this.getLangFromQuery();
   this.lang = this.$i18n.locale;
  }
}
</script>

<style scoped lang="scss">

</style>
