import { createSelector } from "reselect";
import { AppState } from "reducks/store";

const userInfoSelector = (state: AppState) => {
  return state.userInfo;
};
export const getUserInfo = createSelector([userInfoSelector], (state) => state);
