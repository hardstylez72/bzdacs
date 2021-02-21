<template>
  <Dialog v-model="show">
    <v-card>
      <v-card-title class="headline grey lighten-2">{{$t('title')}}</v-card-title>
      <v-card-text>
        <v-form ref="create-user-form" v-model="valid" lazy-validation>
          <v-row>
            <v-col cols="12" sm="4" md="4">
              <v-text-field
                v-model="user.externalId"
                required
                :rules="externalIdRules"
                :label="$t('label.external-id')"
              />
            </v-col>
          </v-row>
        </v-form>

        <v-card-actions>
          <v-spacer />
          <v-btn color="blue darken-1" text @click="close">{{$t('cancel')}}</v-btn>
          <v-btn color="blue darken-1" :disabled="disable" text @click="update">{{$t('update')}}</v-btn>
        </v-card-actions>
      </v-card-text>
    </v-card>
  </Dialog>
</template>

<script lang="ts">
import {
  Component, Vue, Model, Watch, Prop,
} from 'vue-property-decorator';
import { User } from '@/views/user/services/user';
import Dialog from '../../common/components/Dialog.vue';

@Component({
  components: {
    Dialog,
  },
})
export default class UpdateUserDialog extends Vue {
  show = false

  disable = true

  valid = false

  user: User = {
    id: -1,
    externalId: '',
  }

  initialUserState: User = this.user

  @Prop({ required: true }) id!: number

  @Watch('id')
  onIdChange(id: number) {
    this.$store.direct.dispatch.user.GetById(id).then((user) => {
      this.user = user;
      this.initialUserState.externalId = this.user.externalId;
      this.initialUserState.id = this.user.id;
    });
  }

  @Model('change', { default: false, type: Boolean })
  readonly value!: boolean

  @Watch('value')
  protected onChangeValue(value: boolean): void {
    this.show = value;
  }

  @Watch('user', { deep: true })
  onUserChange(u: User) {
    console.log(u);
    this.disable = false;
    if (this.sameUsers(this.initialUserState, u)) {
      this.disable = true;
      return;
    }

    if (!this.sameUsers(this.user, u)) {
      this.disable = false;
    }
  }

  sameUsers(a: User, b: User): boolean {
    return (b.externalId === a.externalId);
  }

  validate() {
    // eslint-disable-next-line @typescript-eslint/ban-ts-ignore
    // @ts-ignore
    this.$refs['create-user-form'].validate();
  }

  rules = [
    (v: string) => !!v || this.$t('required'),
]

  externalIdRules = this.rules

  async update() {
    this.validate();
    if (!this.valid) {
      return;
    }

    if (!this.user.externalId) {
      return;
    }

    await this.$store.direct.dispatch.user.Update(this.user);
    this.$emit('change', false);
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
    "add-btn": "Add user",
    "title": "User creation",
    "label": {
      "external-id": "External Id"
    },
    "cancel": "Cancel",
    "update": "Update",
    "required": "required"
  },
  "ru": {
    "add-btn": "Добавить пользователи",
    "title": "Создание пользователя",
    "label": {
      "external-id": "Идентфикатор"
    },
    "cancel": "Отмена",
    "update": "Обновить",
    "required": "Обязательное поле"
  }
}
</i18n>
