<template>
  <div>
    <artist-toolbar />
    <v-container fluid class="my-4">
      <artist-list />
    </v-container>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import ArtistList from "./ArtistList.vue";
import ArtistToolbar from "./ArtistToolbar.vue";

export default Vue.extend({
  components: {
    ArtistList,
    ArtistToolbar,
  },
  beforeMount() {
    this.$store.dispatch("getArtists");
  },
});
</script>
