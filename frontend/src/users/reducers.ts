import * as Actions from "users/actions";
import { Reducer } from "reducks/reducer";
import { UserInfo } from "users/types";

export const UserReducer = Reducer.case(
  Actions.fetchUserInfoAction,
  (state, payload: UserInfo) => ({
    ...state,
    user: { ...state.userInfo, ...payload },
  })
);
