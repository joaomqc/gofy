<template>
  <v-app-bar dense>
    <v-btn icon @click="refresh">
      <v-icon>mdi-refresh</v-icon>
    </v-btn>
  </v-app-bar>
</template>

<script lang="ts">
import Vue from "vue";

export default Vue.extend({
  methods: {
    refresh() {
      this.$store.dispatch("getArtists");
    },
  },
});
</script>
