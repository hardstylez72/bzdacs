<template>
  <div>
    <h2>{{ $t('user') }} {{ getUser.externalId }}</h2>
    <GroupsBelongUserTable :user-id="getUserId" />

    <UserRoutesTable :items="routes" >
      <template v-slot:top>
        <v-toolbar flat>
          <v-toolbar-title>{{ $t('titleRoute') }}</v-toolbar-title>
          <v-divider class="mx-4" inset vertical/>
          <v-spacer/>
          <AddRoutesButton :user-id="userIdC" />
        </v-toolbar>
      </template>

      <template v-slot:item.actions="{ item }">
        <div class="d-flex">
          <UpdateRouteButton :item="item" :user="getUser"/>
          <OverwriteRouteButton :item="item" :user="getUser"/>
          <DeleteRouteButton :item="item" :user="getUser"/>
        </div>
      </template>

      <template v-slot:item.statuses="{ item }" >
        <div class="d-flex">
          <StatusIconRouteAccess :item="item"/>
          <StatusIconRouteOverwritten :item="item"/>
        </div>
      </template>

    </UserRoutesTable>
  </div>
</template>

<script lang="ts">
import {
  Component, Vue,
} from 'vue-property-decorator';
import { Group } from '@/views/group/entity';
import { User } from '@/views/user/entity';
import { RouteWithGroups } from '@/views/user/services/userroute';
import UserRoutesSelectableTable from '../components/UserRoutesSelectableTable.vue';
import AddRoutesButton from '../components/AddRoutesButton.vue';
import UpdateRouteButton from '../components/UpdateRouteButton.vue';
import StatusIconRouteOverwritten from '../components/StatusIconRouteOverwritten.vue';
import StatusIconRouteAccess from '../components/StatusIconRouteAccess.vue';
import OverwriteRouteButton from '../components/OverwriteRouteButton.vue';
import DeleteRouteButton from '../components/DeleteRouteButton.vue';
import UserRoutesTable from '../components/UserRoutesTable.vue';
import GroupsBelongUserTable from '../components/GroupsBelongUserTable.vue';

@Component({
  components: {
    UserRoutesSelectableTable,
    AddRoutesButton,
    UpdateRouteButton,
    StatusIconRouteOverwritten,
    StatusIconRouteAccess,
    OverwriteRouteButton,
    DeleteRouteButton,
    UserRoutesTable,
    GroupsBelongUserTable,
},
})
export default class UserPage extends Vue {
  userId = Number(this.$route.params.id);

  namespaceId = Number(this.$route.query.namespaceId);

  user: User = {
    externalId: 'Не найден',
    id: -1,
  }

  titleRoute = 'Маршруты'

  selectedGroups: Group[] = []

  async created() {
    this.user = await this.$store.direct.dispatch.user.GetById({ id: this.userId, namespaceId: this.namespaceId });
  }

  get getUser(): User {
    return this.user;
  }

  get getUserId(): number {
    return this.userId;
  }

  get showDeleteGroupButton(): boolean {
    return this.selectedGroups.length > 0;
  }

  async deleteSelectedGroups() {
    const groups = this.selectedGroups;
    const params = groups.map((group) => ({
      groupId: group.id,
      userId: this.userId,
    }));

    await this.$store.direct.dispatch.userGroup.Delete(params);
    this.selectedGroups = [];
    await this.$store.direct.dispatch.userRoute.GetListNotBelongToUser(this.userId);
    await this.$store.direct.dispatch.userRoute.GetListBelongToUser(this.userId);
  }

  get groups(): readonly Group[] {
    return this.$store.direct.getters.userGroup.getGroupsBelongToUser;
  }

  get routes(): readonly RouteWithGroups[] {
    return this.$store.direct.getters.userRoute.getRoutesBelongToUser;
  }
}
</script>

<i18n>
{
  "en": {
    "no-data": "No data",
    "user": "User",
    "titleGroup": "Groups",
    "titleRoute": "Routes",
    "delete-group-btn": "Delete selected groups"
  },
  "ru": {
    "no-data": "Нет данных",
    "user": "Пользователь",
    "titleGroup": "Группы",
    "titleRoute": "Маршруты",
    "delete-group-btn": "Удалить выбранные группы"
  }
}
</i18n>
