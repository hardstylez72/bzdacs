<template>
  <Dialog v-model="show">
    <template v-slot:activator="props">
      <v-btn color="primary" class="mb-2" v-bind="props" v-on="props.on">{{$t('add-btn')}}</v-btn>
    </template>

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
          <v-btn color="blue darken-1" text @click="create">{{$t('save')}}</v-btn>
        </v-card-actions>
      </v-card-text>
    </v-card>
  </Dialog>
</template>

<script lang="ts">
import {
  Component, Vue,
} from 'vue-property-decorator';
import { User } from '@/user/entity';
import Dialog from '../../common/components/Dialog.vue';

@Component({
  components: {
    Dialog,
  },
})
export default class CreateDialog extends Vue {
  show = false

  valid = false

  namespaceId = Number(this.$route.query.namespaceId);

  user: User = {
    id: -1,
    externalId: '',
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

  async create() {
    this.validate();
    if (!this.valid) {
      return;
    }

    if (!this.user.externalId) {
      return;
    }
    await this.$store.direct.dispatch.user.Create({
      namespaceId: this.namespaceId,
      externalId: this.user.externalId,
    });
    this.$emit('userCreated');
    this.user.externalId = '';
    this.show = false;
  }

  close() {
    this.show = false;
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
    "save": "Save",
    "required": "required"
  },
  "ru": {
    "add-btn": "Добавить пользователи",
    "title": "Создание пользователя",
    "label": {
      "external-id": "Идентфикатор"
    },
    "cancel": "Отмена",
    "save": "Создать",
    "required": "Обязательное поле"
  }
}
</i18n>
