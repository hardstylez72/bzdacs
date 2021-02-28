<template>
  <div>
    <h2>{{ $t('user') }} {{ getUser.externalId }}</h2>
    <GroupsBelongUserTable :user-id="getUserId" />
    <RoutesBelongUserTable :user-id="getUserId" />
  </div>
</template>

<script lang="ts">
import {
  Component, Vue,
} from 'vue-property-decorator';
import { User } from '@/views/user/entity';
import GroupsBelongUserTable from '../components/GroupsBelongUserTable.vue';
import RoutesBelongUserTable from '../components/RoutesBelongUserTable.vue';

@Component({
  components: {
    GroupsBelongUserTable,
    RoutesBelongUserTable,
},
})
export default class UserPage extends Vue {
  userId = Number(this.$route.params.id);

  namespaceId = Number(this.$route.query.namespaceId);

  user: User = {
    externalId: this.$t('unknown-user').toString(),
    id: -1,
  }

  async created() {
    this.user = await this.$store.direct.dispatch.user.GetById({ id: this.userId, namespaceId: this.namespaceId });
  }

  get getUser(): User {
    return this.user;
  }

  get getUserId(): number {
    return this.userId;
  }
}
</script>

<i18n>
{
  "en": {
    "unknown-user": "unknown user",
    "user": "User"
  },
  "ru": {
    "unknown-user": "Пользователь не найден",
    "user": "Пользователь"
  }
}
</i18n>
