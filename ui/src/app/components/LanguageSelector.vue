<template>
  <div class="locale-changer">
    <v-select v-model="lang" :items="langs" :label="language" dark filled>
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

  language = 'language'

  lang = 'en'

  defaultLang = 'en'

  @Watch('lang')
  onChange(lang: string) {
    if (lang === 'ru') {
      this.language = 'язык';
    }

    if (this.lang === this.getLangFromQuery()) {
      return;
    }
    this.$router.push({ query: { ...this.$route.query, lang } });
    this.lang = this.getLangFromQuery();
    this.$i18n.locale = lang;
    this.$vuetify.lang.current = lang;
  }

  getLangFromQuery(): string {
    const { lang } = this.$route.query;
    if (!lang) {
      this.$router.push({ query: { ...this.$route.query, lang: this.defaultLang } });
      return this.defaultLang;
    }

    // eslint-disable-next-line @typescript-eslint/ban-ts-ignore
    // @ts-ignore
    if (this.langs.includes(lang)) {
      // eslint-disable-next-line @typescript-eslint/ban-ts-ignore
      // @ts-ignore
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
