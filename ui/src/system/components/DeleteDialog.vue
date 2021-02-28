<script lang="ts">
import {
  Component,
} from 'vue-property-decorator';
import { System } from '@/system/entity';
import SimpleEntityDeleteDialog from '@/base/components/SimpleEntityDeleteDialog.vue';

@Component
export default class DeleteSystemDialog extends SimpleEntityDeleteDialog<System> {
  async deleteEntity(id: number) {
    // must be overwritten!
    await this.$store.direct.dispatch.system.Delete(id);
    this.$emit('itemDeleted', id);
  }
}
</script>

<i18n>
{
  "en": {
    "sure": "Are you sure want to delete that system?",
    "cancel": "Cancel",
    "ok": "OK"
  },
  "ru": {
    "sure": "Вы уверены что хотите удлить систему?",
    "cancel": "Отмена",
    "ok": "ОК"
  }
}
</i18n>
