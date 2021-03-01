<template>
  <div>
    <span class="white--text" data-test="avatar">{{$t('user')}}: <b style="color: black">{{ login }}</b></span>
  </div>
</template>

<script lang="ts">
import {
  Component, Vue, Prop,
} from 'vue-property-decorator';

@Component
export default class Avatar extends Vue {
  @Prop({ type: String }) login!: string
}
</script>

<style scoped lang="scss">

</style>

<i18n>
{
  "en": {
    "user": "User"
  },
  "ru": {
    "user": "Пользователь"
  }
}
</i18n>
