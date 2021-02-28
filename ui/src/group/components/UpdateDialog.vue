<template>
  <Dialog v-model="show">
    <v-card>
      <v-card-title class="headline grey lighten-2">{{$t('editing')}}</v-card-title>
      <v-card-text>
        <v-form ref="create-group-form" v-model="valid" lazy-validation>
          <v-row>
            <v-col cols="12" sm="4" md="4">
              <v-text-field
                v-model="group.code"
                required
                :rules="codeRules"
                :label="$t('label.code')"
              />
            </v-col>
            <v-col cols="12" sm="10" md="10">
              <v-textarea
                v-model="group.description"
                outlined
                required
                :rules="descriptionRules"
                :label="$t('label.desc')"
              />
            </v-col>
          </v-row>
        </v-form>

        <v-card-actions>
          <v-spacer />
          <v-btn color="blue darken-1" text @click="close">{{$t('cancel')}}</v-btn>
          <v-btn color="blue darken-1" text :disabled="disable" @click="updateGroup">{{$t('update')}}</v-btn>
        </v-card-actions>
      </v-card-text>
    </v-card>
  </Dialog>
</template>

<script lang="ts">
import {
  Component, Model, Prop, Vue, Watch,
} from 'vue-property-decorator';
import { Group } from '@/group/entity';

@Component({
  components: {
    Dialog: () => import('../../common/components/Dialog.vue'),
  },

})
export default class UpdateGroupDialog extends Vue {
  disable = false

  @Prop({ required: true }) id!: number

  @Model('change', { default: false, type: Boolean })
  readonly value!: boolean

  @Watch('value')
  protected onChangeValue(value: boolean): void {
    this.show = value;
  }

  @Watch('id')
   protected async onChangeItem(groupId: number): Promise<void> {
    const namespaceId = Number(this.$route.query.namespaceId);
    const g = await this.$store.direct.dispatch.group.GetById({ namespaceId, id: groupId });
    this.group = g;
     Object.assign(this.initialGroupState, g);
  }

  @Watch('group', { deep: true })
  protected onChangeGroup(group: Group): void {
    this.disable = false;
    if (this.groupsSame(this.initialGroupState, group)) {
      console.log('grewgew');
      this.disable = true;
      return;
    }

    if (this.groupsSame(this.group, group)) {
      this.disable = false;
    }
  }

  groupsSame(a: Group, b: Group): boolean {
    return (a.description === b.description
      && a.code === b.code);
  }

  show = false

  valid = false

  group: Group = {
    description: '',
    id: 0,
    code: '',
  }

  initialGroupState: Group = this.group

    validate() {
    // eslint-disable-next-line @typescript-eslint/ban-ts-ignore
    // @ts-ignore
    this.$refs['create-group-form'].validate();
  }

  rules = [
    (v: string) => !!v || 'Обязательное поле',
]

  codeRules = this.rules

  descriptionRules = this.rules

  async updateGroup() {
    this.validate();
    if (!this.valid) {
      return;
    }

    if (!this.group.code || !this.group.description || !this.group.id) {
      return;
    }

    await this.$store.direct.dispatch.group.Update({
      id: this.group.id,
      namespaceId: this.group.namespaceId,
      code: this.group.code,
      description: this.group.description,
    });
    this.$emit('change', false);
    this.$emit('groupUpdated');
    this.show = false;
  }

  close() {
    this.show = false;
    this.$emit('change', false);
  }
}
</script>

<i18n>
{
  "en": {
    "editing": "Group editing",
    "label": {
      "code": "Code",
      "desc": "Description"
    },
    "cancel": "Cancel",
    "update": "Update"
  },
  "ru": {
    "editing": "Редактирование группы",
    "label": {
      "code": "Код",
      "desc": "Описание"
    },
    "cancel": "Отмена",
    "update": "Обновить"
  }
}
</i18n>
