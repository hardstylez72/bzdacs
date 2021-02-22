<script lang="ts">
import {
  Component,
} from 'vue-property-decorator';

import SimpleEntityForm from '@/views/base/components/SimpleEntityForm.vue';
import { Namespace } from '@/views/namespace/entity';

@Component
export default class SystemForm extends SimpleEntityForm<Namespace> {}
</script>

<i18n>
{
  "en": {
    "label": {
      "name": "Enter service name"
    },
    "required": "Required"
  },
  "ru": {
    "label": {
      "name": "Название сервиса"
    },
    "required": "Обязательное поле"
  }
}
</i18n>
