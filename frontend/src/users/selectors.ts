import { createSelector } from "reselect";
import { State } from "reducks/reducer";

const usersSelector = (state: State) => state.userInfo;

export const getUserInfo = createSelector([usersSelector], (state) => state);
