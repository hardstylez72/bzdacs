<template>
  <div style="opacity: 0.7">
    <h1>{{$t('welcome')}}</h1>
    <h2>{{$t('hint')}}</h2>
  </div>
</template>

<script lang="ts">
import {
  Component, Vue,
} from 'vue-property-decorator';

@Component
export default class Home extends Vue {

}
</script>

<style scoped lang="scss">

</style>
<i18n>
{
  "en": {
    "welcome": "Welcome to access control system BZDACS",
    "hint": "<- To start create a system and service"
  },
  "ru": {
    "welcome": "Добро пожаловать в систему управления доступом BZDACS",
    "hint": "<- Для начала работы создайте систему и сервис"
  }
}
</i18n>
