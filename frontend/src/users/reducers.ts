import * as Actions from "users/actions";
import { Reducer } from "reducks/reducer";

export const UserReducer = Reducer.case(
  Actions.getUserInfoAction,
  (state, payload) => ({
    ...state,
    user: { ...state.user, userId: payload },
  })
);
