// types.ts
import { Option } from 'fp-ts/lib/Option';

export type UserId = string;
export type User = {
  id: UserId;
  name: string;
  email: string;
  role: string;
};

export type Users = Record<UserId, User>;

export type PostId = string;
export type Post = {
  id: PostId;
  title: string;
  content: string;
  author: UserId;
};

export type Posts = Record<PostId, Post>;

export type CommentId = string;
export type Comment = {
  id: CommentId;
  content: string;
  author: UserId;
  post: PostId;
};

export type Comments = Record<CommentId, Comment>;

export type DashboardState = {
  users: Users;
  posts: Posts;
  comments: Comments;
};

export type DashboardAction =
  | { type: 'ADD_USER'; payload: User }
  | { type: 'ADD_POST'; payload: Post }
  | { type: 'ADD_COMMENT'; payload: Comment }
  | { type: 'UPDATE_USER'; payload: User }
  | { type: 'UPDATE_POST'; payload: Post }
  | { type: 'UPDATE_COMMENT'; payload: Comment }
  | { type: 'DELETE_USER'; payload: UserId }
  | { type: 'DELETE_POST'; payload: PostId }
  | { type: 'DELETE_COMMENT'; payload: CommentId };

export type DashboardReducer = (state: DashboardState, action: DashboardAction) => DashboardState;