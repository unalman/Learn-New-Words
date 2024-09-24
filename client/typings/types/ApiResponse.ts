export type APIResponse<T> = {
  data: T;
  error: string;
  success: boolean;
};
