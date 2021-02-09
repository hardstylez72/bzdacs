<script lang="ts">
import {
  Component,
} from 'vue-property-decorator';

import SimpleEntityForm from '@/views/base/components/SimpleEntityForm.vue';

@Component
export default class SystemForm extends SimpleEntityForm<System> {}
</script>

<i18n>
{
  "en": {
    "label": {
      "name": "Enter system name"
    },
    "required": "Required"
  },
  "ru": {
    "label": {
      "name": "Название системы"
    },
    "required": "Обязательное поле"
  }
}
</i18n>
