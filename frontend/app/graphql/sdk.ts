import { GraphQLClient } from "graphql-request";
import * as Dom from "graphql-request/dist/types.dom";
import gql from "graphql-tag";

export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = {
  [K in keyof T]: T[K];
};
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & {
  [SubKey in K]?: Maybe<T[SubKey]>;
};
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & {
  [SubKey in K]: Maybe<T[SubKey]>;
};
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
  Date: any;
  Map: any;
  Time: any;
};

export type Mutation = {
  __typename?: "Mutation";
  flamingo?: Maybe<Scalars["String"]>;
  setCatched: Pokemon;
};

export type MutationSetCatchedArgs = {
  catched: Scalars["Boolean"];
  id: Scalars["Int"];
};

export type Pokemon = {
  __typename?: "Pokemon";
  catched: Scalars["Boolean"];
  id: Scalars["Int"];
  name: Scalars["String"];
  type: Array<Scalars["String"]>;
};

export type Query = {
  __typename?: "Query";
  flamingo?: Maybe<Scalars["String"]>;
  pokemon?: Maybe<Array<Pokemon>>;
  total: Scalars["Int"];
  totalCatched: Scalars["Int"];
};

export type QueryPokemonArgs = {
  catched?: InputMaybe<Scalars["Boolean"]>;
  ids?: InputMaybe<Array<Scalars["Int"]>>;
};

export type GetPokemonQueryVariables = Exact<{ [key: string]: never }>;

export type GetPokemonQuery = {
  __typename?: "Query";
  pokemon?: Array<{
    __typename?: "Pokemon";
    id: number;
    name: string;
    type: Array<string>;
    catched: boolean;
  }> | null;
};

export type SetCatchedMutationVariables = Exact<{
  id: Scalars["Int"];
  catched: Scalars["Boolean"];
}>;

export type SetCatchedMutation = {
  __typename?: "Mutation";
  setCatched: { __typename?: "Pokemon"; id: number };
};

export const GetPokemonDocument = gql`
  query getPokemon {
    pokemon {
      id
      name
      type
      catched
    }
  }
`;
export const SetCatchedDocument = gql`
  mutation setCatched($id: Int!, $catched: Boolean!) {
    setCatched(catched: $catched, id: $id) {
      id
    }
  }
`;

export type SdkFunctionWrapper = <T>(
  action: (requestHeaders?: Record<string, string>) => Promise<T>,
  operationName: string,
  operationType?: string
) => Promise<T>;

const defaultWrapper: SdkFunctionWrapper = (
  action,
  _operationName,
  _operationType
) => action();

export function getSdk(
  client: GraphQLClient,
  withWrapper: SdkFunctionWrapper = defaultWrapper
) {
  return {
    getPokemon(
      variables?: GetPokemonQueryVariables,
      requestHeaders?: Dom.RequestInit["headers"]
    ): Promise<GetPokemonQuery> {
      return withWrapper(
        (wrappedRequestHeaders) =>
          client.request<GetPokemonQuery>(GetPokemonDocument, variables, {
            ...requestHeaders,
            ...wrappedRequestHeaders,
          }),
        "getPokemon",
        "query"
      );
    },
    setCatched(
      variables: SetCatchedMutationVariables,
      requestHeaders?: Dom.RequestInit["headers"]
    ): Promise<SetCatchedMutation> {
      return withWrapper(
        (wrappedRequestHeaders) =>
          client.request<SetCatchedMutation>(SetCatchedDocument, variables, {
            ...requestHeaders,
            ...wrappedRequestHeaders,
          }),
        "setCatched",
        "mutation"
      );
    },
  };
}
export type Sdk = ReturnType<typeof getSdk>;
