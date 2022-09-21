import { SharedState } from "./modules/shared";
import { SearchState } from "./modules/search";

interface State extends SharedState, SearchState {}

export default State;
