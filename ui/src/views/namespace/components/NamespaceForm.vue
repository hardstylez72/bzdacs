<template>
  <div>
    <v-form ref="form" v-model="valid" lazy-validation>
      <v-row>
        <v-col cols="12" sm="10" md="10">
          <v-textarea
            v-model="namespace.name"
            outlined
            required
            :rules="nameRules"
            :label="$t('label.namespace')"
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

@Component
export default class NamespaceForm extends Vue {
  valid = true

  namespace: Namespace = {
    id: -1,
    name: '',
  }

  @Model('change', { default: {} })
  readonly value!: Namespace

  @Watch('value', { immediate: true })
  protected onChangeValue(namespace: Namespace): void {
    this.namespace = namespace;
  }

  @Watch('namespace', { deep: true })
  protected onChangeRoute(namespace: Namespace): void {
    if (JSON.stringify(namespace) !== JSON.stringify(this.value)) {
      this.$emit('change', namespace);
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
      "namespace": "Namespace"
    },
    "required": "Required"
  },
  "ru": {
    "label": {
      "namespace": "Пространство имен"
    },
    "required": "Обязательное поле"
  }
}
</i18n>
