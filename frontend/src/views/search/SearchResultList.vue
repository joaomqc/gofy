<template>
  <v-row dense v-if="!!results">
    <v-col cols="6" sm="3" lg="2" xl="2" v-for="item in results" :key="item.id">
      <v-dialog transition="dialog-bottom-transition" max-width="600">
        <template v-slot:activator="{ on, attrs }">
          <v-card
            class="result-container"
            v-bind="attrs"
            v-on="on"
            @click="selectItem(item)"
            :min-height="item.minHeight"
          >
            <img
              class="cover-art"
              referrerpolicy="no-referrer"
              :src="item.image"
            />
            <div class="text-block">
              <v-card-title class="card-title">{{ item.title }}</v-card-title>
              <v-card-subtitle v-if="item.subtitle" class="card-subtitle">{{
                item.subtitle
              }}</v-card-subtitle>
            </div>
          </v-card>
        </template>
        <template v-slot:default="dialog">
          <v-card>
            <v-toolbar dark>Download</v-toolbar>
            <v-card-text>
              <div class="text-h5 pt-6 pb-6">Download {{ type }}</div>
              <div class="text-body-1">
                {{ getDialogBody }}
              </div>
            </v-card-text>
            <v-card-actions class="justify-end">
              <v-btn text @click="dialog.value = false">Close</v-btn>
              <v-btn
                color="green darken-1"
                text
                @click="
                  dialog.value = false;
                  onAdd(false);
                "
                >Download</v-btn
              >
              <v-btn
                color="green darken-1"
                text
                @click="
                  dialog.value = false;
                  onAdd(true);
                "
                >Download and monitor</v-btn
              >
            </v-card-actions>
          </v-card>
        </template>
      </v-dialog>
    </v-col>
  </v-row>
</template>

<script lang="ts">
import Vue from "vue";

export default Vue.extend({
  props: {
    type: String,
  },
  data: () => ({
    selectedItem: null as any | null,
  }),
  computed: {
    results(): any[] {
      let values = [];
      switch (this.type) {
        case "artist":
          values = this.$store.getters.searchedArtists;
          break;
        case "album":
          values = this.$store.getters.searchedAlbums;
          break;
        case "track":
          values = this.$store.getters.searchedTracks;
          break;
        case "playlist":
          values = this.$store.getters.searchedPlaylists;
          break;
      }

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
    getDialogBody(): string {
      let text = "Do you want to download ";

      switch (this.type) {
        case "artist":
          text += `the discography of ${this.selectedItem?.title}`;
          break;
        case "album":
        case "track":
          text += `${this.selectedItem?.title} by ${this.selectedItem?.subtitle}`;
          break;
      }

      return text + "?";
    },
  },
  methods: {
    selectItem(item: any): void {
      this.selectedItem = item;
    },
    async onAdd(monitor: boolean) {
      await this.$store.dispatch("add", {
        id: this.selectedItem.id,
        monitor,
        type: this.type,
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
