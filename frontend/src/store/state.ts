import { SharedState } from "./modules/shared";
import { SearchState } from "./modules/search";
import { ArtistsState } from "./modules/artists";

interface State extends SharedState, SearchState, ArtistsState {}

export default State;
