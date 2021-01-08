<template>
  <div>
    <v-data-table
      v-model="selected"
      :headers="headers"
      :items="items"
      sort-by="calories"
      class="elevation-1"
      show-select
    >
      <template v-slot:no-data>
        Нет данных
      </template>

      <template v-slot:top>
        <slot name="top" />
      </template>

      <template v-slot:item.actions="{ item }">
        <slot name="item.actions" v-bind="{ item }"/>
      </template>

      <template v-slot:item.statuses="{ item }">
        <slot name="item.statuses" v-bind="{ item }"/>
      </template>

      <template v-slot:item.method="{ item }">
        <HttpMethodBox :method="item.method"></HttpMethodBox>
      </template>

    </v-data-table>
  </div>
</template>

<script lang="ts">
import {
  Component, Model, Prop, Vue, Watch,
} from 'vue-property-decorator';
import { DataTableHeader } from 'vuetify';
import HttpMethodBox from './HttpMethodBox.vue';

@Component({
  components: {
    HttpMethodBox,
  },
})
export default class SelectableTable<T> extends Vue {
  @Prop({ default: () => ([]), type: Array }) items!: Array<T>

  protected selected: T[] = []

  @Model('change', { default: () => ([]), type: Array })
  readonly value!: []

  @Watch('value')
  protected onChangeValue(value: T[]): void {
    this.selected = value;
  }

  @Watch('selected')
  protected onChangeIsShowDialog(selected: T[]): void {
      this.$emit('change', selected);
  }

  protected headers: DataTableHeader[] = []
}
</script>

<style scoped lang="scss">
</style>
