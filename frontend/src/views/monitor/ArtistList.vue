<template>
  <v-row dense v-if="!!results">
    <v-col cols="6" sm="3" lg="2" xl="2" v-for="item in results" :key="item.id">
      <v-card class="result-container" :min-height="item.minHeight">
        <img class="cover-art" referrerpolicy="no-referrer" :src="item.image" />
        <div class="text-block">
          <v-card-title class="card-title">{{ item.title }}</v-card-title>
          <v-card-subtitle v-if="item.subtitle" class="card-subtitle">{{
            item.subtitle
          }}</v-card-subtitle>
        </div>
      </v-card>
    </v-col>
  </v-row>
</template>

<script lang="ts">
import Vue from "vue";

export default Vue.extend({
  computed: {
    results(): any[] {
      let values = this.$store.getters.artists;

      if (!values || !values.length) {
        return [];
      }

      return values.map((v: any) => {
        const thumbnail = v.thumbnails?.length
          ? v.thumbnails.sort((a: any, b: any) => a.height > b.height)[0]
          : null;

        const item = {
          id: v.id,
          title: v.name || v.title,
          image: thumbnail?.url,
          minHeight: thumbnail?.url ? null : "100%",
          subtitle: null,
        };
        if (v.artists) {
          item.subtitle = v.artists.map((a: any) => a.name).join("/");
        }
        return item;
      });
    },
  },
});
</script>

<style lang="scss" scopes>
.card-title,
.card-subtitle {
  color: white;
}
.result-container {
  background-color: black !important;
  position: relative;
}
.text-block {
  position: absolute;
  bottom: 0;
  left: 0;
  text-shadow: 2px 2px black;
}
.cover-art {
  vertical-align: middle;
  opacity: 0.7;
  width: 100%;
  height: 100%;
}
</style>
