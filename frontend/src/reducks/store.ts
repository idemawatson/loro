import { combineReducers, createStore, compose, applyMiddleware } from "redux";
import { UserInfoReducer } from "users/reducers";
import { UserInfo } from "users/types";
import thunk from "redux-thunk";

export type AppState = {
  userInfo: UserInfo;
};

const storeEnhancers =
  (window as any).__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ || compose;

const store = createStore(
  combineReducers<AppState>({
    userInfo: UserInfoReducer,
  }),
  storeEnhancers(applyMiddleware(thunk))
);

export default store;
