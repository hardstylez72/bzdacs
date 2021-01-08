<template>
  <div>
    <v-form ref="refss" v-model="valid" lazy-validation>
      <v-row>
        <v-col cols="12" sm="10" md="10">
          <v-textarea
            v-model="tag.name"
            outlined
            required
            :rules="nameRules"
            label="Название"
          />
        </v-col>
      </v-row>
    </v-form>
    <slot name="actions" v-bind="{ ref: this.$refs.refss }"/>
  </div>

</template>

<script lang="ts">
import {
  Component, Model, Prop, Vue, Watch,
} from 'vue-property-decorator';
import { Route } from '@/views/route/service';
import { Tag } from '@/views/tag/service';

@Component
export default class TagForm extends Vue {
  valid = true

  tag: Tag = {
    id: -1,
    name: '',
  }

  @Model('change', { default: {} })
  readonly value!: Tag

  @Watch('value', { immediate: true })
  protected onChangeValue(route: Tag): void {
    this.tag = route;
  }

  @Watch('route', { deep: true })
  protected onChangeRoute(route: Route): void {
    if (JSON.stringify(route) !== JSON.stringify(this.value)) {
      this.$emit('change', route);
    }
  }

  nameRules = [
    (v: string) => !!v || 'Обязательное поле',
  ]
}
</script>

<style scoped lang="scss">

</style>
