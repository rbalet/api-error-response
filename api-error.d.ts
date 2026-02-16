import { ErrorCode } from './error-code.enum.ts'

export type ErrorType =
  | "AUTH"
  | "VALIDATION"
  | "DOMAIN"
  | "CONFLICT"
  | "NOT_FOUND"
  | "RATE_LIMIT"
  | "SYSTEM"
  | "API";

export type ApiError<Codes extends string = ErrorCodeBase> = ValidationError<Codes> | NonValidationError<Codes>;

export interface ApiErrorBase<Codes extends string = ErrorCode> {
  type?: ErrorType;
  code?: Codes;
  message?: string;
  traceId?: string;
  timestamp?: string;
}

export interface ValidationError<Codes extends string = ErrorCodeBase> extends ApiErrorBase<Codes> {
  type: "VALIDATION";
  issues?: ValidationIssue<Codes>[];
}

export interface NonValidationError<Codes extends string = ErrorCodeBase> extends ApiErrorBase<Codes> {
  type: Exclude<ErrorType, "VALIDATION">;
  issues?: never;
}

export interface ValidationIssue<Codes extends string = ErrorCode> {
  code?: Codes;
  path?: (string | number)[];
  message?: string;
  meta?: Record<string, unknown>;
}
