<template>
  <div>
    <v-form ref="form" v-model="valid" lazy-validation>
      <v-row>
        <v-col cols="12" sm="10" md="10">
          <v-text-field
            v-model="local.name"
            outlined
            required
            :rules="rules"
            data-test="name-input"
            :label="$t('label.name')"
          />
        </v-col>
      </v-row>
    </v-form>
    <slot name="actions" v-bind="{ validate: () => {
      this.$refs.form.validate
    } }"/>
  </div>

</template>

<script lang="ts">
import {
  Component, Model, Vue, Watch,
} from 'vue-property-decorator';
import { Namespace } from '@/views/namespace/service';
import { SimpleEntity } from '@/views/base/services/entity';

@Component
export default class NamespaceForm<T extends SimpleEntity> extends Vue {
  valid = true

  local: T = {
    id: -1,
    name: '',
  }

  @Model('change', { default: {} })
  readonly value!: Namespace

  @Watch('value', { immediate: true })
  protected onChangeValue(prop: T): void {
    this.local = prop;
  }

  @Watch('local', { deep: true })
  protected onChangeRoute(local: T): void {
    if (JSON.stringify(local) !== JSON.stringify(this.value)) {
      this.$emit('change', local);
    }
  }

  rules = [
    (v: string) => !!v || this.$t('required'),
  ]
}
</script>

<i18n>
{
  "en": {
    "label": {
      "name": "Namespace"
    },
    "required": "Required"
  },
  "ru": {
    "label": {
      "name": "Пространство имен"
    },
    "required": "Обязательное поле"
  }
}
</i18n>
