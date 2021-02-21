<template>
  <div v-if="method" :style="getStyle(method)" class="">
    {{method}}
  </div>
</template>

<script lang="ts">
import {
  Component, Prop, Vue,
} from 'vue-property-decorator';

@Component

export default class CreateRouteDialog extends Vue {
  @Prop({}) method!: string

  getStyle(method: string): string {
    let style = 'border-radius: 5px; '
      + 'color: white; '
      + 'text-align: center;'
      + ' font-family: \'Roboto\';';
    switch (method.toLowerCase()) {
      case 'post':
        style += 'background: #4FC3F7';
        break;
      case 'get':
        style += 'background:#81C784';
        break;
      case 'put':
        style += 'background: #BA68C8;';
        break;
      case 'delete':
        style += 'background: #BA68C8;';
        break;
      case 'patch':
        style += 'background: #FFCCBC;';
        break;
      default:
        style = '';
    }
    return style;
  }
}
</script>

<style scoped lang="scss">

</style>
