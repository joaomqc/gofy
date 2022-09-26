import { Module } from "vuex";
import State from "../state";
import { performSearch, addArtist } from "@/server/requests";

export interface SearchState {
  lastSearch: string;
  searchedArtists: any[];
  searchedAlbums: any[];
  searchedTracks: any[];
  searchPlaylists: any[];
}

const sharedModule: Module<SearchState, State> = {
  state: {
    lastSearch: "",
    searchedArtists: [],
    searchedAlbums: [],
    searchedTracks: [],
    searchPlaylists: [],
  },
  getters: {
    searchedArtists: (state) => state.searchedArtists,
    searchedAlbums: (state) => state.searchedAlbums,
    searchedTracks: (state) => state.searchedTracks,
    searchedPlaylists: (state) => state.searchPlaylists,
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
      if (searchTerm == state.state.lastSearch) {
        switch (type) {
          case "artist":
            if (state.state.searchedArtists.length) {
              return;
            }
            break;
          case "album":
            if (state.state.searchedAlbums.length) {
              return;
            }
            break;
          case "track":
            if (state.state.searchedTracks.length) {
              return;
            }
            break;
          case "playlist":
            if (state.state.searchPlaylists.length) {
              return;
            }
            break;
        }
      } else {
        state.state.searchedArtists = [];
        state.state.searchedAlbums = [];
        state.state.searchedTracks = [];
        state.state.searchPlaylists = [];
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
      state.lastSearch = searchTerm;
      switch (type) {
        case "artist":
          state.searchedArtists = data;
          break;
        case "album":
          state.searchedAlbums = data;
          break;
        case "track":
          state.searchedTracks = data;
          break;
        case "playlist":
          state.searchPlaylists = data;
          break;
      }
    },
  },
};

export default sharedModule;
