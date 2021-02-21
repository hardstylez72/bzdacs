<template>
  <v-dialog
    v-model="show"
    :max-width="maxWidth"
    persistent
    v-bind="$attrs"
    v-on="$listeners"
  >
    <template v-slot:activator="props">
      <slot name="activator" v-bind="props"/>
    </template>

    <template v-if="isShowDialogContent" v-slot:default>
      <slot name="default" v-bind="{ hideDialog }"/>
    </template>

  </v-dialog>
</template>

<script lang="ts">
import {
  Component, Model, Prop, Vue, Watch,
} from 'vue-property-decorator';

@Component({ inheritAttrs: false })
export default class RoutesTab extends Vue {
  protected show = false;

  protected isShowDialogContent = false;

  @Prop({ type: String, default: '800px' }) readonly maxWidth?: string

  @Model('change', { default: false, type: Boolean })
  readonly value!: boolean

  @Watch('value')
  protected onChangeValue(value: boolean): void {
    this.show = value;
  }

  @Watch('show')
  protected onChangeIsShowDialog(show: boolean): void {
    if (show !== this.value) {
      this.$emit('change', show);
    }

    if (show) {
      this.isShowDialogContent = true;
    } else {
      setTimeout(() => {
        this.isShowDialogContent = false;
      }, 250);
    }
  }

protected hideDialog(): void {
  this.show = false;
}

  close() {
    this.show = false;
  }
}
</script>

<style scoped lang="scss">

</style>
