<template>
  <div>
    <v-form ref="refss" v-model="valid" lazy-validation>
      <v-row>
        <v-col cols="12" sm="10" md="10">
          <v-text-field
            v-model="tag.name"
            outlined
            required
            :rules="nameRules"
            :label="$t('label.tag')"
          />
        </v-col>
      </v-row>
    </v-form>
    <slot name="actions" v-bind="{ ref: this.$refs.refss }"/>
  </div>

</template>

<script lang="ts">
import {
  Component, Model, Vue, Watch,
} from 'vue-property-decorator';
import { Route } from '@/route/entity';
import { Tag } from '@/tag/entity';

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
    (v: string) => !!v || this.$t('required'),
  ]
}
</script>

<i18n>
{
  "en": {
    "label": {
      "tag": "Tag"
    },
    "required": "Required"
  },
  "ru": {
    "label": {
      "tag": "Тег"
    },
    "required": "Обязательное поле"
  }
}
</i18n>
