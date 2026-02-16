export type ApiResponse<T, Codes extends string = ErrorCodeBase> =
  | { data: T; error?: never }
  | { data?: never; error: ApiError<Codes> };
