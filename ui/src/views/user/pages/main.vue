<template>
  <div>
    <h2>Пользователь {{userC.externalId}}</h2>
    <UserGroupsSelectableTable v-model="selectedGroups" :items="groups">
      <template v-slot:top>
        <v-toolbar flat>
          <v-toolbar-title>{{ titleGroup }}</v-toolbar-title>
          <v-divider class="mx-4" inset vertical/>
          <v-spacer/>
          <div>
            <v-btn
              v-if="showDeleteGroupButton"
              color="primary"
              class="mb-2"
              @click="deleteSelectedGroups"
            >
              Удалить выбранные группы
            </v-btn>
          </div>
          <AddGroupsButton :user-id="userIdC" />
        </v-toolbar>
      </template>
    </UserGroupsSelectableTable>

    <UserRoutesTable :items="routes" >
      <template v-slot:top>
        <v-toolbar flat>
          <v-toolbar-title>{{ titleRoute }}</v-toolbar-title>
          <v-divider class="mx-4" inset vertical/>
          <v-spacer/>
          <AddRoutesButton :user-id="userIdC" />
        </v-toolbar>
      </template>

      <template v-slot:item.actions="{ item }">
        <div class="d-flex">
          <UpdateRouteButton :item="item" :user="userC"/>
          <OverwriteRouteButton :item="item" :user="userC"/>
          <DeleteRouteButton :item="item" :user="userC"/>
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
import { Group } from '@/views/group/services/group';
import { User } from '@/views/user/services/user';
import { RouteWithGroups } from '@/views/user/services/userroute';
import UserGroupsSelectableTable from '../components/UserGroupsSelectableTable.vue';
import UserRoutesSelectableTable from '../components/UserRoutesSelectableTable.vue';
import AddGroupsButton from '../components/AddGroupsButton.vue';
import AddRoutesButton from '../components/AddRoutesButton.vue';
import UpdateRouteButton from '../components/UpdateRouteButton.vue';
import StatusIconRouteOverwritten from '../components/StatusIconRouteOverwritten.vue';
import StatusIconRouteAccess from '../components/StatusIconRouteAccess.vue';
import OverwriteRouteButton from '../components/OverwriteRouteButton.vue';
import DeleteRouteButton from '../components/DeleteRouteButton.vue';
import UserRoutesTable from '../components/UserRoutesTable.vue';

@Component({
  components: {
    UserGroupsSelectableTable,
    UserRoutesSelectableTable,
    AddGroupsButton,
    AddRoutesButton,
    UpdateRouteButton,
    StatusIconRouteOverwritten,
    StatusIconRouteAccess,
    OverwriteRouteButton,
    DeleteRouteButton,
    UserRoutesTable,
},
})
export default class UserPage extends Vue {
  userId = Number(this.$route.params.id);

  user: User = {
    externalId: 'Не найден',
    isSystem: false,
    description: '',
    id: -1,
  }

  titleGroup = 'Группы'

  titleRoute = 'Маршруты'

  selectedGroups: Group[] = []

  async mounted() {
    this.$store.direct.dispatch.userGroup.GetListBelongToUser(this.userId);
    this.$store.direct.dispatch.userRoute.GetListBelongToUser(this.userId);
    this.$store.direct.dispatch.user.GetById(this.userId).then((user) => {
       this.user = user;
     });
  }

  get userC(): User {
    return this.user;
  }

  get userIdC(): number {
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

<style scoped lang="scss">

</style>
