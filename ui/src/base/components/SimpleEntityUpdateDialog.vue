<template>
  <Dialog v-model="show">
    <v-card>
      <v-card-title class="headline grey lighten-2">
        {{$t('editing')}}
      </v-card-title>
      <v-card-text>
        <SimpleEntityForm v-model="local">
          <template v-slot:actions="{validate}">
            <v-card-actions>
              <v-spacer />
              <v-btn color="blue darken-1" text @click="close">{{$t('cancel')}}</v-btn>
              <v-btn color="blue darken-1" text :disabled="disableUpdateButton" @click="update(validate)">{{$t('update')}}</v-btn>
            </v-card-actions>
          </template>
        </SimpleEntityForm>
      </v-card-text>
    </v-card>
  </Dialog>
</template>

<script lang="ts">
import {
  Component, Model, Prop, Vue, Watch,
} from 'vue-property-decorator';
import Dialog from '@/common/components/Dialog.vue';
import { SimpleEntity } from '@/base/services/entity';
import SimpleEntityForm from './SimpleEntityForm.vue';

@Component({
  components: {
    Dialog,
    SimpleEntityForm,
  },
})
export default class SimpleEntityUpdateDialog<T extends SimpleEntity> extends Vue {
  show = false

  disableUpdateButton = true

  valid = true

  @Prop({ required: true }) id!: number

  @Model('change', { default: false, type: Boolean })
  readonly value!: boolean

  @Watch('value')
  protected onChangeValue(value: boolean): void {
    this.show = value;
  }

  local: SimpleEntity = {
    name: '',
    id: -1,
  }

   propInitialState: SimpleEntity = this.local

  @Watch('id', { immediate: true })
   protected async onChangeItem(id: number) {
    const item = await this.loadItemWithId(id);

      this.local = {
        name: item.name,
        id: item.id,
      };
      this.propInitialState = {
        name: item.name,
        id: item.id,
      };
  }

  // @ts-ignore
  async loadItemWithId(id: number): Promise<T> {
    console.error('override this method!');
  }

  @Watch('local', { deep: true })
  protected localChanged(local: T): void {
    if (!local.name) {
      this.disableUpdateButton = true;
      return;
    }
    if (this.propInitialState) {
      this.disableUpdateButton = this.areSame(this.propInitialState, local);
    }
  }

  async update(validate: any) {
    if (validate) {
      validate(); // vee-validate specifics
    }

    if (!this.local.id || !this.local.name) {
      return;
    }

    await this.updateValid(this.local);
    this.$emit('change', false);
    this.disableUpdateButton = true;
    this.show = false;
  }

  // @ts-ignore
  async updateValid(entity: SimpleEntity) {
    console.warn('override this method!');
  }

  areSame(a: SimpleEntity, b: SimpleEntity): boolean {
    return (a.name === b.name);
  }

  close() {
    this.$emit('change', false);
    this.show = false;
  }
}
</script>
<i18n>
{
  "en": {
    "editing": "Tag editing",
    "cancel": "Cancel",
    "update": "Update"
  },
  "ru": {
    "editing": "Редактирование тега",
    "cancel": "Отмена",
    "update": "Обновить"
  }
}
</i18n>
