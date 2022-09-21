import { Module } from "vuex";
import State from "../state";
import { performSearch, addArtist } from "@/server/requests";

export interface SearchState {
  data: {
    lastSearch: string;
    artists: any[];
    albums: any[];
    tracks: any[];
    playlists: any[];
  };
}

const sharedModule: Module<SearchState, State> = {
  state: {
    data: {
      lastSearch: "",
      artists: [],
      albums: [],
      tracks: [],
      playlists: [],
    },
  },
  getters: {
    searchedArtists: (state) => state.data.artists,
    searchedAlbums: (state) => state.data.albums,
    searchedTracks: (state) => state.data.tracks,
    searchedPlaylists: (state) => state.data.playlists,
  },
  actions: {
    async add(state, { id, monitor, type }) {
      state.dispatch("setLoading", true);
      await state.dispatch("clearError");

      const error = await addArtist(id, monitor);

      if (error) {
        state.dispatch("setLoading", false);
        await state.dispatch("setError", error);
      } else {
        state.dispatch("setLoading", false);
      }
    },
    async search(state, { searchTerm, type }) {
      if (searchTerm == state.state.data.lastSearch) {
        switch (type) {
          case "artist":
            if (state.state.data.artists.length) {
              return;
            }
            break;
          case "album":
            if (state.state.data.albums.length) {
              return;
            }
            break;
          case "track":
            if (state.state.data.tracks.length) {
              return;
            }
            break;
          case "playlist":
            if (state.state.data.playlists.length) {
              return;
            }
            break;
        }
      } else {
        state.state.data.artists = [];
        state.state.data.albums = [];
        state.state.data.tracks = [];
        state.state.data.playlists = [];
      }

      state.dispatch("setLoading", true);
      await state.dispatch("clearError");

      const { error, data } = await performSearch(searchTerm, type);

      if (error) {
        state.dispatch("setLoading", false);
        await state.dispatch("setError", error);
      } else {
        state.commit("setResult", { data, type, searchTerm });
        state.dispatch("setLoading", false);
      }
    },
  },
  mutations: {
    setResult(state, { data, type, searchTerm }) {
      state.data.lastSearch = searchTerm;
      switch (type) {
        case "artist":
          state.data.artists = data;
          break;
        case "album":
          state.data.albums = data;
          break;
        case "track":
          state.data.tracks = data;
          break;
        case "playlist":
          state.data.playlists = data;
          break;
      }
    },
  },
};

export default sharedModule;
