export interface UserInfo {
  userId: string;
  userName: string;
  numFollowing: number;
  numFollowed: number;
  numPost: number;
}

export interface getUserInfoRequest {
  userId: string;
}
