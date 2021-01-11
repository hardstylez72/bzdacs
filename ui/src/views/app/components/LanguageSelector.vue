<template>
  <div class="locale-changer">
    <v-select v-model="$i18n.locale" :items="langs" label="language" dark filled>
    </v-select>
  </div>
</template>

<script lang="ts">
import {
  Component, Vue,
} from 'vue-property-decorator';

@Component
export default class LanguageSelector extends Vue {
  langs: string[] = ['ru', 'en']
}
</script>

<style scoped lang="scss">

</style>
