<template>
  <div>
    <v-avatar color="#1B5E20">
      <span class="white--text" data-test="avatar">{{ login }}</span>
    </v-avatar>
  </div>
</template>

<script lang="ts">
import {
  Component, Vue, Prop,
} from 'vue-property-decorator';

@Component
export default class Avatar extends Vue {
  @Prop({ type: String }) login!: string
}
</script>

<style scoped lang="scss">

</style>
