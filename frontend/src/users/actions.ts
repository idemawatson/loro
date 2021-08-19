import actionCreatorFactory from "typescript-fsa";

const actionCreator = actionCreatorFactory();

export const getUserInfoAction = actionCreator<string>("GET_USER");
